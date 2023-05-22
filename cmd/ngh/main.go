package main

import (
	"fmt"
	"os"

	"github.com/Tobi-D7/ngh/internal/ngh"
)

// Simple Stupid Deploy Handler

func main() {
	cfg, err := ngh.ParseConfig("ngo_prj.json")
	if err != nil {
		panic(err)
	}
	if len(os.Args) < 2 {
		fmt.Println("No argument Specified!")
		os.Exit(0)
	}
	if os.Args[1] == "clean" {
		ngh.Clean(cfg.OutDir)
	} else if os.Args[1] == "deploy" {
		os.MkdirAll(cfg.OutDir, os.ModePerm)
		for _, obj := range cfg.Deployments {
			ngh.Build(&obj, cfg.OutDir)
		}
	} else {
		fmt.Println("Argument not Supported")
	}
}
