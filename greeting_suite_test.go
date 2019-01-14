package greeting_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConstructors(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Greeting Suite")
}
