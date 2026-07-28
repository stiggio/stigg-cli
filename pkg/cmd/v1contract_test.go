// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stiggio/stigg-cli/internal/mocktest"
	"github.com/stiggio/stigg-cli/internal/requestflag"
)

func TestV1ContractsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:contracts", "create",
			"--customer-id", "customerId",
			"--subscription", "{existingSubscriptionId: existingSubscriptionId, newSubscription: {customerId: customerId, planId: planId, id: id, addons: [{id: id, quantity: 0}], appliedCoupon: {billingCouponId: billingCouponId, configuration: {startDate: '2019-12-27T18:11:19.117Z'}, couponId: couponId, discount: {amountsOff: [{amount: 0, currency: usd}], description: description, durationInMonths: 1, name: name, percentOff: 1}, promotionCode: promotionCode}, awaitPaymentConfirmation: true, billingCountryCode: billingCountryCode, billingCycleAnchor: UNCHANGED, billingId: billingId, billingInformation: {billingAddress: {city: city, country: country, line1: line1, line2: line2, postalCode: postalCode, state: state}, chargeOnBehalfOfAccount: chargeOnBehalfOfAccount, integrationId: integrationId, invoiceDaysUntilDue: 0, isBackdated: true, isInvoicePaid: true, metadata: {foo: string}, prorationBehavior: INVOICE_IMMEDIATELY, taxIds: [{type: type, value: value}], taxPercentage: 0, taxRateIds: [string]}, billingPeriod: MONTHLY, budget: {hasSoftLimit: true, limit: 0}, cancellationDate: '2019-12-27T18:11:19.117Z', charges: [{id: id, quantity: 0, type: FEATURE}], checkoutOptions: {cancelUrl: https://example.com, successUrl: https://example.com, allowPromoCodes: true, allowTaxIdCollection: true, collectBillingAddress: true, collectPhoneNumber: true, referenceId: referenceId}, entitlements: [{id: id, type: FEATURE, hasSoftLimit: true, hasUnlimitedUsage: true, monthlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, resetPeriod: YEAR, usageLimit: 0, weeklyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, yearlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}}], metadata: {foo: string}, minimumSpend: {amount: 0, currency: usd}, payingCustomerId: payingCustomerId, paymentCollectionMethod: CHARGE, priceOverrides: [{addonId: addonId, amount: 0, baseCharge: true, billingCountryCode: billingCountryCode, blockSize: 0, creditGrantCadence: BEGINNING_OF_BILLING_PERIOD, creditRate: {amount: 1, currencyId: currencyId, costFormula: costFormula}, currency: usd, featureId: featureId, tiers: [{flatPrice: {amount: 0, currency: usd}, unitPrice: {amount: 0, currency: usd}, upTo: 0}]}], resourceId: resourceId, salesforceId: salesforceId, scheduleStrategy: END_OF_BILLING_PERIOD, startDate: '2019-12-27T18:11:19.117Z', trialOverrideConfiguration: {isTrial: true, trialEndBehavior: CONVERT_TO_PAID, trialEndDate: '2019-12-27T18:11:19.117Z'}, unitQuantity: 0}}",
			"--activation-end-date", "'2019-12-27T18:11:19.117Z'",
			"--activation-start-date", "'2019-12-27T18:11:19.117Z'",
			"--name", "name",
			"--po-number", "poNumber",
			"--setup-billing=true",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1ContractsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:contracts", "create",
			"--customer-id", "customerId",
			"--subscription.existing-subscription-id", "existingSubscriptionId",
			"--subscription.new-subscription", "{customerId: customerId, planId: planId, id: id, addons: [{id: id, quantity: 0}], appliedCoupon: {billingCouponId: billingCouponId, configuration: {startDate: '2019-12-27T18:11:19.117Z'}, couponId: couponId, discount: {amountsOff: [{amount: 0, currency: usd}], description: description, durationInMonths: 1, name: name, percentOff: 1}, promotionCode: promotionCode}, awaitPaymentConfirmation: true, billingCountryCode: billingCountryCode, billingCycleAnchor: UNCHANGED, billingId: billingId, billingInformation: {billingAddress: {city: city, country: country, line1: line1, line2: line2, postalCode: postalCode, state: state}, chargeOnBehalfOfAccount: chargeOnBehalfOfAccount, integrationId: integrationId, invoiceDaysUntilDue: 0, isBackdated: true, isInvoicePaid: true, metadata: {foo: string}, prorationBehavior: INVOICE_IMMEDIATELY, taxIds: [{type: type, value: value}], taxPercentage: 0, taxRateIds: [string]}, billingPeriod: MONTHLY, budget: {hasSoftLimit: true, limit: 0}, cancellationDate: '2019-12-27T18:11:19.117Z', charges: [{id: id, quantity: 0, type: FEATURE}], checkoutOptions: {cancelUrl: https://example.com, successUrl: https://example.com, allowPromoCodes: true, allowTaxIdCollection: true, collectBillingAddress: true, collectPhoneNumber: true, referenceId: referenceId}, entitlements: [{id: id, type: FEATURE, hasSoftLimit: true, hasUnlimitedUsage: true, monthlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, resetPeriod: YEAR, usageLimit: 0, weeklyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, yearlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}}], metadata: {foo: string}, minimumSpend: {amount: 0, currency: usd}, payingCustomerId: payingCustomerId, paymentCollectionMethod: CHARGE, priceOverrides: [{addonId: addonId, amount: 0, baseCharge: true, billingCountryCode: billingCountryCode, blockSize: 0, creditGrantCadence: BEGINNING_OF_BILLING_PERIOD, creditRate: {amount: 1, currencyId: currencyId, costFormula: costFormula}, currency: usd, featureId: featureId, tiers: [{flatPrice: {amount: 0, currency: usd}, unitPrice: {amount: 0, currency: usd}, upTo: 0}]}], resourceId: resourceId, salesforceId: salesforceId, scheduleStrategy: END_OF_BILLING_PERIOD, startDate: '2019-12-27T18:11:19.117Z', trialOverrideConfiguration: {isTrial: true, trialEndBehavior: CONVERT_TO_PAID, trialEndDate: '2019-12-27T18:11:19.117Z'}, unitQuantity: 0}",
			"--activation-end-date", "'2019-12-27T18:11:19.117Z'",
			"--activation-start-date", "'2019-12-27T18:11:19.117Z'",
			"--name", "name",
			"--po-number", "poNumber",
			"--setup-billing=true",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"customerId: customerId\n" +
			"subscriptions:\n" +
			"  - existingSubscriptionId: existingSubscriptionId\n" +
			"    newSubscription:\n" +
			"      customerId: customerId\n" +
			"      planId: planId\n" +
			"      id: id\n" +
			"      addons:\n" +
			"        - id: id\n" +
			"          quantity: 0\n" +
			"      appliedCoupon:\n" +
			"        billingCouponId: billingCouponId\n" +
			"        configuration:\n" +
			"          startDate: '2019-12-27T18:11:19.117Z'\n" +
			"        couponId: couponId\n" +
			"        discount:\n" +
			"          amountsOff:\n" +
			"            - amount: 0\n" +
			"              currency: usd\n" +
			"          description: description\n" +
			"          durationInMonths: 1\n" +
			"          name: name\n" +
			"          percentOff: 1\n" +
			"        promotionCode: promotionCode\n" +
			"      awaitPaymentConfirmation: true\n" +
			"      billingCountryCode: billingCountryCode\n" +
			"      billingCycleAnchor: UNCHANGED\n" +
			"      billingId: billingId\n" +
			"      billingInformation:\n" +
			"        billingAddress:\n" +
			"          city: city\n" +
			"          country: country\n" +
			"          line1: line1\n" +
			"          line2: line2\n" +
			"          postalCode: postalCode\n" +
			"          state: state\n" +
			"        chargeOnBehalfOfAccount: chargeOnBehalfOfAccount\n" +
			"        integrationId: integrationId\n" +
			"        invoiceDaysUntilDue: 0\n" +
			"        isBackdated: true\n" +
			"        isInvoicePaid: true\n" +
			"        metadata:\n" +
			"          foo: string\n" +
			"        prorationBehavior: INVOICE_IMMEDIATELY\n" +
			"        taxIds:\n" +
			"          - type: type\n" +
			"            value: value\n" +
			"        taxPercentage: 0\n" +
			"        taxRateIds:\n" +
			"          - string\n" +
			"      billingPeriod: MONTHLY\n" +
			"      budget:\n" +
			"        hasSoftLimit: true\n" +
			"        limit: 0\n" +
			"      cancellationDate: '2019-12-27T18:11:19.117Z'\n" +
			"      charges:\n" +
			"        - id: id\n" +
			"          quantity: 0\n" +
			"          type: FEATURE\n" +
			"      checkoutOptions:\n" +
			"        cancelUrl: https://example.com\n" +
			"        successUrl: https://example.com\n" +
			"        allowPromoCodes: true\n" +
			"        allowTaxIdCollection: true\n" +
			"        collectBillingAddress: true\n" +
			"        collectPhoneNumber: true\n" +
			"        referenceId: referenceId\n" +
			"      entitlements:\n" +
			"        - id: id\n" +
			"          type: FEATURE\n" +
			"          hasSoftLimit: true\n" +
			"          hasUnlimitedUsage: true\n" +
			"          monthlyResetPeriodConfiguration:\n" +
			"            accordingTo: SubscriptionStart\n" +
			"          resetPeriod: YEAR\n" +
			"          usageLimit: 0\n" +
			"          weeklyResetPeriodConfiguration:\n" +
			"            accordingTo: SubscriptionStart\n" +
			"          yearlyResetPeriodConfiguration:\n" +
			"            accordingTo: SubscriptionStart\n" +
			"      metadata:\n" +
			"        foo: string\n" +
			"      minimumSpend:\n" +
			"        amount: 0\n" +
			"        currency: usd\n" +
			"      payingCustomerId: payingCustomerId\n" +
			"      paymentCollectionMethod: CHARGE\n" +
			"      priceOverrides:\n" +
			"        - addonId: addonId\n" +
			"          amount: 0\n" +
			"          baseCharge: true\n" +
			"          billingCountryCode: billingCountryCode\n" +
			"          blockSize: 0\n" +
			"          creditGrantCadence: BEGINNING_OF_BILLING_PERIOD\n" +
			"          creditRate:\n" +
			"            amount: 1\n" +
			"            currencyId: currencyId\n" +
			"            costFormula: costFormula\n" +
			"          currency: usd\n" +
			"          featureId: featureId\n" +
			"          tiers:\n" +
			"            - flatPrice:\n" +
			"                amount: 0\n" +
			"                currency: usd\n" +
			"              unitPrice:\n" +
			"                amount: 0\n" +
			"                currency: usd\n" +
			"              upTo: 0\n" +
			"      resourceId: resourceId\n" +
			"      salesforceId: salesforceId\n" +
			"      scheduleStrategy: END_OF_BILLING_PERIOD\n" +
			"      startDate: '2019-12-27T18:11:19.117Z'\n" +
			"      trialOverrideConfiguration:\n" +
			"        isTrial: true\n" +
			"        trialEndBehavior: CONVERT_TO_PAID\n" +
			"        trialEndDate: '2019-12-27T18:11:19.117Z'\n" +
			"      unitQuantity: 0\n" +
			"activationEndDate: '2019-12-27T18:11:19.117Z'\n" +
			"activationStartDate: '2019-12-27T18:11:19.117Z'\n" +
			"name: name\n" +
			"poNumber: poNumber\n" +
			"setupBilling: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:contracts", "create",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})
}

func TestV1ContractsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:contracts", "retrieve",
			"--id", "x",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})
}

func TestV1ContractsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:contracts", "update",
			"--id", "x",
			"--activation-end-date", "'2019-12-27T18:11:19.117Z'",
			"--activation-start-date", "'2019-12-27T18:11:19.117Z'",
			"--name", "name",
			"--po-number", "poNumber",
			"--setup-billing=true",
			"--subscription-id", "NxI",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"activationEndDate: '2019-12-27T18:11:19.117Z'\n" +
			"activationStartDate: '2019-12-27T18:11:19.117Z'\n" +
			"name: name\n" +
			"poNumber: poNumber\n" +
			"setupBilling: true\n" +
			"subscriptionIds:\n" +
			"  - NxI\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:contracts", "update",
			"--id", "x",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})
}

func TestV1ContractsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:contracts", "list",
			"--max-items", "10",
			"--after", "after",
			"--before", "before",
			"--customer-external-id", "customerExternalId",
			"--limit", "1",
			"--name", "name",
			"--state", "state",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})
}

func TestV1ContractsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:contracts", "delete",
			"--id", "x",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})
}
