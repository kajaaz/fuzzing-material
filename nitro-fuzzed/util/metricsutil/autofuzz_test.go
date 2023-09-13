package util

// Edit if desired. Code generated by "go-fuzz-fill-utils ./metricsutil/".

import (
	"testing"

	"github.com/offchainlabs/nitro/util/metricsutil"
)

func Fuzz_CanonicalizeMetricName(f *testing.F) {
	f.Fuzz(func(t *testing.T, metric string) {
		metricsutil.CanonicalizeMetricName(metric)
	})
}
