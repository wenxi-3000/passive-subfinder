package virustotal

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/passive-subfinder/libs"
	"github.com/passive-subfinder/utils"
)

type subdomains struct {
	Resresults []string `json:"subdomains"`
}

func Virustotal(options libs.Options) {
	log.Println("==Virustotal==")
	var results []string
	requestUrl := "https://www.virustotal.com/vtapi/v2/domain/report?apikey=" + libs.VirustotalKey + "&domain=" + options.Domain
	resp, err := http.Get(requestUrl)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	var respSubdomain subdomains
	err = json.NewDecoder(resp.Body).Decode(&respSubdomain)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(respSubdomain.Resresults)
	for _, record := range respSubdomain.Resresults {
		//result := record + "." + domain
		results = append(results, record)
	}

	results = utils.RemoveDuplicateElement(results)
	utils.SaveTmp(results, "virustotal_domain.txt", options.TmpPath)
}
