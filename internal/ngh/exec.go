package ngh

import (
	"fmt"
	"os"
	"os/exec"
)

func Build(d *Deployment, ddir string) {
	os.Setenv("GOOS", d.OS)
	os.Setenv("GOOARCH", d.ARCH)
	exec_name := d.Name
	dep_path := ddir + "/" + exec_name + "/" + d.OS + "/" + d.ARCH
	os.MkdirAll(dep_path, os.ModePerm)
	if d.OS == "windows" {
		exec_name = fmt.Sprintf("%s.exe", exec_name)
	}
	fmt.Println("Building: " + exec_name + "...")
	fmt.Println("OS: " + d.OS)
	fmt.Println("ARCH: " + d.ARCH)
	cmd := exec.Command("go", "build", "-o", dep_path+"/"+exec_name, d.Path+"/"+d.Main)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("CMD_OUT: " + string(output))
		fmt.Println("Error:")
		panic(err)
	}
}
