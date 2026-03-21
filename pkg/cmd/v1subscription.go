// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stiggio/stigg-cli/internal/apiquery"
	"github.com/stiggio/stigg-cli/internal/requestflag"
	"github.com/stiggio/stigg-go"
	"github.com/stiggio/stigg-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var v1SubscriptionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a subscription by its unique identifier, including plan details,\nbilling period, status, and add-ons.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleV1SubscriptionsRetrieve,
	HideHelpCommand: true,
}

var v1SubscriptionsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates an active subscription's properties including billing period, add-ons,\nunit quantities, and discounts.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "addon",
			BodyPath: "addons",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "applied-coupon",
			BodyPath: "appliedCoupon",
		},
		&requestflag.Flag[bool]{
			Name:     "await-payment-confirmation",
			Usage:    "Await payment confirmation",
			BodyPath: "awaitPaymentConfirmation",
		},
		&requestflag.Flag[string]{
			Name:     "billing-cycle-anchor",
			Usage:    `Allowed values: "UNCHANGED", "NOW".`,
			BodyPath: "billingCycleAnchor",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "billing-information",
			BodyPath: "billingInformation",
		},
		&requestflag.Flag[string]{
			Name:     "billing-period",
			Usage:    `Allowed values: "MONTHLY", "ANNUALLY".`,
			BodyPath: "billingPeriod",
		},
		&requestflag.Flag[any]{
			Name:     "budget",
			BodyPath: "budget",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "charge",
			BodyPath: "charges",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "entitlement",
			BodyPath: "entitlements",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Additional metadata for the subscription",
			BodyPath: "metadata",
		},
		&requestflag.Flag[any]{
			Name:     "minimum-spend",
			Usage:    "Minimum spend amount",
			BodyPath: "minimumSpend",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "price-override",
			BodyPath: "priceOverrides",
		},
		&requestflag.Flag[string]{
			Name:     "promotion-code",
			Usage:    "Promotion code",
			BodyPath: "promotionCode",
		},
		&requestflag.Flag[string]{
			Name:     "schedule-strategy",
			Usage:    `Allowed values: "END_OF_BILLING_PERIOD", "END_OF_BILLING_MONTH", "IMMEDIATE".`,
			BodyPath: "scheduleStrategy",
		},
		&requestflag.Flag[any]{
			Name:     "trial-end-date",
			Usage:    "Subscription trial end date",
			BodyPath: "trialEndDate",
		},
	},
	Action:          handleV1SubscriptionsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"addon": {
		&requestflag.InnerFlag[string]{
			Name:       "addon.id",
			Usage:      "Addon ID",
			InnerField: "id",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "addon.quantity",
			Usage:      "Number of addon instances",
			InnerField: "quantity",
		},
	},
	"applied-coupon": {
		&requestflag.InnerFlag[string]{
			Name:       "applied-coupon.billing-coupon-id",
			InnerField: "billingCouponId",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "applied-coupon.configuration",
			InnerField: "configuration",
		},
		&requestflag.InnerFlag[string]{
			Name:       "applied-coupon.coupon-id",
			Usage:      "Stigg coupon ID",
			InnerField: "couponId",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "applied-coupon.discount",
			InnerField: "discount",
		},
		&requestflag.InnerFlag[any]{
			Name:       "applied-coupon.promotion-code",
			InnerField: "promotionCode",
		},
	},
	"billing-information": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "billing-information.billing-address",
			Usage:      "Physical address",
			InnerField: "billingAddress",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-information.charge-on-behalf-of-account",
			InnerField: "chargeOnBehalfOfAccount",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-information.coupon-id",
			InnerField: "couponId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-information.integration-id",
			InnerField: "integrationId",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "billing-information.invoice-days-until-due",
			InnerField: "invoiceDaysUntilDue",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "billing-information.is-backdated",
			InnerField: "isBackdated",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "billing-information.is-invoice-paid",
			InnerField: "isInvoicePaid",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "billing-information.metadata",
			Usage:      "Additional metadata for the subscription",
			InnerField: "metadata",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-information.proration-behavior",
			Usage:      `Allowed values: "INVOICE_IMMEDIATELY", "CREATE_PRORATIONS", "NONE".`,
			InnerField: "prorationBehavior",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "billing-information.tax-ids",
			InnerField: "taxIds",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "billing-information.tax-percentage",
			InnerField: "taxPercentage",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "billing-information.tax-rate-ids",
			InnerField: "taxRateIds",
		},
	},
	"budget": {
		&requestflag.InnerFlag[bool]{
			Name:       "budget.has-soft-limit",
			Usage:      "Whether the budget is a soft limit",
			InnerField: "hasSoftLimit",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "budget.limit",
			Usage:      "Maximum spending limit",
			InnerField: "limit",
		},
	},
	"charge": {
		&requestflag.InnerFlag[string]{
			Name:       "charge.id",
			Usage:      "Charge ID",
			InnerField: "id",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "charge.quantity",
			InnerField: "quantity",
		},
		&requestflag.InnerFlag[string]{
			Name:       "charge.type",
			Usage:      `Allowed values: "FEATURE", "CREDIT".`,
			InnerField: "type",
		},
	},
	"minimum-spend": {
		&requestflag.InnerFlag[float64]{
			Name:       "minimum-spend.amount",
			Usage:      "The price amount",
			InnerField: "amount",
		},
		&requestflag.InnerFlag[string]{
			Name:       "minimum-spend.currency",
			Usage:      "The price currency",
			InnerField: "currency",
		},
	},
	"price-override": {
		&requestflag.InnerFlag[string]{
			Name:       "price-override.addon-id",
			Usage:      "Addon ID",
			InnerField: "addonId",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "price-override.amount",
			Usage:      "The price amount",
			InnerField: "amount",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "price-override.base-charge",
			Usage:      "Whether this is a base charge override",
			InnerField: "baseCharge",
		},
		&requestflag.InnerFlag[string]{
			Name:       "price-override.currency",
			Usage:      "The price currency",
			InnerField: "currency",
		},
		&requestflag.InnerFlag[string]{
			Name:       "price-override.currency-id",
			Usage:      "The corresponding custom currency id of the recurring credits price",
			InnerField: "currencyId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "price-override.feature-id",
			Usage:      "Feature ID",
			InnerField: "featureId",
		},
	},
})

