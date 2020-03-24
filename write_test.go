package meili_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/fortytw2/meili"
)

func TestAddDocumentToIndex(t *testing.T) {
	client := getCleanTestInstance(t)

	err := client.CreateIndex(context.TODO(), "test", "test", &meili.IndexSettings{
		PrimaryKey: "id",
	})
	if err != nil {
		t.Fatalf("could not create index: %s", err)
	}

	document, _ := json.Marshal(map[string]interface{}{
		"id":   1,
		"name": "potatoes",
	})

	_, err = client.WriteOne(context.TODO(), "test", json.RawMessage(document))
	if err != nil {
		t.Fatalf("could not add documen to index: %s", err)
	}

	// TODO: this should use SynchronousWrite
	time.Sleep(time.Second)

	results, err := client.Search(context.TODO(), "test", "potatoes")
	if err != nil {
		t.Fatalf("could not add documen to index: %s", err)
	}

	if len(results) != 1 {
		t.Fatalf("did not get one result")
	}
}
