package censys

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/passive-subfinder/libs"
	"github.com/passive-subfinder/utils"

	"github.com/corpix/uarand"
)

type resultsq struct {
	Data []string `json:"parsed.names"`
}

type response struct {
	Results  []resultsq `json:"results"`
	Metadata struct {
		Page  int `json:page`
		Pages int `json:"pages"`
	} `json:"metadata"`
}

type BasicAuth struct {
	Username string
	Password string
}

func Censys(options libs.Options) {
	log.Println("==Censys==")
	UID := libs.CensysUid
	SECRET := libs.CensysSecret
	delay := 3
	page := 1
	var censysResponse response
	// var results []string
	var results []string
	for {
		var data = []byte(`{"query":"` + options.Domain + `", "page":` + strconv.Itoa(page) + `, "fields":["parsed.names"], "flatten":true}`)
		time.Sleep(time.Second * time.Duration(delay))
		resp, err := HTTPRequest(
			"POST",
			"https://www.censys.io/api/v1/search/certificates",
			"",
			map[string]string{"Content-Type": "application/json", "Accept": "application/json"},
			bytes.NewReader(data),
			BasicAuth{Username: UID, Password: SECRET},
		)
		if err != nil {
			log.Println(err)
			continue
		}

		if resp.Status == "200 OK" {

			err = json.NewDecoder(resp.Body).Decode(&censysResponse)
			if err != nil {
				log.Println(err)
			}
			//fmt.Println(censysResponse.Results)
			// fmt.Println("page: ", censysResponse.Metadata.Page)
			if len(censysResponse.Results) == 0 {
				//fmt.Println(len(censysResponse.Results))
				break
			}
			// // strconv.
			for _, record := range censysResponse.Results {
				for _, i := range record.Data {
					//fmt.Println(i)
					results = append(results, i)
				}

				// results = append(results, record.Data[i])
				// utils.SaveTmp(results, "censys_domain.txt")
			}
			// utils.SaveTmp(results, "censys_domain.txt")
			//fmt.Println(json.NewDecoder(resp.Body))

			//fmt.Println(len(response.Results))
			resp.Body.Close()
			//fmt.Println(page)
			page++
		} else {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
			}
			log.Println(string(body))
			break
		}

	}
	results = utils.RemoveDuplicateElement(results)
	utils.SaveTmp(results, "censys_domain.txt", options.TmpPath)
}

func HTTPRequest(method, requestURL, cookies string, headers map[string]string, body io.Reader, basicAuth BasicAuth) (*http.Response, error) {
	req, err := http.NewRequest(method, requestURL, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", uarand.GetRandom())
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en")
	req.Header.Set("Connection", "close")

	if basicAuth.Username != "" || basicAuth.Password != "" {
		req.SetBasicAuth(basicAuth.Username, basicAuth.Password)
	}

	if cookies != "" {
		req.Header.Set("Cookie", cookies)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// proxy, _ := url.Parse("http://127.0.0.1:8080")
	// tr := &http.Transport{
	// 	Proxy:           http.ProxyURL(proxy),
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }

	client := &http.Client{
		Timeout: time.Second * 10, //超时时间
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
