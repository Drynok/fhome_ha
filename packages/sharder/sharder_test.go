package sharder_test

import (
	"testing"

	"github.com/Drynok/fhome_ha/packages/sharder"
	"github.com/stretchr/testify/suite"
)

type shardTestSuite struct {
	suite.Suite
}

func (s *shardTestSuite) TestShard() {
	list := []int{1, 2, 3, 4, 5}
	numShards := 3

	shards, err := sharder.Shard(list, numShards)
	s.Assert().NoError(err)

	expectedShards := []sharder.ShardItem[int]{
		{ID: 0, Values: []int{1, 4}},
		{ID: 1, Values: []int{2, 5}},
		{ID: 2, Values: []int{3}},
	}

	s.Assert().Equal(len(shards), len(expectedShards))

	for i := range shards {
		s.Assert().Equal(shards[i].ID, expectedShards[i].ID)
		s.Assert().Equal(shards[i].Values, expectedShards[i].Values)
	}
}

func TestSharder(t *testing.T) {
	suite.Run(t, new(shardTestSuite))
}
