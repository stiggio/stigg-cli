// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/stigg-cli/internal/mocktest"
	"github.com/stainless-sdks/stigg-cli/internal/requestflag"
)

func TestV1PlansCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:plans", "create",
		"--api-key", "string",
		"--id", "id",
		"--display-name", "displayName",
		"--product-id", "productId",
		"--billing-id", "billingId",
		"--default-trial-config", "{duration: 0, units: DAY, budget: {hasSoftLimit: true, limit: 0}, trialEndBehavior: CONVERT_TO_PAID}",
		"--description", "description",
		"--metadata", "{foo: string}",
		"--parent-plan-id", "parentPlanId",
		"--pricing-type", "FREE",
		"--status", "DRAFT",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(v1PlansCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:plans", "create",
		"--api-key", "string",
		"--id", "id",
		"--display-name", "displayName",
		"--product-id", "productId",
		"--billing-id", "billingId",
		"--default-trial-config.duration", "0",
		"--default-trial-config.units", "DAY",
		"--default-trial-config.budget", "{hasSoftLimit: true, limit: 0}",
		"--default-trial-config.trial-end-behavior", "CONVERT_TO_PAID",
		"--description", "description",
		"--metadata", "{foo: string}",
		"--parent-plan-id", "parentPlanId",
		"--pricing-type", "FREE",
		"--status", "DRAFT",
	)
}

func TestV1PlansRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:plans", "retrieve",
		"--api-key", "string",
		"--id", "x",
	)
}

func TestV1PlansUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:plans", "update",
		"--api-key", "string",
		"--id", "x",
		"--billing-id", "billingId",
		"--compatible-addon-id", "[string]",
		"--default-trial-config", "{duration: 0, units: DAY, budget: {hasSoftLimit: true, limit: 0}, trialEndBehavior: CONVERT_TO_PAID}",
		"--description", "description",
		"--display-name", "displayName",
		"--metadata", "{foo: string}",
		"--parent-plan-id", "parentPlanId",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(v1PlansUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:plans", "update",
		"--api-key", "string",
		"--id", "x",
		"--billing-id", "billingId",
		"--compatible-addon-id", "[string]",
		"--default-trial-config.duration", "0",
		"--default-trial-config.units", "DAY",
		"--default-trial-config.budget", "{hasSoftLimit: true, limit: 0}",
		"--default-trial-config.trial-end-behavior", "CONVERT_TO_PAID",
		"--description", "description",
		"--display-name", "displayName",
		"--metadata", "{foo: string}",
		"--parent-plan-id", "parentPlanId",
	)
}

func TestV1PlansList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:plans", "list",
		"--api-key", "string",
		"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--created-at", "{gt: '2019-12-27T18:11:19.117Z', gte: '2019-12-27T18:11:19.117Z', lt: '2019-12-27T18:11:19.117Z', lte: '2019-12-27T18:11:19.117Z'}",
		"--limit", "1",
		"--product-id", "productId",
		"--status", "status",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(v1PlansList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:plans", "list",
		"--api-key", "string",
		"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--created-at.gt", "2019-12-27T18:11:19.117Z",
		"--created-at.gte", "2019-12-27T18:11:19.117Z",
		"--created-at.lt", "2019-12-27T18:11:19.117Z",
		"--created-at.lte", "2019-12-27T18:11:19.117Z",
		"--limit", "1",
		"--product-id", "productId",
		"--status", "status",
	)
}

func TestV1PlansArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:plans", "archive",
		"--api-key", "string",
		"--id", "x",
	)
}

func TestV1PlansCreateDraft(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:plans", "create-draft",
		"--api-key", "string",
		"--id", "x",
	)
}

func TestV1PlansPublish(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:plans", "publish",
		"--api-key", "string",
		"--id", "x",
		"--migration-type", "NEW_CUSTOMERS",
	)
}

func TestV1PlansRemoveDraft(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:plans", "remove-draft",
		"--api-key", "string",
		"--id", "x",
	)
}