var v1SubscriptionsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Retrieves a paginated list of subscriptions, with optional filters for customer,\nstatus, and plan.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "after",
			Usage:     "Return items that come after this cursor",
			QueryPath: "after",
		},
		&requestflag.Flag[string]{
			Name:      "before",
			Usage:     "Return items that come before this cursor",
			QueryPath: "before",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "created-at",
			Usage:     "Filter by creation date using range operators: gt, gte, lt, lte",
			QueryPath: "createdAt",
		},
		&requestflag.Flag[string]{
			Name:      "customer-id",
			Usage:     "Filter by customer ID",
			QueryPath: "customerId",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "plan-id",
			Usage:     "Filter by plan ID",
			QueryPath: "planId",
		},
		&requestflag.Flag[string]{
			Name:      "pricing-type",
			Usage:     "Filter by pricing type. Supports comma-separated values for multiple types",
			QueryPath: "pricingType",
		},
		&requestflag.Flag[string]{
			Name:      "resource-id",
			Usage:     "Filter by resource ID",
			QueryPath: "resourceId",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter by subscription status. Supports comma-separated values for multiple statuses",
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleV1SubscriptionsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"created-at": {
		&requestflag.InnerFlag[any]{
			Name:       "created-at.gt",
			Usage:      "Greater than the specified createdAt value",
			InnerField: "gt",
		},
		&requestflag.InnerFlag[any]{
			Name:       "created-at.gte",
			Usage:      "Greater than or equal to the specified createdAt value",
			InnerField: "gte",
		},
		&requestflag.InnerFlag[any]{
			Name:       "created-at.lt",
			Usage:      "Less than the specified createdAt value",
			InnerField: "lt",
		},
		&requestflag.InnerFlag[any]{
			Name:       "created-at.lte",
			Usage:      "Less than or equal to the specified createdAt value",
			InnerField: "lte",
		},
	},
})

