package base

type Point struct {
	A int32
}

func PointRes() *Point {
	var res *Point
	res.A = 1
	return res
}
