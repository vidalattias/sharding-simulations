package main

import (
	"fmt"

	alavl "github.com/ancientlore/go-avltree"
)

type Shard struct {
	id                  uint
	childs              [](*Shard)
	parent              *Shard
	participating_nodes [](*Shard)
	dissemination_rate  float64
	transactions        [](*Transaction)
	referencer          *Shard
	next_reference      float64
	to_validate         *alavl.Tree
	capacity            float64
	allocated           float64
	proofs_to_process   []*Transaction
	depth               uint
	valid               bool
}

func print_shard(s *Shard, depth uint) {
	for i := 0; uint(i) < depth; i++ {
		fmt.Printf("\t")
	}

	//fmt.Println(s.id, ";", s.allocated)

	for _, c := range s.childs {
		print_shard(c, depth+1)
	}
}
