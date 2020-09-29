package tree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func DeepSum(rootNode *Node) map[int]int {
	ret := make(map[int]int)

	queue := make([]*Node, 0)
	queue = append(queue, rootNode)
	deep := 0
	thisDeepCount := 0 //本层需要处理的节点数
	for len(queue) > 0 {
		if thisDeepCount == 0 {
			deep++
			thisDeepCount = len(queue)
		}

		elem := queue[0]
		queue = queue[1:]
		ret[deep] += elem.Value
		thisDeepCount--

		if elem.Left != nil {
			queue = append(queue, elem.Left)
		}
		if elem.Right != nil {
			queue = append(queue, elem.Right)
		}
	}

	return ret
}
