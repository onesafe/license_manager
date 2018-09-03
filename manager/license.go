package manager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"

	"github.com/onesafe/license_manager/cipher"
	"github.com/onesafe/license_manager/db"
	"github.com/onesafe/license_manager/modules"
	"github.com/onesafe/license_manager/swagtype"
	"github.com/onesafe/license_manager/utils"
	"github.com/onesafe/license_manager/views"
)

type LicenseManager struct {
	Manager
}

func GetLicenseManager() *LicenseManager {
	return &LicenseManager{
		Manager: NewManager(),
	}
}

func (m *LicenseManager) RegisterPath() error {
	m.apiRouter.Register("POST", "/license-manager/v1/licenses/upload", m.handlerLicenseUpload)
	m.apiRouter.Register("GET", "/license-manager/v1/licenses", m.handlerListLicenses)

	m.apiRouter.Register("POST", "/license-manager/v1/daslicense", m.handlerGenDasLicense)

	m.apiRouter.Register("GET", "/license-manager/v1/rsakeys", m.handlerGenRSAKeys)

	return nil
}

// @Summary Generate Das License
// @Description Generate Das License
// @Accept  json
// @Produce  json
// @Param body body swagtype.DasLicense true "Das license"
// @Success 200 {object} utils.Response
// @Router /daslicense [post]
func (m *LicenseManager) handlerGenDasLicense(ctx *gin.Context) {

	var inputdata swagtype.DasLicense
	err := ctx.ShouldBindJSON(&inputdata)
	if err != nil {
		utils.BadRequestResp(ctx, "Invalid Json data: "+err.Error())
	}

	DiDiExpiredDate, _ := time.Parse("2006-01-02", inputdata.DiDiExpiredDate)
	ExpiredDate, _ := time.Parse("2006-01-02", inputdata.ExpiredDate)

	license := views.License{
		IssuedDate:        int64(time.Now().UnixNano() / 1000000),
		ExpiredDate:       int64(ExpiredDate.UnixNano() / 1000000),
		Product:           inputdata.Product,
		VersionsSupported: inputdata.VersionsSupported,
	}
	Component := views.ComponentLicense{
		DiDiExpiredDate: int64(DiDiExpiredDate.UnixNano() / 1000000),
		License:         license,
	}
	data := views.DasLicense{
		ComponentLicense: Component,
		MaxCpuCores:      inputdata.MaxCpuCores,
		MaxMemoryBytes:   inputdata.MaxMemoryBytes,
	}

	// convert license to []byte
	buf, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))

	// encrypt license
	encryptedLicense, err := cipher.EncryptLicense(buf)
	if err != nil {
		utils.BadRequestResp(ctx, "Encrypt License error: "+err.Error())
	}
	log.Println(encryptedLicense)

	utils.OkResp(ctx, "Encrypted License Success", encryptedLicense)
}

/*
curl -X POST http://localhost:8080/licenses/upload -F "file=@/Users/name/Documents/licensefile"
 -H "Content-Type: multipart/form-data"
*/
// @Summary Upload License File
// @Description Upload License File
// @Accept  multipart/form-data
// @Produce  json
// @Param file formData file true "license file"
// @Success 200 {object} utils.Response
// @Router /licenses/upload [post]
func (m *LicenseManager) handlerLicenseUpload(ctx *gin.Context) {
	licenseFile, err := ctx.FormFile("file")
	if err != nil {
		msg := "Please upload license file"
		fmt.Println(msg)
		utils.BadRequestResp(ctx, msg)
		return
	}
	log.Println("license file Name: " + licenseFile.Filename)

	// tempfile to store upload file content
	tmpfile, err := ioutil.TempFile("", "tmpfile")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	// save file content and read
	dst := tmpfile.Name()
	ctx.SaveUploadedFile(licenseFile, dst)
	data, err := ioutil.ReadFile(dst)
	if err != nil {
		msg := "file content read error"
		fmt.Println(msg)
		utils.BadRequestResp(ctx, msg)
		return
	}

	// print license file content
	log.Println("license content is: \n" + string(data))

	// Decrypt License
	var lcs *views.License
	lcs, err = cipher.DecryptLicense(string(data))
	if err != nil {
		msg := "Decrypt License error"
		fmt.Println(msg)
		utils.BadRequestResp(ctx, msg+err.Error())
		return
	}
	log.Println("license parse finished")

	// Register License
	err = registerLicense(lcs, string(data))
	if err != nil {
		msg := "Register License error"
		utils.BadRequestResp(ctx, msg+err.Error())
		return
	}

	utils.OkResp(ctx, fmt.Sprintf("%s uploaded!", licenseFile.Filename), nil)
}

// @Summary Get all licenses
// @Description Get all licenses
// @Accept  json
// @Produce  json
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Success 200 {object} utils.Response
// @Router /licenses [get]
func (m *LicenseManager) handlerListLicenses(ctx *gin.Context) {
	log.Println("Get all licenses")
	page := ctx.Query("page")
	size := ctx.Query("size")

	if page == "" {
		page = "0"
	}
	if size == "" {
		size = "10"
	}
	maps := make(map[string]interface{})

	lr := &db.License_record{}
	lr.GetLicenses(com.StrTo(page).MustInt(), com.StrTo(size).MustInt(), maps)
	fmt.Println(lr)

	utils.OkResp(ctx, "OK", lr)
}

// @Summary Generate rsa keys
// @Description Generate rsa keys
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.Response
// @Router /rsakeys [get]
func (m *LicenseManager) handlerGenRSAKeys(ctx *gin.Context) {
	log.Println("Generate RSA private key and public key")
	privateKey, publicKey, err := cipher.GenRsaKeys(2048)
	if err != nil {
		utils.BadRequestResp(ctx, "generate RSA private key and public key error"+err.Error())
	}

	rsakeys := modules.RSAKeys{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}
	utils.OkResp(ctx, "Generate RSA keys Success", rsakeys)
}
