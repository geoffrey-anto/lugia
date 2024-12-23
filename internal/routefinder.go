package server

type RouteFinder struct {
	store Store

	Graph       [][][]int
	NodesToName []string
	NameToNodes map[string]int

	NoOfNodes int
}

func NewRouteFinder(nodes []string, edges [][]int, store Store) *RouteFinder {
	rf := new(RouteFinder)

	noOfNodes := len(nodes)

	rf.Graph = make([][][]int, noOfNodes)

	for _, edge := range edges {
		rf.Graph[edge[0]] = append(rf.Graph[edge[0]], []int{edge[1], 1})
		rf.Graph[edge[1]] = append(rf.Graph[edge[1]], []int{edge[0], 1})
	}

	rf.NameToNodes = map[string]int{}

	for i, node := range nodes {
		rf.NameToNodes[node] = i
	}

	rf.NoOfNodes = noOfNodes
	rf.NodesToName = nodes
	rf.store = store

	return rf
}

//
