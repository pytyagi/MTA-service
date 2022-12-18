package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mta-hosting-optimizer/application"
	"net/http"
)

type IpConfig struct {
	Id       int64  `json:"id"`
	Hostname string `json:"hostname"`
	Ip       string `json:"ip"`
	Active   bool   `json:"active"`
}
type Pair struct {
	first, second int64
}

type IpDataServiceInterface interface {
	GetServiceIpData() (*[]IpConfig, error)
	GetInefficientServers(ipData *[]IpConfig) []string
}

type IpDataService struct {
	HTTPClient *http.Client
}

func NewIpDataService(httpClient *http.Client) *IpDataService {
	return &IpDataService{HTTPClient: httpClient}
}
func (m *IpDataService) GetServiceIpData() (*[]IpConfig, error) {
	log.Println("Getting IpConfig from mock services")

	req, err := http.NewRequest(application.Cfg.Method, application.Cfg.Url, nil)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	res, err := m.HTTPClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Total Server assuming to be 35 as described in the document shared
	ServerData := make([]IpConfig, 35)
	err = json.Unmarshal(body, &ServerData)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &ServerData, nil
}

func (m *IpDataService) GetInefficientServers(ipData *[]IpConfig) []string {
	hostnameMap := make(map[string]Pair)
	for _, val := range *ipData {
		p := hostnameMap[val.Hostname]
		if val.Active {
			hostnameMap[val.Hostname] = Pair{p.first + 1, p.second}
		} else {
			hostnameMap[val.Hostname] = Pair{p.first, p.second + 1}
		}
	}
	var Hosts []string
	for index, val := range hostnameMap {
		if val.first <= application.Cfg.X {
			Hosts = append(Hosts, index)
		}
	}
	return Hosts
}
