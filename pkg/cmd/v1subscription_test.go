// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stiggio/stigg-cli/internal/mocktest"
	"github.com/stiggio/stigg-cli/internal/requestflag"
)

func TestV1SubscriptionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:subscriptions", "retrieve",
			"--api-key", "string",
			"--id", "x",
		)
	})
}

func TestV1SubscriptionsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:subscriptions", "update",
			"--api-key", "string",
			"--id", "x",
			"--addon", "{id: id, quantity: 0}",
			"--applied-coupon", "{billingCouponId: billingCouponId, configuration: {startDate: '2019-12-27T18:11:19.117Z'}, couponId: couponId, discount: {amountsOff: [{amount: 0, currency: usd}], description: description, durationInMonths: 1, name: name, percentOff: 1}, promotionCode: promotionCode}",
			"--await-payment-confirmation=true",
			"--billing-cycle-anchor", "UNCHANGED",
			"--billing-information", "{billingAddress: {city: city, country: country, line1: line1, line2: line2, postalCode: postalCode, state: state}, chargeOnBehalfOfAccount: chargeOnBehalfOfAccount, couponId: couponId, integrationId: integrationId, invoiceDaysUntilDue: 0, isBackdated: true, isInvoicePaid: true, metadata: {foo: bar}, prorationBehavior: INVOICE_IMMEDIATELY, taxIds: [{type: type, value: value}], taxPercentage: 0, taxRateIds: [string]}",
			"--billing-period", "MONTHLY",
			"--budget", "{hasSoftLimit: true, limit: 0}",
			"--charge", "{id: id, quantity: 1, type: FEATURE}",
			"--entitlement", "{id: id, type: FEATURE, hasSoftLimit: true, hasUnlimitedUsage: true, monthlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, resetPeriod: YEAR, usageLimit: 0, weeklyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, yearlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}}",
			"--metadata", "{foo: string}",
			"--minimum-spend", "{amount: 0, currency: usd}",
			"--price-override", "{addonId: addonId, amount: 0, baseCharge: true, currency: usd, currencyId: currencyId, featureId: featureId}",
			"--promotion-code", "promotionCode",
			"--schedule-strategy", "END_OF_BILLING_PERIOD",
			"--trial-end-date", "'2019-12-27T18:11:19.117Z'",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1SubscriptionsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "v1:subscriptions", "update",
			"--api-key", "string",
			"--id", "x",
			"--addon.id", "id",
			"--addon.quantity", "0",
			"--applied-coupon.billing-coupon-id", "billingCouponId",
			"--applied-coupon.configuration", "{startDate: '2019-12-27T18:11:19.117Z'}",
			"--applied-coupon.coupon-id", "couponId",
			"--applied-coupon.discount", "{amountsOff: [{amount: 0, currency: usd}], description: description, durationInMonths: 1, name: name, percentOff: 1}",
			"--applied-coupon.promotion-code", "promotionCode",
			"--await-payment-confirmation=true",
			"--billing-cycle-anchor", "UNCHANGED",
			"--billing-information.billing-address", "{city: city, country: country, line1: line1, line2: line2, postalCode: postalCode, state: state}",
			"--billing-information.charge-on-behalf-of-account", "chargeOnBehalfOfAccount",
			"--billing-information.coupon-id", "couponId",
			"--billing-information.integration-id", "integrationId",
			"--billing-information.invoice-days-until-due", "0",
			"--billing-information.is-backdated=true",
			"--billing-information.is-invoice-paid=true",
			"--billing-information.metadata", "{foo: bar}",
			"--billing-information.proration-behavior", "INVOICE_IMMEDIATELY",
			"--billing-information.tax-ids", "[{type: type, value: value}]",
			"--billing-information.tax-percentage", "0",
			"--billing-information.tax-rate-ids", "[string]",
			"--billing-period", "MONTHLY",
			"--budget.has-soft-limit=true",
			"--budget.limit", "0",
			"--charge.id", "id",
			"--charge.quantity", "1",
			"--charge.type", "FEATURE",
			"--entitlement", "{id: id, type: FEATURE, hasSoftLimit: true, hasUnlimitedUsage: true, monthlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, resetPeriod: YEAR, usageLimit: 0, weeklyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, yearlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}}",
			"--metadata", "{foo: string}",
			"--minimum-spend.amount", "0",
			"--minimum-spend.currency", "usd",
			"--price-override.addon-id", "addonId",
			"--price-override.amount", "0",
			"--price-override.base-charge=true",
			"--price-override.currency", "usd",
			"--price-override.currency-id", "currencyId",
			"--price-override.feature-id", "featureId",
			"--promotion-code", "promotionCode",
			"--schedule-strategy", "END_OF_BILLING_PERIOD",
			"--trial-end-date", "'2019-12-27T18:11:19.117Z'",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"addons:\n" +
			"  - id: id\n" +
			"    quantity: 0\n" +
			"appliedCoupon:\n" +
			"  billingCouponId: billingCouponId\n" +
			"  configuration:\n" +
			"    startDate: '2019-12-27T18:11:19.117Z'\n" +
			"  couponId: couponId\n" +
			"  discount:\n" +
			"    amountsOff:\n" +
			"      - amount: 0\n" +
			"        currency: usd\n" +
			"    description: description\n" +
			"    durationInMonths: 1\n" +
			"    name: name\n" +
			"    percentOff: 1\n" +
			"  promotionCode: promotionCode\n" +
			"awaitPaymentConfirmation: true\n" +
			"billingCycleAnchor: UNCHANGED\n" +
			"billingInformation:\n" +
			"  billingAddress:\n" +
			"    city: city\n" +
			"    country: country\n" +
			"    line1: line1\n" +
			"    line2: line2\n" +
			"    postalCode: postalCode\n" +
			"    state: state\n" +
			"  chargeOnBehalfOfAccount: chargeOnBehalfOfAccount\n" +
			"  couponId: couponId\n" +
			"  integrationId: integrationId\n" +
			"  invoiceDaysUntilDue: 0\n" +
			"  isBackdated: true\n" +
			"  isInvoicePaid: true\n" +
			"  metadata:\n" +
			"    foo: bar\n" +
			"  prorationBehavior: INVOICE_IMMEDIATELY\n" +
			"  taxIds:\n" +
			"    - type: type\n" +
			"      value: value\n" +
			"  taxPercentage: 0\n" +
			"  taxRateIds:\n" +
			"    - string\n" +
			"billingPeriod: MONTHLY\n" +
			"budget:\n" +
			"  hasSoftLimit: true\n" +
			"  limit: 0\n" +
			"charges:\n" +
			"  - id: id\n" +
			"    quantity: 1\n" +
			"    type: FEATURE\n" +
			"entitlements:\n" +
			"  - id: id\n" +
			"    type: FEATURE\n" +
			"    hasSoftLimit: true\n" +
			"    hasUnlimitedUsage: true\n" +
			"    monthlyResetPeriodConfiguration:\n" +
			"      accordingTo: SubscriptionStart\n" +
			"    resetPeriod: YEAR\n" +
			"    usageLimit: 0\n" +
			"    weeklyResetPeriodConfiguration:\n" +
			"      accordingTo: SubscriptionStart\n" +
			"    yearlyResetPeriodConfiguration:\n" +
			"      accordingTo: SubscriptionStart\n" +
			"metadata:\n" +
			"  foo: string\n" +
			"minimumSpend:\n" +
			"  amount: 0\n" +
			"  currency: usd\n" +
			"priceOverrides:\n" +
			"  - addonId: addonId\n" +
			"    amount: 0\n" +
			"    baseCharge: true\n" +
			"    currency: usd\n" +
			"    currencyId: currencyId\n" +
			"    featureId: featureId\n" +
			"promotionCode: promotionCode\n" +
			"scheduleStrategy: END_OF_BILLING_PERIOD\n" +
			"trialEndDate: '2019-12-27T18:11:19.117Z'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1:subscriptions", "update",
			"--api-key", "string",
			"--id", "x",
		)
	})
}

