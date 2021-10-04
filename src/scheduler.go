package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat/distuv"
)

func schedule_shard(s *Shard) {
	for _, c := range s.childs {
		schedule_shard(c)
	}

	fmt.Println("Scheduling shard", s.id, " depth ", s.depth)

	schedule_rand := distuv.Exponential{Rate: float64(s.capacity)}
	proof_rand := distuv.Exponential{Rate: float64(1 / g_period)}

	exp_normal := distuv.Normal{Mu: float64(s.dissemination_rate), Sigma: float64(s.dissemination_rate) / 10}

	var schedule_time float64 = g_start_time

	for schedule_time < g_start_time+g_duration {
		var next_tx = pick_proof(s, schedule_time)

		if next_tx == nil {
			if s.is_issuing {
				next_tx = &Transaction{
					issuer:         s,
					id:             new_tx_id(),
					timestamp:      schedule_time,
					time_validated: -1.,
					is_proof:       false,
					validator:      nil,
					scenario:       uint(g_scenario_index),
					time_seen:      schedule_time + exp_normal.Rand(),
				}
				g_transactions = append(g_transactions, next_tx)
			}
		}

		if next_tx != nil {
			s.to_validate_slice = append(s.to_validate_slice, next_tx)

			if s == g_root {
				next_tx.time_validated = schedule_time
			}
		}

		schedule_time += schedule_rand.Rand()
	}

	if s.parent != nil {
		var proof_time float64 = g_start_time

		for proof_time < g_start_time+g_duration {
			var new_proof *Transaction = &Transaction{
				issuer:         s,
				id:             new_tx_id(),
				timestamp:      proof_time,
				time_seen:      proof_time + exp_normal.Rand(),
				time_validated: -1,
				is_proof:       true,
				validator:      nil,
				scenario:       uint(g_scenario_index),
			}

			g_transactions = append(g_transactions, new_proof)

			var count_chosen uint = 0

			for len(s.to_validate_slice) > 0 && s.to_validate_slice[0].time_seen < new_proof.timestamp {

				s.to_validate_slice[0].validator = new_proof
				count_chosen++

				s.to_validate_slice = s.to_validate_slice[1:]
			}

			if count_chosen > 0 {
				s.parent.proofs_to_process = append(s.parent.proofs_to_process, new_proof)
			}
			proof_time += proof_rand.Rand()
		}
	}
}
