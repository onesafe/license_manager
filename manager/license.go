package manager

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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
	m.apiRouter.Register("POST", "/licenses/upload", m.handlerLicenseUpload)
	m.apiRouter.Register("GET", "/licenses", m.handlerListLicenses)

	return nil
}

/*
curl -X POST http://localhost:8080/licenses/upload -F "file=@/Users/name/Documents/licensefile"
 -H "Content-Type: multipart/form-data"
*/
func (m *LicenseManager) handlerLicenseUpload(ctx *gin.Context) {
	licenseFile, err := ctx.FormFile("file")
	if err != nil {
		msg := "Please upload license file"
		fmt.Println(msg)
		ctx.String(http.StatusBadRequest, msg)
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
		ctx.String(http.StatusBadRequest, msg)
	}

	// print license file content
	log.Println("license content is: \n" + string(data))

	ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", licenseFile.Filename))
}

func (m *LicenseManager) handlerListLicenses(ctx *gin.Context) {
	log.Println("Get all licenses")

	ctx.String(http.StatusOK, "OK")
}