var v1SubscriptionsCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancels an active subscription, either immediately or at a specified time such\nas end of billing period.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "cancellation-action",
			Usage:    "Action on cancellation (downgrade or revoke)",
			BodyPath: "cancellationAction",
		},
		&requestflag.Flag[string]{
			Name:     "cancellation-time",
			Usage:    "When to cancel (immediate, period end, or date)",
			BodyPath: "cancellationTime",
		},
		&requestflag.Flag[any]{
			Name:     "end-date",
			Usage:    "Subscription end date",
			BodyPath: "endDate",
		},
		&requestflag.Flag[bool]{
			Name:     "prorate",
			Usage:    "If set, enables or disables prorating of credits on subscription cancellation.",
			BodyPath: "prorate",
		},
	},
	Action:          handleV1SubscriptionsCancel,
	HideHelpCommand: true,
}

var v1SubscriptionsDelegate = cli.Command{
	Name:    "delegate",
	Usage:   "Delegates the payment responsibility of a subscription to a different customer.\nThe delegated customer will be billed for this subscription.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "target-customer-id",
			Usage:    "The unique identifier of the customer who will assume payment responsibility for this subscription. This customer must already exist in your Stigg account and have a valid payment method if the subscription requires payment.",
			Required: true,
			BodyPath: "targetCustomerId",
		},
	},
	Action:          handleV1SubscriptionsDelegate,
	HideHelpCommand: true,
}

var v1SubscriptionsImport = requestflag.WithInnerFlags(cli.Command{
	Name:    "import",
	Usage:   "Imports multiple subscriptions in bulk. Used for migrating subscription data\nfrom external systems.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "subscription",
			Usage:    "List of subscription objects to import",
			Required: true,
			BodyPath: "subscriptions",
		},
		&requestflag.Flag[any]{
			Name:     "integration-id",
			Usage:    "Integration ID to use for importing subscriptions",
			BodyPath: "integrationId",
		},
	},
	Action:          handleV1SubscriptionsImport,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"subscription": {
		&requestflag.InnerFlag[string]{
			Name:       "subscription.id",
			Usage:      "Subscription ID",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "subscription.customer-id",
			Usage:      "Customer ID",
			InnerField: "customerId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "subscription.plan-id",
			Usage:      "Plan ID",
			InnerField: "planId",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "subscription.addons",
			InnerField: "addons",
		},
		&requestflag.InnerFlag[any]{
			Name:       "subscription.billing-id",
			Usage:      "Billing ID",
			InnerField: "billingId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "subscription.billing-period",
			Usage:      "Billing period (MONTHLY or ANNUALLY)",
			InnerField: "billingPeriod",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "subscription.charges",
			InnerField: "charges",
		},
		&requestflag.InnerFlag[any]{
			Name:       "subscription.end-date",
			Usage:      "Subscription end date",
			InnerField: "endDate",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "subscription.metadata",
			Usage:      "Additional metadata for the subscription",
			InnerField: "metadata",
		},
		&requestflag.InnerFlag[any]{
			Name:       "subscription.resource-id",
			Usage:      "Resource ID",
			InnerField: "resourceId",
		},
		&requestflag.InnerFlag[any]{
			Name:       "subscription.start-date",
			Usage:      "Subscription start date",
			InnerField: "startDate",
		},
	},
})

