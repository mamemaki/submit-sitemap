# submit-sitemap
submit-sitemap is a command line tool to submit a sitemap to search engines (Google/Bing).

## Usage
To submit a sitemap to Google:
```
submit-sitemap -f=https://example.com/sitemap.xml -t=google
```

To submit a sitemap to Google and Bing:
```
submit-sitemap -f=https://example.com/sitemap.xml -t=google,bing
```

## Command Line Options

The following are submit-sitemap's command-line options:

```
Usage:
  submit-sitemap [flags]

Flags:
  -f, --feedurl string           The sitemap URL to submit. e.g. https://example.com/sitemap.xml
  -t, --target strings           The target search engines(google/bing) (default [])
      --dry-run                  Output the operations but do not execute anything
      --verbose                  Enable verbose output mode
      --google-siteurl string    The URL of the property as defined in Search Console. For example: https://example.com/ (URL-prefix property), or sc-domain:example.com (Domain property). If omitted, we will guess it from feedurl. If your site in Google Search Console is Domain property, you have to use this option explicitly.
      --google-feedpath string   FeedUrl for Google. If omitted, we will guess it from feedurl. e.g. https://example.com/sitemap.xml
      --bing-siteurl string      Site URL for Bing. If omitted, we will guess it from feedurl. e.g. http://example.com
      --bing-feedpath string     FeedUrl for Bing. If omitted, we will guess it from feedurl. e.g. https://example.com/sitemap.xml
  -h, --help                     help for submit-sitemap
  -v, --version                  version for submit-sitemap
```

## Google
We use the Google Search Console API to submit a sitemap to the site property(URL-prefix or Domain).

### Prerequirements

- Your Site property(either URL prefix or domain) on Google Search Console
- GCP Project to be used for sitemap submission

### Authentication
We use Service Account authentication. other authentication method is not supported.

To authenticate your Service Account when submit a sitemap, set the contents of the credential JSON to the environment variable(`GOOGLE_APPLICATION_CREDENTIALS_JSON`).

### How to create a Service Account on your GCP project
1. Open [Service account](https://search.google.com/search-console/users) in your GCP project
1. Create a Service account and save the credential JSON file\
   All default. no roles, no permissions.

### How to enable "Search Console API" on your GCP project
1. Open [Search Console API](https://console.cloud.google.com/apis/library/searchconsole.googleapis.com) in your GCP project
1. Click "Enable" button

### How to link Service Account and Site property on your Google Search Console
1. Open [Users and privileges](https://search.google.com/search-console/users) in Google Search Console
1. Add your Service Account email as a user with "full" privileges.

### If your site is a Domain property
If your site on Google Search Console is Domain property, you have to use the `--google-siteurl` option explicitly as shown below.

To submit a sitemap to Google that uses Domain property:
```
submit-sitemap -f=https://example.com/sitemap.xml -t=google --google-siteurl=sc-domain:example.com
```

## Bing
We use the Bing Webmaster API to submit a sitemap to the site on Bing.

To authenticate, set the API Key to the environment variable(`BING_APIKEY`).

### How to generate API Key
Please refer to [Getting Access to the Bing Webmaster Tools API](https://learn.microsoft.com/en-us/bingwebmaster/getting-access#using-api-key).

## FAQ

### Why not use sitemap ping?
Google and Bing have stopped supporting sitemap ping due to spam issues([ref](https://developers.google.com/search/blog/2023/06/sitemaps-lastmod-ping), [ref2](https://blogs.bing.com/webmaster/may-2022/Spring-cleaning-Removed-Bing-anonymous-sitemap-submission)). Therefore, if we want to submit a sitemap, we must use an authenticated API.

### Why not use Indexing API or IndexNow? what is the difference between them?
[Indexing API](https://developers.google.com/search/apis/indexing-api/v3/reference/indexing/rest/v3/urlNotifications/publish) and [IndexNow](https://www.indexnow.org/) are per-URL index request APIs. These APIs and sitemap submissions have different purposes as follows.

| | Request unit | Indexing speed | Purpose |
| ------------- | ------------- | ------------- | ------------- |
| **Indexing API and IndexNow**  | URL | in few mins | Request indexing of specific URLs as soon as possible
| **Submit a sitemap**  | Sitemap | in few days | Let search engines know that the site has been updated

## Contributing Guidelines

Contributions are welcome via pull requests. Please see [here](docs/CONTRIBUTING.md) for more information.
