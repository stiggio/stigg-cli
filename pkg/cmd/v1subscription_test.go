// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/stigg-cli/internal/mocktest"
	"github.com/stainless-sdks/stigg-cli/internal/requestflag"
)

func TestV1SubscriptionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:subscriptions", "retrieve",
		"--api-key", "string",
		"--id", "x",
	)
}

func TestV1SubscriptionsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:subscriptions", "update",
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
		"--metadata", "{foo: string}",
		"--minimum-spend", "{amount: 0, currency: usd}",
		"--price-override", "{addonId: addonId, amount: 0, baseCharge: true, currency: usd, currencyId: currencyId, featureId: featureId}",
		"--promotion-code", "promotionCode",
		"--schedule-strategy", "END_OF_BILLING_PERIOD",
		"--subscription-entitlement", "{id: id, featureId: featureId, hasSoftLimit: true, hasUnlimitedUsage: true, monthlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, resetPeriod: YEAR, usageLimit: 0, weeklyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, yearlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}}",
		"--trial-end-date", "'2019-12-27T18:11:19.117Z'",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(v1SubscriptionsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:subscriptions", "update",
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
		"--subscription-entitlement.id", "id",
		"--subscription-entitlement.feature-id", "featureId",
		"--subscription-entitlement.has-soft-limit=true",
		"--subscription-entitlement.has-unlimited-usage=true",
		"--subscription-entitlement.monthly-reset-period-configuration", "{accordingTo: SubscriptionStart}",
		"--subscription-entitlement.reset-period", "YEAR",
		"--subscription-entitlement.usage-limit", "0",
		"--subscription-entitlement.weekly-reset-period-configuration", "{accordingTo: SubscriptionStart}",
		"--subscription-entitlement.yearly-reset-period-configuration", "{accordingTo: SubscriptionStart}",
		"--trial-end-date", "'2019-12-27T18:11:19.117Z'",
	)
}

func TestV1SubscriptionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:subscriptions", "list",
		"--api-key", "string",
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

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(v1SubscriptionsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:subscriptions", "list",
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
}

func TestV1SubscriptionsCancel(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:subscriptions", "cancel",
		"--api-key", "string",
		"--id", "x",
		"--cancellation-action", "DEFAULT",
		"--cancellation-time", "END_OF_BILLING_PERIOD",
		"--end-date", "'2019-12-27T18:11:19.117Z'",
		"--prorate=true",
	)
}

func TestV1SubscriptionsDelegate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:subscriptions", "delegate",
		"--api-key", "string",
		"--id", "x",
		"--target-customer-id", "targetCustomerId",
	)
}

func TestV1SubscriptionsImport(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:subscriptions", "import",
		"--api-key", "string",
		"--subscription", "{id: id, customerId: customerId, planId: planId, addons: [{id: id, quantity: 0}], billingId: billingId, billingPeriod: MONTHLY, charges: [{id: id, quantity: 1, type: FEATURE}], endDate: '2019-12-27T18:11:19.117Z', metadata: {foo: string}, resourceId: resourceId, startDate: '2019-12-27T18:11:19.117Z'}",
		"--integration-id", "integrationId",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(v1SubscriptionsImport)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:subscriptions", "import",
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
}

func TestV1SubscriptionsMigrate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:subscriptions", "migrate",
		"--api-key", "string",
		"--id", "x",
		"--subscription-migration-time", "END_OF_BILLING_PERIOD",
	)
}

func TestV1SubscriptionsPreview(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:subscriptions", "preview",
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

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(v1SubscriptionsPreview)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:subscriptions", "preview",
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
}

func TestV1SubscriptionsProvision(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:subscriptions", "provision",
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
		"--metadata", "{foo: string}",
		"--minimum-spend", "{amount: 0, currency: usd}",
		"--paying-customer-id", "payingCustomerId",
		"--payment-collection-method", "CHARGE",
		"--price-override", "{addonId: addonId, amount: 0, baseCharge: true, billingCountryCode: billingCountryCode, blockSize: 0, creditGrantCadence: BEGINNING_OF_BILLING_PERIOD, creditRate: {amount: 1, currencyId: currencyId, costFormula: costFormula}, currency: usd, featureId: featureId, tiers: [{flatPrice: {amount: 0, currency: usd}, unitPrice: {amount: 0, currency: usd}, upTo: 0}]}",
		"--resource-id", "resourceId",
		"--salesforce-id", "salesforceId",
		"--schedule-strategy", "END_OF_BILLING_PERIOD",
		"--start-date", "'2019-12-27T18:11:19.117Z'",
		"--subscription-entitlement", "{featureId: featureId, usageLimit: 0, isGranted: true}",
		"--trial-override-configuration", "{isTrial: true, trialEndBehavior: CONVERT_TO_PAID, trialEndDate: '2019-12-27T18:11:19.117Z'}",
		"--unit-quantity", "1",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(v1SubscriptionsProvision)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:subscriptions", "provision",
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
		"--subscription-entitlement.feature-id", "featureId",
		"--subscription-entitlement.usage-limit", "0",
		"--subscription-entitlement.is-granted=true",
		"--trial-override-configuration.is-trial=true",
		"--trial-override-configuration.trial-end-behavior", "CONVERT_TO_PAID",
		"--trial-override-configuration.trial-end-date", "2019-12-27T18:11:19.117Z",
		"--unit-quantity", "1",
	)
}

func TestV1SubscriptionsTransfer(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:subscriptions", "transfer",
		"--api-key", "string",
		"--id", "x",
		"--destination-resource-id", "destinationResourceId",
	)
}
