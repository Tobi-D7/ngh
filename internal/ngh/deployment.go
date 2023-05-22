package ngh

type Deployment struct {
	Path string `json:"path"`
	Name string `json;"name"`
	Main string `json:"main"`
	OS   string `json:"os"`
	ARCH string `json:"arch"`
}
