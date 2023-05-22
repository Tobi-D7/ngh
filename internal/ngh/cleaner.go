package ngh

import "os"

func Clean(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		panic(err)
	}
}
