package views

import (
	"fmt"
)

type License struct {
	Product           string   `json:"product"`
	VersionsSupported []string `json:"versionsSupported"`
	IssuedDate        int64    `json:"issuedDate"`
	ExpiredDate       int64    `json:"expiredDate"`
}

func (l *License) IsExpired() (valid bool, err error) {
	valid = false

	if l.ExpiredDate <= nowMs() {
		err = fmt.Errorf("%s license is expired.", l.Product)
		return
	}

	valid = true
	return
}
