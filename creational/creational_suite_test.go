package creational_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCreational(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Creational Suite")
}
