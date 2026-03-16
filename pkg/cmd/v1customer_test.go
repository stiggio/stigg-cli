// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stiggio/stigg-cli/internal/mocktest"
	"github.com/stiggio/stigg-cli/internal/requestflag"
)

func TestV1CustomersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:customers", "retrieve",
			"--api-key", "string",
			"--id", "x",
		)
	})
}

func TestV1CustomersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:customers", "update",
			"--api-key", "string",
			"--id", "x",
			"--billing-currency", "usd",
			"--billing-id", "billingId",
			"--coupon-id", "''",
			"--email", "dev@stainless.com",
			"--integration", "{id: id, syncedEntityId: syncedEntityId, vendorIdentifier: AUTH0}",
			"--language", "language",
			"--metadata", "{foo: string}",
			"--name", "name",
			"--passthrough", "{stripe: {billingAddress: {city: city, country: country, line1: line1, line2: line2, postalCode: postalCode, state: state}, customerName: customerName, invoiceCustomFields: {foo: string}, metadata: {foo: string}, paymentMethodId: paymentMethodId, shippingAddress: {city: city, country: country, line1: line1, line2: line2, postalCode: postalCode, state: state}, taxIds: [{type: type, value: value}]}, zuora: {billingAddress: {city: city, country: country, line1: line1, line2: line2, postalCode: postalCode, state: state}, currency: usd, metadata: {foo: string}, paymentMethodId: paymentMethodId}}",
			"--timezone", "timezone",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1CustomersUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "v1:customers", "update",
			"--api-key", "string",
			"--id", "x",
			"--billing-currency", "usd",
			"--billing-id", "billingId",
			"--coupon-id", "''",
			"--email", "dev@stainless.com",
			"--integration.id", "id",
			"--integration.synced-entity-id", "syncedEntityId",
			"--integration.vendor-identifier", "AUTH0",
			"--language", "language",
			"--metadata", "{foo: string}",
			"--name", "name",
			"--passthrough.stripe", "{billingAddress: {city: city, country: country, line1: line1, line2: line2, postalCode: postalCode, state: state}, customerName: customerName, invoiceCustomFields: {foo: string}, metadata: {foo: string}, paymentMethodId: paymentMethodId, shippingAddress: {city: city, country: country, line1: line1, line2: line2, postalCode: postalCode, state: state}, taxIds: [{type: type, value: value}]}",
			"--passthrough.zuora", "{billingAddress: {city: city, country: country, line1: line1, line2: line2, postalCode: postalCode, state: state}, currency: usd, metadata: {foo: string}, paymentMethodId: paymentMethodId}",
			"--timezone", "timezone",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"billingCurrency: usd\n" +
			"billingId: billingId\n" +
			"couponId: ''\n" +
			"email: dev@stainless.com\n" +
			"integrations:\n" +
			"  - id: id\n" +
			"    syncedEntityId: syncedEntityId\n" +
			"    vendorIdentifier: AUTH0\n" +
			"language: language\n" +
			"metadata:\n" +
			"  foo: string\n" +
			"name: name\n" +
			"passthrough:\n" +
			"  stripe:\n" +
			"    billingAddress:\n" +
			"      city: city\n" +
			"      country: country\n" +
			"      line1: line1\n" +
			"      line2: line2\n" +
			"      postalCode: postalCode\n" +
			"      state: state\n" +
			"    customerName: customerName\n" +
			"    invoiceCustomFields:\n" +
			"      foo: string\n" +
			"    metadata:\n" +
			"      foo: string\n" +
			"    paymentMethodId: paymentMethodId\n" +
			"    shippingAddress:\n" +
			"      city: city\n" +
			"      country: country\n" +
			"      line1: line1\n" +
			"      line2: line2\n" +
			"      postalCode: postalCode\n" +
			"      state: state\n" +
			"    taxIds:\n" +
			"      - type: type\n" +
			"        value: value\n" +
			"  zuora:\n" +
			"    billingAddress:\n" +
			"      city: city\n" +
			"      country: country\n" +
			"      line1: line1\n" +
			"      line2: line2\n" +
			"      postalCode: postalCode\n" +
			"      state: state\n" +
			"    currency: usd\n" +
			"    metadata:\n" +
			"      foo: string\n" +
			"    paymentMethodId: paymentMethodId\n" +
			"timezone: timezone\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1:customers", "update",
			"--api-key", "string",
			"--id", "x",
		)
	})
}

