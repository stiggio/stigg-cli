// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/stainless-sdks/stigg-cli/internal/autocomplete"
	"github.com/stainless-sdks/stigg-cli/internal/requestflag"
	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
)

var (
	Command            *cli.Command
	CommandErrorBuffer bytes.Buffer
)

func init() {
	Command = &cli.Command{
		Name:      "stigg",
		Usage:     "CLI for the stigg API",
		Suggest:   true,
		Version:   Version,
		ErrWriter: &CommandErrorBuffer,
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
			&requestflag.Flag[string]{
				Name:    "api-key",
				Sources: cli.EnvVars("STIGG_API_KEY"),
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
				Name:     "v1:events:credits",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1EventsCreditsGetAutoRecharge,
					&v1EventsCreditsGetUsage,
					&v1EventsCreditsListLedger,
				},
			},
			{
				Name:     "v1:events:credits:grants",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1EventsCreditsGrantsCreate,
					&v1EventsCreditsGrantsList,
					&v1EventsCreditsGrantsVoid,
				},
			},
			{
				Name:     "v1:features",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1FeaturesArchiveFeature,
					&v1FeaturesCreateFeature,
					&v1FeaturesListFeatures,
					&v1FeaturesRetrieveFeature,
					&v1FeaturesUnarchiveFeature,
					&v1FeaturesUpdateFeature,
				},
			},
			{
				Name:     "v1:addons",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1AddonsCreate,
					&v1AddonsRetrieve,
					&v1AddonsUpdate,
					&v1AddonsList,
					&v1AddonsArchive,
					&v1AddonsCreateDraft,
					&v1AddonsPublish,
					&v1AddonsRemoveDraft,
					&v1AddonsSetPricing,
				},
			},
			{
				Name:     "v1:addons:entitlements",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1AddonsEntitlementsCreate,
					&v1AddonsEntitlementsUpdate,
					&v1AddonsEntitlementsList,
					&v1AddonsEntitlementsDelete,
				},
			},
			{
				Name:     "v1:plans",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1PlansCreate,
					&v1PlansRetrieve,
					&v1PlansUpdate,
					&v1PlansList,
					&v1PlansArchive,
					&v1PlansCreateDraft,
					&v1PlansPublish,
					&v1PlansRemoveDraft,
					&v1PlansSetPricing,
				},
			},
			{
				Name:     "v1:plans:entitlements",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1PlansEntitlementsCreate,
					&v1PlansEntitlementsUpdate,
					&v1PlansEntitlementsList,
					&v1PlansEntitlementsDelete,
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
