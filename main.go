package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// This is to work around bugs in `zig fmt`, which can't parse filepaths on Windows
func runZigFmt(paths []string) {
	fixedArgs := []string{"fmt"}
	for _, arg := range paths {
		fixedArgs = append(fixedArgs, filepath.FromSlash(arg))
	}
	cmd := exec.Command("zig", fixedArgs...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Command (%s) errored\n", cmd)
		log.Fatal(err)
	}
}

func fixZigFiles() error {
	files := []string{}
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// TODO: Exclude files ignored by gitignore and the like
		if info.IsDir() && info.Name() == "zig-cache" {
			return filepath.SkipDir
		}
		if filepath.Ext(path) == ".zig" {
			files = append(files, path)
		}
		return nil
	})
	if len(files) > 0 {
		runZigFmt(files)
	}
	return err
}

func fixAll() error {
	err := fixZigFiles()
	return err
}

func main() {
	var zigFmt bool
	flag.BoolVar(&zigFmt, "zig-fmt", false, "run zig fmt on files")
	flag.Parse()
	if zigFmt {
		runZigFmt(flag.Args())
	} else {
		err := fixAll()
		if err != nil {
			log.Fatal(err)
		}
	}
}
