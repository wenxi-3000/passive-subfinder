package utils

import (
	"bufio"
	"log"
	"os"
	"path"
	"regexp"
)

//存入临时文件,重置文件
func SaveTmp(results []string, saveFile string, tmpPath string) {
	//fmt.Print(tmpPath)
	tmpFilePath := path.Join(tmpPath, saveFile)
	tmpFile, err := os.OpenFile(tmpFilePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer tmpFile.Close()
	writer := bufio.NewWriter(tmpFile)

	for i := 0; i < len(results); i++ {
		// fmt.Println(results[i])
		writer.WriteString(results[i] + "\n")
	}
	writer.Flush()
}

// //存入临时文件,追加
// func SaveTmpAdd(results []string, saveFile string, resultP) {
// 	//fmt.Print(tmpPath)
// 	tmpFile, err := os.OpenFile(tmpPath+saveFile, os.O_WRONLY|os.O_CREATE, 0666)
// 	if err != nil {
// 		log.Println(nil)
// 		os.Exit(1)
// 	}
// 	defer tmpFile.Close()
// 	writer := bufio.NewWriter(tmpFile)

// 	for i := 0; i < len(results); i++ {
// 		// fmt.Println(results[i])
// 		writer.WriteString(results[i] + "\n")
// 	}
// 	writer.Flush()
// }

func SaveResult(results []string, fileName string, jobPath string) {
	resultFilePath := path.Join(jobPath, fileName)
	resultFile, err := os.OpenFile(resultFilePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(nil)
		os.Exit(1)
	}
	defer resultFile.Close()
	writer := bufio.NewWriter(resultFile)

	for i := 0; i < len(results); i++ {
		// fmt.Println(results[i])
		writer.WriteString(results[i] + "\n")
	}
	writer.Flush()
}

// // //存入结果文件
// // func SaveResult(results []string, saveFile string) {
// // 	//存入临时文件
// // 	resultPath := "/Users/shadowflow/pentest/shadowflow/github.com/github.com/passive-subfinder/results"
// // 	fmt.Print(resultPath)
// // 	tmpFile, err := os.OpenFilePath+"/"+saveFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
// // 	if err != nil {
// // 		fmt.Println(nil)
// // 	}
// // 	defer tmpFile.Close()
// // 	writer := bufio.NewWriter(tmpFile)

// // 	for i := 0; i < len(results); i++ {
// // 		fmt.Println(results[i])
// // 		writer.WriteString(results[i] + "\n")
// // 	}
// // 	writer.Flush()
// // }

//去重
func RemoveDuplicateElement(languages []string) []string {
	result := make([]string, 0, len(languages))
	temp := map[string]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

//匹配域名
func GetSubomains(source string, domain string) []string {
	//strings.Replace(domain,'.', '`\.`')
	reg := `(?:[a-z0-9](?:[a-z0-9\-]{0,61}[a-z0-9])?\.){0,}` + domain
	//results_domains = re.findall(regexp, str(source), re.I)
	var linkFinderRegex = regexp.MustCompile(reg)
	matchs := linkFinderRegex.FindAllString(source, -1)
	//fmt.Println(matchs)
	return matchs
}

//匹配域名不包含本身
func GetSubomainsNot(source string, domain string) []string {
	//strings.Replace(domain,'.', '`\.`')
	reg := `(?:[a-z0-9](?:[a-z0-9\-]{0,61}[a-z0-9])?\.){1,}` + domain
	//results_domains = re.findall(regexp, str(source), re.I)
	var linkFinderRegex = regexp.MustCompile(reg)
	matchs := linkFinderRegex.FindAllString(source, -1)
	//fmt.Println(matchs)
	return matchs
}

// func dnsInfo() {
// 	// it requires a list of resolvers
// 	resolvers := []string{"8.8.8.8:53", "8.8.4.4:53"}
// 	retries := 2
// 	hostname := "hackerone.com"
// 	dnsClient := retryabledns.New(resolvers, retries)

// 	ips, err := dnsClient.Resolve(hostname)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	log.Println(ips)

// 	// Query Types: dns.TypeA, dns.TypeNS, dns.TypeCNAME, dns.TypeSOA, dns.TypePTR, dns.TypeMX
// 	// dns.TypeTXT, dns.TypeAAAA (from github.com/miekg/dns)
// 	dnsResponses, err := dnsClient.Query(hostname, dns.TypeA)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	log.Println(dnsResponses)
// }

// 判断文件是否存在
func FolderExists(foldername string) bool {
	if _, err := os.Stat(foldername); os.IsNotExist(err) {
		return false
	}
	return true
}
