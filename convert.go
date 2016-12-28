package main

import (
	"fmt"
	"sync"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"path/filepath"
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err)
	} else {
		var wg sync.WaitGroup
		for _, f := range files {
			file_path := f.Name()
			ext := filepath.Ext(file_path)
			if ext == ".png" {
				s := strings.Split(file_path, ".")
				filename := s[0]
				fmt.Println(filename)
				wg.Add(1)
				name := filename
				go func() {
					exec.Command("convert", name+".png", name+".jpeg").Run()
					exec.Command("convert", name+".jpeg", "eps2:"+name+".eps").Run()
					if err := os.Remove(name+".jpeg"); err != nil {
						fmt.Println(err)
					}
					wg.Done()
				}()
			}
		}
		wg.Wait()
	}
}
