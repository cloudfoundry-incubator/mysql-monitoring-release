package metrics_test

import (
. "github.com/onsi/ginkgo"
. "github.com/onsi/gomega"

"testing"
)

func TestMetrics(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "acceptance_tests/metrics")
}
