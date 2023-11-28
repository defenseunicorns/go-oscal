package uuid

import (
	"testing"
)

func TestDeterministicUUIDGeneration(t *testing.T) {
	uuid1 := NewUUIDWithSource("https://raw.githubusercontent.com/GSA/fedramp-automation/93ca0e20ff5e54fc04140613476fba80f08e3c7d/dist/content/rev5/baselines/json/FedRAMP_rev5_HIGH-baseline-resolved-profile_catalog.json")
	uuid2 := "d2afb4c4-2cd8-5305-a6cc-d1bc7b388d0c"
	uuid3 := NewUUIDWithSource("Testing345")

	if uuid1 != uuid2 {
		t.Fatal("UUIDs are not deterministic")
	}

	if uuid2 == uuid3 {
		t.Fatal("UUIDs are not deterministic")
	}
}
