package tree

func PostOrderTraversal(root *Node) (ret []int) {
	var cur, rm *Node = root, nil

	for cur != nil {
		if cur.Right == nil {
			ret = append(ret, cur.Val)
			cur = cur.Left
		}else{
			rm = cur.Right
			for rm.Left != nil && rm.Left != cur {
				rm = rm.Left
			}

			if rm.Left == nil {
				rm.Left = cur

				ret = append(ret, cur.Val)
				cur = cur.Right
			}else{
				rm.Left = nil
				cur = cur.Left
			}
		}
	}

	for i, j := 0, len(ret)-1; i<j ; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	return
}