func TestV1SubscriptionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:subscriptions", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--created-at", "{gt: '2019-12-27T18:11:19.117Z', gte: '2019-12-27T18:11:19.117Z', lt: '2019-12-27T18:11:19.117Z', lte: '2019-12-27T18:11:19.117Z'}",
			"--customer-id", "customerId",
			"--limit", "1",
			"--plan-id", "planId",
			"--pricing-type", "pricingType",
			"--resource-id", "resourceId",
			"--status", "status",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1SubscriptionsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "v1:subscriptions", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--created-at.gt", "2019-12-27T18:11:19.117Z",
			"--created-at.gte", "2019-12-27T18:11:19.117Z",
			"--created-at.lt", "2019-12-27T18:11:19.117Z",
			"--created-at.lte", "2019-12-27T18:11:19.117Z",
			"--customer-id", "customerId",
			"--limit", "1",
			"--plan-id", "planId",
			"--pricing-type", "pricingType",
			"--resource-id", "resourceId",
			"--status", "status",
		)
	})
}

func TestV1SubscriptionsCancel(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:subscriptions", "cancel",
			"--api-key", "string",
			"--id", "x",
			"--cancellation-action", "DEFAULT",
			"--cancellation-time", "END_OF_BILLING_PERIOD",
			"--end-date", "'2019-12-27T18:11:19.117Z'",
			"--prorate=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"cancellationAction: DEFAULT\n" +
			"cancellationTime: END_OF_BILLING_PERIOD\n" +
			"endDate: '2019-12-27T18:11:19.117Z'\n" +
			"prorate: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1:subscriptions", "cancel",
			"--api-key", "string",
			"--id", "x",
		)
	})
}

