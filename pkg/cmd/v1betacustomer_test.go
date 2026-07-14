// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stiggio/stigg-cli/internal/mocktest"
)

func TestV1BetaCustomersRetrieveGovernance(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1-beta:customers", "retrieve-governance",
			"--id", "id",
			"--after", "after",
			"--currency-id", "string",
			"--entity-id-search", "x",
			"--entity-type-id", "string",
			"--feature-id", "string",
			"--limit", "1",
			"--min-utilization", "0",
			"--order", "asc",
			"--scope", "all",
			"--sort-by", "utilization",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})
}
