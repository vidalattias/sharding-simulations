package main

import (
	"fmt"
	"math/rand"
	"time"
)

var g_shards []*Shard
var g_parents map[*Shard]*Shard
var g_root *Shard
var g_childs map[*Shard]([]*Shard)
var g_transactions []*Transaction

var g_available_throughput float64 = 0

var g_start_time float64 = 0

var g_leaf_model bool
var g_width uint
var g_period float64
var g_duration float64
var g_dissemination_rate float64
var g_depth uint
var g_shard_capacity float64
var g_active_shards []*Shard
var g_analyzers []Analyzer
var g_ref_of_message map[uint]uint
var g_validation_time map[uint]float64
var g_scenario Scenarios
var g_scenario_index int

var g_rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func run() {
	import_scenarios()

	for g_scenario_index, scenar := range g_scenario.Scenarios {
		init_scenario(scenar)
		generate_network()
		print_network(g_scenario_index)
		schedule_shard(g_root)

		g_start_time += scenar.Duration

		fmt.Println()
		g_analyzers[g_scenario_index].analyse_txs(g_scenario_index)
	}
}

func main() {

	run()

	fmt.Println("FIN")

	fmt.Println("Len = ", len(g_scenario.Scenarios))

	//g_analyzers[g_scenario_index].PrintTotalThroughput()
}
