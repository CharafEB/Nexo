package midel

type StructBuild struct {
	File string `json:"file"`
	Folder string `json:"folder"`
}

type Blueprints struct {
	Name string `json:"name"`
	Libs string `json:"libs"`
	StructType string `json:"struct"`
	StructBuild StructBuild `json:"structbuild"`
}