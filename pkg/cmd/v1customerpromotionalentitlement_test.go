// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/stigg-cli/internal/mocktest"
	"github.com/stainless-sdks/stigg-cli/internal/requestflag"
)

func TestV1CustomersPromotionalEntitlementsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:customers:promotional-entitlements", "create",
		"--id", "x",
		"--promotional-entitlement", "{customEndDate: '2019-12-27T18:11:19.117Z', enumValues: [string], featureId: featureId, hasSoftLimit: true, hasUnlimitedUsage: true, isVisible: true, monthlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, period: 1 week, resetPeriod: YEAR, usageLimit: -9007199254740991, weeklyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, yearlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(v1CustomersPromotionalEntitlementsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:customers:promotional-entitlements", "create",
		"--id", "x",
		"--promotional-entitlement.custom-end-date", "2019-12-27T18:11:19.117Z",
		"--promotional-entitlement.enum-values", "[string]",
		"--promotional-entitlement.feature-id", "featureId",
		"--promotional-entitlement.has-soft-limit=true",
		"--promotional-entitlement.has-unlimited-usage=true",
		"--promotional-entitlement.is-visible=true",
		"--promotional-entitlement.monthly-reset-period-configuration", "{accordingTo: SubscriptionStart}",
		"--promotional-entitlement.period", "1 week",
		"--promotional-entitlement.reset-period", "YEAR",
		"--promotional-entitlement.usage-limit", "-9007199254740991",
		"--promotional-entitlement.weekly-reset-period-configuration", "{accordingTo: SubscriptionStart}",
		"--promotional-entitlement.yearly-reset-period-configuration", "{accordingTo: SubscriptionStart}",
	)
}

func TestV1CustomersPromotionalEntitlementsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:customers:promotional-entitlements", "list",
		"--id", "x",
		"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--created-at", "{gt: '2019-12-27T18:11:19.117Z', gte: '2019-12-27T18:11:19.117Z', lt: '2019-12-27T18:11:19.117Z', lte: '2019-12-27T18:11:19.117Z'}",
		"--limit", "1",
		"--status", "status",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(v1CustomersPromotionalEntitlementsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:customers:promotional-entitlements", "list",
		"--id", "x",
		"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--created-at.gt", "2019-12-27T18:11:19.117Z",
		"--created-at.gte", "2019-12-27T18:11:19.117Z",
		"--created-at.lt", "2019-12-27T18:11:19.117Z",
		"--created-at.lte", "2019-12-27T18:11:19.117Z",
		"--limit", "1",
		"--status", "status",
	)
}

func TestV1CustomersPromotionalEntitlementsRevoke(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:customers:promotional-entitlements", "revoke",
		"--id", "id",
		"--feature-id", "featureId",
	)
}
