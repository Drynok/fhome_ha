package sharder_test

import (
	"testing"

	"github.com/Drynok/fhome_ha/packages/sharder"
)

func TestShard(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	numShards := 3

	shards, err := sharder.Shard(list, numShards)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expectedShards := []sharder.ShardItem[int]{
		{ID: 0, Values: []int{1, 4}},
		{ID: 1, Values: []int{2, 5}},
		{ID: 2, Values: []int{3}},
	}

	if len(shards) != len(expectedShards) {
		t.Errorf("expected %d shards but got %d", len(expectedShards), len(shards))
	}

	for i := range shards {
		if shards[i].ID != expectedShards[i].ID {
			t.Errorf("expected shard ID %d but got %d", expectedShards[i].ID, shards[i].ID)
		}

		if !equalValues(shards[i].Values, expectedShards[i].Values) {
			t.Errorf("expected shard values %v but got %v", expectedShards[i].Values, shards[i].Values)
		}
	}
}

func equalValues(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
