package views

import (
	"fmt"
)

type ComponentLicense struct {
	License

	DiDiExpiredDate int64 `json:"DiDiExpiredDate"`
}

type DasLicense struct {
	ComponentLicense

	MaxCpuCores    int64 `json:"maxCpuCores"`
	MaxMemoryBytes int64 `json:"maxMemoryBytes"`
}

type DwsLicense struct {
	ComponentLicense

	MaxCpuCores       int64    `json:"maxCpuCores"`
	MaxMemoryBytes    int64    `json:"maxMemoryBytes"`
	MaxGpuUnits       int64    `json:"maxGpuUnits"`
	AllowAll          bool     `json:"allowAll"`
	OperatorWhitelist []string `json:"operatorWhitelist"`
}

func (cl *ComponentLicense) IsComponentExpired() (valid bool, err error) {
	valid = false

	if cl.DiDiExpiredDate <= nowMs() {
		err = fmt.Errorf("%s license is expired.", cl.Product)
		return
	}

	valid = true
	return
}
