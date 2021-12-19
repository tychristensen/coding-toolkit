package graphs_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGraphs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Graphs Suite")
}