var v1SubscriptionsMigrate = cli.Command{
	Name:    "migrate",
	Usage:   "Migrates a subscription to the latest published version of its plan or add-ons.\nHandles prorated charges or credits automatically.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "subscription-migration-time",
			Usage:    "When to migrate (immediate or period end)",
			BodyPath: "subscriptionMigrationTime",
		},
	},
	Action:          handleV1SubscriptionsMigrate,
	HideHelpCommand: true,
}

var v1SubscriptionsPreview = requestflag.WithInnerFlags(cli.Command{
	Name:    "preview",
	Usage:   "Previews the pricing impact of creating or updating a subscription without\nmaking changes. Returns estimated costs, taxes, and proration details.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "customer-id",
			Usage:    "Customer ID",
			Required: true,
			BodyPath: "customerId",
		},
		&requestflag.Flag[string]{
			Name:     "plan-id",
			Usage:    "Plan ID",
			Required: true,
			BodyPath: "planId",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "addon",
			Usage:    "Addons to include",
			BodyPath: "addons",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "applied-coupon",
			Usage:    "Coupon or discount to apply",
			BodyPath: "appliedCoupon",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "billable-feature",
			Usage:    "Billable features with quantities",
			BodyPath: "billableFeatures",
		},
		&requestflag.Flag[string]{
			Name:     "billing-country-code",
			Usage:    "ISO 3166-1 country code for localization",
			BodyPath: "billingCountryCode",
		},
		&requestflag.Flag[string]{
			Name:     "billing-cycle-anchor",
			Usage:    "Billing cycle anchor behavior for the subscription",
			BodyPath: "billingCycleAnchor",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "billing-information",
			Usage:    "Billing and tax configuration",
			BodyPath: "billingInformation",
		},
		&requestflag.Flag[string]{
			Name:     "billing-period",
			Usage:    "Billing period (MONTHLY or ANNUALLY)",
			BodyPath: "billingPeriod",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "charge",
			Usage:    "One-time or recurring charges",
			BodyPath: "charges",
		},
		&requestflag.Flag[string]{
			Name:     "paying-customer-id",
			Usage:    "Paying customer ID for delegated billing",
			BodyPath: "payingCustomerId",
		},
		&requestflag.Flag[string]{
			Name:     "resource-id",
			Usage:    "Resource ID for multi-instance subscriptions",
			BodyPath: "resourceId",
		},
		&requestflag.Flag[string]{
			Name:     "schedule-strategy",
			Usage:    "When to apply subscription changes",
			BodyPath: "scheduleStrategy",
		},
		&requestflag.Flag[any]{
			Name:     "start-date",
			Usage:    "Subscription start date",
			BodyPath: "startDate",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "trial-override-configuration",
			Usage:    "Trial period override settings",
			BodyPath: "trialOverrideConfiguration",
		},
		&requestflag.Flag[int64]{
			Name:     "unit-quantity",
			Usage:    "Unit quantity for per-unit pricing",
			BodyPath: "unitQuantity",
		},
	},
	Action:          handleV1SubscriptionsPreview,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"addon": {
		&requestflag.InnerFlag[string]{
			Name:       "addon.id",
			Usage:      "Addon ID",
			InnerField: "id",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "addon.quantity",
			Usage:      "Number of addon instances",
			InnerField: "quantity",
		},
	},
	"applied-coupon": {
		&requestflag.InnerFlag[string]{
			Name:       "applied-coupon.billing-coupon-id",
			Usage:      "Billing provider coupon ID",
			InnerField: "billingCouponId",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "applied-coupon.configuration",
			Usage:      "Coupon timing configuration",
			InnerField: "configuration",
		},
		&requestflag.InnerFlag[string]{
			Name:       "applied-coupon.coupon-id",
			Usage:      "Stigg coupon ID",
			InnerField: "couponId",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "applied-coupon.discount",
			Usage:      "Ad-hoc discount configuration",
			InnerField: "discount",
		},
		&requestflag.InnerFlag[string]{
			Name:       "applied-coupon.promotion-code",
			Usage:      "Promotion code to apply",
			InnerField: "promotionCode",
		},
	},
	"billable-feature": {
		&requestflag.InnerFlag[string]{
			Name:       "billable-feature.feature-id",
			Usage:      "Feature ID",
			InnerField: "featureId",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "billable-feature.quantity",
			Usage:      "Quantity of feature units",
			InnerField: "quantity",
		},
	},
	"billing-information": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "billing-information.billing-address",
			Usage:      "Billing address",
			InnerField: "billingAddress",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-information.charge-on-behalf-of-account",
			Usage:      "Connected account ID for platform billing",
			InnerField: "chargeOnBehalfOfAccount",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-information.integration-id",
			Usage:      "Billing integration ID",
			InnerField: "integrationId",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "billing-information.invoice-days-until-due",
			Usage:      "Days until invoice is due",
			InnerField: "invoiceDaysUntilDue",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "billing-information.is-backdated",
			Usage:      "Whether subscription is backdated",
			InnerField: "isBackdated",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "billing-information.is-invoice-paid",
			Usage:      "Whether invoice is already paid",
			InnerField: "isInvoicePaid",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "billing-information.metadata",
			Usage:      "Additional billing metadata",
			InnerField: "metadata",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-information.proration-behavior",
			Usage:      "Proration behavior",
			InnerField: "prorationBehavior",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "billing-information.tax-ids",
			Usage:      "Customer tax IDs",
			InnerField: "taxIds",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "billing-information.tax-percentage",
			Usage:      "Tax percentage to apply",
			InnerField: "taxPercentage",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "billing-information.tax-rate-ids",
			Usage:      "Tax rate IDs from billing provider",
			InnerField: "taxRateIds",
		},
	},
	"charge": {
		&requestflag.InnerFlag[string]{
			Name:       "charge.id",
			Usage:      "Charge ID",
			InnerField: "id",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "charge.quantity",
			Usage:      "Charge quantity",
			InnerField: "quantity",
		},
		&requestflag.InnerFlag[string]{
			Name:       "charge.type",
			Usage:      "Charge type",
			InnerField: "type",
		},
	},
	"trial-override-configuration": {
		&requestflag.InnerFlag[bool]{
			Name:       "trial-override-configuration.is-trial",
			Usage:      "Whether to start as trial",
			InnerField: "isTrial",
		},
		&requestflag.InnerFlag[string]{
			Name:       "trial-override-configuration.trial-end-behavior",
			Usage:      "Behavior when trial ends",
			InnerField: "trialEndBehavior",
		},
		&requestflag.InnerFlag[any]{
			Name:       "trial-override-configuration.trial-end-date",
			Usage:      "Trial end date",
			InnerField: "trialEndDate",
		},
	},
})

