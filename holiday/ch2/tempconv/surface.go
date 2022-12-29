package tempconv

import (
	"fmt"
	"image"
)

func main() {
	const size = 300

	pic := image.NewGray(image.Rect(0, 0, size, size))
	fmt.Println(pic)

}
