// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stiggio/stigg-cli/internal/mocktest"
	"github.com/stiggio/stigg-cli/internal/requestflag"
)

func TestV1PlansCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:plans", "create",
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
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1PlansCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "v1:plans", "create",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"id: id\n" +
			"displayName: displayName\n" +
			"productId: productId\n" +
			"billingId: billingId\n" +
			"defaultTrialConfig:\n" +
			"  duration: 0\n" +
			"  units: DAY\n" +
			"  budget:\n" +
			"    hasSoftLimit: true\n" +
			"    limit: 0\n" +
			"  trialEndBehavior: CONVERT_TO_PAID\n" +
			"description: description\n" +
			"metadata:\n" +
			"  foo: string\n" +
			"parentPlanId: parentPlanId\n" +
			"pricingType: FREE\n" +
			"status: DRAFT\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1:plans", "create",
			"--api-key", "string",
		)
	})
}

func TestV1PlansRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:plans", "retrieve",
			"--api-key", "string",
			"--id", "x",
		)
	})
}

func TestV1PlansUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:plans", "update",
			"--api-key", "string",
			"--id", "x",
			"--billing-id", "billingId",
			"--charges", "{pricingType: FREE, billingId: billingId, minimumSpend: [{billingPeriod: MONTHLY, minimum: {amount: 0, currency: usd}}], overageBillingPeriod: ON_SUBSCRIPTION_RENEWAL, overagePricingModels: [{billingModel: FLAT_FEE, pricePeriods: [{billingPeriod: MONTHLY, billingCountryCode: billingCountryCode, blockSize: 0, creditGrantCadence: BEGINNING_OF_BILLING_PERIOD, creditRate: {amount: 1, currencyId: currencyId, costFormula: costFormula}, price: {amount: 0, currency: usd}, tiers: [{flatPrice: {amount: 0, currency: usd}, unitPrice: {amount: 0, currency: usd}, upTo: 0}]}], billingCadence: RECURRING, entitlement: {featureId: featureId, hasSoftLimit: true, hasUnlimitedUsage: true, monthlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, resetPeriod: YEAR, usageLimit: 0, weeklyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, yearlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}}, featureId: featureId, topUpCustomCurrencyId: topUpCustomCurrencyId}], pricingModels: [{billingModel: FLAT_FEE, pricePeriods: [{billingPeriod: MONTHLY, billingCountryCode: billingCountryCode, blockSize: 0, creditGrantCadence: BEGINNING_OF_BILLING_PERIOD, creditRate: {amount: 1, currencyId: currencyId, costFormula: costFormula}, price: {amount: 0, currency: usd}, tiers: [{flatPrice: {amount: 0, currency: usd}, unitPrice: {amount: 0, currency: usd}, upTo: 0}]}], billingCadence: RECURRING, featureId: featureId, maxUnitQuantity: 1, minUnitQuantity: 1, monthlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, resetPeriod: YEAR, tiersMode: VOLUME, topUpCustomCurrencyId: topUpCustomCurrencyId, weeklyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, yearlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}}]}",
			"--compatible-addon-id", "[string]",
			"--default-trial-config", "{duration: 0, units: DAY, budget: {hasSoftLimit: true, limit: 0}, trialEndBehavior: CONVERT_TO_PAID}",
			"--description", "description",
			"--display-name", "displayName",
			"--metadata", "{foo: string}",
			"--parent-plan-id", "parentPlanId",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1PlansUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "v1:plans", "update",
			"--api-key", "string",
			"--id", "x",
			"--billing-id", "billingId",
			"--charges.pricing-type", "FREE",
			"--charges.billing-id", "billingId",
			"--charges.minimum-spend", "[{billingPeriod: MONTHLY, minimum: {amount: 0, currency: usd}}]",
			"--charges.overage-billing-period", "ON_SUBSCRIPTION_RENEWAL",
			"--charges.overage-pricing-models", "[{billingModel: FLAT_FEE, pricePeriods: [{billingPeriod: MONTHLY, billingCountryCode: billingCountryCode, blockSize: 0, creditGrantCadence: BEGINNING_OF_BILLING_PERIOD, creditRate: {amount: 1, currencyId: currencyId, costFormula: costFormula}, price: {amount: 0, currency: usd}, tiers: [{flatPrice: {amount: 0, currency: usd}, unitPrice: {amount: 0, currency: usd}, upTo: 0}]}], billingCadence: RECURRING, entitlement: {featureId: featureId, hasSoftLimit: true, hasUnlimitedUsage: true, monthlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, resetPeriod: YEAR, usageLimit: 0, weeklyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, yearlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}}, featureId: featureId, topUpCustomCurrencyId: topUpCustomCurrencyId}]",
			"--charges.pricing-models", "[{billingModel: FLAT_FEE, pricePeriods: [{billingPeriod: MONTHLY, billingCountryCode: billingCountryCode, blockSize: 0, creditGrantCadence: BEGINNING_OF_BILLING_PERIOD, creditRate: {amount: 1, currencyId: currencyId, costFormula: costFormula}, price: {amount: 0, currency: usd}, tiers: [{flatPrice: {amount: 0, currency: usd}, unitPrice: {amount: 0, currency: usd}, upTo: 0}]}], billingCadence: RECURRING, featureId: featureId, maxUnitQuantity: 1, minUnitQuantity: 1, monthlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, resetPeriod: YEAR, tiersMode: VOLUME, topUpCustomCurrencyId: topUpCustomCurrencyId, weeklyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, yearlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}}]",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"billingId: billingId\n" +
			"charges:\n" +
			"  pricingType: FREE\n" +
			"  billingId: billingId\n" +
			"  minimumSpend:\n" +
			"    - billingPeriod: MONTHLY\n" +
			"      minimum:\n" +
			"        amount: 0\n" +
			"        currency: usd\n" +
			"  overageBillingPeriod: ON_SUBSCRIPTION_RENEWAL\n" +
			"  overagePricingModels:\n" +
			"    - billingModel: FLAT_FEE\n" +
			"      pricePeriods:\n" +
			"        - billingPeriod: MONTHLY\n" +
			"          billingCountryCode: billingCountryCode\n" +
			"          blockSize: 0\n" +
			"          creditGrantCadence: BEGINNING_OF_BILLING_PERIOD\n" +
			"          creditRate:\n" +
			"            amount: 1\n" +
			"            currencyId: currencyId\n" +
			"            costFormula: costFormula\n" +
			"          price:\n" +
			"            amount: 0\n" +
			"            currency: usd\n" +
			"          tiers:\n" +
			"            - flatPrice:\n" +
			"                amount: 0\n" +
			"                currency: usd\n" +
			"              unitPrice:\n" +
			"                amount: 0\n" +
			"                currency: usd\n" +
			"              upTo: 0\n" +
			"      billingCadence: RECURRING\n" +
			"      entitlement:\n" +
			"        featureId: featureId\n" +
			"        hasSoftLimit: true\n" +
			"        hasUnlimitedUsage: true\n" +
			"        monthlyResetPeriodConfiguration:\n" +
			"          accordingTo: SubscriptionStart\n" +
			"        resetPeriod: YEAR\n" +
			"        usageLimit: 0\n" +
			"        weeklyResetPeriodConfiguration:\n" +
			"          accordingTo: SubscriptionStart\n" +
			"        yearlyResetPeriodConfiguration:\n" +
			"          accordingTo: SubscriptionStart\n" +
			"      featureId: featureId\n" +
			"      topUpCustomCurrencyId: topUpCustomCurrencyId\n" +
			"  pricingModels:\n" +
			"    - billingModel: FLAT_FEE\n" +
			"      pricePeriods:\n" +
			"        - billingPeriod: MONTHLY\n" +
			"          billingCountryCode: billingCountryCode\n" +
			"          blockSize: 0\n" +
			"          creditGrantCadence: BEGINNING_OF_BILLING_PERIOD\n" +
			"          creditRate:\n" +
			"            amount: 1\n" +
			"            currencyId: currencyId\n" +
			"            costFormula: costFormula\n" +
			"          price:\n" +
			"            amount: 0\n" +
			"            currency: usd\n" +
			"          tiers:\n" +
			"            - flatPrice:\n" +
			"                amount: 0\n" +
			"                currency: usd\n" +
			"              unitPrice:\n" +
			"                amount: 0\n" +
			"                currency: usd\n" +
			"              upTo: 0\n" +
			"      billingCadence: RECURRING\n" +
			"      featureId: featureId\n" +
			"      maxUnitQuantity: 1\n" +
			"      minUnitQuantity: 1\n" +
			"      monthlyResetPeriodConfiguration:\n" +
			"        accordingTo: SubscriptionStart\n" +
			"      resetPeriod: YEAR\n" +
			"      tiersMode: VOLUME\n" +
			"      topUpCustomCurrencyId: topUpCustomCurrencyId\n" +
			"      weeklyResetPeriodConfiguration:\n" +
			"        accordingTo: SubscriptionStart\n" +
			"      yearlyResetPeriodConfiguration:\n" +
			"        accordingTo: SubscriptionStart\n" +
			"compatibleAddonIds:\n" +
			"  - string\n" +
			"defaultTrialConfig:\n" +
			"  duration: 0\n" +
			"  units: DAY\n" +
			"  budget:\n" +
			"    hasSoftLimit: true\n" +
			"    limit: 0\n" +
			"  trialEndBehavior: CONVERT_TO_PAID\n" +
			"description: description\n" +
			"displayName: displayName\n" +
			"metadata:\n" +
			"  foo: string\n" +
			"parentPlanId: parentPlanId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1:plans", "update",
			"--api-key", "string",
			"--id", "x",
		)
	})
}

func TestV1PlansList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:plans", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--created-at", "{gt: '2019-12-27T18:11:19.117Z', gte: '2019-12-27T18:11:19.117Z', lt: '2019-12-27T18:11:19.117Z', lte: '2019-12-27T18:11:19.117Z'}",
			"--limit", "1",
			"--product-id", "productId",
			"--status", "status",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1PlansList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "v1:plans", "list",
			"--api-key", "string",
			"--max-items", "10",
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
	})
}

func TestV1PlansArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:plans", "archive",
			"--api-key", "string",
			"--id", "x",
		)
	})
}

func TestV1PlansCreateDraft(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:plans", "create-draft",
			"--api-key", "string",
			"--id", "x",
		)
	})
}

func TestV1PlansPublish(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:plans", "publish",
			"--api-key", "string",
			"--id", "x",
			"--migration-type", "NEW_CUSTOMERS",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("migrationType: NEW_CUSTOMERS")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1:plans", "publish",
			"--api-key", "string",
			"--id", "x",
		)
	})
}

func TestV1PlansRemoveDraft(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:plans", "remove-draft",
			"--api-key", "string",
			"--id", "x",
		)
	})
}
