package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
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
		if tx.time_seen <= time {
			s.proofs_to_process = append(s.proofs_to_process[:i], s.proofs_to_process[i+1:]...)

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

func write_file(name string, content string) {
	f, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(content)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
