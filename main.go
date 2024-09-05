package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"

	"github.com/mamemaki/submit-sitemap/internal/bing"
	"github.com/mamemaki/submit-sitemap/internal/flagutil"
	"github.com/mamemaki/submit-sitemap/internal/google"
	"github.com/spf13/cobra"
)

var (
	// These three variables are set from ldflags
	version = "(devel)"
	commit  = "none"
	date    = "unknown"
)

func getVersion() string {
	if version == "(devel)" {
		// for `go install` or `go run`
		if buildInfo, ok := debug.ReadBuildInfo(); ok {
			version = buildInfo.Main.Version
		}
	}

	ver := version
	ver += fmt.Sprintf("\nRevision: %s", commit)
	ver += fmt.Sprintf("\nBuilt at %s", date)
	return ver
}

type CommandOptions struct {
	FeedUrl string
	Target  []string
	DryRun  bool
	Verbose bool

	GoogleSiteUrl  string
	GoogleFeedPath string

	BingSiteUrl  string
	BingFeedPath string
}

func initParseFlags(cmdOptions *CommandOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-sitemap",
		Short: "Submit a sitemap to search engines(Google/Bing)",
		Long: `submit-sitemap is a command line tool to submit a sitemap to search engines (Google/Bing).

Google:
  We use the Google Search Console API to submit a sitemap to the site property(URL-prefix or Domain).
	To authenticate your Service Account, please set the contents of the credential JSON to the environment variable(GOOGLE_APPLICATION_CREDENTIALS_JSON).
  If your site is a Domain property, you have to use --google-siteurl option explicitly.

Bing:
  We use the Bing Webmaster API to submit a sitemap to the site.
  To authenticate, please set the API Key to the environment variable(BING_APIKEY).

For more information, please read https://github.com/mamemaki/submit-sitemap

Environment variables:
  GOOGLE_APPLICATION_CREDENTIALS_JSON   The content of Service Account credential JSON for Google Console API
  BING_APIKEY                           The API Key for Bing Webmaster API
`,
		Example: `  submit-sitemap -f=https://example.com/sitemap.xml -t=bing
  submit-sitemap -f=https://example.com/sitemap.xml -t=google,bing
  submit-sitemap -f=https://example.com/sitemap.xml -t=google --google-siteurl=sc-domain:example.com`,
		Version: getVersion(),
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(cmdOptions)
		},
	}

	cmd.Flags().StringVarP(&cmdOptions.FeedUrl, "feedurl", "f", "", "The sitemap URL to submit. e.g. https://example.com/sitemap.xml")
	cmd.MarkFlagRequired("feedurl")
	cmd.Flags().VarP(flagutil.NewChoiceSet([]string{"google", "bing"}, []string{}, &cmdOptions.Target), "target", "t", "The target search engines(google/bing)")
	cmd.MarkFlagRequired("target")
	cmd.Flags().BoolVar(&cmdOptions.DryRun, "dry-run", false, "Output the operations but do not execute anything")
	cmd.Flags().BoolVarP(&cmdOptions.Verbose, "verbose", "", false, "Enable verbose output mode")

	cmd.Flags().StringVar(&cmdOptions.GoogleSiteUrl, "google-siteurl", "", "The URL of the property as defined in Search Console. For example: https://example.com/ (URL-prefix property), or sc-domain:example.com (Domain property). If omitted, we will guess it from feedurl. If your site in Google Search Console is Domain property, you have to use this option explicitly.")
	cmd.Flags().StringVar(&cmdOptions.GoogleFeedPath, "google-feedpath", "", "FeedUrl for Google. If omitted, we will guess it from feedurl. e.g. https://example.com/sitemap.xml")

	cmd.Flags().StringVar(&cmdOptions.BingSiteUrl, "bing-siteurl", "", "Site URL for Bing. If omitted, we will guess it from feedurl. e.g. http://example.com")
	cmd.Flags().StringVar(&cmdOptions.BingFeedPath, "bing-feedpath", "", "FeedUrl for Bing. If omitted, we will guess it from feedurl. e.g. https://example.com/sitemap.xml")

	cmd.Flags().SortFlags = false

	return cmd
}

func Run(cmdOptions *CommandOptions) error {
	ctx := context.Background()

	if cmdOptions.Verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	} else {
		slog.SetLogLoggerLevel(slog.LevelInfo)
	}

	for _, target := range cmdOptions.Target {
		if target == "google" {
			params, err := CreateGoogleSubmitSitemapParams(cmdOptions)
			if err != nil {
				return err
			}
			err = google.SubmitSitemap(ctx, params)
			if err != nil {
				return err
			}
		} else if target == "bing" {
			params, err := CreateBingSubmitSitemapParams(cmdOptions)
			if err != nil {
				return err
			}
			err = bing.SubmitSitemap(ctx, params)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("Unknown target(%s)", target)
		}
	}

	return nil
}

func main() {
	cmdOptions := CommandOptions{}
	cmd := initParseFlags(&cmdOptions)
	if err := cmd.Execute(); err != nil {
		slog.Error("Exit on error", slog.Any("error", err))
		os.Exit(1)
	}
}
