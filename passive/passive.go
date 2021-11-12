package passive

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/passive-subfinder/libs"
	"github.com/passive-subfinder/passive/sources/alienvault"
	"github.com/passive-subfinder/passive/sources/censys"
	"github.com/passive-subfinder/passive/sources/fofa"
	"github.com/passive-subfinder/passive/sources/qianxun"
	"github.com/passive-subfinder/passive/sources/securitytrails"
	"github.com/passive-subfinder/passive/sources/virustotal"
	"github.com/passive-subfinder/utils"
)

var wg sync.WaitGroup

func Passive(options libs.Options) {
	log.Println("开始收集子域名：" + options.Domain)
	tasks := []string{
		"alienvault",
		"fofa",
		"sensys",
		"qianxun",
		"securitytrails",
		"virustotal",
	}

	//并发查询
	for _, task := range tasks {
		wg.Add(1)
		go doTask(task, options)
	}
	wg.Wait()

	//读取结果文件，去重，清洗
	results := GetPassiveResult(options.Domain, options.TmpPath)
	for _, result := range results {
		fmt.Println(result)
	}

	//保存结果
	utils.SaveResult(results, "passive_sudomain.txt", options.JobPath)
	log.Println("被动收集的域名已完成,结果保存在: ", options.JobPath)

	os.Exit(0)
}

func doTask(task string, options libs.Options) {
	switch {
	case task == "alienvault":
		alienvault.Alienvault(options)
		wg.Done()

	case task == "fofa":
		fofa.Fofa(options)
		wg.Done()

	case task == "sensys":
		censys.Censys(options)
		wg.Done()

	case task == "qianxun":
		qianxun.Qianxun(options)
		wg.Done()

	case task == "securitytrails":
		securitytrails.Securitytrails(options)
		wg.Done()

	case task == "virustotal":
		virustotal.Virustotal(options)
		wg.Done()

	}
}

func GetPassiveResult(domain string, tmpPath string) []string {
	var results []string
	files, _ := ioutil.ReadDir(tmpPath)
	for _, f := range files {
		filePath := path.Join(tmpPath, f.Name())
		file, err := os.OpenFile(filePath, os.O_RDONLY, 444)
		if err != nil {
			log.Println(err)
		}
		bufFile := bufio.NewReader(file)
		for {
			line, err := bufFile.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				} else {
					log.Println(err)
				}
			}

			domainline := strings.TrimSpace(line)
			if strings.Contains(domainline, "*.") {
				domainline = strings.Replace(domainline, "*.", "", 1)
			}
			if strings.Contains(domainline, domain) {
				results = append(results, domainline)
			}

		}
		file.Close()
	}

	results = utils.RemoveDuplicateElement(results)
	return results

}
