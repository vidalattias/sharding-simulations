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
	to_validate_slice   []*Transaction
	capacity            float64
	proofs_to_process   []*Transaction
	depth               uint
	is_issuing          bool
	active_childs       int
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

func print_network(id int) {
	ret_str := "digraph{"

	ret_str += g_root.print_shard()

	ret_str += "}"

	write_file(fmt.Sprintf("data/scenario_%d.dot", id), ret_str)
}

func (s Shard) print_shard() string {
	ret_str := ""

	if s.is_issuing {
		ret_str = fmt.Sprintf("%d[color=lightgreen style=filled label=\"%d\\n%d\"]", s.id, s.depth, int(s.capacity))
	} else {
		ret_str = fmt.Sprintf("%d[color=yellow style=filled label=\"%d\\n%d\"]", s.id, s.depth, int(s.capacity))
	}

	for _, c := range s.childs {
		ret_str += fmt.Sprintf("%d -> %d\n", s.id, c.id)
		ret_str += c.print_shard()
	}

	return ret_str
}
