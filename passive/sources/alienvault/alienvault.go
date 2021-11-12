package alienvault

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/passive-subfinder/libs"
	"github.com/passive-subfinder/utils"
)

//const SUBRE = `(?i)(([a-zA-Z0-9]{1}|[_a-zA-Z0-9]{1}[_a-zA-Z0-9-]{0,61}[a-zA-Z0-9]{1})[.]{1})+`

type alienvaultResponse struct {
	Detail     string `json:"detail"`
	Error      string `json:"error"`
	PassiveDNS []struct {
		Hostname string `json:"hostname"`
		Ip       string `json:"address"`
	} `json:"passive_dns"`
}

// Source is the passive scraping agent
type alienvaultResult struct {
	Domain string
	Ip     string
}

func Alienvault(options libs.Options) {
	//创建临时文件
	log.Println("==Alienvault==")
	resp, err := http.Get("https://otx.alienvault.com/api/v1/indicators/domain/" + options.Domain + "/passive_dns")
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	// var response alienvaultResponse

	// err = json.NewDecoder(resp.Body).Decode(&response)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	var results []string
	// for _, record := range response.PassiveDNS {
	// 	recordStruct := alienvaultResult{
	// 		Domain: record.Hostname,
	// 		Ip:     record.Ip,
	// 	}
	// 	result, err := json.Marshal((recordStruct))
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	results = append(results, string(result))
	// }
	// fmt.Println(results)
	// utils.SaveTmp(results, "alienvault.txt")
	// log.Println(domain)
	// fmt.Println(resp.Body)
	// body, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	matchs := utils.GetSubomains(string(body), options.Domain)
	for _, sub := range matchs {
		results = append(results, sub)
	}
	results = utils.RemoveDuplicateElement(results)
	utils.SaveTmp(results, "alienvault_domain.txt", options.TmpPath)
	time.Sleep(time.Second * 2)

}

// func main() {
// 	Alienvault("lu.com")
// }

//func GetSubdomains(source, domain string) []string {
//	var subs []string
//	fmt.Println(source)
//
//	for _, match := range re.FindAllStringSubmatch(source, -1) {
//		subs = append(subs, match[0])
//	}
//	return subs
//}

// func GetSubomains(source string, domain string) []string {
// 	//strings.Replace(domain,'.', '`\.`')
// 	reg := `(?:[a-z0-9](?:[a-z0-9\-]{0,61}[a-z0-9])?\.){0,}` + domain
// 	//results_domains = re.findall(regexp, str(source), re.I)
// 	var linkFinderRegex = regexp.MustCompile(reg)
// 	matchs := linkFinderRegex.FindAllString(source, -1)
// 	//fmt.Println(matchs)
// 	return matchs
// }
