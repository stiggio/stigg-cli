// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/stainless-sdks/stigg-cli/internal/autocomplete"
	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
)

var (
	Command *cli.Command
)

func init() {
	Command = &cli.Command{
		Name:    "stigg",
		Usage:   "CLI for the stigg API",
		Suggest: true,
		Version: Version,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
			&cli.StringFlag{
				Name:        "base-url",
				DefaultText: "url",
				Usage:       "Override the base URL for API requests",
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "The format for displaying response data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "format-error",
				Usage: "The format for displaying error data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "transform",
				Usage: "The GJSON transformation for data output.",
			},
			&cli.StringFlag{
				Name:  "transform-error",
				Usage: "The GJSON transformation for errors.",
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "v1:customers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1CustomersRetrieve,
					&v1CustomersUpdate,
					&v1CustomersList,
					&v1CustomersArchive,
					&v1CustomersImport,
					&v1CustomersListResources,
					&v1CustomersProvision,
					&v1CustomersUnarchive,
				},
			},
			{
				Name:     "v1:customers:payment-method",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1CustomersPaymentMethodAttach,
					&v1CustomersPaymentMethodDetach,
				},
			},
			{
				Name:     "v1:customers:promotional-entitlements",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1CustomersPromotionalEntitlementsCreate,
					&v1CustomersPromotionalEntitlementsList,
					&v1CustomersPromotionalEntitlementsRevoke,
				},
			},
			{
				Name:     "v1:subscriptions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1SubscriptionsRetrieve,
					&v1SubscriptionsUpdate,
					&v1SubscriptionsList,
					&v1SubscriptionsCancel,
					&v1SubscriptionsDelegate,
					&v1SubscriptionsImport,
					&v1SubscriptionsMigrate,
					&v1SubscriptionsPreview,
					&v1SubscriptionsProvision,
					&v1SubscriptionsTransfer,
				},
			},
			{
				Name:     "v1:subscriptions:future-update",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1SubscriptionsFutureUpdateCancelPendingPayment,
					&v1SubscriptionsFutureUpdateCancelSchedule,
				},
			},
			{
				Name:     "v1:subscriptions:usage",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1SubscriptionsUsageChargeUsage,
					&v1SubscriptionsUsageSync,
				},
			},
			{
				Name:     "v1:subscriptions:invoice",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1SubscriptionsInvoiceMarkAsPaid,
				},
			},
			{
				Name:     "v1:coupons",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1CouponsCreate,
					&v1CouponsRetrieve,
					&v1CouponsList,
					&v1CouponsArchiveCoupon,
					&v1CouponsUpdateCoupon,
				},
			},
			{
				Name:     "v1:events",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1EventsReport,
				},
			},
			{
				Name:     "v1:events:features",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1EventsFeaturesArchiveFeature,
					&v1EventsFeaturesCreateFeature,
					&v1EventsFeaturesListFeatures,
					&v1EventsFeaturesRetrieveFeature,
					&v1EventsFeaturesUnarchiveFeature,
					&v1EventsFeaturesUpdateFeature,
				},
			},
			{
				Name:     "v1:events:addons",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1EventsAddonsArchiveAddon,
					&v1EventsAddonsCreateAddon,
					&v1EventsAddonsListAddons,
					&v1EventsAddonsPublishAddon,
					&v1EventsAddonsRetrieveAddon,
					&v1EventsAddonsSetPricing,
					&v1EventsAddonsUpdateAddon,
				},
			},
			{
				Name:     "v1:events:addons:draft",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1EventsAddonsDraftCreateAddonDraft,
					&v1EventsAddonsDraftRemoveAddonDraft,
				},
			},
			{
				Name:     "v1:events:addons:entitlements",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1EventsAddonsEntitlementsCreate,
					&v1EventsAddonsEntitlementsUpdate,
					&v1EventsAddonsEntitlementsList,
					&v1EventsAddonsEntitlementsDelete,
				},
			},
			{
				Name:     "v1:events:plans",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1EventsPlansCreate,
					&v1EventsPlansRetrieve,
					&v1EventsPlansUpdate,
					&v1EventsPlansList,
					&v1EventsPlansArchive,
					&v1EventsPlansPublish,
					&v1EventsPlansSetPricing,
				},
			},
			{
				Name:     "v1:events:plans:draft",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1EventsPlansDraftCreate,
					&v1EventsPlansDraftRemove,
				},
			},
			{
				Name:     "v1:events:plans:entitlements",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1EventsPlansEntitlementsCreate,
					&v1EventsPlansEntitlementsUpdate,
					&v1EventsPlansEntitlementsList,
					&v1EventsPlansEntitlementsDelete,
				},
			},
			{
				Name:     "v1:usage",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1UsageHistory,
					&v1UsageReport,
				},
			},
			{
				Name:     "v1:products",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1ProductsArchiveProduct,
					&v1ProductsCreateProduct,
					&v1ProductsDuplicateProduct,
					&v1ProductsListProducts,
					&v1ProductsUnarchiveProduct,
					&v1ProductsUpdateProduct,
				},
			},
			{
				Name:            "@manpages",
				Usage:           "Generate documentation for 'man'",
				UsageText:       "stigg @manpages [-o stigg.1] [--gzip]",
				Hidden:          true,
				Action:          generateManpages,
				HideHelpCommand: true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "write manpages to the given folder",
						Value:   "man",
					},
					&cli.BoolFlag{
						Name:    "gzip",
						Aliases: []string{"z"},
						Usage:   "output gzipped manpage files to .gz",
						Value:   true,
					},
					&cli.BoolFlag{
						Name:    "text",
						Aliases: []string{"z"},
						Usage:   "output uncompressed text files",
						Value:   false,
					},
				},
			},
			{
				Name:            "__complete",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.ExecuteShellCompletion,
			},
			{
				Name:            "@completion",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.OutputCompletionScript,
			},
		},
		HideHelpCommand: true,
	}
}

func generateManpages(ctx context.Context, c *cli.Command) error {
	manpage, err := docs.ToManWithSection(Command, 1)
	if err != nil {
		return err
	}
	dir := c.String("output")
	err = os.MkdirAll(filepath.Join(dir, "man1"), 0755)
	if err != nil {
		// handle error
	}
	if c.Bool("text") {
		file, err := os.Create(filepath.Join(dir, "man1", "stigg.1"))
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(manpage); err != nil {
			return err
		}
	}
	if c.Bool("gzip") {
		file, err := os.Create(filepath.Join(dir, "man1", "stigg.1.gz"))
		if err != nil {
			return err
		}
		defer file.Close()
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		_, err = gzWriter.Write([]byte(manpage))
		if err != nil {
			return err
		}
	}
	fmt.Printf("Wrote manpages to %s\n", dir)
	return nil
}
