package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"sort"

	"github.com/tsavola/gols/internal/goversion"
)

func main() {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		home := os.Getenv("HOME")
		if home == "" {
			log.Fatal("HOME not set")
		}
		gopath = path.Join(home, "go")
	}
	gobin := path.Join(gopath, "bin")

	dir, err := os.Open(gobin)
	if err != nil {
		log.Fatal(err)
	}

	names, err := dir.Readdirnames(-1)
	if err != nil {
		log.Fatal(err)
	}

	sort.Strings(names)

	for _, name := range names {
		printIfGo(path.Join(gobin, name))
	}
}

func printIfGo(path string) {
	fi, err := os.Stat(path)
	if err != nil {
		return
	}

	version, ok := goversion.Report(path, path, fi)
	if ok {
		fmt.Printf("%v\t%v\n", version, path)
	}
}
