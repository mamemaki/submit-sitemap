package bing

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

type SubmitSitemapParams struct {
	DryRun   bool
	SiteUrl  string
	FeedPath string
}

func SubmitSitemap(ctx context.Context, params *SubmitSitemapParams) error {
	// Load API Key
	apikey := os.Getenv("BING_APIKEY")
	if len(apikey) == 0 {
		return fmt.Errorf("Missing API Key")
	}

	// Submit sitemap
	slog.Info("Bing: Submit sitemap..")
	slog.Info(fmt.Sprintf("Bing: siteUrl: %s", params.SiteUrl))
	slog.Info(fmt.Sprintf("Bing: feedpath: %s", params.FeedPath))

	url := fmt.Sprintf("https://ssl.bing.com/webmaster/api.svc/json/SubmitFeed?apikey=%s", apikey)
	body := fmt.Sprintf(`{"siteUrl":"%s", "feedUrl":"%s"}`, params.SiteUrl, params.FeedPath)
	slog.Debug(fmt.Sprintf("Bing: request: %s", body))
	if !params.DryRun {
		buf := bytes.NewBufferString(body)
		res, err := http.Post(url, "application/json; charset=utf-8", buf)
		if err != nil {
			return fmt.Errorf("failed to submit feed: %w", err)
		}
		defer res.Body.Close()

		if res.StatusCode == http.StatusOK {
			bodyBytes, err := io.ReadAll(res.Body)
			if err == nil {
				bodyString := string(bodyBytes)
				slog.Debug(fmt.Sprintf("Bing: response: %s", bodyString))
			}
		} else {
			return fmt.Errorf("failed to submit feed: %d", res.StatusCode)
		}

		slog.Info("Bing: Submit sitemap succeeded")
	} else {
		slog.Info("Bing: Skips submission due to dry-run mode")
	}

	return nil
}
