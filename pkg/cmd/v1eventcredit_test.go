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
			t,
			"--api-key", "string",
			"v1:events:credits", "get-auto-recharge",
			"--currency-id", "currencyId",
			"--customer-id", "customerId",
		)
	})
}

func TestV1EventsCreditsGetUsage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:events:credits", "get-usage",
			"--customer-id", "customerId",
			"--currency-id", "currencyId",
			"--end-date", "'2019-12-27T18:11:19.117Z'",
			"--resource-id", "resourceId",
			"--start-date", "'2019-12-27T18:11:19.117Z'",
			"--time-range", "LAST_DAY",
		)
	})
}

func TestV1EventsCreditsListLedger(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:events:credits", "list-ledger",
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
