package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Browser struct {
	cookies []*http.Cookie
	client  *http.Client
}

//初始化
func NewBrowser() *Browser {
	hc := &Browser{}
	hc.client = &http.Client{}
	//为所有重定向的请求增加cookie
	hc.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) > 0 {
			for _, v := range hc.GetCookie() {
				req.AddCookie(v)
			}
		}
		return nil
	}
	return hc
}

//设置代理地址
func (self *Browser) SetProxyUrl(proxyUrl string) {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(proxyUrl)
	}
	transport := &http.Transport{Proxy: proxy}
	self.client.Transport = transport
}

//设置请求cookie
func (self *Browser) AddCookie(cookies []*http.Cookie) {
	self.cookies = append(self.cookies, cookies...)
}

//获取当前所有的cookie
func (self *Browser) GetCookie() []*http.Cookie {
	return self.cookies
}

//发送Get请求
func (self *Browser) HttpGet(requestUrl string) ([]byte, int) {
	request, _ := http.NewRequest("GET", requestUrl, nil)
	self.setRequestCookie(request)
	response, _ := self.client.Do(request)
	defer response.Body.Close()

	data, _ := ioutil.ReadAll(response.Body)
	return data, response.StatusCode
}

//发送Post请求
func (self *Browser) HttpPost(requestUrl string, params map[string]string) []byte {
	postData := self.encodeParams(params)
	request, _ := http.NewRequest("POST", requestUrl, strings.NewReader(postData))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	self.setRequestCookie(request)

	response, _ := self.client.Do(request)
	defer response.Body.Close()

	//保存响应的 cookie
	respCks := response.Cookies()
	self.cookies = append(self.cookies, respCks...)

	data, _ := ioutil.ReadAll(response.Body)
	return data
}

//发送Post json请求
func (self *Browser) HttpPostJson(requestUrl string, params map[string]string) []byte {
	postData, error := json.Marshal(params)
	if error != nil {
		return []byte(error.Error())
	}
	postData1 := string(postData)
	request, _ := http.NewRequest("POST", requestUrl, strings.NewReader(postData1))
	request.Header.Set("Content-Type", "application/json")
	self.setRequestCookie(request)

	response, _ := self.client.Do(request)
	defer response.Body.Close()

	//保存响应的 cookie
	respCks := response.Cookies()
	self.cookies = append(self.cookies, respCks...)

	data, _ := ioutil.ReadAll(response.Body)
	return data
}

//为请求设置 cookie
func (self *Browser) setRequestCookie(request *http.Request) {
	for _, v := range self.cookies {
		request.AddCookie(v)
	}
}

//参数 encode
func (self *Browser) encodeParams(params map[string]string) string {
	paramsData := url.Values{}
	for k, v := range params {
		paramsData.Set(k, v)
	}
	return paramsData.Encode()
}
