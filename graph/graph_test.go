package graph

import (
	"testing"
)

func TestMain(t *testing.T) {
	g := create_example_lgraph()
	g.Print()
	// println("Bfs.............")
	// bfs_v := g.Bfs()
	// for i := 0; i < len(bfs_v); i++ {
	// 	println(bfs_v[i].Data)
	// }
	// bfs_res := []string{"A", "B", "C", "E", "F", "D", "G"}
}

func create_example_lgraph() *Graph {
	vexs := []string{"A", "B", "C", "D", "E", "F", "G"}
	edges := [][2]string{
		{"A", "B"},
		{"B", "C"},
		{"B", "E"},
		{"B", "F"},
		{"C", "E"},
		{"D", "C"},
		{"E", "B"},
		{"E", "D"},
		{"F", "G"}}
	vlen := len(vexs)
	elen := len(edges)

	pG := Create()
	// 初始化"顶点数"和"边数"
	pG.vexnum = vlen
	pG.edgnum = elen
	// 初始化"邻接表"的顶点
	for i := 0; i < pG.vexnum; i++ {
		pG.InsertVertex(*NewVertex(vexs[i]))
	}
	// 初始化"邻接表"的边
	for i := 0; i < pG.edgnum; i++ {
		// 读取边的起始顶点和结束顶点
		pG.InsertEdge(*NewVertex(edges[i][0]), *NewVertex(edges[i][1]))
	}
	return pG
}
