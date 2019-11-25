package version

type Version struct {
	Name  string `json:"name"`
	Version string `json:"version"`
}
  
var CurrentVersion = &Version{
	Name:  "Farm Monitor API",
	Version: "0.0.1",
}