func TestV1SubscriptionsDelegate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:subscriptions", "delegate",
			"--api-key", "string",
			"--id", "x",
			"--target-customer-id", "targetCustomerId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("targetCustomerId: targetCustomerId")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1:subscriptions", "delegate",
			"--api-key", "string",
			"--id", "x",
		)
	})
}

func TestV1SubscriptionsImport(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:subscriptions", "import",
			"--api-key", "string",
			"--subscription", "{id: id, customerId: customerId, planId: planId, addons: [{id: id, quantity: 0}], billingId: billingId, billingPeriod: MONTHLY, charges: [{id: id, quantity: 1, type: FEATURE}], endDate: '2019-12-27T18:11:19.117Z', metadata: {foo: string}, resourceId: resourceId, startDate: '2019-12-27T18:11:19.117Z'}",
			"--integration-id", "integrationId",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1SubscriptionsImport)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "v1:subscriptions", "import",
			"--api-key", "string",
			"--subscription.id", "id",
			"--subscription.customer-id", "customerId",
			"--subscription.plan-id", "planId",
			"--subscription.addons", "[{id: id, quantity: 0}]",
			"--subscription.billing-id", "billingId",
			"--subscription.billing-period", "MONTHLY",
			"--subscription.charges", "[{id: id, quantity: 1, type: FEATURE}]",
			"--subscription.end-date", "2019-12-27T18:11:19.117Z",
			"--subscription.metadata", "{foo: string}",
			"--subscription.resource-id", "resourceId",
			"--subscription.start-date", "2019-12-27T18:11:19.117Z",
			"--integration-id", "integrationId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"subscriptions:\n" +
			"  - id: id\n" +
			"    customerId: customerId\n" +
			"    planId: planId\n" +
			"    addons:\n" +
			"      - id: id\n" +
			"        quantity: 0\n" +
			"    billingId: billingId\n" +
			"    billingPeriod: MONTHLY\n" +
			"    charges:\n" +
			"      - id: id\n" +
			"        quantity: 1\n" +
			"        type: FEATURE\n" +
			"    endDate: '2019-12-27T18:11:19.117Z'\n" +
			"    metadata:\n" +
			"      foo: string\n" +
			"    resourceId: resourceId\n" +
			"    startDate: '2019-12-27T18:11:19.117Z'\n" +
			"integrationId: integrationId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1:subscriptions", "import",
			"--api-key", "string",
		)
	})
}

