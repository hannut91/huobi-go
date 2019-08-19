package huobi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/hannut91/gocryptotrader/common"

	"github.com/hannut91/huobi-go/types"
	"github.com/hannut91/huobi-go/utils"
)

const (
	version = "1"
	baseURL = "api.huobi.pro"
	api     = "https://api.huobi.pro"
)

type Client struct {
	Key    string
	Secret string

	AccountID int
}

func CreateHuobiClient(key, secret string) *Client {
	return &Client{Key: key, Secret: secret}
}

func (c *Client) SetAccountID(id int) {
	c.AccountID = id
}

func (c *Client) Symbols() (symbols []*types.Symbol, err error) {
	endpoint := "/common/symbols"
	apiURL := api + "/v" + version + endpoint

	var res struct {
		types.Response
		Symbols []*types.Symbol `json:"data"`
	}
	err = utils.HTTP(&utils.Options{
		URL: apiURL,
	}, &res)
	if err != nil {
		return
	}

	if res.ErrorMessage != "" {
		err = errors.New(res.ErrorMessage)
		return
	}

	symbols = res.Symbols
	return
}

func (c *Client) Depth(
	symbol, step string,
	count int,
) (tick *types.Tick, err error) {
	endpoint := "/market/depth"
	apiURL := api + endpoint

	var res struct {
		types.Response
		Tick *types.Tick `json:"tick"`
	}
	err = utils.HTTP(&utils.Options{
		URL: apiURL,
		QueryParams: map[string]string{
			"symbol": symbol,
			"type":   step,
			"depth":  strconv.Itoa(count),
		},
	}, &res)
	if err != nil {
		return
	}

	if res.ErrorMessage != "" {
		err = errors.New(res.ErrorMessage)
		return
	}

	tick = res.Tick
	return
}

func (c *Client) Accounts() (accounts []*types.Account, err error) {
	values := url.Values{}
	values.Set("AccessKeyId", c.Key)
	values.Set("SignatureMethod", "HmacSHA256")
	values.Set("SignatureVersion", "2")
	values.Set("Timestamp", time.Now().UTC().Format("2006-01-02T15:04:05"))

	method := "GET"
	endpoint := fmt.Sprintf("/v%s%s", version, "/account/accounts")
	payload := fmt.Sprintf("%s\n%s\n%s\n%s",
		method, baseURL, endpoint, values.Encode())

	hmac := utils.HMAC([]byte(payload), []byte(c.Secret))
	signature := utils.Base64Encode(hmac)
	values.Set("Signature", signature)

	var res struct {
		types.Response
		Accounts []*types.Account `json:"data"`
	}
	err = utils.HTTP(&utils.Options{
		URL: api + endpoint + "?" + values.Encode(),
		Headers: map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		},
	}, &res)
	if err != nil {
		return
	}

	if res.ErrorMessage != "" {
		err = errors.New(res.ErrorMessage)
		return
	}

	accounts = res.Accounts
	return
}

func (c *Client) Balance() (balance map[string]*types.Balance, err error) {
	values := url.Values{}
	values.Set("AccessKeyId", c.Key)
	values.Set("SignatureMethod", "HmacSHA256")
	values.Set("SignatureVersion", "2")
	values.Set("Timestamp", time.Now().UTC().Format("2006-01-02T15:04:05"))

	method := "GET"
	endpoint := fmt.Sprintf("/v%s%s/%s/balance", version, "/account/accounts",
		strconv.Itoa(c.AccountID))
	payload := fmt.Sprintf("%s\n%s\n%s\n%s",
		method, baseURL, endpoint, values.Encode())

	hmac := utils.HMAC([]byte(payload), []byte(c.Secret))
	signature := utils.Base64Encode(hmac)
	values.Set("Signature", signature)

	var res struct {
		types.Response
		BalanceResponses *types.BalanceResponse `json:"data"`
	}
	err = utils.HTTP(&utils.Options{
		URL: api + endpoint + "?" + values.Encode(),
		Headers: map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		},
	}, &res)
	if err != nil {
		return
	}

	if res.ErrorMessage != "" {
		err = errors.New(res.ErrorMessage)
		return
	}

	balance = make(map[string]*types.Balance)

	for _, v := range res.BalanceResponses.List {
		if balance[v.Currency] == nil {
			balance[v.Currency] = new(types.Balance)
		}

		if v.Type == "frozen" {
			balance[v.Currency].Locked = v.Balance
		} else {
			balance[v.Currency].Available = v.Balance
		}
	}

	return
}

func (c *Client) Order(
	side, symbol, price, amount string,
) (orderID string, err error) {
	values := url.Values{}
	values.Set("AccessKeyId", c.Key)
	values.Set("SignatureMethod", "HmacSHA256")
	values.Set("SignatureVersion", "2")
	values.Set("Timestamp", time.Now().UTC().Format("2006-01-02T15:04:05"))

	method := "POST"
	endpoint := fmt.Sprintf("/v%s%s", version, "/order/orders/place")
	payload := fmt.Sprintf("%s\n%s\n%s\n%s",
		method, baseURL, endpoint, values.Encode())

	hmac := utils.HMAC([]byte(payload), []byte(c.Secret))
	signature := utils.Base64Encode(hmac)
	values.Set("Signature", signature)

	data := map[string]interface{}{
		"account-id": strconv.Itoa(c.AccountID),
		"type":       side,
		"symbol":     symbol,
		"price":      price,
		"amount":     amount,
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		return
	}

	var res struct {
		types.Response
		OrderID string `json:"data"`
	}
	err = utils.HTTP(&utils.Options{
		Method: method,
		URL:    api + endpoint + "?" + values.Encode(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Data: bytes,
	}, &res)
	if err != nil {
		return
	}

	if res.ErrorMessage != "" {
		err = errors.New(res.ErrorMessage)
		return
	}

	orderID = res.OrderID
	return
}

func (c *Client) AggregateBalance() (
	result []types.AggregateBalance,
	err error,
) {
	values := url.Values{}
	values.Set("AccessKeyId", c.Key)
	values.Set("SignatureMethod", "HmacSHA256")
	values.Set("SignatureVersion", "2")
	values.Set("Timestamp", time.Now().UTC().Format("2006-01-02T15:04:05"))

	method := "GET"
	endpoint := fmt.Sprintf("/v%s%s", version, "/subuser/aggregate-balance")
	payload := fmt.Sprintf("%s\napi.huobi.pro\n%s\n%s",
		method, endpoint, values.Encode())

	headers := make(map[string]string)

	if method == http.MethodGet {
		headers["Content-Type"] = "application/x-www-form-urlencoded"
	} else {
		headers["Content-Type"] = "application/json"
	}

	hmac := common.GetHMAC(common.HashSHA256, []byte(payload), []byte(c.Secret))
	signature := common.Base64Encode(hmac)
	values.Set("Signature", signature)

	urlPath := common.EncodeURLValues(
		fmt.Sprintf("%s%s", api, endpoint), values,
	)

	res, err := http.Get(urlPath)
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	defer res.Body.Close()

	type response struct {
		types.Response
		Balances []types.AggregateBalance `json:"data"`
	}

	var r response

	err = json.Unmarshal(data, &r)
	if err != nil {
		return
	}

	if r.ErrorMessage != "" {
		err = errors.New(r.ErrorMessage)
		return
	}

	return r.Balances, nil
}
