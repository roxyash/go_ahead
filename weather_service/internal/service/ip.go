package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type IpService struct {
}

func NewIpService() *IpService {
	return &IpService{}
}

type IP struct {
	Query string
}

func (s *IpService) Get() (string, error) {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return "", nil
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", nil
	}
	var ip IP
	json.Unmarshal(body, &ip)
	return ip.Query, nil
}
