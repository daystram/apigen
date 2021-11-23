package internal

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/daystram/apigen/internal/definition"
)

const (
	ExitSuccess = 0
	ExitError   = 1
)

type APIGen struct {
	Src    string
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
	}
	a.Src = fs.Arg(0)

	return a.Run()
}

func (a *APIGen) parseFlags(args []string) (*flag.FlagSet, error) {
	fs := flag.NewFlagSet("apigen", flag.ExitOnError)
	fs.SetOutput(a.StdErr)
	fs.Usage = func() {
		fmt.Fprintln(a.StdErr, "apigen is an HTTP Rest API server generator.")
		fmt.Fprintln(a.StdErr, "Usage: apigen [definition]")
		fs.PrintDefaults()
	}
	err := fs.Parse(args)
	if err != nil {
		return nil, err
	}
	return fs, nil
}

func (a *APIGen) Run() int {
	p := definition.NewParser()
	_, err := p.ParseFile(a.Src)
	if err != nil {
		fmt.Fprintln(a.StdErr, "Error:", err)
		return ExitError
	}

	return ExitSuccess
}
