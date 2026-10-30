// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stiggio/stigg-cli/internal/mocktest"
	"github.com/stiggio/stigg-cli/internal/requestflag"
)

func TestV1CustomersEntitiesEntitlementsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:customers:entities:entitlements", "list",
			"--max-items", "10",
			"--id", "id",
			"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--currency-id", "currencyId",
			"--entity-id", "entityId",
			"--feature-id", "featureId",
			"--limit", "1",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})
}

func TestV1CustomersEntitiesEntitlementsUpsert(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:customers:entities:entitlements", "upsert",
			"--id", "id",
			"--entity-entitlement", "{entityId: workspace-001, cadence: P1M, currencyId: currencyId, featureId: compute-minutes, parentId: parentId, scopeEntityIds: [NxI], usageLimit: 1000}",
			"--entity-entitlement", "{entityId: workspace-002, cadence: P1M, currencyId: cred-type-tokens, featureId: featureId, parentId: workspace-001, scopeEntityIds: [user-1], usageLimit: 2000}",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1CustomersEntitiesEntitlementsUpsert)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:customers:entities:entitlements", "upsert",
			"--id", "id",
			"--entity-entitlement.entity-id", "workspace-001",
			"--entity-entitlement.cadence", "P1M",
			"--entity-entitlement.currency-id", "currencyId",
			"--entity-entitlement.feature-id", "compute-minutes",
			"--entity-entitlement.parent-id", "parentId",
			"--entity-entitlement.scope-entity-ids", "[NxI]",
			"--entity-entitlement.usage-limit", "1000",
			"--entity-entitlement.entity-id", "workspace-002",
			"--entity-entitlement.cadence", "P1M",
			"--entity-entitlement.currency-id", "cred-type-tokens",
			"--entity-entitlement.feature-id", "featureId",
			"--entity-entitlement.parent-id", "workspace-001",
			"--entity-entitlement.scope-entity-ids", "[user-1]",
			"--entity-entitlement.usage-limit", "2000",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"entityEntitlements:\n" +
			"  - entityId: workspace-001\n" +
			"    cadence: P1M\n" +
			"    currencyId: currencyId\n" +
			"    featureId: compute-minutes\n" +
			"    parentId: parentId\n" +
			"    scopeEntityIds:\n" +
			"      - NxI\n" +
			"    usageLimit: 1000\n" +
			"  - entityId: workspace-002\n" +
			"    cadence: P1M\n" +
			"    currencyId: cred-type-tokens\n" +
			"    featureId: featureId\n" +
			"    parentId: workspace-001\n" +
			"    scopeEntityIds:\n" +
			"      - user-1\n" +
			"    usageLimit: 2000\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:customers:entities:entitlements", "upsert",
			"--id", "id",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})
}
