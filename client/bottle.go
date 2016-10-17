package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/net/websocket"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// CreateBottlePayload is the bottle create action payload.
type CreateBottlePayload struct {
	Color     string  `form:"color" json:"color" xml:"color"`
	Country   *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	Name      string  `form:"name" json:"name" xml:"name"`
	Region    *string `form:"region,omitempty" json:"region,omitempty" xml:"region,omitempty"`
	Review    *string `form:"review,omitempty" json:"review,omitempty" xml:"review,omitempty"`
	Sweetness *int    `form:"sweetness,omitempty" json:"sweetness,omitempty" xml:"sweetness,omitempty"`
	Varietal  string  `form:"varietal" json:"varietal" xml:"varietal"`
	Vineyard  string  `form:"vineyard" json:"vineyard" xml:"vineyard"`
	Vintage   int     `form:"vintage" json:"vintage" xml:"vintage"`
}

// CreateBottlePath computes a request path to the create action of bottle.
func CreateBottlePath(accountID string) string {
	return fmt.Sprintf("/cellar/accounts/%v/bottles", accountID)
}

// Record new bottle
func (c *Client) CreateBottle(ctx context.Context, path string, payload *CreateBottlePayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateBottleRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateBottleRequest create the request corresponding to the create action endpoint of the bottle resource.
func (c *Client) NewCreateBottleRequest(ctx context.Context, path string, payload *CreateBottlePayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// DeleteBottlePath computes a request path to the delete action of bottle.
func DeleteBottlePath(accountID string, bottleID int) string {
	return fmt.Sprintf("/cellar/accounts/%v/bottles/%v", accountID, bottleID)
}

// DeleteBottle makes a request to the delete action endpoint of the bottle resource
func (c *Client) DeleteBottle(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteBottleRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteBottleRequest create the request corresponding to the delete action endpoint of the bottle resource.
func (c *Client) NewDeleteBottleRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListBottlePath computes a request path to the list action of bottle.
func ListBottlePath(accountID string) string {
	return fmt.Sprintf("/cellar/accounts/%v/bottles", accountID)
}

// List all bottles in account optionally filtering by year
func (c *Client) ListBottle(ctx context.Context, path string, years []int) (*http.Response, error) {
	req, err := c.NewListBottleRequest(ctx, path, years)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBottleRequest create the request corresponding to the list action endpoint of the bottle resource.
func (c *Client) NewListBottleRequest(ctx context.Context, path string, years []int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if years != nil {
		tmp15 := make([]string, len(years))
		for i, e := range years {
			tmp16 := strconv.Itoa(e)
			tmp15[i] = tmp16
		}
		tmp14 := strings.Join(tmp15, ",")
		values.Set("years", tmp14)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// RateBottlePayload is the bottle rate action payload.
type RateBottlePayload struct {
	// Rating of bottle between 1 and 5
	Rating int `form:"rating" json:"rating" xml:"rating"`
}

// RateBottlePath computes a request path to the rate action of bottle.
func RateBottlePath(accountID string, bottleID int) string {
	return fmt.Sprintf("/cellar/accounts/%v/bottles/%v/actions/rate", accountID, bottleID)
}

// RateBottle makes a request to the rate action endpoint of the bottle resource
func (c *Client) RateBottle(ctx context.Context, path string, payload *RateBottlePayload, contentType string) (*http.Response, error) {
	req, err := c.NewRateBottleRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRateBottleRequest create the request corresponding to the rate action endpoint of the bottle resource.
func (c *Client) NewRateBottleRequest(ctx context.Context, path string, payload *RateBottlePayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// ShowBottlePath computes a request path to the show action of bottle.
func ShowBottlePath(accountID string, bottleID int) string {
	return fmt.Sprintf("/cellar/accounts/%v/bottles/%v", accountID, bottleID)
}

// Retrieve bottle with given id
func (c *Client) ShowBottle(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowBottleRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowBottleRequest create the request corresponding to the show action endpoint of the bottle resource.
func (c *Client) NewShowBottleRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateBottlePayload is the bottle update action payload.
type UpdateBottlePayload struct {
	Color     *string `form:"color,omitempty" json:"color,omitempty" xml:"color,omitempty"`
	Country   *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	Name      *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Region    *string `form:"region,omitempty" json:"region,omitempty" xml:"region,omitempty"`
	Review    *string `form:"review,omitempty" json:"review,omitempty" xml:"review,omitempty"`
	Sweetness *int    `form:"sweetness,omitempty" json:"sweetness,omitempty" xml:"sweetness,omitempty"`
	Varietal  *string `form:"varietal,omitempty" json:"varietal,omitempty" xml:"varietal,omitempty"`
	Vineyard  *string `form:"vineyard,omitempty" json:"vineyard,omitempty" xml:"vineyard,omitempty"`
	Vintage   *int    `form:"vintage,omitempty" json:"vintage,omitempty" xml:"vintage,omitempty"`
}

// UpdateBottlePath computes a request path to the update action of bottle.
func UpdateBottlePath(accountID string, bottleID int) string {
	return fmt.Sprintf("/cellar/accounts/%v/bottles/%v", accountID, bottleID)
}

// UpdateBottle makes a request to the update action endpoint of the bottle resource
func (c *Client) UpdateBottle(ctx context.Context, path string, payload *UpdateBottlePayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateBottleRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateBottleRequest create the request corresponding to the update action endpoint of the bottle resource.
func (c *Client) NewUpdateBottleRequest(ctx context.Context, path string, payload *UpdateBottlePayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PATCH", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// WatchBottlePath computes a request path to the watch action of bottle.
func WatchBottlePath(accountID string, bottleID int) string {
	return fmt.Sprintf("/cellar/accounts/%v/bottles/%v/watch", accountID, bottleID)
}

// Retrieve bottle with given id
func (c *Client) WatchBottle(ctx context.Context, path string) (*websocket.Conn, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "ws"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	return websocket.Dial(u.String(), "", u.String())
}