func TestV1CustomersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:customers", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--created-at", "{gt: '2019-12-27T18:11:19.117Z', gte: '2019-12-27T18:11:19.117Z', lt: '2019-12-27T18:11:19.117Z', lte: '2019-12-27T18:11:19.117Z'}",
			"--email", "email",
			"--limit", "1",
			"--name", "name",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1CustomersList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "v1:customers", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--created-at.gt", "2019-12-27T18:11:19.117Z",
			"--created-at.gte", "2019-12-27T18:11:19.117Z",
			"--created-at.lt", "2019-12-27T18:11:19.117Z",
			"--created-at.lte", "2019-12-27T18:11:19.117Z",
			"--email", "email",
			"--limit", "1",
			"--name", "name",
		)
	})
}

func TestV1CustomersArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:customers", "archive",
			"--api-key", "string",
			"--id", "x",
		)
	})
}

func TestV1CustomersImport(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:customers", "import",
			"--api-key", "string",
			"--customer", "{id: id, email: dev@stainless.com, name: name, billingId: billingId, metadata: {foo: string}, paymentMethodId: paymentMethodId, salesforceId: salesforceId, updatedAt: '2019-12-27T18:11:19.117Z'}",
			"--integration-id", "integrationId",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1CustomersImport)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "v1:customers", "import",
			"--api-key", "string",
			"--customer.id", "id",
			"--customer.email", "dev@stainless.com",
			"--customer.name", "name",
			"--customer.billing-id", "billingId",
			"--customer.metadata", "{foo: string}",
			"--customer.payment-method-id", "paymentMethodId",
			"--customer.salesforce-id", "salesforceId",
			"--customer.updated-at", "2019-12-27T18:11:19.117Z",
			"--integration-id", "integrationId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"customers:\n" +
			"  - id: id\n" +
			"    email: dev@stainless.com\n" +
			"    name: name\n" +
			"    billingId: billingId\n" +
			"    metadata:\n" +
			"      foo: string\n" +
			"    paymentMethodId: paymentMethodId\n" +
			"    salesforceId: salesforceId\n" +
			"    updatedAt: '2019-12-27T18:11:19.117Z'\n" +
			"integrationId: integrationId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1:customers", "import",
			"--api-key", "string",
		)
	})
}

func TestV1CustomersListResources(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:customers", "list-resources",
			"--api-key", "string",
			"--max-items", "10",
			"--id", "x",
			"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--limit", "1",
		)
	})
}

