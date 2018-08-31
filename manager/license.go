package manager

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/onesafe/license_manager/utils"
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

	return nil
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
// @Success 200 {string} string	"ok"
// @Router /licenses/upload [post]
func (m *LicenseManager) handlerLicenseUpload(ctx *gin.Context) {
	licenseFile, err := ctx.FormFile("file")
	if err != nil {
		msg := "Please upload license file"
		fmt.Println(msg)
		utils.BadRequestResp(ctx, msg)
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
	}

	// print license file content
	log.Println("license content is: \n" + string(data))

	utils.OkResp(ctx, fmt.Sprintf("%s uploaded!", licenseFile.Filename), nil)
}

// @Summary Get all licenses
// @Description Get all licenses
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /licenses [get]
func (m *LicenseManager) handlerListLicenses(ctx *gin.Context) {
	log.Println("Get all licenses")

	utils.OkResp(ctx, "OK", nil)
}
