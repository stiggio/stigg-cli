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
			"--assignment", "{capabilityId: compute-minutes, entityId: workspace-001, cadence: MONTH, usageLimit: 1000}",
			"--assignment", "{capabilityId: compute-minutes, entityId: workspace-002, cadence: MONTH, usageLimit: 2000}",
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
			"--assignment.capability-id", "compute-minutes",
			"--assignment.entity-id", "workspace-001",
			"--assignment.cadence", "MONTH",
			"--assignment.usage-limit", "1000",
			"--assignment.capability-id", "compute-minutes",
			"--assignment.entity-id", "workspace-002",
			"--assignment.cadence", "MONTH",
			"--assignment.usage-limit", "2000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"assignments:\n" +
			"  - capabilityId: compute-minutes\n" +
			"    entityId: workspace-001\n" +
			"    cadence: MONTH\n" +
			"    usageLimit: 1000\n" +
			"  - capabilityId: compute-minutes\n" +
			"    entityId: workspace-002\n" +
			"    cadence: MONTH\n" +
			"    usageLimit: 2000\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1-beta:customers:assignments", "upsert",
			"--id", "id",
		)
	})
}
