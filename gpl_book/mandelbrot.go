package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height          = 200, 200 // increase those values to see why i abandon goroutines concurrency
	)
	//	var wg sync.WaitGroup
	// wg.Add(height * width)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
			// go func(a, b int, c complex128) {
			// 	img.Set(a, b, mandelbrot(c))
			// 	wg.Done()
			// }(px, py, z)
		}
	}
	//	wg.Wait()
	png.Encode(os.Stdout, img)

}

func mandelbrot(z complex128) color.Color {
	const (
		iterations = 200
		contrast   = 15
	)
	//	time.Sleep(time.Millisecond)

	var val complex128
	for n := uint8(0); n < iterations; n++ {
		val = val*val + z
		if cmplx.Abs(val) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
