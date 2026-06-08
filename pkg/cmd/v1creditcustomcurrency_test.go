// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stiggio/stigg-cli/internal/mocktest"
	"github.com/stiggio/stigg-cli/internal/requestflag"
)

func TestV1CreditsCustomCurrenciesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:credits:custom-currencies", "create",
			"--id", "id",
			"--display-name", "displayName",
			"--description", "description",
			"--metadata", "{foo: string}",
			"--symbol", "symbol",
			"--units", "{plural: plural, singular: singular}",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1CreditsCustomCurrenciesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:credits:custom-currencies", "create",
			"--id", "id",
			"--display-name", "displayName",
			"--description", "description",
			"--metadata", "{foo: string}",
			"--symbol", "symbol",
			"--units.plural", "plural",
			"--units.singular", "singular",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"id: id\n" +
			"displayName: displayName\n" +
			"description: description\n" +
			"metadata:\n" +
			"  foo: string\n" +
			"symbol: symbol\n" +
			"units:\n" +
			"  plural: plural\n" +
			"  singular: singular\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:credits:custom-currencies", "create",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})
}

func TestV1CreditsCustomCurrenciesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:credits:custom-currencies", "update",
			"--currency-id", "currencyId",
			"--description", "description",
			"--display-name", "displayName",
			"--metadata", "{foo: string}",
			"--symbol", "symbol",
			"--units", "{plural: plural, singular: singular}",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1CreditsCustomCurrenciesUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:credits:custom-currencies", "update",
			"--currency-id", "currencyId",
			"--description", "description",
			"--display-name", "displayName",
			"--metadata", "{foo: string}",
			"--symbol", "symbol",
			"--units.plural", "plural",
			"--units.singular", "singular",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: description\n" +
			"displayName: displayName\n" +
			"metadata:\n" +
			"  foo: string\n" +
			"symbol: symbol\n" +
			"units:\n" +
			"  plural: plural\n" +
			"  singular: singular\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:credits:custom-currencies", "update",
			"--currency-id", "currencyId",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})
}

func TestV1CreditsCustomCurrenciesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:credits:custom-currencies", "list",
			"--max-items", "10",
			"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--limit", "1",
			"--status", "ACTIVE",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})
}

func TestV1CreditsCustomCurrenciesArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:credits:custom-currencies", "archive",
			"--currency-id", "currencyId",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})
}

func TestV1CreditsCustomCurrenciesListAssociatedEntities(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:credits:custom-currencies", "list-associated-entities",
			"--currency-id", "currencyId",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})
}

func TestV1CreditsCustomCurrenciesUnarchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:credits:custom-currencies", "unarchive",
			"--currency-id", "currencyId",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})
}
