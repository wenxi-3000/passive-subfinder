package fofa

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/passive-subfinder/libs"
	"github.com/passive-subfinder/utils"
)

type fofaResponse struct {
	Results []string `json:"results"`
}

func Fofa(options libs.Options) {
	log.Println("==Fofa==")
	email := libs.FofaEmail
	key := libs.FofaKey
	page := 1
	qbase64 := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("domain=\"%s\"", options.Domain)))

	//fmt.Println(url)

	var response fofaResponse
	var results []string
	for {
		//最多只能查看10000条数据，如需获取 更多数据，请在搜索界面点击"下载数据"链接进行下载。
		if page >= 1001 {
			break
		}

		url := "https://fofa.so/api/v1/search/all?full=true&fields=host&page=" + strconv.Itoa(page) + "&size=10&email=" + email + "&key=" + key + "&qbase64=" + qbase64
		// fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
		}
		// body, err := ioutil.ReadAll(resp.Body)
		// fmt.Println(string(body))
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			log.Println(err)
		}
		//fmt.Println(response.Results)
		if len(response.Results) == 0 {
			break
		}

		for _, records := range response.Results {
			lines := utils.GetSubomains(records, options.Domain)
			for _, line := range lines {
				//fmt.Println(line)
				results = append(results, line)
			}

		}

		//fmt.Println(json.NewDecoder(resp.Body))
		//fmt.Print(response.Results)
		//fmt.Println(len(response.Results))

		resp.Body.Close()
		//fmt.Println("page: ", page)
		page++
		time.Sleep(1 * time.Second)

	}
	// for i := range results {
	// 	iprecords, _ := dns.ReverseAddr(results[i])
	// 	for _, ip := range iprecords {
	// 		fmt.Print(ip)
	// 	}
	// }
	results = utils.RemoveDuplicateElement(results)
	utils.SaveTmp(results, "fofa_domain.txt", options.TmpPath)

}