var v1SubscriptionsProvision = requestflag.WithInnerFlags(cli.Command{
	Name:    "provision",
	Usage:   "Creates a new subscription for an existing customer. When payment is required\nand no payment method exists, returns a checkout URL.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "customer-id",
			Usage:    "Customer ID to provision the subscription for",
			Required: true,
			BodyPath: "customerId",
		},
		&requestflag.Flag[string]{
			Name:     "plan-id",
			Usage:    "Plan ID to provision",
			Required: true,
			BodyPath: "planId",
		},
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "Unique identifier for the subscription",
			BodyPath: "id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "addon",
			BodyPath: "addons",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "applied-coupon",
			Usage:    "Coupon configuration",
			BodyPath: "appliedCoupon",
		},
		&requestflag.Flag[bool]{
			Name:     "await-payment-confirmation",
			Usage:    "Whether to wait for payment confirmation before returning the subscription",
			Default:  true,
			BodyPath: "awaitPaymentConfirmation",
		},
		&requestflag.Flag[any]{
			Name:     "billing-country-code",
			Usage:    "The ISO 3166-1 alpha-2 country code for billing",
			BodyPath: "billingCountryCode",
		},
		&requestflag.Flag[string]{
			Name:     "billing-cycle-anchor",
			Usage:    "Billing cycle anchor behavior for the subscription",
			BodyPath: "billingCycleAnchor",
		},
		&requestflag.Flag[any]{
			Name:     "billing-id",
			Usage:    "External billing system identifier",
			BodyPath: "billingId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "billing-information",
			BodyPath: "billingInformation",
		},
		&requestflag.Flag[string]{
			Name:     "billing-period",
			Usage:    "Billing period (MONTHLY or ANNUALLY)",
			BodyPath: "billingPeriod",
		},
		&requestflag.Flag[any]{
			Name:     "budget",
			BodyPath: "budget",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "charge",
			BodyPath: "charges",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "checkout-options",
			Usage:    "Checkout page configuration for payment collection",
			BodyPath: "checkoutOptions",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "entitlement",
			BodyPath: "entitlements",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Additional metadata for the subscription",
			BodyPath: "metadata",
		},
		&requestflag.Flag[any]{
			Name:     "minimum-spend",
			Usage:    "Minimum spend amount",
			BodyPath: "minimumSpend",
		},
		&requestflag.Flag[any]{
			Name:     "paying-customer-id",
			Usage:    "Optional paying customer ID for split billing scenarios",
			BodyPath: "payingCustomerId",
		},
		&requestflag.Flag[string]{
			Name:     "payment-collection-method",
			Usage:    "How payments should be collected for this subscription",
			Default:  "CHARGE",
			BodyPath: "paymentCollectionMethod",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "price-override",
			BodyPath: "priceOverrides",
		},
		&requestflag.Flag[any]{
			Name:     "resource-id",
			Usage:    "Optional resource ID for multi-instance subscriptions",
			BodyPath: "resourceId",
		},
		&requestflag.Flag[any]{
			Name:     "salesforce-id",
			Usage:    "Salesforce ID",
			BodyPath: "salesforceId",
		},
		&requestflag.Flag[string]{
			Name:     "schedule-strategy",
			Usage:    "Strategy for scheduling subscription changes",
			BodyPath: "scheduleStrategy",
		},
		&requestflag.Flag[any]{
			Name:     "start-date",
			Usage:    "Subscription start date",
			BodyPath: "startDate",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "trial-override-configuration",
			Usage:    "Trial period override settings",
			BodyPath: "trialOverrideConfiguration",
		},
		&requestflag.Flag[int64]{
			Name:     "unit-quantity",
			Usage:    "Unit quantity",
			BodyPath: "unitQuantity",
		},
	},
	Action:          handleV1SubscriptionsProvision,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"addon": {
		&requestflag.InnerFlag[string]{
			Name:       "addon.id",
			Usage:      "Addon ID",
			InnerField: "id",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "addon.quantity",
			Usage:      "Number of addon instances",
			InnerField: "quantity",
		},
	},
	"applied-coupon": {
		&requestflag.InnerFlag[string]{
			Name:       "applied-coupon.billing-coupon-id",
			Usage:      "Billing provider coupon ID",
			InnerField: "billingCouponId",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "applied-coupon.configuration",
			Usage:      "Coupon timing configuration",
			InnerField: "configuration",
		},
		&requestflag.InnerFlag[string]{
			Name:       "applied-coupon.coupon-id",
			Usage:      "Stigg coupon ID",
			InnerField: "couponId",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "applied-coupon.discount",
			Usage:      "Ad-hoc discount configuration",
			InnerField: "discount",
		},
		&requestflag.InnerFlag[string]{
			Name:       "applied-coupon.promotion-code",
			Usage:      "Promotion code to apply",
			InnerField: "promotionCode",
		},
	},
	"billing-information": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "billing-information.billing-address",
			Usage:      "Billing address for the subscription",
			InnerField: "billingAddress",
		},
		&requestflag.InnerFlag[any]{
			Name:       "billing-information.charge-on-behalf-of-account",
			Usage:      "Stripe Connect account to charge on behalf of",
			InnerField: "chargeOnBehalfOfAccount",
		},
		&requestflag.InnerFlag[any]{
			Name:       "billing-information.integration-id",
			Usage:      "Billing integration identifier",
			InnerField: "integrationId",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "billing-information.invoice-days-until-due",
			Usage:      "Number of days until invoice is due",
			InnerField: "invoiceDaysUntilDue",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "billing-information.is-backdated",
			Usage:      "Whether the subscription is backdated",
			InnerField: "isBackdated",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "billing-information.is-invoice-paid",
			Usage:      "Whether the invoice is marked as paid",
			InnerField: "isInvoicePaid",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "billing-information.metadata",
			Usage:      "Additional metadata for the subscription",
			InnerField: "metadata",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-information.proration-behavior",
			Usage:      "How to handle proration for billing changes",
			InnerField: "prorationBehavior",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "billing-information.tax-ids",
			Usage:      "Customer tax identification numbers",
			InnerField: "taxIds",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "billing-information.tax-percentage",
			Usage:      "Tax percentage (0-100)",
			InnerField: "taxPercentage",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "billing-information.tax-rate-ids",
			Usage:      "Tax rate identifiers to apply",
			InnerField: "taxRateIds",
		},
	},
	"budget": {
		&requestflag.InnerFlag[bool]{
			Name:       "budget.has-soft-limit",
			Usage:      "Whether the budget is a soft limit",
			InnerField: "hasSoftLimit",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "budget.limit",
			Usage:      "Maximum spending limit",
			InnerField: "limit",
		},
	},
	"charge": {
		&requestflag.InnerFlag[string]{
			Name:       "charge.id",
			Usage:      "Charge ID",
			InnerField: "id",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "charge.quantity",
			Usage:      "Charge quantity",
			InnerField: "quantity",
		},
		&requestflag.InnerFlag[string]{
			Name:       "charge.type",
			Usage:      "Charge type",
			InnerField: "type",
		},
	},
	"checkout-options": {
		&requestflag.InnerFlag[string]{
			Name:       "checkout-options.cancel-url",
			Usage:      "URL to redirect to if checkout is canceled",
			InnerField: "cancelUrl",
		},
		&requestflag.InnerFlag[string]{
			Name:       "checkout-options.success-url",
			Usage:      "URL to redirect to after successful checkout",
			InnerField: "successUrl",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "checkout-options.allow-promo-codes",
			Usage:      "Allow promotional codes during checkout",
			InnerField: "allowPromoCodes",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "checkout-options.allow-tax-id-collection",
			Usage:      "Allow tax ID collection during checkout",
			InnerField: "allowTaxIdCollection",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "checkout-options.collect-billing-address",
			Usage:      "Collect billing address during checkout",
			InnerField: "collectBillingAddress",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "checkout-options.collect-phone-number",
			Usage:      "Collect phone number during checkout",
			InnerField: "collectPhoneNumber",
		},
		&requestflag.InnerFlag[any]{
			Name:       "checkout-options.reference-id",
			Usage:      "Optional reference ID for the checkout session",
			InnerField: "referenceId",
		},
	},
	"minimum-spend": {
		&requestflag.InnerFlag[float64]{
			Name:       "minimum-spend.amount",
			Usage:      "The price amount",
			InnerField: "amount",
		},
		&requestflag.InnerFlag[string]{
			Name:       "minimum-spend.currency",
			Usage:      "The price currency",
			InnerField: "currency",
		},
	},
	"price-override": {
		&requestflag.InnerFlag[any]{
			Name:       "price-override.addon-id",
			Usage:      "Addon identifier for the price override",
			InnerField: "addonId",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "price-override.amount",
			Usage:      "The price amount",
			InnerField: "amount",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "price-override.base-charge",
			Usage:      "Whether this is a base charge override",
			InnerField: "baseCharge",
		},
		&requestflag.InnerFlag[string]{
			Name:       "price-override.billing-country-code",
			Usage:      "The billing country code of the price",
			InnerField: "billingCountryCode",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "price-override.block-size",
			Usage:      "Block size for pricing",
			InnerField: "blockSize",
		},
		&requestflag.InnerFlag[string]{
			Name:       "price-override.credit-grant-cadence",
			Usage:      `Allowed values: "BEGINNING_OF_BILLING_PERIOD", "MONTHLY".`,
			InnerField: "creditGrantCadence",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "price-override.credit-rate",
			InnerField: "creditRate",
		},
		&requestflag.InnerFlag[string]{
			Name:       "price-override.currency",
			Usage:      "The price currency",
			InnerField: "currency",
		},
		&requestflag.InnerFlag[any]{
			Name:       "price-override.feature-id",
			Usage:      "Feature identifier for the price override",
			InnerField: "featureId",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "price-override.tiers",
			Usage:      "Pricing tiers configuration",
			InnerField: "tiers",
		},
	},
	"trial-override-configuration": {
		&requestflag.InnerFlag[bool]{
			Name:       "trial-override-configuration.is-trial",
			Usage:      "Whether the subscription should start with a trial period",
			InnerField: "isTrial",
		},
		&requestflag.InnerFlag[string]{
			Name:       "trial-override-configuration.trial-end-behavior",
			Usage:      "Behavior when trial ends: CONVERT_TO_PAID or CANCEL_SUBSCRIPTION",
			InnerField: "trialEndBehavior",
		},
		&requestflag.InnerFlag[any]{
			Name:       "trial-override-configuration.trial-end-date",
			Usage:      "Custom trial end date",
			InnerField: "trialEndDate",
		},
	},
})

