// Package sharder splits input JSON into parts.
package sharder

// TODO: make dependency injectable

// ShardItem
type ShardItem[T any] struct {
	ID     int
	Values []T
}

func Shard[T any](list []T, numShards int) ([]ShardItem[T], error) {
	shards := make([]ShardItem[T], numShards)

	for i, item := range list {
		shardIndex := i % numShards
		shards[shardIndex].ID = shardIndex
		shards[shardIndex].Values = append(shards[shardIndex].Values, item)
	}

	return shards, nil
}
