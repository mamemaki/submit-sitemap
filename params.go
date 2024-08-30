package main

import (
	"fmt"
	"net/url"

	"github.com/mamemaki/submit-sitemap/internal/bing"
	"github.com/mamemaki/submit-sitemap/internal/google"
)

// stripUrlPath strip path, query and fragment from url
func stripPathFromUrl(u *url.URL) {
	u.Path = ""
	u.RawQuery = ""
	u.Fragment = ""
}

func CreateGoogleSubmitSitemapParams(cmdOptions *CommandOptions) (*google.SubmitSitemapParams, error) {
	params := google.SubmitSitemapParams{
		DryRun:   cmdOptions.DryRun,
		SiteUrl:  cmdOptions.GoogleSiteUrl,
		FeedPath: cmdOptions.GoogleFeedPath,
	}
	if params.SiteUrl == "" {
		u, err := url.ParseRequestURI(cmdOptions.FeedUrl)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse FeedUrl(%s)", cmdOptions.FeedUrl)
		}
		stripPathFromUrl(u)
		params.SiteUrl = u.String() + "/"
	}
	if params.FeedPath == "" {
		params.FeedPath = cmdOptions.FeedUrl
	}
	return &params, nil
}

func CreateBingSubmitSitemapParams(cmdOptions *CommandOptions) (*bing.SubmitSitemapParams, error) {
	params := bing.SubmitSitemapParams{
		DryRun:   cmdOptions.DryRun,
		SiteUrl:  cmdOptions.BingSiteUrl,
		FeedPath: cmdOptions.BingFeedPath,
	}
	if params.SiteUrl == "" {
		u, err := url.ParseRequestURI(cmdOptions.FeedUrl)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse FeedUrl(%s)", cmdOptions.FeedUrl)
		}
		stripPathFromUrl(u)
		params.SiteUrl = u.String()
	}
	if params.FeedPath == "" {
		params.FeedPath = cmdOptions.FeedUrl
	}
	return &params, nil
}
