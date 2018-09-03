package swagtype

type DasLicense struct {
	Product           string   `json:"product"`
	VersionsSupported []string `json:"versionsSupported"`
	ExpiredDate       string   `json:"expiredDate"`
	MaxCpuCores       int64    `json:"maxCpuCores"`
	MaxMemoryBytes    int64    `json:"maxMemoryBytes"`
	DiDiExpiredDate   string   `json:"DiDiExpiredDate"`
}
