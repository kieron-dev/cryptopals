package ctr_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCtr(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ctr Suite")
}
