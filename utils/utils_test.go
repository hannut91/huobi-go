package utils

import (
	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utils", func() {
	Describe("CreateParams", func() {
		It("returns params", func() {
			params := map[string]string{
				"data": "hello",
			}

			result := CreateParams(params)

			Expect(result).To(Equal("?data=hello"))
		})
	})

	Describe("CreateQueryString", func() {
		method := "GET"
		url := "api.huobi.pro"
		path := "/v1/order/orders"
		params := map[string]string{
			"data": "hello",
		}

		It("returns query string of map", func() {
			result := CreateQueryString(method, url, path, params)

			Expect(result).
				To(Equal("GET\napi.huobi.pro\n/v1/order/orders\ndata=hello"))
		})
	})

	Describe("CreateTimestamp", func() {
		It("returns time with format YYYY-MM-DDThh:mm:ss", func() {
			now := CreateTimestamp()

			log.Println(now)
			Expect(len(now)).To(Equal(19))
		})
	})

	Describe("CreateSignature", func() {
		result := "1b2c16b75bd2a870c114153ccda5bcfca63314bc722fa160d690de133ccbb9db"

		It("returns signature", func() {
			signature := CreateSignature("secret", "data")

			Expect(signature).To(Equal(result))
		})
	})
})
