package meili_test

import (
	"context"
	"testing"
)

func TestSearchEmptyIndex(t *testing.T) {
	client := getCleanTestInstance(t)

	err := client.CreateIndex(context.TODO(), "test", "test", nil)
	if err != nil {
		t.Fatalf("could not create test index: %s", err)
	}

	hits, err := client.Search(context.TODO(), "test", "")
	if err != nil {
		t.Fatalf("could not search test index: %s", err)
	}

	if len(hits) != 0 {
		t.Fatalf("did not get 0 hits, got '%d' hits", len(hits))
	}
}

func TestSearchNoIndex(t *testing.T) {
	client := getCleanTestInstance(t)

	_, err := client.Search(context.TODO(), "test", "")
	if err == nil {
		t.Fatalf("did not get an error on searching a nonexistant index: %s", err)
	}
}
