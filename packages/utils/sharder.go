// Package shareder shards JSON into parts for later processing.
package sharder

import (
	"encoding/json"
)

type Sharder interface {
	Shard(string, int) error
}

type Shard struct {
	ID     int
	Values []string
}

type JsonSharder struct {
}

// Shard func ...
func (j *JsonSharder) Shard(inputJSON string, numShards int) ([]Shard, error) {
	var data []string
	err := json.Unmarshal([]byte(inputJSON), &data)
	if err != nil {
		return nil, err
	}

	shards := make([]Shard, numShards)
	for i := 0; i < len(data); i++ {
		shardIndex := i % numShards
		shards[shardIndex].ID = shardIndex
		shards[shardIndex].Values = append(shards[shardIndex].Values, data[i])
	}

	return shards, nil
}

// New function is constructor for json sharder.
func New() *JsonSharder {
	return &JsonSharder{}
}
