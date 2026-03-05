// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/stigg-cli/internal/mocktest"
	"github.com/stainless-sdks/stigg-cli/internal/requestflag"
)

func TestV1AddonsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:addons", "create",
		"--api-key", "string",
		"--id", "id",
		"--display-name", "displayName",
		"--product-id", "productId",
		"--billing-id", "billingId",
		"--description", "description",
		"--max-quantity", "0",
		"--metadata", "{foo: string}",
		"--pricing-type", "FREE",
		"--status", "DRAFT",
	)
}

func TestV1AddonsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:addons", "retrieve",
		"--api-key", "string",
		"--id", "x",
	)
}

func TestV1AddonsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:addons", "update",
		"--api-key", "string",
		"--id", "x",
		"--billing-id", "billingId",
		"--dependency", "[string]",
		"--description", "description",
		"--display-name", "displayName",
		"--max-quantity", "0",
		"--metadata", "{foo: string}",
		"--status", "DRAFT",
	)
}

func TestV1AddonsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:addons", "list",
		"--api-key", "string",
		"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--created-at", "{gt: '2019-12-27T18:11:19.117Z', gte: '2019-12-27T18:11:19.117Z', lt: '2019-12-27T18:11:19.117Z', lte: '2019-12-27T18:11:19.117Z'}",
		"--limit", "1",
		"--product-id", "productId",
		"--status", "status",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(v1AddonsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:addons", "list",
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

func TestV1AddonsArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:addons", "archive",
		"--api-key", "string",
		"--id", "x",
	)
}

func TestV1AddonsCreateDraft(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:addons", "create-draft",
		"--api-key", "string",
		"--id", "x",
	)
}

func TestV1AddonsPublish(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:addons", "publish",
		"--api-key", "string",
		"--id", "x",
		"--migration-type", "NEW_CUSTOMERS",
	)
}

func TestV1AddonsRemoveDraft(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:addons", "remove-draft",
		"--api-key", "string",
		"--id", "x",
	)
}

func TestV1AddonsSetPricing(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:addons", "set-pricing",
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
	requestflag.CheckInnerFlags(v1AddonsSetPricing)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:addons", "set-pricing",
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
