// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stiggio/stigg-cli/internal/mocktest"
	"github.com/stiggio/stigg-cli/internal/requestflag"
)

func TestV1BetaEntityTypesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1-beta:entity-types", "list",
			"--max-items", "10",
			"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--limit", "1",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})
}

func TestV1BetaEntityTypesUpsert(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1-beta:entity-types", "upsert",
			"--type", "{id: org, attributionKeys: [organizationId], displayName: Organization}",
			"--type", "{id: team, attributionKeys: [teamId], displayName: Team}",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1BetaEntityTypesUpsert)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1-beta:entity-types", "upsert",
			"--type.id", "org",
			"--type.attribution-keys", "[organizationId]",
			"--type.display-name", "Organization",
			"--type.id", "team",
			"--type.attribution-keys", "[teamId]",
			"--type.display-name", "Team",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"types:\n" +
			"  - id: org\n" +
			"    attributionKeys:\n" +
			"      - organizationId\n" +
			"    displayName: Organization\n" +
			"  - id: team\n" +
			"    attributionKeys:\n" +
			"      - teamId\n" +
			"    displayName: Team\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1-beta:entity-types", "upsert",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})
}
