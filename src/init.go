package main

import alavl "github.com/ancientlore/go-avltree"

func init_simulation() {
	g_root = &Shard{
		id:                 new_shard_id(),
		parent:             nil,
		dissemination_rate: g_dissemination_rate,
		next_reference:     random_delay(),
		depth:              0,
	}

	g_analyzer.time_throughput = make(map[uint][]uint)
	g_analyzer.messages = make(map[uint]uint)
	g_analyzer.references = make(map[uint]uint)
	g_ref_of_message = make(map[uint]uint)
	g_validation_time = make(map[uint]float64)
	g_shards = append(g_shards, g_root)
	g_root.to_validate = alavl.New(compareTx, alavl.AllowDuplicates)

	import_scenarios()
}
