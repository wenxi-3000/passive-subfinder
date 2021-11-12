package main

import (
	"flag"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/passive-subfinder/libs"
	"github.com/passive-subfinder/passive"
	"github.com/passive-subfinder/utils"
)

var options = libs.Options{}
var cliDomain = flag.String("d", "example.com", "Input Your Domain")

func init() {
	//设置日志格式
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	flag.Parse()
	options.Domain = *cliDomain
	//目录初始化
	initPath()
}

func main() {
	passive.Passive(options)

}

func initPath() {
	//当前文件路径
	_, currentFile, _, _ := runtime.Caller(0)
	//当前文件目录
	currentPath := path.Dir(currentFile)
	//结果文件路径
	options.ResultPath = path.Join(currentPath, "results")
	//不存在结果文件夹就创建
	if !utils.FolderExists(options.ResultPath) {
		os.MkdirAll(options.ResultPath, 0750)
	}

	// fmt.Println(options.ResultPath)
	// files, _ := ioutil.ReadDir(options.ResultPath)
	// for _, file := range files {
	// 	println(file.Name())
	// }

	//当前任务保存目录
	options.JobPath = path.Join(options.ResultPath, options.Domain)
	if !utils.FolderExists(options.JobPath) {
		os.MkdirAll(options.JobPath, 0750)
	}
	log.Println("结果保存目录：", options.JobPath)

	//临时文件目录
	options.TmpPath = path.Join(options.JobPath, "tmp")
	if !utils.FolderExists(options.TmpPath) {
		os.MkdirAll(options.TmpPath, 0750)
	}
	//log.Println("临时文件目录：", options.TmpPath)
}
