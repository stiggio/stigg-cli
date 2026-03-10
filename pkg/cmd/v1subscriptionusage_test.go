// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stiggio/stigg-cli/internal/mocktest"
)

func TestV1SubscriptionsUsageChargeUsage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:subscriptions:usage", "charge-usage",
			"--api-key", "string",
			"--id", "x",
			"--until-date", "'2019-12-27T18:11:19.117Z'",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("untilDate: '2019-12-27T18:11:19.117Z'")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1:subscriptions:usage", "charge-usage",
			"--api-key", "string",
			"--id", "x",
		)
	})
}

func TestV1SubscriptionsUsageSync(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:subscriptions:usage", "sync",
			"--api-key", "string",
			"--id", "x",
		)
	})
}
