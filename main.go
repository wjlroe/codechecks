package main

import (
	"bytes"
	"flag"
	"log"
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
		log.Fatal(err)
	}
}

func main() {
	var zigFmt bool
	flag.BoolVar(&zigFmt, "zig-fmt", false, "run zig fmt on files")
	flag.Parse()
	if zigFmt {
		runZigFmt(flag.Args())
	} else {
		log.Fatal("Unknown mode")
	}
}
