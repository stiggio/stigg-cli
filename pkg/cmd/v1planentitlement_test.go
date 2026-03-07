// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/stigg-cli/internal/mocktest"
	"github.com/stainless-sdks/stigg-cli/internal/requestflag"
)

func TestV1PlansEntitlementsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:plans:entitlements", "create",
			"--api-key", "string",
			"--plan-id", "planId",
			"--entitlement", "{credit: {amount: 1, cadence: MONTH, customCurrencyId: customCurrencyId, behavior: Increment, description: description, displayNameOverride: displayNameOverride, hiddenFromWidgets: [PAYWALL], isCustom: true, isGranted: true, order: 0}, feature: {featureId: featureId, behavior: Increment, description: description, displayNameOverride: displayNameOverride, enumValues: [string], hasSoftLimit: true, hasUnlimitedUsage: true, hiddenFromWidgets: [PAYWALL], isCustom: true, isGranted: true, monthlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, order: 0, resetPeriod: YEAR, usageLimit: 0, weeklyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, yearlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1PlansEntitlementsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "v1:plans:entitlements", "create",
			"--api-key", "string",
			"--plan-id", "planId",
			"--entitlement.credit", "{amount: 1, cadence: MONTH, customCurrencyId: customCurrencyId, behavior: Increment, description: description, displayNameOverride: displayNameOverride, hiddenFromWidgets: [PAYWALL], isCustom: true, isGranted: true, order: 0}",
			"--entitlement.feature", "{featureId: featureId, behavior: Increment, description: description, displayNameOverride: displayNameOverride, enumValues: [string], hasSoftLimit: true, hasUnlimitedUsage: true, hiddenFromWidgets: [PAYWALL], isCustom: true, isGranted: true, monthlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, order: 0, resetPeriod: YEAR, usageLimit: 0, weeklyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, yearlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"entitlements:\n" +
			"  - credit:\n" +
			"      amount: 1\n" +
			"      cadence: MONTH\n" +
			"      customCurrencyId: customCurrencyId\n" +
			"      behavior: Increment\n" +
			"      description: description\n" +
			"      displayNameOverride: displayNameOverride\n" +
			"      hiddenFromWidgets:\n" +
			"        - PAYWALL\n" +
			"      isCustom: true\n" +
			"      isGranted: true\n" +
			"      order: 0\n" +
			"    feature:\n" +
			"      featureId: featureId\n" +
			"      behavior: Increment\n" +
			"      description: description\n" +
			"      displayNameOverride: displayNameOverride\n" +
			"      enumValues:\n" +
			"        - string\n" +
			"      hasSoftLimit: true\n" +
			"      hasUnlimitedUsage: true\n" +
			"      hiddenFromWidgets:\n" +
			"        - PAYWALL\n" +
			"      isCustom: true\n" +
			"      isGranted: true\n" +
			"      monthlyResetPeriodConfiguration:\n" +
			"        accordingTo: SubscriptionStart\n" +
			"      order: 0\n" +
			"      resetPeriod: YEAR\n" +
			"      usageLimit: 0\n" +
			"      weeklyResetPeriodConfiguration:\n" +
			"        accordingTo: SubscriptionStart\n" +
			"      yearlyResetPeriodConfiguration:\n" +
			"        accordingTo: SubscriptionStart\n")
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
			"--credit", "{amount: 1, behavior: Increment, cadence: MONTH, description: description, displayNameOverride: displayNameOverride, hiddenFromWidgets: [PAYWALL], isCustom: true, isGranted: true, order: 0}",
			"--feature", "{behavior: Increment, description: description, displayNameOverride: displayNameOverride, enumValues: [string], hasSoftLimit: true, hasUnlimitedUsage: true, hiddenFromWidgets: [PAYWALL], isCustom: true, isGranted: true, monthlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, order: 0, resetPeriod: YEAR, usageLimit: 0, weeklyResetPeriodConfiguration: {accordingTo: SubscriptionStart}, yearlyResetPeriodConfiguration: {accordingTo: SubscriptionStart}}",
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
			"--credit.amount", "1",
			"--credit.behavior", "Increment",
			"--credit.cadence", "MONTH",
			"--credit.description", "description",
			"--credit.display-name-override", "displayNameOverride",
			"--credit.hidden-from-widgets", "[PAYWALL]",
			"--credit.is-custom=true",
			"--credit.is-granted=true",
			"--credit.order", "0",
			"--feature.behavior", "Increment",
			"--feature.description", "description",
			"--feature.display-name-override", "displayNameOverride",
			"--feature.enum-values", "[string]",
			"--feature.has-soft-limit=true",
			"--feature.has-unlimited-usage=true",
			"--feature.hidden-from-widgets", "[PAYWALL]",
			"--feature.is-custom=true",
			"--feature.is-granted=true",
			"--feature.monthly-reset-period-configuration", "{accordingTo: SubscriptionStart}",
			"--feature.order", "0",
			"--feature.reset-period", "YEAR",
			"--feature.usage-limit", "0",
			"--feature.weekly-reset-period-configuration", "{accordingTo: SubscriptionStart}",
			"--feature.yearly-reset-period-configuration", "{accordingTo: SubscriptionStart}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"credit:\n" +
			"  amount: 1\n" +
			"  behavior: Increment\n" +
			"  cadence: MONTH\n" +
			"  description: description\n" +
			"  displayNameOverride: displayNameOverride\n" +
			"  hiddenFromWidgets:\n" +
			"    - PAYWALL\n" +
			"  isCustom: true\n" +
			"  isGranted: true\n" +
			"  order: 0\n" +
			"feature:\n" +
			"  behavior: Increment\n" +
			"  description: description\n" +
			"  displayNameOverride: displayNameOverride\n" +
			"  enumValues:\n" +
			"    - string\n" +
			"  hasSoftLimit: true\n" +
			"  hasUnlimitedUsage: true\n" +
			"  hiddenFromWidgets:\n" +
			"    - PAYWALL\n" +
			"  isCustom: true\n" +
			"  isGranted: true\n" +
			"  monthlyResetPeriodConfiguration:\n" +
			"    accordingTo: SubscriptionStart\n" +
			"  order: 0\n" +
			"  resetPeriod: YEAR\n" +
			"  usageLimit: 0\n" +
			"  weeklyResetPeriodConfiguration:\n" +
			"    accordingTo: SubscriptionStart\n" +
			"  yearlyResetPeriodConfiguration:\n" +
			"    accordingTo: SubscriptionStart\n")
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
