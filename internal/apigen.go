package internal

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	ExitSuccess = 0
	ExitError   = 1
)

type APIGen struct {
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

	_, err := a.parseFlags(args)
	if err != nil {
		fmt.Fprintln(a.StdErr, "Error:", err)
	}

	return ExitSuccess
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
