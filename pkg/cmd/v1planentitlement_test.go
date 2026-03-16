// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stiggio/stigg-cli/internal/mocktest"
	"github.com/stiggio/stigg-cli/internal/requestflag"
)

func TestV1PlansEntitlementsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:plans:entitlements", "create",
			"--api-key", "string",
			"--plan-id", "planId",
			"--entitlement", "{id: id, type: FEATURE, behavior: Increment, description: description, displayNameOverride: displayNameOverride, enumValues: [string], hasSoftLimit: true, hasUnlimitedUsage: true, hiddenFromWidgets: [PAYWALL], isCustom: true, isGranted: true, monthlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, order: 0, resetPeriod: YEAR, usageLimit: 0, weeklyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, yearlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"entitlements:\n" +
			"  - id: id\n" +
			"    type: FEATURE\n" +
			"    behavior: Increment\n" +
			"    description: description\n" +
			"    displayNameOverride: displayNameOverride\n" +
			"    enumValues:\n" +
			"      - string\n" +
			"    hasSoftLimit: true\n" +
			"    hasUnlimitedUsage: true\n" +
			"    hiddenFromWidgets:\n" +
			"      - PAYWALL\n" +
			"    isCustom: true\n" +
			"    isGranted: true\n" +
			"    monthlyResetPeriodConfiguration:\n" +
			"      accordingTo: SubscriptionStart\n" +
			"    order: 0\n" +
			"    resetPeriod: YEAR\n" +
			"    usageLimit: 0\n" +
			"    weeklyResetPeriodConfiguration:\n" +
			"      accordingTo: SubscriptionStart\n" +
			"    yearlyResetPeriodConfiguration:\n" +
			"      accordingTo: SubscriptionStart\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1:plans:entitlements", "create",
			"--api-key", "string",
			"--plan-id", "planId",
		)
	})
}

func TestV1PlansEntitlementsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:plans:entitlements", "update",
			"--api-key", "string",
			"--plan-id", "planId",
			"--id", "id",
			"--type", "FEATURE",
			"--behavior", "Increment",
			"--description", "description",
			"--display-name-override", "displayNameOverride",
			"--enum-value", "string",
			"--has-soft-limit=true",
			"--has-unlimited-usage=true",
			"--hidden-from-widget", "PAYWALL",
			"--is-custom=true",
			"--is-granted=true",
			"--monthly-reset-period-configuration", "{accordingTo: SubscriptionStart}",
			"--order", "0",
			"--reset-period", "YEAR",
			"--usage-limit", "0",
			"--weekly-reset-period-configuration", "{accordingTo: SubscriptionStart}",
			"--yearly-reset-period-configuration", "{accordingTo: SubscriptionStart}",
			"--amount", "1",
			"--cadence", "MONTH",
			"--dependency-feature-id", "dependencyFeatureId",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1PlansEntitlementsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "v1:plans:entitlements", "update",
			"--api-key", "string",
			"--plan-id", "planId",
			"--id", "id",
			"--type", "FEATURE",
			"--behavior", "Increment",
			"--description", "description",
			"--display-name-override", "displayNameOverride",
			"--enum-value", "string",
			"--has-soft-limit=true",
			"--has-unlimited-usage=true",
			"--hidden-from-widget", "PAYWALL",
			"--is-custom=true",
			"--is-granted=true",
			"--monthly-reset-period-configuration.according-to", "SubscriptionStart",
			"--order", "0",
			"--reset-period", "YEAR",
			"--usage-limit", "0",
			"--weekly-reset-period-configuration.according-to", "SubscriptionStart",
			"--yearly-reset-period-configuration.according-to", "SubscriptionStart",
			"--amount", "1",
			"--cadence", "MONTH",
			"--dependency-feature-id", "dependencyFeatureId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"type: FEATURE\n" +
			"behavior: Increment\n" +
			"description: description\n" +
			"displayNameOverride: displayNameOverride\n" +
			"enumValues:\n" +
			"  - string\n" +
			"hasSoftLimit: true\n" +
			"hasUnlimitedUsage: true\n" +
			"hiddenFromWidgets:\n" +
			"  - PAYWALL\n" +
			"isCustom: true\n" +
			"isGranted: true\n" +
			"monthlyResetPeriodConfiguration:\n" +
			"  accordingTo: SubscriptionStart\n" +
			"order: 0\n" +
			"resetPeriod: YEAR\n" +
			"usageLimit: 0\n" +
			"weeklyResetPeriodConfiguration:\n" +
			"  accordingTo: SubscriptionStart\n" +
			"yearlyResetPeriodConfiguration:\n" +
			"  accordingTo: SubscriptionStart\n" +
			"amount: 1\n" +
			"cadence: MONTH\n" +
			"dependencyFeatureId: dependencyFeatureId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1:plans:entitlements", "update",
			"--api-key", "string",
			"--plan-id", "planId",
			"--id", "id",
		)
	})
}

func TestV1PlansEntitlementsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:plans:entitlements", "list",
			"--api-key", "string",
			"--plan-id", "planId",
		)
	})
}

func TestV1PlansEntitlementsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:plans:entitlements", "delete",
			"--api-key", "string",
			"--plan-id", "planId",
			"--id", "id",
		)
	})
}
