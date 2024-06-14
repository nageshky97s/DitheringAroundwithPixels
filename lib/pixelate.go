package algos

import (
	"fmt"
	"math"
)

func BoxblurMirror(pixels [][]([]int), blocksize int) {
	width := len(pixels[0])
	height := len(pixels)

	if blocksize > width || blocksize > height {
		fmt.Println("The block size is greater than the image size")
		return
	}
	blurr := make([][][]int, len(pixels))

	for i := range blurr {
		blurr[i] = make([][]int, len(pixels[i]))
		for j := range pixels[i] {
			blurr[i][j] = make([]int, len(pixels[i][j]))
		}

	}

	for i := 0; i < height; i = i + blocksize {

		for j := 0; j < width; j = j + blocksize {

			sum_R := 0
			sum_G := 0
			sum_B := 0
			z := 0
			y := 0
			for out := 0; out < blocksize; out++ {
				y = 0
				for in := 0; in < blocksize; in++ {

					if i+out < height && j+in < width {
						sum_R += pixels[i+out][j+in][0]
						sum_G += pixels[i+out][j+in][1]
						sum_B += pixels[i+out][j+in][2]
					} else if i+out >= height && j+in < width {
						sum_R += pixels[i-1-z][j+in][0]
						sum_G += pixels[i-1-z][j+in][1]
						sum_B += pixels[i-1-z][j+in][2]

					} else if i+out < height && j+in >= width {
						sum_R += pixels[i+out][j-1-y][0]
						sum_G += pixels[i+out][j-1-y][1]
						sum_B += pixels[i+out][j-1-y][2]
						y++

					} else if i+out >= height && j+in >= width {
						sum_R += pixels[i-1-z][j-1-y][0]
						sum_G += pixels[i-1-z][j-1-y][1]
						sum_B += pixels[i-1-z][j-1-y][2]
						y++

					}

				}
				if i+out >= height {
					z++
				}

			}
			sum_R /= (blocksize * blocksize)
			sum_G /= (blocksize * blocksize)
			sum_B /= (blocksize * blocksize)
			for out := 0; out < blocksize && i+out < height; out++ {
				for in := 0; in < blocksize && j+in < width; in++ {

					blurr[i+out][j+in][0] = sum_R
					blurr[i+out][j+in][1] = sum_G
					blurr[i+out][j+in][2] = sum_B
					blurr[i+out][j+in][3] = pixels[i+out][j+in][3]

				}

			}

		}

	}

	// writeImage(blurr, width, height, "BoxBlurrMirror")
}

func NearestNeighbourScaling(pixels [][]([]int), scale float64) [][][]int {

	width := len(pixels[0])
	heigth := len(pixels)
	tarwidth := int(scale * float64(width))
	tarheigth := int(scale * float64(heigth))
	scaledImage := make([][][]int, tarheigth)
	for i := range scaledImage {
		scaledImage[i] = make([][]int, tarwidth)
		for j := range scaledImage[i] {
			scaledImage[i][j] = make([]int, 4)
		}
	}

	for i := 0; i < tarheigth; i++ {
		srcX := int(math.Round(float64(i) / float64(tarheigth) * float64(heigth)))
		srcX = int(math.Min(float64(srcX), float64(heigth-1)))
		for j := 0; j < tarwidth; j++ {
			srcY := int(math.Round(float64(j) / float64(tarwidth) * float64(width)))
			srcY = int(math.Min(float64(srcY), float64(width-1)))

			for z := 0; z < 4; z++ {
				scaledImage[i][j][z] = pixels[srcX][srcY][z]
			}
		}
	}
	// writeImage(scaledImage, tarwidth, tarheigth, "PointFilterScaled")
	return scaledImage
}
