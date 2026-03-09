// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/stigg-cli/internal/mocktest"
	"github.com/stainless-sdks/stigg-cli/internal/requestflag"
)

func TestV1EventsCreditsGrantsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:events:credits:grants", "create",
			"--api-key", "string",
			"--amount", "0",
			"--currency-id", "currencyId",
			"--customer-id", "customerId",
			"--display-name", "displayName",
			"--grant-type", "PAID",
			"--await-payment-confirmation=true",
			"--billing-information", "{billingAddress: {city: city, country: country, line1: line1, line2: line2, postalCode: postalCode, state: state}, invoiceDaysUntilDue: 0, isInvoicePaid: true}",
			"--comment", "comment",
			"--cost", "{amount: 0, currency: usd}",
			"--effective-at", "'2019-12-27T18:11:19.117Z'",
			"--expire-at", "'2019-12-27T18:11:19.117Z'",
			"--metadata", "{foo: string}",
			"--payment-collection-method", "CHARGE",
			"--priority", "0",
			"--resource-id", "resourceId",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1EventsCreditsGrantsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "v1:events:credits:grants", "create",
			"--api-key", "string",
			"--amount", "0",
			"--currency-id", "currencyId",
			"--customer-id", "customerId",
			"--display-name", "displayName",
			"--grant-type", "PAID",
			"--await-payment-confirmation=true",
			"--billing-information.billing-address", "{city: city, country: country, line1: line1, line2: line2, postalCode: postalCode, state: state}",
			"--billing-information.invoice-days-until-due", "0",
			"--billing-information.is-invoice-paid=true",
			"--comment", "comment",
			"--cost.amount", "0",
			"--cost.currency", "usd",
			"--effective-at", "'2019-12-27T18:11:19.117Z'",
			"--expire-at", "'2019-12-27T18:11:19.117Z'",
			"--metadata", "{foo: string}",
			"--payment-collection-method", "CHARGE",
			"--priority", "0",
			"--resource-id", "resourceId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"amount: 0\n" +
			"currencyId: currencyId\n" +
			"customerId: customerId\n" +
			"displayName: displayName\n" +
			"grantType: PAID\n" +
			"awaitPaymentConfirmation: true\n" +
			"billingInformation:\n" +
			"  billingAddress:\n" +
			"    city: city\n" +
			"    country: country\n" +
			"    line1: line1\n" +
			"    line2: line2\n" +
			"    postalCode: postalCode\n" +
			"    state: state\n" +
			"  invoiceDaysUntilDue: 0\n" +
			"  isInvoicePaid: true\n" +
			"comment: comment\n" +
			"cost:\n" +
			"  amount: 0\n" +
			"  currency: usd\n" +
			"effectiveAt: '2019-12-27T18:11:19.117Z'\n" +
			"expireAt: '2019-12-27T18:11:19.117Z'\n" +
			"metadata:\n" +
			"  foo: string\n" +
			"paymentCollectionMethod: CHARGE\n" +
			"priority: 0\n" +
			"resourceId: resourceId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1:events:credits:grants", "create",
			"--api-key", "string",
		)
	})
}

func TestV1EventsCreditsGrantsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:events:credits:grants", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--customer-id", "customerId",
			"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--created-at", "{gt: '2019-12-27T18:11:19.117Z', gte: '2019-12-27T18:11:19.117Z', lt: '2019-12-27T18:11:19.117Z', lte: '2019-12-27T18:11:19.117Z'}",
			"--currency-id", "currencyId",
			"--limit", "1",
			"--resource-id", "resourceId",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1EventsCreditsGrantsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "v1:events:credits:grants", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--customer-id", "customerId",
			"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--created-at.gt", "2019-12-27T18:11:19.117Z",
			"--created-at.gte", "2019-12-27T18:11:19.117Z",
			"--created-at.lt", "2019-12-27T18:11:19.117Z",
			"--created-at.lte", "2019-12-27T18:11:19.117Z",
			"--currency-id", "currencyId",
			"--limit", "1",
			"--resource-id", "resourceId",
		)
	})
}

func TestV1EventsCreditsGrantsVoid(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:events:credits:grants", "void",
			"--api-key", "string",
			"--id", "x",
		)
	})
}
