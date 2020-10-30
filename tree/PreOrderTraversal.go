package tree

func PreOrderTraversal(root *Node) (ret []int) {

	var cur, midOrderPre *Node = root, nil
	for cur != nil {
		if cur.Left == nil {
			ret = append(ret, cur.Val)
			cur = cur.Right
		}else{
			midOrderPre = cur.Left
			for midOrderPre.Right != nil && midOrderPre.Right != cur {
				midOrderPre = midOrderPre.Right
			}

			if midOrderPre.Right == nil {
				midOrderPre.Right = cur
				ret = append(ret, cur.Val)
				cur = cur.Left
			}else{
				midOrderPre.Right = nil
				cur = cur.Right
			}
		}
	}

	return
}