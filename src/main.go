package main

import (
	"fmt"
	"math/rand"
	"time"

	alavl "github.com/ancientlore/go-avltree"
)

var g_shards []*Shard
var g_parents map[*Shard]*Shard
var g_root *Shard
var g_childs map[*Shard]([]*Shard)
var g_transactions []*Transaction

var g_available_throughput float64 = 0

var g_start_time float64 = 0

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

func init_simulation() {
	g_root = &Shard{
		id:                 new_shard_id(),
		parent:             nil,
		dissemination_rate: g_dissemination_rate,
		next_reference:     random_delay(),
		capacity:           g_shard_capacity,
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

func run() {
	for i, scenar := range g_scenario.Scenarios {
		fmt.Println()
		fmt.Println("\r\tPreparing scenario no.", i, ";", scenar.Lambda)
		g_duration = scenar.Duration
		g_width = scenar.Width
		g_dissemination_rate = scenar.Dissemination_rate
		g_shard_capacity = scenar.Shard_capacity
		g_lambda = scenar.Lambda
		g_period = scenar.Period

		for _, s := range g_shards {
			s.capacity = scenar.Shard_capacity
			s.allocated = 0
		}

		make_schedule()
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
