package main

import (
	"math"
	"math/rand"
)

var shard_counter uint = 0
var tx_counter uint = 0
var node_counter uint = 0

func new_shard_id() uint {
	shard_counter++
	return shard_counter
}

func new_tx_id() uint {
	tx_counter++
	return tx_counter
}

func new_node_id() uint {
	node_counter++
	return node_counter
}

func random_delay() float64 {
	var ret = math.Abs(rand.NormFloat64() * float64(g_dissemination_rate) / 100.)
	return ret
}

func pick_proof(s *Shard, time float64) *Transaction {
	for i, tx := range s.proofs_to_process {
		if tx.timestamp+s.dissemination_rate < time {
			s.proofs_to_process = append(s.proofs_to_process[:i], s.proofs_to_process[i+1:]...)

			/*
				for _, c := range tx.tx_validated {
					validate_cone(c, time)
				}
			*/

			return tx
		}
	}
	return nil
}

func get_validation_time(tx *Transaction, depth int) float64 {
	if tx.time_validated == -1 {
		if tx.validator == nil {
			return -1
		} else {
			return get_validation_time(tx.validator, depth+1)
		}
	} else {
		return tx.time_validated
	}
}

type PairTxTime struct {
	time float64
	tx   *Transaction
}

func compareTx(a interface{}, b interface{}) int {
	if a.(*Transaction).timestamp < b.(*Transaction).timestamp {
		return -1
	} else if a.(*Transaction).timestamp > b.(*Transaction).timestamp {
		return 1
	}
	return 0
}

func compareTime(a interface{}, b interface{}) int {
	if a.(float64) < b.(float64) {
		return -1
	} else if a.(float64) > b.(float64) {
		return 1
	}
	return 0
}
