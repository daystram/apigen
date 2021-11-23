package generator


type FileGroup []File

type File struct {
	Name string
	Dir  string
	AST  *ast.File
}
