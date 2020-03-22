package meili_test

import (
	"context"
	"os"
	"testing"

	"github.com/fortytw2/meili"
)

func getCleanTestInstance(t *testing.T) *meili.Client {
	addr, ok := os.LookupEnv("MEILI_ADDR")
	if !ok {
		t.Fatal("cannot find meili instance, make sure MEILI_ADDR is set")
	}

	client, err := meili.NewClient(addr, meili.WithNoKeys())
	if err != nil {
		t.Fatalf("could not instantiate meili client: %s", err)
	}

	err = client.WipeForTests()
	if err != nil {
		t.Fatalf("could not wipe meili instance for tests: %s", err)
	}

	return client
}

func TestListIndexEmpty(t *testing.T) {
	client := getCleanTestInstance(t)

	indexes, err := client.ListIndexes(context.TODO())
	if err != nil {
		t.Fatalf("could not list indexes: %s", err)
	}

	if len(indexes) != 0 {
		t.Fatal("got more than one index when expecting none")
	}
}

func TestCreateAndListIndex(t *testing.T) {
	client := getCleanTestInstance(t)

	err := client.CreateIndex(context.TODO(), "test", "test", nil)
	if err != nil {
		t.Fatalf("could not create index: %s", err)
	}

	indexes, err := client.ListIndexes(context.TODO())
	if err != nil {
		t.Fatalf("could not list indexes: %s", err)
	}

	if len(indexes) != 1 {
		t.Fatal("did not get exactly one index")
	}
}
