// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stiggio/stigg-cli/internal/mocktest"
	"github.com/stiggio/stigg-cli/internal/requestflag"
)

func TestV1EventsCreditsCustomCurrenciesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:events:credits:custom-currencies", "create",
			"--id", "id",
			"--display-name", "displayName",
			"--description", "description",
			"--metadata", "{foo: string}",
			"--symbol", "symbol",
			"--units", "{plural: plural, singular: singular}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1EventsCreditsCustomCurrenciesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:events:credits:custom-currencies", "create",
			"--id", "id",
			"--display-name", "displayName",
			"--description", "description",
			"--metadata", "{foo: string}",
			"--symbol", "symbol",
			"--units.plural", "plural",
			"--units.singular", "singular",
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
			"v1:events:credits:custom-currencies", "create",
		)
	})
}

func TestV1EventsCreditsCustomCurrenciesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:events:credits:custom-currencies", "update",
			"--currency-id", "currencyId",
			"--description", "description",
			"--display-name", "displayName",
			"--metadata", "{foo: string}",
			"--symbol", "symbol",
			"--units", "{plural: plural, singular: singular}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1EventsCreditsCustomCurrenciesUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:events:credits:custom-currencies", "update",
			"--currency-id", "currencyId",
			"--description", "description",
			"--display-name", "displayName",
			"--metadata", "{foo: string}",
			"--symbol", "symbol",
			"--units.plural", "plural",
			"--units.singular", "singular",
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
			"v1:events:credits:custom-currencies", "update",
			"--currency-id", "currencyId",
		)
	})
}

func TestV1EventsCreditsCustomCurrenciesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:events:credits:custom-currencies", "list",
			"--max-items", "10",
			"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--limit", "1",
			"--status", "ACTIVE",
		)
	})
}

func TestV1EventsCreditsCustomCurrenciesArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:events:credits:custom-currencies", "archive",
			"--currency-id", "currencyId",
		)
	})
}

func TestV1EventsCreditsCustomCurrenciesListAssociatedEntities(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:events:credits:custom-currencies", "list-associated-entities",
			"--currency-id", "currencyId",
		)
	})
}

func TestV1EventsCreditsCustomCurrenciesUnarchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:events:credits:custom-currencies", "unarchive",
			"--currency-id", "currencyId",
		)
	})
}
