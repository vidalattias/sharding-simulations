package main

import (
	"fmt"
	"math"

	alavl "github.com/ancientlore/go-avltree"
	"gonum.org/v1/gonum/stat/distuv"
)

func reset_shards() {
	for _, s := range g_shards {
		s.is_issuing = false
		s.is_stamping = false
		s.capacity = g_shard_capacity
		s.allocated = 0
	}
}

func make_schedule() {
	// g_available_throughput is only updated in new_shard()
	g_available_throughput = g_shard_capacity
	for g_available_throughput < g_lambda {
		fmt.Println("test -> ", g_available_throughput)
		new_shard()
		fmt.Println("test -> ", g_available_throughput, "\n")
	}

	g_active_shards = []*Shard{}

	var allocated float64 = g_lambda
	var queue = []*Shard{g_root}
	var current *Shard

	for allocated > 0 && len(queue) > 0 {
		current = queue[0]
		queue = queue[1:]

		current.allocated = math.Min(allocated, current.capacity)

		allocated -= current.allocated

		fmt.Println(current.id, " ", current.allocated)

		g_active_shards = append(g_active_shards, current)

		for _, child := range current.childs {
			if child.cardinal <= int(g_width) {
				queue = append(queue, child)
			}
		}
	}

	for _, s := range g_shards {
		if s.allocated > 0 {
			s.is_stamping = true
			if g_leaf_model {
				cnt := 0

				for _, c := range s.childs {
					if c.allocated > 0 {
						cnt++
					}
				}
				if cnt == 0 {
					s.is_issuing = true
				}
			} else {
				s.is_issuing = true
			}
		}
	}
}

func new_shard() {
	var queue = []*Shard{g_root}
	var current *Shard

	for len(queue) > 0 {
		current = queue[0]
		queue = queue[1:]

		if len(current.childs) < int(g_width) {
			var new_shard *Shard = &Shard{
				id:                 new_shard_id(),
				parent:             current,
				dissemination_rate: g_dissemination_rate,
				next_reference:     g_period,
				capacity:           g_shard_capacity,
				depth:              current.depth + 1,
				cardinal:           len(current.childs) + 1,
			}

			fmt.Println("\tlast = ", new_shard.id)

			new_shard.to_validate = alavl.New(compareTx, alavl.AllowDuplicates)

			g_shards = append(g_shards, new_shard)
			current.childs = append(current.childs, new_shard)

			g_available_throughput += new_shard.capacity

			g_analyzer.time_throughput[new_shard.id] = []uint{}

			return
		}

		for _, child := range current.childs {
			queue = append(queue, child)
		}
	}

	return
}

func schedule_shard(s *Shard) {
	for _, c := range s.childs {
		if c.allocated > 0 {
			schedule_shard(c)
		}
	}

	fmt.Printf("\rScheduling shard %d;%d\n", s.id, s.depth)

	exp_rand := distuv.Exponential{Rate: float64(s.allocated)}

	var schedule_time float64 = g_start_time

	//fmt.Println("Scheduling messages")
	for schedule_time < g_start_time+g_duration {
		var next_tx = pick_proof(s, schedule_time)

		if next_tx == nil {
			next_tx = &Transaction{
				issuer:         s,
				id:             new_tx_id(),
				timestamp:      schedule_time,
				time_validated: -1.,
				is_proof:       false,
				validator:      nil,
			}

			g_analyzer.messages[uint(schedule_time)]++

		} else {
			g_analyzer.references[uint(schedule_time)]++
		}

		s.to_validate.Add(next_tx)

		g_transactions = append(g_transactions, next_tx)

		if s == g_root {
			next_tx.time_validated = schedule_time
		}

		g_transactions = append(g_transactions, next_tx)
		schedule_time += exp_rand.Rand()
	}

	//fmt.Println("Scheduling proofs")

	if s.parent != nil {
		var proof_time float64 = g_start_time

		for proof_time < g_start_time+g_duration {
			var new_proof *Transaction = &Transaction{
				issuer:         &Shard{},
				id:             new_tx_id(),
				timestamp:      proof_time,
				time_validated: -1,
				is_proof:       true,
				validator:      nil,
			}

			for s.to_validate.Len() > 0 && s.to_validate.Data()[0].(*Transaction).timestamp < new_proof.timestamp {

				s.to_validate.Data()[0].(*Transaction).validator = new_proof

				s.to_validate.RemoveAt(0)
			}

			s.parent.proofs_to_process = append(s.parent.proofs_to_process, new_proof)
			proof_time += g_period
		}
	}
}
