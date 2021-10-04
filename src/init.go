package main

import (
	"fmt"

	alavl "github.com/ancientlore/go-avltree"
)

func init_scenario(scenar Scenario) {
	fmt.Println("\r\tPreparing scenario no.", g_scenario_index, ";", scenar.Depth)

	g_duration = scenar.Duration
	g_width = scenar.Width
	g_dissemination_rate = scenar.Dissemination_rate
	g_depth = scenar.Depth
	g_shard_capacity = scenar.Shard_capacity
	g_period = scenar.Period
	g_leaf_model = scenar.LeafModel

	g_analyzers = append(g_analyzers, Analyzer{})

	g_analyzers[g_scenario_index].time_throughput = make(map[uint][]uint)
	g_analyzers[g_scenario_index].messages = make(map[uint]uint)
	g_analyzers[g_scenario_index].references = make(map[uint]uint)
	g_ref_of_message = make(map[uint]uint)
	g_validation_time = make(map[uint]float64)

	shard_counter = 0
	tx_counter = 0
	node_counter = 0

}

func generate_layers(s *Shard, depth uint) *Shard {
	new_shard := new_shard(s, depth)

	if depth < g_depth {
		for i := 0; i < int(g_width); i++ {
			generate_layers(new_shard, depth+1)
		}
	}

	return new_shard
}

func new_shard(s *Shard, depth uint) *Shard {
	var new_shard *Shard
	new_shard = &Shard{
		id:                 new_shard_id(),
		parent:             s,
		dissemination_rate: g_dissemination_rate,
		next_reference:     g_period,
		capacity:           g_shard_capacity,
		depth:              depth,
	}

	new_shard.to_validate = alavl.New(compareTx, alavl.AllowDuplicates)

	g_shards = append(g_shards, new_shard)
	if s != nil {
		s.childs = append(s.childs, new_shard)
	}

	g_analyzers[g_scenario_index].time_throughput[new_shard.id] = []uint{}

	if (g_leaf_model && depth == g_depth) || !g_leaf_model {
		new_shard.is_issuing = true
	}

	return new_shard
}

func generate_network() {
	g_root = generate_layers(nil, 0)

	g_root.to_validate = alavl.New(compareTx, alavl.AllowDuplicates)
}
