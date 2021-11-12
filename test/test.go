package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/github.com/passive-subfinder/libs"
)

func main() {
	//qianxun.Qianxun("teambition.com")
	GetProjectAbsPath()
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("xxx")
}

func GetProjectAbsPath() {
	programPath, _ := filepath.Abs(os.Args[0])

	fmt.Println("programPath:", programPath)

	projectAbsPath := path.Dir(path.Dir(programPath))

	fmt.Println("PROJECT_ABS_PATH:", projectAbsPath)

	BinPathPrefix := filepath.Join(projectAbsPath, "bin")
	fmt.Println(BinPathPrefix)

	file, _ := exec.LookPath(os.Args[0])
	fmt.Println(file)

}

func Testx(opt *libs.Options) {
	log.Println("logxx")

}
