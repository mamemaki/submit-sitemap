package google

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/searchconsole/v1"
)

type SubmitSitemapParams struct {
	DryRun   bool
	SiteUrl  string
	FeedPath string
}

func SubmitSitemap(ctx context.Context, params *SubmitSitemapParams) error {
	// Load credentials
	var opt option.ClientOption = nil
	credJson := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS_JSON")
	if len(credJson) > 0 {
		opt = option.WithCredentialsJSON([]byte(credJson))
	} else {
		credJson, err := os.ReadFile("./credentials.json")
		if err == nil {
			opt = option.WithCredentialsJSON([]byte(credJson))
		}
	}

	searchconsoleService, err := searchconsole.NewService(ctx, opt)
	if err != nil {
		return err
	}

	// Submit sitemap
	slog.Info("Google: Submit sitemap..")
	slog.Info(fmt.Sprintf("Google: siteUrl: %s", params.SiteUrl))
	slog.Info(fmt.Sprintf("Google: feedpath: %s", params.FeedPath))
	if !params.DryRun {
		sitemapService := searchconsole.NewSitemapsService(searchconsoleService)
		err = sitemapService.Submit(params.SiteUrl, params.FeedPath).Do()
		if err != nil {
			return err
		}

		slog.Info("Google: Submit sitemap succeeded")
	} else {
		slog.Info("Google: Skips submission due to dry-run mode")
	}

	return nil
}
