package main

import (
	"os"

	"github.com/daystram/apigen/internal"
)

func main() {
	os.Exit(internal.Main(os.Args[1:]))
}
