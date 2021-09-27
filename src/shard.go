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
	is_issuing          bool
	is_stamping         bool
	active_childs       int
	cardinal            int
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

	write_file(fmt.Sprintf("scenario_%d.dot", id), ret_str)
}

func (s Shard) print_shard() string {
	ret_str := ""

	fmt.Println(s.id, " ", s.allocated)

	if s.is_issuing {
		ret_str = fmt.Sprintf("%d[color=lightgreen style=filled label=\"%d\\n%d\"]", s.id, s.id, int(s.allocated))
	} else if s.is_stamping {
		ret_str = fmt.Sprintf("%d[color=yellow style=filled label=\"%d\\n%d\"]", s.id, s.id, int(s.allocated))
	} else {
		ret_str = fmt.Sprintf("%d[color=red style=filled label=\"%d\\n%d\"]", s.id, s.id, int(s.allocated))
	}

	for _, c := range s.childs {
		ret_str += fmt.Sprintf("%d -> %d\n", s.id, c.id)
		ret_str += c.print_shard()
	}

	return ret_str
}