func TestV1SubscriptionsMigrate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:subscriptions", "migrate",
			"--api-key", "string",
			"--id", "x",
			"--subscription-migration-time", "END_OF_BILLING_PERIOD",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("subscriptionMigrationTime: END_OF_BILLING_PERIOD")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1:subscriptions", "migrate",
			"--api-key", "string",
			"--id", "x",
		)
	})
}

func TestV1SubscriptionsPreview(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:subscriptions", "preview",
			"--api-key", "string",
			"--customer-id", "customerId",
			"--plan-id", "planId",
			"--addon", "{id: id, quantity: 0}",
			"--applied-coupon", "{billingCouponId: billingCouponId, configuration: {startDate: '2019-12-27T18:11:19.117Z'}, couponId: couponId, discount: {amountsOff: [{amount: 0, currency: usd}], description: description, durationInMonths: 1, name: name, percentOff: 1}, promotionCode: promotionCode}",
			"--billable-feature", "{featureId: featureId, quantity: 1}",
			"--billing-country-code", "billingCountryCode",
			"--billing-cycle-anchor", "UNCHANGED",
			"--billing-information", "{billingAddress: {city: city, country: country, line1: line1, line2: line2, postalCode: postalCode, state: state}, chargeOnBehalfOfAccount: chargeOnBehalfOfAccount, integrationId: integrationId, invoiceDaysUntilDue: 0, isBackdated: true, isInvoicePaid: true, metadata: {}, prorationBehavior: INVOICE_IMMEDIATELY, taxIds: [{type: type, value: value}], taxPercentage: 0, taxRateIds: [string]}",
			"--billing-period", "MONTHLY",
			"--charge", "{id: id, quantity: 1, type: FEATURE}",
			"--paying-customer-id", "payingCustomerId",
			"--resource-id", "resourceId",
			"--schedule-strategy", "END_OF_BILLING_PERIOD",
			"--start-date", "'2019-12-27T18:11:19.117Z'",
			"--trial-override-configuration", "{isTrial: true, trialEndBehavior: CONVERT_TO_PAID, trialEndDate: '2019-12-27T18:11:19.117Z'}",
			"--unit-quantity", "1",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1SubscriptionsPreview)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "v1:subscriptions", "preview",
			"--api-key", "string",
			"--customer-id", "customerId",
			"--plan-id", "planId",
			"--addon.id", "id",
			"--addon.quantity", "0",
			"--applied-coupon.billing-coupon-id", "billingCouponId",
			"--applied-coupon.configuration", "{startDate: '2019-12-27T18:11:19.117Z'}",
			"--applied-coupon.coupon-id", "couponId",
			"--applied-coupon.discount", "{amountsOff: [{amount: 0, currency: usd}], description: description, durationInMonths: 1, name: name, percentOff: 1}",
			"--applied-coupon.promotion-code", "promotionCode",
			"--billable-feature.feature-id", "featureId",
			"--billable-feature.quantity", "1",
			"--billing-country-code", "billingCountryCode",
			"--billing-cycle-anchor", "UNCHANGED",
			"--billing-information.billing-address", "{city: city, country: country, line1: line1, line2: line2, postalCode: postalCode, state: state}",
			"--billing-information.charge-on-behalf-of-account", "chargeOnBehalfOfAccount",
			"--billing-information.integration-id", "integrationId",
			"--billing-information.invoice-days-until-due", "0",
			"--billing-information.is-backdated=true",
			"--billing-information.is-invoice-paid=true",
			"--billing-information.metadata", "{}",
			"--billing-information.proration-behavior", "INVOICE_IMMEDIATELY",
			"--billing-information.tax-ids", "[{type: type, value: value}]",
			"--billing-information.tax-percentage", "0",
			"--billing-information.tax-rate-ids", "[string]",
			"--billing-period", "MONTHLY",
			"--charge.id", "id",
			"--charge.quantity", "1",
			"--charge.type", "FEATURE",
			"--paying-customer-id", "payingCustomerId",
			"--resource-id", "resourceId",
			"--schedule-strategy", "END_OF_BILLING_PERIOD",
			"--start-date", "'2019-12-27T18:11:19.117Z'",
			"--trial-override-configuration.is-trial=true",
			"--trial-override-configuration.trial-end-behavior", "CONVERT_TO_PAID",
			"--trial-override-configuration.trial-end-date", "2019-12-27T18:11:19.117Z",
			"--unit-quantity", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"customerId: customerId\n" +
			"planId: planId\n" +
			"addons:\n" +
			"  - id: id\n" +
			"    quantity: 0\n" +
			"appliedCoupon:\n" +
			"  billingCouponId: billingCouponId\n" +
			"  configuration:\n" +
			"    startDate: '2019-12-27T18:11:19.117Z'\n" +
			"  couponId: couponId\n" +
			"  discount:\n" +
			"    amountsOff:\n" +
			"      - amount: 0\n" +
			"        currency: usd\n" +
			"    description: description\n" +
			"    durationInMonths: 1\n" +
			"    name: name\n" +
			"    percentOff: 1\n" +
			"  promotionCode: promotionCode\n" +
			"billableFeatures:\n" +
			"  - featureId: featureId\n" +
			"    quantity: 1\n" +
			"billingCountryCode: billingCountryCode\n" +
			"billingCycleAnchor: UNCHANGED\n" +
			"billingInformation:\n" +
			"  billingAddress:\n" +
			"    city: city\n" +
			"    country: country\n" +
			"    line1: line1\n" +
			"    line2: line2\n" +
			"    postalCode: postalCode\n" +
			"    state: state\n" +
			"  chargeOnBehalfOfAccount: chargeOnBehalfOfAccount\n" +
			"  integrationId: integrationId\n" +
			"  invoiceDaysUntilDue: 0\n" +
			"  isBackdated: true\n" +
			"  isInvoicePaid: true\n" +
			"  metadata: {}\n" +
			"  prorationBehavior: INVOICE_IMMEDIATELY\n" +
			"  taxIds:\n" +
			"    - type: type\n" +
			"      value: value\n" +
			"  taxPercentage: 0\n" +
			"  taxRateIds:\n" +
			"    - string\n" +
			"billingPeriod: MONTHLY\n" +
			"charges:\n" +
			"  - id: id\n" +
			"    quantity: 1\n" +
			"    type: FEATURE\n" +
			"payingCustomerId: payingCustomerId\n" +
			"resourceId: resourceId\n" +
			"scheduleStrategy: END_OF_BILLING_PERIOD\n" +
			"startDate: '2019-12-27T18:11:19.117Z'\n" +
			"trialOverrideConfiguration:\n" +
			"  isTrial: true\n" +
			"  trialEndBehavior: CONVERT_TO_PAID\n" +
			"  trialEndDate: '2019-12-27T18:11:19.117Z'\n" +
			"unitQuantity: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1:subscriptions", "preview",
			"--api-key", "string",
		)
	})
}

