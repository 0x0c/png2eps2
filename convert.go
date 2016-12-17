package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"path/filepath"
)

func main() {
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		file_path := f.Name()
		ext := filepath.Ext(file_path)
		if ext == ".png" {
			s := strings.Split(file_path, ".")
			filename := s[0]
			fmt.Println(filename)
			exec.Command("convert", filename+".png", filename+".jpeg").Run()
			exec.Command("convert", filename+".jpeg", "eps2:"+filename+".eps").Run()
			if err := os.Remove(filename+".jpeg"); err != nil {
				fmt.Println(err)
			}
		}
	}
}
