package huobi

import (
	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	key    = ""
	secret = ""
)

var _ = Describe("HuobiGo", func() {
	client := CreateHuobiClient(key, secret)

	BeforeEach(func() {
		accounts, err := client.Accounts()
		if err != nil {
			log.Panicln(err)
		}

		client.SetAccountID(accounts[0].ID)
	})

	Describe("Symbols", func() {
		It("returns symbols", func() {
			symbols, err := client.Symbols()
			if err != nil {
				log.Panicln(err)
			}

			Expect(len(symbols) > 0).To(BeTrue())
		})
	})

	Describe("Depth", func() {
		symbol := "ethbtc"
		step := "step0"
		count := 5

		It("returns depth", func() {
			tick, err := client.Depth(symbol, step, count)
			if err != nil {
				log.Panicln(err)
			}

			Expect(tick.Bids).To(HaveLen(5))
		})
	})

	Describe("Balance", func() {
		It("returns balance", func() {
			balance, err := client.Balance()
			if err != nil {
				return
			}

			Expect(len(balance) > 0).To(BeTrue())
		})
	})

	// Describe("Order", func() {
	// 	side := "buy-limit"
	// 	symbol := "ethbtc"
	// 	price := "0.0018"
	// 	amount := "0.1"

	// 	It("places order", func() {
	// 		orderID, err := client.Order(side, symbol, price, amount)
	// 		if err != nil {
	// 			log.Panicln(err)
	// 		}

	// 		Expect(len(orderID) > 0).To(BeTrue())
	// 	})
	// })

	Describe("Accounts", func() {
		It("returns accounts", func() {
			result, err := client.Accounts()
			if err != nil {
				log.Panicln(err)
			}

			Expect(len(result) > 0).To(BeTrue())
		})
	})
})
