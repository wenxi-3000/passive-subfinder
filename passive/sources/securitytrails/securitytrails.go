package securitytrails

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

func Securitytrails(options libs.Options) {
	log.Println("==Securitytrails==")
	var results []string

	requestUrl := "https://api.securitytrails.com/v1/domain/" + options.Domain + "/subdomains?apikey=" + libs.SecuritytrailsKey

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
		result := record + "." + options.Domain
		results = append(results, result)
	}

	results = utils.RemoveDuplicateElement(results)
	utils.SaveTmp(results, "securitytrails_domain.txt", options.TmpPath)
}
