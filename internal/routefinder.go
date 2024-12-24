package server

type Position struct {
	Position int
	End      int
}

type Vertex struct {
	Index int
	Cost  float64
}

type RouteFinder struct {
	store Store

	Graph       [][]Vertex
	NodesToName []string
	NameToNodes map[string]int

	Nodes     [][]string
	Positions map[string]Position

	NoOfNodes int

	ShortestPaths [][]int
}

func NewRouteFinder(nodes []string, edges [][]int, store Store) *RouteFinder {
	rf := new(RouteFinder)

	noOfNodes := len(nodes)

	rf.Graph = make([][]Vertex, noOfNodes)
	rf.Nodes = make([][]string, noOfNodes)

	for _, edge := range edges {
		rf.Graph[edge[0]] = append(rf.Graph[edge[0]], Vertex{edge[1], 1.0})
		rf.Graph[edge[1]] = append(rf.Graph[edge[1]], Vertex{edge[0], 1.0})
	}

	rf.NameToNodes = map[string]int{}

	for i, node := range nodes {
		rf.NameToNodes[node] = i
		rf.Nodes[i] = []string{}
	}

	rf.NoOfNodes = noOfNodes
	rf.NodesToName = nodes
	rf.store = store
	rf.Positions = make(map[string]Position)

	return rf
}

func (rf *RouteFinder) Add(id string, start, end int) (string, bool) {
	if _, ok := rf.Positions[id]; ok {
		return "User already on route", false
	}

	rf.Positions[id] = Position{
		Position: start,
		End:      end,
	}

	rf.Nodes[start] = append(rf.Nodes[start], id)

	return "User add on route", true
}

func (rf *RouteFinder) Update(id string, next int) (string, bool) {
	if _, ok := rf.Positions[id]; !ok {
		return "User not on route", false
	}

	currentPosition := rf.Positions[id]

	if currentPosition.End == next {
		rf.Nodes[currentPosition.Position] = DeleteElement(rf.Nodes[currentPosition.Position], id)

		delete(rf.Positions, id)

		return "User reached the destination", false
	}

	rf.Nodes[currentPosition.Position] = DeleteElement(rf.Nodes[currentPosition.Position], id)

	rf.Nodes[next] = append(rf.Nodes[next], id)

	currentPosition.Position = next

	rf.Positions[id] = currentPosition

	return "Updated position", true
}

func (rf *RouteFinder) End(id string) (string, bool) {
	if _, ok := rf.Positions[id]; !ok {
		return "User not on route", false
	}

	currentPosition := rf.Positions[id]

	rf.Nodes[currentPosition.Position] = DeleteElement(rf.Nodes[currentPosition.Position], id)

	delete(rf.Positions, id)

	return "User reached the destination", false
}

// func (rf *RouteFinder) GetPath(id string) ([]string, bool) {
// 	if _, ok := rf.Positions[id]; !ok {
// 		return []string{}, false
// 	}

// 	currentPosition := rf.Positions[id]

// 	//

// }

// func (rf *RouteFinder) Search() (bool, int, float64)

func DeleteElement(slice []string, value string) []string {
	for i, v := range slice {
		if v == value {
			// Remove the element at index i
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
