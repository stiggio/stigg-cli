// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stiggio/stigg-cli/internal/mocktest"
)

func TestV1EventsBetaCustomersEntitlementsCheck(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:events:beta:customers:entitlements", "check",
			"--id", "x",
			"--currency-id", "x",
			"--dimensions", "{foo: string}",
			"--feature-id", "x",
			"--requested-usage", "0",
			"--requested-value", "string",
			"--resource-id", "x",
		)
	})
}
