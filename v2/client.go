package surflinef

import (
	"fmt"
	"net/http"
	"net/url"
)

// ConditionsBaseURL is the base URL for the regional conditions service
const ConditionsBaseURL = "https://services.surfline.com/kbyg/regions/forecasts/conditions"

// TidesBaseURL is the base URL for the spot tides service
const TidesBaseURL = "https://services.surfline.com/kbyg/spots/forecasts/tides"

// TaxonomyBaseURL is the base URL for the taxonomy service
const TaxonomyBaseURL = "https://services.surfline.com/taxonomy"

// WaveBaseURL is the base URL for the wave/swell service
const WaveBaseURL = "https://services.surfline.com/kbyg/spots/forecasts/wave"

const LoginBaseURL = "https://services.surfline.com/trusted/token"

const defaultUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"

// Client is the SurflineF HTTP Client.
type Client struct {
	BaseURL    *url.URL
	HttpClient *http.Client // If nil, uses http.DefaultClient
	UserAgent  string       // If empty, uses a browser-like default
}

// FullURL formats the query string and Client BaseUrl
func (c *Client) FullURL(qs string) string {
	return fmt.Sprintf("%s?%s", c.BaseURL, qs)
}

func (c *Client) httpClient() *http.Client {
	if c.HttpClient != nil {
		return c.HttpClient
	}
	return http.DefaultClient
}

func (c *Client) userAgent() string {
	if c.UserAgent != "" {
		return c.UserAgent
	}
	return defaultUserAgent
}

// Get performs an HTTP GET with browser-like headers
func (c *Client) Get(u *url.URL) (*http.Response, error) {
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.userAgent())
	req.Header.Set("Accept", "application/json")
	return c.httpClient().Do(req)
}
