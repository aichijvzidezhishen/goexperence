package math

// 近似
func Approximate(x float32) int {
	var r int
	res := int(x * 2)
	if res%2 == 0 {
		r = int(x)
	} else {
		r = int(x + 1)
	}

	return r

}

// 合并表记录
func MergeSameindex() {

}
