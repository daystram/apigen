package internal

import (
	"flag"
	"fmt"
	"go/token"
	"io"
	"os"
	"path/filepath"

	"github.com/daystram/apigen/internal/definition"
	"github.com/daystram/apigen/internal/generator"
	"github.com/daystram/apigen/internal/writer"
)

const (
	ExitSuccess = 0
	ExitError   = 1
)

var (
	ErrUsage = fmt.Errorf("incorrect usage")
)

type APIGen struct {
	Src    string
	Pkg    string
	StdOut io.Writer
	StdErr io.Writer
	StdIn  io.Reader
}

func Main(args []string) int {
	a := &APIGen{
		StdOut: os.Stdout,
		StdErr: os.Stderr,
		StdIn:  os.Stdin,
	}

	fs, err := a.parseFlags(args)
	if err != nil {
		fmt.Fprintln(a.StdErr, "Error:", err)
		return ExitError
	}
	a.Src = fs.Arg(0)
	a.Pkg = filepath.Clean(fs.Arg(1))

	return a.Run()
}

func (a *APIGen) parseFlags(args []string) (*flag.FlagSet, error) {
	fs := flag.NewFlagSet("apigen", flag.ExitOnError)
	fs.SetOutput(a.StdErr)
	fs.Usage = func() {
		fmt.Fprintln(a.StdErr, "apigen is an HTTP Rest API server generator.")
		fmt.Fprintln(a.StdErr, "Usage: apigen [definition] [packagepath]")
		fs.PrintDefaults()
	}
	err := fs.Parse(args)
	if err != nil {
		return nil, err
	}

	if fs.NArg() != 2 {
		return nil, fmt.Errorf("expected 2 arguments, got %d", fs.NArg())
	}

	return fs, nil
}

func (a *APIGen) Run() int {
	parser := definition.NewParser()
	def, err := parser.ParseFile(a.Src)
	if err != nil {
		a.printErr(err)
		return ExitError
	}

	fg, err := generator.Generate(def, a.Pkg)
	if err != nil {
		a.printErr(err)
		return ExitError
	}

	w := writer.NewWriter(token.NewFileSet(), filepath.Base(a.Pkg))
	for _, f := range fg {
		err = w.Write(f)
		if err != nil {
			a.printErr(err)
			return ExitError
		}
	}
	err = w.InitMod(a.Pkg)
	if err != nil {
		a.printErr(err)
		return ExitError
	}
	return ExitSuccess
}

func (a *APIGen) printErr(err error) {
	fmt.Fprintln(a.StdErr, "Error:", err)
}
