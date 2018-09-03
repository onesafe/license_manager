package manager

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/onesafe/license_manager/cipher"
	"github.com/onesafe/license_manager/manager/v1"
	"github.com/onesafe/license_manager/modules"
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
	m.apiRouter.Register("POST", "/license-manager/v1/licenses/upload", v1.LicenseUpload)
	m.apiRouter.Register("GET", "/license-manager/v1/licenses", v1.ListLicenses)
	m.apiRouter.Register("POST", "/license-manager/v1/daslicense", v1.GenDasLicense)

	m.apiRouter.Register("GET", "/license-manager/v1/rsakeys", m.handlerGenRSAKeys)

	return nil
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
