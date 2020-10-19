package sort


func Insertion(data []int){
	lo := 0
	hi := len(data)

	for i := lo+1; i<hi; i++ {
		for j := i; j > lo && data[j] < data[j-1] ; j-- {
			data[j-1], data[j] = data[j], data[j-1]
		}
	}
}
