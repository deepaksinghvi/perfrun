package utils

import (
	"bytes"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const (
	getAction = "GET"
	postAction = "POST"
	contentType = "Content-Type"
)
func GetRequestWithClient(url string, client http.Client) ([]byte, error) {
	method := getAction
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Error(err)
	}
	res, err := client.Do(req)
	if nil != err {
		return nil,err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func PostRequest(url string, payload *bytes.Buffer, contentTypeString string, client http.Client) (*http.Response, error) {
	method := postAction
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set(contentType, contentTypeString)
	res, err := client.Do(req)
	if(nil != err ) {
		return nil, errors.New("Request Error")
	}
	return res,nil;
}

func GetClient() *http.Client {
	proxyEnabled := IsProxyEnabled()
	client := &http.Client{}
	if proxyEnabled {
		proxyUrl, _ := url.Parse(GetProxy())
		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	}
	return client
}

func IsProxyEnabled() bool {
	proxy := false
	proxyStr := os.Getenv("PROXY_ENABLED")
	if len(proxyStr) == 0 {
		log.Info("IsProxyEnabled: false")
		return false
	}
	proxy, _ = strconv.ParseBool(proxyStr)
	return proxy
}

func GetProxy() string {
	host := GetEnv("PROXY_HOST","172.20.110.83")
	port := GetEnv("PROXY_PORT","4000")
	return "http://" + host + ":" + port
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}