func TestV1SubscriptionsProvision(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:subscriptions", "provision",
			"--api-key", "string",
			"--customer-id", "customerId",
			"--plan-id", "planId",
			"--id", "id",
			"--addon", "{id: id, quantity: 0}",
			"--applied-coupon", "{billingCouponId: billingCouponId, configuration: {startDate: '2019-12-27T18:11:19.117Z'}, couponId: couponId, discount: {amountsOff: [{amount: 0, currency: usd}], description: description, durationInMonths: 1, name: name, percentOff: 1}, promotionCode: promotionCode}",
			"--await-payment-confirmation=true",
			"--billing-country-code", "billingCountryCode",
			"--billing-cycle-anchor", "UNCHANGED",
			"--billing-id", "billingId",
			"--billing-information", "{billingAddress: {city: city, country: country, line1: line1, line2: line2, postalCode: postalCode, state: state}, chargeOnBehalfOfAccount: chargeOnBehalfOfAccount, integrationId: integrationId, invoiceDaysUntilDue: 0, isBackdated: true, isInvoicePaid: true, metadata: {foo: string}, prorationBehavior: INVOICE_IMMEDIATELY, taxIds: [{type: type, value: value}], taxPercentage: 0, taxRateIds: [string]}",
			"--billing-period", "MONTHLY",
			"--budget", "{hasSoftLimit: true, limit: 0}",
			"--charge", "{id: id, quantity: 1, type: FEATURE}",
			"--checkout-options", "{cancelUrl: https://example.com, successUrl: https://example.com, allowPromoCodes: true, allowTaxIdCollection: true, collectBillingAddress: true, collectPhoneNumber: true, referenceId: referenceId}",
			"--entitlement", "{id: id, type: FEATURE, hasSoftLimit: true, hasUnlimitedUsage: true, monthlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, resetPeriod: YEAR, usageLimit: 0, weeklyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, yearlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}}",
			"--metadata", "{foo: string}",
			"--minimum-spend", "{amount: 0, currency: usd}",
			"--paying-customer-id", "payingCustomerId",
			"--payment-collection-method", "CHARGE",
			"--price-override", "{addonId: addonId, amount: 0, baseCharge: true, billingCountryCode: billingCountryCode, blockSize: 0, creditGrantCadence: BEGINNING_OF_BILLING_PERIOD, creditRate: {amount: 1, currencyId: currencyId, costFormula: costFormula}, currency: usd, featureId: featureId, tiers: [{flatPrice: {amount: 0, currency: usd}, unitPrice: {amount: 0, currency: usd}, upTo: 0}]}",
			"--resource-id", "resourceId",
			"--salesforce-id", "salesforceId",
			"--schedule-strategy", "END_OF_BILLING_PERIOD",
			"--start-date", "'2019-12-27T18:11:19.117Z'",
			"--trial-override-configuration", "{isTrial: true, trialEndBehavior: CONVERT_TO_PAID, trialEndDate: '2019-12-27T18:11:19.117Z'}",
			"--unit-quantity", "1",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1SubscriptionsProvision)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "v1:subscriptions", "provision",
			"--api-key", "string",
			"--customer-id", "customerId",
			"--plan-id", "planId",
			"--id", "id",
			"--addon.id", "id",
			"--addon.quantity", "0",
			"--applied-coupon.billing-coupon-id", "billingCouponId",
			"--applied-coupon.configuration", "{startDate: '2019-12-27T18:11:19.117Z'}",
			"--applied-coupon.coupon-id", "couponId",
			"--applied-coupon.discount", "{amountsOff: [{amount: 0, currency: usd}], description: description, durationInMonths: 1, name: name, percentOff: 1}",
			"--applied-coupon.promotion-code", "promotionCode",
			"--await-payment-confirmation=true",
			"--billing-country-code", "billingCountryCode",
			"--billing-cycle-anchor", "UNCHANGED",
			"--billing-id", "billingId",
			"--billing-information.billing-address", "{city: city, country: country, line1: line1, line2: line2, postalCode: postalCode, state: state}",
			"--billing-information.charge-on-behalf-of-account", "chargeOnBehalfOfAccount",
			"--billing-information.integration-id", "integrationId",
			"--billing-information.invoice-days-until-due", "0",
			"--billing-information.is-backdated=true",
			"--billing-information.is-invoice-paid=true",
			"--billing-information.metadata", "{foo: string}",
			"--billing-information.proration-behavior", "INVOICE_IMMEDIATELY",
			"--billing-information.tax-ids", "[{type: type, value: value}]",
			"--billing-information.tax-percentage", "0",
			"--billing-information.tax-rate-ids", "[string]",
			"--billing-period", "MONTHLY",
			"--budget.has-soft-limit=true",
			"--budget.limit", "0",
			"--charge.id", "id",
			"--charge.quantity", "1",
			"--charge.type", "FEATURE",
			"--checkout-options.cancel-url", "https://example.com",
			"--checkout-options.success-url", "https://example.com",
			"--checkout-options.allow-promo-codes=true",
			"--checkout-options.allow-tax-id-collection=true",
			"--checkout-options.collect-billing-address=true",
			"--checkout-options.collect-phone-number=true",
			"--checkout-options.reference-id", "referenceId",
			"--entitlement", "{id: id, type: FEATURE, hasSoftLimit: true, hasUnlimitedUsage: true, monthlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, resetPeriod: YEAR, usageLimit: 0, weeklyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, yearlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}}",
			"--metadata", "{foo: string}",
			"--minimum-spend.amount", "0",
			"--minimum-spend.currency", "usd",
			"--paying-customer-id", "payingCustomerId",
			"--payment-collection-method", "CHARGE",
			"--price-override.addon-id", "addonId",
			"--price-override.amount", "0",
			"--price-override.base-charge=true",
			"--price-override.billing-country-code", "billingCountryCode",
			"--price-override.block-size", "0",
			"--price-override.credit-grant-cadence", "BEGINNING_OF_BILLING_PERIOD",
			"--price-override.credit-rate", "{amount: 1, currencyId: currencyId, costFormula: costFormula}",
			"--price-override.currency", "usd",
			"--price-override.feature-id", "featureId",
			"--price-override.tiers", "[{flatPrice: {amount: 0, currency: usd}, unitPrice: {amount: 0, currency: usd}, upTo: 0}]",
			"--resource-id", "resourceId",
			"--salesforce-id", "salesforceId",
			"--schedule-strategy", "END_OF_BILLING_PERIOD",
			"--start-date", "'2019-12-27T18:11:19.117Z'",
			"--trial-override-configuration.is-trial=true",
			"--trial-override-configuration.trial-end-behavior", "CONVERT_TO_PAID",
			"--trial-override-configuration.trial-end-date", "2019-12-27T18:11:19.117Z",
			"--unit-quantity", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"customerId: customerId\n" +
			"planId: planId\n" +
			"id: id\n" +
			"addons:\n" +
			"  - id: id\n" +
			"    quantity: 0\n" +
			"appliedCoupon:\n" +
			"  billingCouponId: billingCouponId\n" +
			"  configuration:\n" +
			"    startDate: '2019-12-27T18:11:19.117Z'\n" +
			"  couponId: couponId\n" +
			"  discount:\n" +
			"    amountsOff:\n" +
			"      - amount: 0\n" +
			"        currency: usd\n" +
			"    description: description\n" +
			"    durationInMonths: 1\n" +
			"    name: name\n" +
			"    percentOff: 1\n" +
			"  promotionCode: promotionCode\n" +
			"awaitPaymentConfirmation: true\n" +
			"billingCountryCode: billingCountryCode\n" +
			"billingCycleAnchor: UNCHANGED\n" +
			"billingId: billingId\n" +
			"billingInformation:\n" +
			"  billingAddress:\n" +
			"    city: city\n" +
			"    country: country\n" +
			"    line1: line1\n" +
			"    line2: line2\n" +
			"    postalCode: postalCode\n" +
			"    state: state\n" +
			"  chargeOnBehalfOfAccount: chargeOnBehalfOfAccount\n" +
			"  integrationId: integrationId\n" +
			"  invoiceDaysUntilDue: 0\n" +
			"  isBackdated: true\n" +
			"  isInvoicePaid: true\n" +
			"  metadata:\n" +
			"    foo: string\n" +
			"  prorationBehavior: INVOICE_IMMEDIATELY\n" +
			"  taxIds:\n" +
			"    - type: type\n" +
			"      value: value\n" +
			"  taxPercentage: 0\n" +
			"  taxRateIds:\n" +
			"    - string\n" +
			"billingPeriod: MONTHLY\n" +
			"budget:\n" +
			"  hasSoftLimit: true\n" +
			"  limit: 0\n" +
			"charges:\n" +
			"  - id: id\n" +
			"    quantity: 1\n" +
			"    type: FEATURE\n" +
			"checkoutOptions:\n" +
			"  cancelUrl: https://example.com\n" +
			"  successUrl: https://example.com\n" +
			"  allowPromoCodes: true\n" +
			"  allowTaxIdCollection: true\n" +
			"  collectBillingAddress: true\n" +
			"  collectPhoneNumber: true\n" +
			"  referenceId: referenceId\n" +
			"entitlements:\n" +
			"  - id: id\n" +
			"    type: FEATURE\n" +
			"    hasSoftLimit: true\n" +
			"    hasUnlimitedUsage: true\n" +
			"    monthlyResetPeriodConfiguration:\n" +
			"      accordingTo: SubscriptionStart\n" +
			"    resetPeriod: YEAR\n" +
			"    usageLimit: 0\n" +
			"    weeklyResetPeriodConfiguration:\n" +
			"      accordingTo: SubscriptionStart\n" +
			"    yearlyResetPeriodConfiguration:\n" +
			"      accordingTo: SubscriptionStart\n" +
			"metadata:\n" +
			"  foo: string\n" +
			"minimumSpend:\n" +
			"  amount: 0\n" +
			"  currency: usd\n" +
			"payingCustomerId: payingCustomerId\n" +
			"paymentCollectionMethod: CHARGE\n" +
			"priceOverrides:\n" +
			"  - addonId: addonId\n" +
			"    amount: 0\n" +
			"    baseCharge: true\n" +
			"    billingCountryCode: billingCountryCode\n" +
			"    blockSize: 0\n" +
			"    creditGrantCadence: BEGINNING_OF_BILLING_PERIOD\n" +
			"    creditRate:\n" +
			"      amount: 1\n" +
			"      currencyId: currencyId\n" +
			"      costFormula: costFormula\n" +
			"    currency: usd\n" +
			"    featureId: featureId\n" +
			"    tiers:\n" +
			"      - flatPrice:\n" +
			"          amount: 0\n" +
			"          currency: usd\n" +
			"        unitPrice:\n" +
			"          amount: 0\n" +
			"          currency: usd\n" +
			"        upTo: 0\n" +
			"resourceId: resourceId\n" +
			"salesforceId: salesforceId\n" +
			"scheduleStrategy: END_OF_BILLING_PERIOD\n" +
			"startDate: '2019-12-27T18:11:19.117Z'\n" +
			"trialOverrideConfiguration:\n" +
			"  isTrial: true\n" +
			"  trialEndBehavior: CONVERT_TO_PAID\n" +
			"  trialEndDate: '2019-12-27T18:11:19.117Z'\n" +
			"unitQuantity: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1:subscriptions", "provision",
			"--api-key", "string",
		)
	})
}

func TestV1SubscriptionsTransfer(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:subscriptions", "transfer",
			"--api-key", "string",
			"--id", "x",
			"--destination-resource-id", "destinationResourceId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("destinationResourceId: destinationResourceId")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1:subscriptions", "transfer",
			"--api-key", "string",
			"--id", "x",
		)
	})
}
