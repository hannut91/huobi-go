package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Options struct {
	Method      string
	URL         string
	Data        []byte
	Headers     map[string]string
	QueryParams map[string]string
}

func CreateQueryString(
	method string,
	baseURL string,
	path string,
	params map[string]string,
) string {
	var str string

	for k, v := range params {
		str += k + "=" + v + "&"
	}

	result := fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		method,
		baseURL,
		path,
		str[:len(str)-1],
	)
	return result
}

func CreateParams(params map[string]string) string {
	str := "?"

	for k, v := range params {
		str += k + "=" + v + "&"
	}

	return str[:len(str)-1]
}

func CreateTimestamp() string {
	now := time.Now()

	return url.QueryEscape(now.Format("2006-01-02T15:04:05"))
}

func CreateSignature(secret string, data string) string {
	h := hmac.New(sha256.New, []byte(secret))

	h.Write([]byte(data))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func HMAC(data, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}

func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func HTTP(options *Options, res interface{}) (err error) {
	resp, err := request(options)
	if err != nil {
		return
	}

	err = ReadBodyJSON(resp, res)
	return
}

func request(options *Options) (resp *http.Response, err error) {
	if options.Method == "" {
		options.Method = "GET"
	}

	req, err := http.NewRequest(options.Method, options.URL,
		bytes.NewBuffer(options.Data))
	if err != nil {
		return
	}

	if options.QueryParams != nil {
		values := url.Values{}
		for k, v := range options.QueryParams {
			values.Set(k, v)
		}
		req.URL.RawQuery = values.Encode()
	}

	for k, v := range options.Headers {
		req.Header.Set(k, v)
	}

	resp, err = http.DefaultClient.Do(req)
	return
}

func ReadBodyJSON(resp *http.Response, data interface{}) (err error) {
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &data)
	return
}
