// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stiggio/stigg-cli/internal/mocktest"
	"github.com/stiggio/stigg-cli/internal/requestflag"
)

func TestV1BetaCustomersAssignmentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1-beta:customers:assignments", "list",
			"--max-items", "10",
			"--id", "id",
			"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--capability-id", "capabilityId",
			"--entity-id", "entityId",
			"--limit", "1",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})
}

func TestV1BetaCustomersAssignmentsUpsert(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1-beta:customers:assignments", "upsert",
			"--id", "id",
			"--assignment", "{entityId: workspace-001, cadence: MONTH, currencyId: currencyId, featureId: compute-minutes, parentId: parentId, scopeEntityIds: [NxI], usageLimit: 1000}",
			"--assignment", "{entityId: workspace-002, cadence: MONTH, currencyId: cred-type-tokens, featureId: featureId, parentId: workspace-001, scopeEntityIds: [user-1], usageLimit: 2000}",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1BetaCustomersAssignmentsUpsert)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1-beta:customers:assignments", "upsert",
			"--id", "id",
			"--assignment.entity-id", "workspace-001",
			"--assignment.cadence", "MONTH",
			"--assignment.currency-id", "currencyId",
			"--assignment.feature-id", "compute-minutes",
			"--assignment.parent-id", "parentId",
			"--assignment.scope-entity-ids", "[NxI]",
			"--assignment.usage-limit", "1000",
			"--assignment.entity-id", "workspace-002",
			"--assignment.cadence", "MONTH",
			"--assignment.currency-id", "cred-type-tokens",
			"--assignment.feature-id", "featureId",
			"--assignment.parent-id", "workspace-001",
			"--assignment.scope-entity-ids", "[user-1]",
			"--assignment.usage-limit", "2000",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"assignments:\n" +
			"  - entityId: workspace-001\n" +
			"    cadence: MONTH\n" +
			"    currencyId: currencyId\n" +
			"    featureId: compute-minutes\n" +
			"    parentId: parentId\n" +
			"    scopeEntityIds:\n" +
			"      - NxI\n" +
			"    usageLimit: 1000\n" +
			"  - entityId: workspace-002\n" +
			"    cadence: MONTH\n" +
			"    currencyId: cred-type-tokens\n" +
			"    featureId: featureId\n" +
			"    parentId: workspace-001\n" +
			"    scopeEntityIds:\n" +
			"      - user-1\n" +
			"    usageLimit: 2000\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1-beta:customers:assignments", "upsert",
			"--id", "id",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})
}
