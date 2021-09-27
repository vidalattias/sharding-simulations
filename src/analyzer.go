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

	fmt.Println("Keys : ", keys)

	sort.Ints(keys)

	fmt.Println("Keys : ", keys)

	for _, k := range keys {

		//fmt.Println(k)
		messages = append(messages, analyzer.messages[uint(k)])
		references = append(references, analyzer.references[uint(k)])
		ratio = append(ratio, float64(analyzer.messages[uint(k)])/float64(analyzer.messages[uint(k)]+analyzer.references[uint(k)]))

		ratio_string += fmt.Sprintf("%f\n", float64(analyzer.references[uint(k)])/float64(analyzer.messages[uint(k)]+analyzer.references[uint(k)]))
		messages_string += fmt.Sprintf("%d\n", analyzer.messages[uint(k)])
		total_string += fmt.Sprintf("%d\n", analyzer.messages[uint(k)]+analyzer.references[uint(k)])

	}

	//fmt.Println(messages)
	//fmt.Println(references)
	//fmt.Println(ratio)

	write_file("../total_throughput.txt", total_string)
	write_file("../messages.txt", messages_string)
	write_file("../ratio.txt", ratio_string)
}

func (analyzer Analyzer) analyse_txs() {

	var ret_str strings.Builder

	ratio_pre := 0.

	fmt.Printf("")

	for i, tx := range g_transactions {
		ratio := float64(i) / float64(len(g_transactions))

		if ratio-ratio_pre > 0.001 {
			ratio_pre = ratio
			fmt.Printf("\r%f", ratio)
		}

		val_time := get_validation_time(tx, 1)

		if val_time != -1 {
			ret_str.WriteString(fmt.Sprintf("%f;%f;%d;%t\n", tx.timestamp, val_time-tx.timestamp, tx.issuer.depth, tx.is_proof))
		}
	}

	fmt.Println()

	write_file("../txs.txt", ret_str.String())
}
