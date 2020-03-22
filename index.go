package meili

import (
	"context"
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

// ListIndexes lists all Indexes within the connected Meili instance
func (c *Client) ListIndexes(ctx context.Context) ([]ListIndexElement, error) {
	key, err := c.getMasterOrPrivateKey()
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest(http.MethodPost, c.address, nil)
	if err != nil {
		return nil, err
	}

	r.Header.Set("X-Meili-API-Key", key)

	return nil, nil
}

func (c *Client) SetIndexSettings(ctx context.Context, uid string, is *IndexSettings) error {
	return nil
}
