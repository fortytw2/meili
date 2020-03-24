package meili

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// MigrateIndexes adds or removes indexes and settings to ensure
// the requested ones exist
// func (c *Client) MigrateIndexes

// IndexSettings define how Meili handles each index
type IndexSettings struct {
	PrimaryKey     string
	DistinctFields []string
}

// CreateIndex creates a new unique index with the given settings
func (c *Client) CreateIndex(ctx context.Context, name, uid string, is *IndexSettings) error {
	key, err := c.getMasterOrPrivateKey()
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"uid":  uid,
		"name": name,
	}

	if is != nil {
		if is.PrimaryKey != "" {
			data["primaryKey"] = is.PrimaryKey
		}
	}

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	r, err := http.NewRequest(http.MethodPost, c.getRoute(listIndexRoute), bytes.NewReader(body))
	if err != nil {
		return err
	}

	r = r.WithContext(ctx)
	r.Header.Set("X-Meili-API-Key", key)

	resp, err := c.makeRequest(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("meili: could not create index, got an unexpected status code '%d' (wanted a 201)", resp.StatusCode)
	}

	return nil
}

// a ListIndexElement is returned from (c *Client).ListIndexes
type ListIndexElement struct {
	Name       string    `json:"name"`
	UID        string    `json:"uid"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	PrimaryKey string    `json:"primaryKey,omitempty"`
}

const listIndexRoute = "/indexes"

// ListIndexes lists all Indexes within the connected Meili instance
func (c *Client) ListIndexes(ctx context.Context) ([]ListIndexElement, error) {
	key, err := c.getMasterOrPrivateKey()
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest(http.MethodGet, c.getRoute(listIndexRoute), nil)
	if err != nil {
		return nil, err
	}

	r = r.WithContext(ctx)
	r.Header.Set("X-Meili-API-Key", key)

	resp, err := c.makeRequest(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out []ListIndexElement

	dec := json.NewDecoder(resp.Body)
	dec.DisallowUnknownFields()
	err = dec.Decode(&out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

const deleteIndexRoutePrefix = "/indexes/"

func (c *Client) DeleteIndex(ctx context.Context, uid string) error {
	key, err := c.getMasterOrPrivateKey()
	if err != nil {
		return err
	}

	r, err := http.NewRequest(http.MethodDelete, c.getRoute(deleteIndexRoutePrefix+uid), nil)
	if err != nil {
		return err
	}

	r = r.WithContext(ctx)
	r.Header.Set("X-Meili-API-Key", key)

	resp, err := c.makeRequest(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("meili: could not delete index, got an unexpected status code '%d' (wanted a 204)", resp.StatusCode)
	}

	return nil
}

func (c *Client) SetIndexSettings(ctx context.Context, uid string, is *IndexSettings) error {
	return nil
}