func TestV1PlansSetPricing(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:plans", "set-pricing",
		"--api-key", "string",
		"--id", "x",
		"--pricing-type", "FREE",
		"--billing-id", "billingId",
		"--minimum-spend", "[{billingPeriod: MONTHLY, minimum: {amount: 0, currency: usd}}]",
		"--overage-billing-period", "ON_SUBSCRIPTION_RENEWAL",
		"--overage-pricing-model", "{billingModel: FLAT_FEE, pricePeriods: [{billingPeriod: MONTHLY, billingCountryCode: billingCountryCode, blockSize: 0, creditGrantCadence: BEGINNING_OF_BILLING_PERIOD, creditRate: {amount: 1, currencyId: currencyId, costFormula: costFormula}, price: {amount: 0, currency: usd}, tiers: [{flatPrice: {amount: 0, currency: usd}, unitPrice: {amount: 0, currency: usd}, upTo: 0}]}], billingCadence: RECURRING, entitlement: {featureId: featureId, hasSoftLimit: true, hasUnlimitedUsage: true, monthlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, resetPeriod: YEAR, usageLimit: 0, weeklyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, yearlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}}, featureId: featureId, topUpCustomCurrencyId: topUpCustomCurrencyId}",
		"--pricing-model", "{billingModel: FLAT_FEE, pricePeriods: [{billingPeriod: MONTHLY, billingCountryCode: billingCountryCode, blockSize: 0, creditGrantCadence: BEGINNING_OF_BILLING_PERIOD, creditRate: {amount: 1, currencyId: currencyId, costFormula: costFormula}, price: {amount: 0, currency: usd}, tiers: [{flatPrice: {amount: 0, currency: usd}, unitPrice: {amount: 0, currency: usd}, upTo: 0}]}], billingCadence: RECURRING, featureId: featureId, maxUnitQuantity: 1, minUnitQuantity: 1, monthlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, resetPeriod: YEAR, tiersMode: VOLUME, topUpCustomCurrencyId: topUpCustomCurrencyId, weeklyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, yearlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(v1PlansSetPricing)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:plans", "set-pricing",
		"--api-key", "string",
		"--id", "x",
		"--pricing-type", "FREE",
		"--billing-id", "billingId",
		"--minimum-spend.billing-period", "MONTHLY",
		"--minimum-spend.minimum", "{amount: 0, currency: usd}",
		"--overage-billing-period", "ON_SUBSCRIPTION_RENEWAL",
		"--overage-pricing-model.billing-model", "FLAT_FEE",
		"--overage-pricing-model.price-periods", "[{billingPeriod: MONTHLY, billingCountryCode: billingCountryCode, blockSize: 0, creditGrantCadence: BEGINNING_OF_BILLING_PERIOD, creditRate: {amount: 1, currencyId: currencyId, costFormula: costFormula}, price: {amount: 0, currency: usd}, tiers: [{flatPrice: {amount: 0, currency: usd}, unitPrice: {amount: 0, currency: usd}, upTo: 0}]}]",
		"--overage-pricing-model.billing-cadence", "RECURRING",
		"--overage-pricing-model.entitlement", "{featureId: featureId, hasSoftLimit: true, hasUnlimitedUsage: true, monthlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, resetPeriod: YEAR, usageLimit: 0, weeklyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, yearlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}}",
		"--overage-pricing-model.feature-id", "featureId",
		"--overage-pricing-model.top-up-custom-currency-id", "topUpCustomCurrencyId",
		"--pricing-model.billing-model", "FLAT_FEE",
		"--pricing-model.price-periods", "[{billingPeriod: MONTHLY, billingCountryCode: billingCountryCode, blockSize: 0, creditGrantCadence: BEGINNING_OF_BILLING_PERIOD, creditRate: {amount: 1, currencyId: currencyId, costFormula: costFormula}, price: {amount: 0, currency: usd}, tiers: [{flatPrice: {amount: 0, currency: usd}, unitPrice: {amount: 0, currency: usd}, upTo: 0}]}]",
		"--pricing-model.billing-cadence", "RECURRING",
		"--pricing-model.feature-id", "featureId",
		"--pricing-model.max-unit-quantity", "1",
		"--pricing-model.min-unit-quantity", "1",
		"--pricing-model.monthly-reset-period-configuration", "{accordingTo: SubscriptionStart}",
		"--pricing-model.reset-period", "YEAR",
		"--pricing-model.tiers-mode", "VOLUME",
		"--pricing-model.top-up-custom-currency-id", "topUpCustomCurrencyId",
		"--pricing-model.weekly-reset-period-configuration", "{accordingTo: SubscriptionStart}",
		"--pricing-model.yearly-reset-period-configuration", "{accordingTo: SubscriptionStart}",
	)
}