func TestV1CustomersProvision(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:customers", "provision",
			"--api-key", "string",
			"--id", "id",
			"--billing-currency", "usd",
			"--billing-id", "billingId",
			"--coupon-id", "''",
			"--default-payment-method", "{billingId: billingId, cardExpiryMonth: 0, cardExpiryYear: 0, cardLast4Digits: cardLast4Digits, type: CARD}",
			"--email", "dev@stainless.com",
			"--integration", "{id: id, syncedEntityId: syncedEntityId, vendorIdentifier: AUTH0}",
			"--language", "language",
			"--metadata", "{foo: string}",
			"--name", "name",
			"--passthrough", "{stripe: {billingAddress: {city: city, country: country, line1: line1, line2: line2, postalCode: postalCode, state: state}, customerName: customerName, invoiceCustomFields: {foo: string}, metadata: {foo: string}, paymentMethodId: paymentMethodId, shippingAddress: {city: city, country: country, line1: line1, line2: line2, postalCode: postalCode, state: state}, taxIds: [{type: type, value: value}]}, zuora: {billingAddress: {city: city, country: country, line1: line1, line2: line2, postalCode: postalCode, state: state}, currency: usd, metadata: {foo: string}, paymentMethodId: paymentMethodId}}",
			"--timezone", "timezone",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1CustomersProvision)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "v1:customers", "provision",
			"--api-key", "string",
			"--id", "id",
			"--billing-currency", "usd",
			"--billing-id", "billingId",
			"--coupon-id", "''",
			"--default-payment-method.billing-id", "billingId",
			"--default-payment-method.card-expiry-month", "0",
			"--default-payment-method.card-expiry-year", "0",
			"--default-payment-method.card-last4-digits", "cardLast4Digits",
			"--default-payment-method.type", "CARD",
			"--email", "dev@stainless.com",
			"--integration.id", "id",
			"--integration.synced-entity-id", "syncedEntityId",
			"--integration.vendor-identifier", "AUTH0",
			"--language", "language",
			"--metadata", "{foo: string}",
			"--name", "name",
			"--passthrough.stripe", "{billingAddress: {city: city, country: country, line1: line1, line2: line2, postalCode: postalCode, state: state}, customerName: customerName, invoiceCustomFields: {foo: string}, metadata: {foo: string}, paymentMethodId: paymentMethodId, shippingAddress: {city: city, country: country, line1: line1, line2: line2, postalCode: postalCode, state: state}, taxIds: [{type: type, value: value}]}",
			"--passthrough.zuora", "{billingAddress: {city: city, country: country, line1: line1, line2: line2, postalCode: postalCode, state: state}, currency: usd, metadata: {foo: string}, paymentMethodId: paymentMethodId}",
			"--timezone", "timezone",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"id: id\n" +
			"billingCurrency: usd\n" +
			"billingId: billingId\n" +
			"couponId: ''\n" +
			"defaultPaymentMethod:\n" +
			"  billingId: billingId\n" +
			"  cardExpiryMonth: 0\n" +
			"  cardExpiryYear: 0\n" +
			"  cardLast4Digits: cardLast4Digits\n" +
			"  type: CARD\n" +
			"email: dev@stainless.com\n" +
			"integrations:\n" +
			"  - id: id\n" +
			"    syncedEntityId: syncedEntityId\n" +
			"    vendorIdentifier: AUTH0\n" +
			"language: language\n" +
			"metadata:\n" +
			"  foo: string\n" +
			"name: name\n" +
			"passthrough:\n" +
			"  stripe:\n" +
			"    billingAddress:\n" +
			"      city: city\n" +
			"      country: country\n" +
			"      line1: line1\n" +
			"      line2: line2\n" +
			"      postalCode: postalCode\n" +
			"      state: state\n" +
			"    customerName: customerName\n" +
			"    invoiceCustomFields:\n" +
			"      foo: string\n" +
			"    metadata:\n" +
			"      foo: string\n" +
			"    paymentMethodId: paymentMethodId\n" +
			"    shippingAddress:\n" +
			"      city: city\n" +
			"      country: country\n" +
			"      line1: line1\n" +
			"      line2: line2\n" +
			"      postalCode: postalCode\n" +
			"      state: state\n" +
			"    taxIds:\n" +
			"      - type: type\n" +
			"        value: value\n" +
			"  zuora:\n" +
			"    billingAddress:\n" +
			"      city: city\n" +
			"      country: country\n" +
			"      line1: line1\n" +
			"      line2: line2\n" +
			"      postalCode: postalCode\n" +
			"      state: state\n" +
			"    currency: usd\n" +
			"    metadata:\n" +
			"      foo: string\n" +
			"    paymentMethodId: paymentMethodId\n" +
			"timezone: timezone\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1:customers", "provision",
			"--api-key", "string",
		)
	})
}

func TestV1CustomersUnarchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:customers", "unarchive",
			"--api-key", "string",
			"--id", "x",
		)
	})
}
