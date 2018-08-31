package views

type DiDiLicense struct {
	License

	Key     string `json:"key"`
	Version string `json:"version"`
	Edition string `json:"edition"`
	Type    string `json:"type"`
}
