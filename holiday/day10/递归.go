package main

func main() {
	
}
int sum = 0;
// 非递归
func GetSum(n int)(sum int){
	for i := n; i > 0; i-- {
		sum+=i
	}
	return sum
}
// 递归
func GetSumRecursively(n int)(sum int){
	if n>1 {
		sum =n + getSumRecursively(n-1);
		return sum
	}else{
		return sum
	}
}

