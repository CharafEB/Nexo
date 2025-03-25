package midel

type StructBuildF struct {
	File string `json:"file"`
	Folder string `json:"folder"`
}
type Structrs struct {
	StructCall string `json:"struct"`
	StructBuild StructBuildF `json:"structbuild"`

}

type Blueprints struct {
	Name string `json:"name"`
	Libs string `json:"libs"`
	StructType string `json:"struct"`
}