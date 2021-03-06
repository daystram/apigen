package writer

import (
	"go/printer"
	"go/token"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/daystram/apigen/internal/generator"
)

type Writer struct {
	fs      *token.FileSet
	rootDir string
}

func NewWriter(fs *token.FileSet, rootDir string) *Writer {
	return &Writer{fs: fs, rootDir: rootDir}
}

func (w *Writer) Write(f generator.File) error {
	dir := path.Join(w.rootDir, f.Dir)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	out, err := os.OpenFile(path.Join(dir, f.Name), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer out.Close()

	return printer.Fprint(out, w.fs, f.AST)
}

func (w *Writer) InitMod(pkg string) error {
	cmd := exec.Command("go", "mod", "init", pkg)
	cmd.Dir = "./" + filepath.Base(w.rootDir)
	err := cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = filepath.Base(w.rootDir)
	return cmd.Run()
}
