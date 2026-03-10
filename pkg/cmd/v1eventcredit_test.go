// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stiggio/stigg-cli/internal/mocktest"
)

func TestV1EventsCreditsGetAutoRecharge(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:events:credits", "get-auto-recharge",
			"--api-key", "string",
			"--currency-id", "currencyId",
			"--customer-id", "customerId",
		)
	})
}

func TestV1EventsCreditsGetUsage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:events:credits", "get-usage",
			"--api-key", "string",
			"--customer-id", "customerId",
			"--currency-id", "currencyId",
			"--resource-id", "resourceId",
			"--time-range", "LAST_DAY",
		)
	})
}

func TestV1EventsCreditsListLedger(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:events:credits", "list-ledger",
			"--api-key", "string",
			"--max-items", "10",
			"--customer-id", "customerId",
			"--after", "after",
			"--before", "before",
			"--currency-id", "currencyId",
			"--limit", "1",
			"--resource-id", "resourceId",
		)
	})
}
