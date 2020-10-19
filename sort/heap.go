package sort

func Heap (data []int) {
	hi := len(data)

	for i := (hi-1)/2; i>=0; i-- {
		siftDown(data, i, hi)
	}
	//至此data树及所有子树都是大顶

	for i := hi-1; i>=0; i-- {
		data[0], data[i] = data[i], data[0]      //将堆顶即最大值与最后一个交换，完成最大值正确排序
		siftDown(data, 0, i)                  //下沉新的堆顶
	}
}

func siftDown(data []int, lo, hi int){    //lo为需要下沉的节点，hi为堆节点个数
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {    //左孩子不存在结束
			break
		}
		if child+1 < hi && data[child] < data[child+1] {    //右孩子存在且大于左孩子， child为两个孩子节点中大的一个
			child++
		}
		if data[root] >= data[child] {    //父节点大于孩子中大的那个，结束下沉
			return
		}
		data[root], data[child] = data[child], data[root]  //下沉
		root = child     //继续子树下沉
	}
}