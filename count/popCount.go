package count

var pc [256]byte

func init(){
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)      //前7位的个数加上最后一位是否是1，i/2写成i>>1更容易理解
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))]+
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