var v1SubscriptionsTransfer = cli.Command{
	Name:    "transfer",
	Usage:   "Transfers a subscription to a different resource ID. Used for multi-resource\nproducts where subscriptions apply to specific entities like websites or apps.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "destination-resource-id",
			Usage:    "Resource ID to transfer the subscription to",
			Required: true,
			BodyPath: "destinationResourceId",
		},
	},
	Action:          handleV1SubscriptionsTransfer,
	HideHelpCommand: true,
}

func handleV1SubscriptionsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Subscriptions.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:subscriptions retrieve", obj, format, transform)
}

func handleV1SubscriptionsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1SubscriptionUpdateParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Subscriptions.Update(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:subscriptions update", obj, format, transform)
}

func handleV1SubscriptionsList(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1SubscriptionListParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.V1.Subscriptions.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "v1:subscriptions list", obj, format, transform)
	} else {
		iter := client.V1.Subscriptions.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "v1:subscriptions list", iter, format, transform, maxItems)
	}
}

func handleV1SubscriptionsCancel(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1SubscriptionCancelParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Subscriptions.Cancel(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:subscriptions cancel", obj, format, transform)
}

func handleV1SubscriptionsDelegate(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1SubscriptionDelegateParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Subscriptions.Delegate(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:subscriptions delegate", obj, format, transform)
}

func handleV1SubscriptionsImport(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1SubscriptionImportParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Subscriptions.Import(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:subscriptions import", obj, format, transform)
}

func handleV1SubscriptionsMigrate(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1SubscriptionMigrateParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Subscriptions.Migrate(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:subscriptions migrate", obj, format, transform)
}

func handleV1SubscriptionsPreview(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1SubscriptionPreviewParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Subscriptions.Preview(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:subscriptions preview", obj, format, transform)
}

func handleV1SubscriptionsProvision(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1SubscriptionProvisionParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Subscriptions.Provision(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:subscriptions provision", obj, format, transform)
}

func handleV1SubscriptionsTransfer(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1SubscriptionTransferParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Subscriptions.Transfer(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:subscriptions transfer", obj, format, transform)
}
