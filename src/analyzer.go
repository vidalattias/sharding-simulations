package main

import (
	"fmt"
	"sort"
	"strings"
)

type Analyzer struct {
	time_throughput map[uint][]uint
	messages        map[uint]uint
	references      map[uint]uint
}

func (analyzer Analyzer) PrintTotalThroughput() {
	var messages []uint
	var references []uint
	var ratio []float64

	var keys []int

	for k := range analyzer.messages {
		keys = append(keys, int(k))
	}

	var ratio_string = ""
	var messages_string = ""
	var total_string = ""

	sort.Ints(keys)

	for _, k := range keys {

		//fmt.Println(k)
		messages = append(messages, analyzer.messages[uint(k)])
		references = append(references, analyzer.references[uint(k)])
		ratio = append(ratio, float64(analyzer.messages[uint(k)])/float64(analyzer.messages[uint(k)]+analyzer.references[uint(k)]))

		ratio_string += fmt.Sprintf("%f\n", float64(analyzer.references[uint(k)])/float64(analyzer.messages[uint(k)]+analyzer.references[uint(k)]))
		messages_string += fmt.Sprintf("%d\n", analyzer.messages[uint(k)])
		total_string += fmt.Sprintf("%d\n", analyzer.messages[uint(k)]+analyzer.references[uint(k)])

	}

	write_file("data/messages.txt", messages_string)
	write_file("data/total_throughput.txt", total_string)
	write_file("data/ratio.txt", ratio_string)
}

func (analyzer Analyzer) analyse_txs(i int) {
	var ret_str strings.Builder

	for _, tx := range g_transactions {

		val_time := get_validation_time(tx, 1)

		if val_time != -1 {
			ret_str.WriteString(fmt.Sprintf("%d;%f;%f;%d;%t\n", tx.id, tx.timestamp, val_time-tx.timestamp, tx.issuer.depth, tx.is_proof))
		} else {
			ret_str.WriteString(fmt.Sprintf("%d;%f;NULL;%d;%t\n", tx.id, tx.timestamp, tx.issuer.depth, tx.is_proof))
		}
	}

	fmt.Println()

	write_file(fmt.Sprintf("data/txs_%d.txt", i), ret_str.String())
}
