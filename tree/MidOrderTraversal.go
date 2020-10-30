package tree

func MidOrderTraversal(root *Node) (ret []int) {
	var cur, pre *Node = root, nil

	for cur != nil {
		if cur.Left == nil {
			ret = append(ret, cur.Val)
			cur = cur.Right
		}else{
			pre = cur.Left
			if pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}

			if pre.Right == nil {
				pre.Right = cur
				cur = cur.Left
			}else{
				pre.Right = nil
				ret = append(ret, cur.Val)
				cur = cur.Right
			}
		}
	}

	return
}
