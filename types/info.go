package types

type PackageInfo struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Tree *Tree  `json:"tree"`
}
