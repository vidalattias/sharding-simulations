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
var g_lambda float64
var g_shard_capacity float64
var g_active_shards []*Shard
var g_analyzer Analyzer
var g_ref_of_message map[uint]uint
var g_validation_time map[uint]float64
var g_scenario Scenarios

var g_rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func run() {
	for i, scenar := range g_scenario.Scenarios {
		fmt.Println()
		fmt.Println("\r\tPreparing scenario no.", i, ";", scenar.Lambda)
		g_duration = scenar.Duration
		g_width = scenar.Width
		g_dissemination_rate = scenar.Dissemination_rate
		g_lambda = scenar.Lambda
		g_shard_capacity = scenar.Shard_capacity
		g_period = scenar.Period
		g_leaf_model = scenar.LeafModel

		for _, s := range g_shards {
			s.capacity = scenar.Shard_capacity
			s.allocated = 0
		}

		reset_shards()
		make_schedule()

		print_network(i)

		fmt.Println()
		schedule_shard(g_root)

		g_start_time += scenar.Duration
		fmt.Println()
	}
}

func main() {
	init_simulation()

	run()

	fmt.Println("FIN")

	g_analyzer.analyse_txs()

	g_analyzer.PrintTotalThroughput()
}
