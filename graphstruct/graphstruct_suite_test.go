package graphs_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGraphstruct(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Graphstruct Suite")
}
