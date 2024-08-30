package main

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mamemaki/submit-sitemap/internal/bing"
	"github.com/mamemaki/submit-sitemap/internal/google"
)

func TestCreateGoogleSubmitSitemapParams(t *testing.T) {
	testCases := []struct {
		name       string
		cmdOptions *CommandOptions
		params     *google.SubmitSitemapParams
		errMsg     string
	}{
		{
			name:       "Guessing",
			cmdOptions: &CommandOptions{FeedUrl: "https://example.com/sitemap.xml"},
			params:     &google.SubmitSitemapParams{SiteUrl: "https://example.com/", FeedPath: "https://example.com/sitemap.xml"},
		},
		{
			name:       "Guessing - Complicated url",
			cmdOptions: &CommandOptions{FeedUrl: "https://user@example.com/def/sitemap.xml#abc"},
			params:     &google.SubmitSitemapParams{SiteUrl: "https://user@example.com/", FeedPath: "https://user@example.com/def/sitemap.xml#abc"},
		},
		{
			name:       "Guessing - Relative url",
			cmdOptions: &CommandOptions{FeedUrl: "abc/sitemap.xml"},
			errMsg:     "Failed to parse FeedUrl",
		},
		{
			name:       "Not guessing - Domain property",
			cmdOptions: &CommandOptions{FeedUrl: "https://example.com/sitemap.xml", GoogleSiteUrl: "sc-domain:example.com"},
			params:     &google.SubmitSitemapParams{SiteUrl: "sc-domain:example.com", FeedPath: "https://example.com/sitemap.xml"},
		},
		{
			name:       "Not guessing - Url-prefix property",
			cmdOptions: &CommandOptions{FeedUrl: "https://example.com/sitemap.xml", GoogleSiteUrl: "https://example.com/", GoogleFeedPath: "https://example.com/sitemap.xml"},
			params:     &google.SubmitSitemapParams{SiteUrl: "https://example.com/", FeedPath: "https://example.com/sitemap.xml"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			params, err := CreateGoogleSubmitSitemapParams(tc.cmdOptions)
			if tc.errMsg == "" {
				if !cmp.Equal(params, tc.params) {
					t.Errorf("not match")
				}
			} else {
				if !strings.Contains(err.Error(), tc.errMsg) {
					t.Errorf("error message not contains '%s'", tc.errMsg)
				}
			}
		})
	}
}

func TestCreateBingSubmitSitemapParams(t *testing.T) {
	testCases := []struct {
		name       string
		cmdOptions *CommandOptions
		params     *bing.SubmitSitemapParams
		errMsg     string
	}{
		{
			name:       "Guessing",
			cmdOptions: &CommandOptions{FeedUrl: "https://example.com/sitemap.xml"},
			params:     &bing.SubmitSitemapParams{SiteUrl: "https://example.com", FeedPath: "https://example.com/sitemap.xml"},
		},
		{
			name:       "Guessing - Complicated url",
			cmdOptions: &CommandOptions{FeedUrl: "https://user@example.com/def/sitemap.xml#abc"},
			params:     &bing.SubmitSitemapParams{SiteUrl: "https://user@example.com", FeedPath: "https://user@example.com/def/sitemap.xml#abc"},
		},
		{
			name:       "Guessing - Relative url",
			cmdOptions: &CommandOptions{FeedUrl: "abc/sitemap.xml"},
			errMsg:     "Failed to parse FeedUrl",
		},
		{
			name:       "Not guessing",
			cmdOptions: &CommandOptions{FeedUrl: "https://example.com/sitemap.xml", BingSiteUrl: "https://example.com", BingFeedPath: "https://example.com/sitemap.xml"},
			params:     &bing.SubmitSitemapParams{SiteUrl: "https://example.com", FeedPath: "https://example.com/sitemap.xml"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			params, err := CreateBingSubmitSitemapParams(tc.cmdOptions)
			if tc.errMsg == "" {
				if !cmp.Equal(params, tc.params) {
					t.Errorf("not match")
				}
			} else {
				if !strings.Contains(err.Error(), tc.errMsg) {
					t.Errorf("error message not contains '%s'", tc.errMsg)
				}
			}
		})
	}
}
