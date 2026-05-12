// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stiggio/stigg-cli/internal/mocktest"
)

func TestV1CreditsAutoRechargeGetAutoRecharge(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:credits:auto-recharge", "get-auto-recharge",
			"--currency-id", "currencyId",
			"--customer-id", "customerId",
		)
	})
}
