package version

type Version struct {
	Name  string `json:"name"`
	Version string `json:"version"`
}
  
var CurrentVersion = &Version{
	Name:  "Farm Monitor API",
	Version: "v0.0.1",
}