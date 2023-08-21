package sharder_test

import (
	"reflect"
	"testing"
)

func TestShardJSON(t *testing.T) {
	inputJSON := `["value1", "value2", "value3", "value4", "value5"]`
	numShards := 3

	expectedShards := []Shard{
		{ID: 0, Values: []string{"value1", "value4"}},
		{ID: 1, Values: []string{"value2", "value5"}},
		{ID: 2, Values: []string{"value3"}},
	}

	shards, err := ShardJSON(inputJSON, numShards)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(shards) != len(expectedShards) {
		t.Errorf("Expected %d shards, but got %d", len(expectedShards), len(shards))
	}

	for i := range shards {
		if !reflect.DeepEqual(shards[i], expectedShards[i]) {
			t.Errorf("Expected shard %+v, but got %+v", expectedShards[i], shards[i])
		}
	}
}
