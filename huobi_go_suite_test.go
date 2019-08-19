package huobi

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHuobiGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HuobiGo Suite")
}
