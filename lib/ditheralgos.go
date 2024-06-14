package algos

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

const (
	R pixel = iota
	G
	B
	A
)

type pixel uint8

func JarvisJudiceNinke(pixels [][]([]int)) {
	width := len(pixels[0])
	height := len(pixels)

	newpixel := 0
	oldpixel := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			for z := 0; z < 3; z++ { //z<3 because we dont want to change the transperancy or the 'A' value
				oldpixel = int(pixels[i][j][z])

				if oldpixel >= 128 {
					newpixel = 255
				} else {
					newpixel = 0
				}
				// newpixel = closestpixel(oldpixel)
				errpixel := oldpixel - newpixel
				pixels[i][j][z] = newpixel
				if j+1 < width {

					pixels[i][j+1][z] += (errpixel * 7) / 48
				}
				if j+2 < width {

					pixels[i][j+2][z] += (errpixel * 5) / 48
				}
				if j-2 >= 0 && i+1 < height {

					pixels[i+1][j-2][z] += (errpixel * 3) / 48
				}
				if j-1 >= 0 && i+1 < height {

					pixels[i+1][j-1][z] += (errpixel * 5) / 48
				}
				if i+1 < height {

					pixels[i+1][j][z] += ((errpixel * 7) / 48)

				}
				if j+1 < width && i+1 < height {

					pixels[i+1][j+1][z] += (errpixel * 5) / 48
				}
				if j+2 < width && i+1 < height {

					pixels[i+1][j+2][z] += (errpixel * 3) / 48
				}
				if j-2 >= 0 && i+2 < height {

					pixels[i+2][j-2][z] += (errpixel) / 48
				}
				if j-1 >= 0 && i+2 < height {

					pixels[i+2][j-1][z] += (errpixel * 3) / 48
				}
				if i+2 < height {

					pixels[i+2][j][z] += ((errpixel * 5) / 48)

				}
				if j+1 < width && i+2 < height {

					pixels[i+2][j+1][z] += (errpixel * 3) / 48
				}
				if j+2 < width && i+2 < height {

					pixels[i+2][j+2][z] += (errpixel) / 48
				}

			}

		}

	}

	// writeImage(pixels, width, height, "JarvisJudiceNinke")

}

func Stucki(pixels [][]([]int)) {
	width := len(pixels[0])
	height := len(pixels)

	newpixel := 0
	oldpixel := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			for z := 0; z < 3; z++ { //z<3 because we dont want to change the transperancy or the 'A' value
				oldpixel = int(pixels[i][j][z])

				if oldpixel >= 128 {
					newpixel = 255
				} else {
					newpixel = 0
				}

				errpixel := oldpixel - newpixel
				pixels[i][j][z] = newpixel
				if j+1 < width {

					pixels[i][j+1][z] += (errpixel * 8) / 42
				}
				if j+2 < width {

					pixels[i][j+2][z] += (errpixel * 4) / 42
				}
				if j-2 >= 0 && i+1 < height {

					pixels[i+1][j-2][z] += (errpixel * 2) / 42
				}
				if j-1 >= 0 && i+1 < height {

					pixels[i+1][j-1][z] += (errpixel * 4) / 42
				}
				if i+1 < height {

					pixels[i+1][j][z] += ((errpixel * 8) / 42)

				}
				if j+1 < width && i+1 < height {

					pixels[i+1][j+1][z] += (errpixel * 4) / 42
				}
				if j+2 < width && i+1 < height {

					pixels[i+1][j+2][z] += (errpixel * 2) / 42
				}
				if j-2 >= 0 && i+2 < height {

					pixels[i+2][j-2][z] += (errpixel) / 42
				}
				if j-1 >= 0 && i+2 < height {

					pixels[i+2][j-1][z] += (errpixel * 2) / 42
				}
				if i+2 < height {

					pixels[i+2][j][z] += ((errpixel * 4) / 42)

				}
				if j+1 < width && i+2 < height {

					pixels[i+2][j+1][z] += (errpixel * 2) / 42
				}
				if j+2 < width && i+2 < height {

					pixels[i+2][j+2][z] += (errpixel) / 42
				}

			}

		}

	}

	// writeImage(pixels, width, height, "Stucki")

}
func Atkinson(pixels [][]([]int)) {
	width := len(pixels[0])
	height := len(pixels)
	newpixel := 0
	oldpixel := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			for z := 0; z < 3; z++ { //z<3 because we dont want to change the transperancy or the 'A' value
				oldpixel = int(pixels[i][j][z])

				if oldpixel >= 128 {
					newpixel = 255
				} else {
					newpixel = 0
				}

				errpixel := oldpixel - newpixel
				pixels[i][j][z] = newpixel
				if j+1 < width {

					pixels[i][j+1][z] += (errpixel) / 8
				}
				if j+2 < width {

					pixels[i][j+2][z] += (errpixel) / 8
				}
				if i+1 < height && j-1 >= 0 {

					pixels[i+1][j-1][z] += ((errpixel) / 8)
				}
				if i+1 < height {

					pixels[i+1][j][z] += ((errpixel) / 8)
				}
				if i+1 < height && j+1 < width {

					pixels[i+1][j+1][z] += ((errpixel) / 8)
				}
				if i+2 < height {

					pixels[i+2][j][z] += ((errpixel) / 8)
				}

			}

		}

	}

	// writeImage(pixels, width, height, "Atkinson")
}
func Burkes(pixels [][]([]int)) {
	width := len(pixels[0])
	height := len(pixels)
	newpixel := 0
	oldpixel := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			for z := 0; z < 3; z++ { //z<3 because we dont want to change the transperancy or the 'A' value
				oldpixel = int(pixels[i][j][z])

				if oldpixel >= 128 {
					newpixel = 255
				} else {
					newpixel = 0
				}

				errpixel := oldpixel - newpixel

				pixels[i][j][z] = newpixel
				if j+1 < width {

					pixels[i][j+1][z] += (errpixel * 8) / 32
				}
				if j+2 < width {

					pixels[i][j+2][z] += (errpixel * 4) / 32
				}
				if i+1 < height && j-2 >= 0 {

					pixels[i+1][j-2][z] += ((errpixel * 2) / 32)
				}
				if i+1 < height && j-1 >= 0 {

					pixels[i+1][j-1][z] += ((errpixel * 4) / 32)
				}
				if i+1 < height {

					pixels[i+1][j][z] += ((errpixel * 8) / 32)
				}
				if i+1 < height && j+1 < width {

					pixels[i+1][j+1][z] += ((errpixel * 4) / 32)
				}
				if i+1 < height && j+2 < width {

					pixels[i+1][j+2][z] += ((errpixel * 2) / 32)
				}

			}

		}

	}

	// writeImage(pixels, width, height, "Burkes")

}

func Sierra(pixels [][]([]int)) {
	width := len(pixels[0])
	height := len(pixels)
	newpixel := 0
	oldpixel := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			for z := 0; z < 3; z++ { //z<3 because we dont want to change the transperancy or the 'A' value
				oldpixel = int(pixels[i][j][z])

				if oldpixel >= 128 {
					newpixel = 255
				} else {
					newpixel = 0
				}

				errpixel := oldpixel - newpixel

				pixels[i][j][z] = newpixel
				if j+1 < width {

					pixels[i][j+1][z] += (errpixel * 5) / 32
				}
				if j+2 < width {

					pixels[i][j+2][z] += (errpixel * 3) / 32
				}
				if i+1 < height && j-2 >= 0 {

					pixels[i+1][j-2][z] += ((errpixel * 2) / 32)
				}
				if i+1 < height && j-1 >= 0 {

					pixels[i+1][j-1][z] += ((errpixel * 4) / 32)
				}
				if i+1 < height {

					pixels[i+1][j][z] += ((errpixel * 5) / 32)
				}
				if i+1 < height && j+1 < width {

					pixels[i+1][j+1][z] += ((errpixel * 4) / 32)
				}
				if i+1 < height && j+2 < width {

					pixels[i+1][j+2][z] += ((errpixel * 2) / 32)
				}
				if i+2 < height && j-1 >= 0 {

					pixels[i+2][j-1][z] += ((errpixel * 2) / 32)
				}
				if i+2 < height {

					pixels[i+2][j][z] += ((errpixel * 3) / 32)
				}
				if i+2 < height && j+1 < width {

					pixels[i+2][j+1][z] += ((errpixel * 2) / 32)
				}

			}

		}

	}

	// writeImage(pixels, width, height, "Sierra")

}
func Sierra2Row(pixels [][]([]int)) {
	width := len(pixels[0])
	height := len(pixels)
	newpixel := 0
	oldpixel := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			for z := 0; z < 3; z++ { //z<3 because we dont want to change the transperancy or the 'A' value
				oldpixel = int(pixels[i][j][z])

				if oldpixel >= 128 {
					newpixel = 255
				} else {
					newpixel = 0
				}

				errpixel := oldpixel - newpixel

				pixels[i][j][z] = newpixel
				if j+1 < width {

					pixels[i][j+1][z] += (errpixel * 4) / 16
				}
				if j+2 < width {

					pixels[i][j+2][z] += (errpixel * 3) / 16
				}
				if i+1 < height && j-2 >= 0 {

					pixels[i+1][j-2][z] += ((errpixel * 1) / 16)
				}
				if i+1 < height && j-1 >= 0 {

					pixels[i+1][j-1][z] += ((errpixel * 2) / 16)
				}
				if i+1 < height {

					pixels[i+1][j][z] += ((errpixel * 3) / 16)
				}
				if i+1 < height && j+1 < width {

					pixels[i+1][j+1][z] += ((errpixel * 2) / 16)
				}
				if i+1 < height && j+2 < width {

					pixels[i+1][j+2][z] += ((errpixel * 1) / 16)
				}

			}

		}

	}

	// writeImage(pixels, width, height, "Sierra2Row")

}
func SierraLite(pixels [][]([]int)) {
	width := len(pixels[0])
	height := len(pixels)
	newpixel := 0
	oldpixel := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			for z := 0; z < 3; z++ { //z<3 because we dont want to change the transperancy or the 'A' value
				oldpixel = int(pixels[i][j][z])

				if oldpixel >= 128 {
					newpixel = 255
				} else {
					newpixel = 0
				}

				errpixel := oldpixel - newpixel

				pixels[i][j][z] = newpixel
				if j+1 < width {

					pixels[i][j+1][z] += (errpixel * 2) / 4
				}

				if i+1 < height && j-1 >= 0 {

					pixels[i+1][j-1][z] += ((errpixel * 1) / 4)
				}
				if i+1 < height {

					pixels[i+1][j][z] += ((errpixel * 1) / 4)
				}

			}

		}

	}

	// writeImage(pixels, width, height, "SierraLite")

}

func Bayer(pixels [][]([]int), mapsize int) {

	baemat := createBayerMat(mapsize, true)
	width := len(pixels[0])
	height := len(pixels)
	// fmt.Println(width)
	// fmt.Println(height)
	// fmt.Printf("%v", baemat)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			for z := 0; z < 3; z++ {
				if pixels[i][j][z] >= baemat[i%(len(baemat))][j%(len(baemat))] {
					pixels[i][j][z] = 255
				} else {
					pixels[i][j][z] = 0
				}
			}
		}
	}

	// writeImage(pixels, width, height, "Bayer"+strconv.Itoa(mapsize))

}

func createBayerMat(sizebay int, scale bool) [][]int {

	matprev := [][]int{{0}}
	mat := [][]int{{0}}
	for i := 0; i < sizebay; i++ {
		power := int(math.Pow(2, float64(i+1)))
		mat = make([][]int, power)
		for i := range mat {
			mat[i] = make([]int, power)
		}

		for j := 0; j < len(matprev); j++ {

			for k := 0; k < len(matprev[0]); k++ {

				mat[j][k] = matprev[j][k] * 4
				mat[j][k+len(matprev)] = matprev[j][k]*4 + 2
				mat[j+len(matprev)][k] = matprev[j][k]*4 + 3
				mat[j+len(matprev)][k+len(matprev)] = matprev[j][k]*4 + 1
			}
		}

		matprev = mat
	}
	if scale {
		scaleto256Normalizze(mat)
	}

	return mat
}

func scaleto256Normalizze(mat [][]int) {

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			mat[i][j] = int(math.Round(float64(255 * (mat[i][j]) / (len(mat) * len(mat)))))
		}
	}

}

func writeImage(pixels [][]([]int), width int, height int, name string) {
	outImage := image.NewRGBA(image.Rect(0, 0, width, height))

	f, err := os.Create("results/" + name + ".png")
	if err != nil {
		fmt.Println("Failed to Create Image")
		os.Exit(1)
	}
	defer f.Close()
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			outImage.SetRGBA(x, y, color.RGBA{uint8(pixels[y][x][R]), uint8(pixels[y][x][G]), uint8(pixels[y][x][B]), uint8(pixels[y][x][A])})
		}
	}

	png.Encode(f, outImage)

}

// func closestpixel(oldpixel int) int {
// 	return int(math.Round(float64(oldpixel)/255)) * 255
// }

func Floydsteinberg(pixels [][]([]int)) {
	width := len(pixels[0])
	height := len(pixels)
	newpixel := 0
	oldpixel := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			for z := 0; z < 3; z++ { //z<3 because we dont want to change the transperancy or the 'A' value
				oldpixel = int(pixels[i][j][z])

				if oldpixel >= 128 {
					newpixel = 255
				} else {
					newpixel = 0
				}

				errpixel := oldpixel - newpixel
				pixels[i][j][z] = newpixel
				if j+1 < width {
					//pixels[i][j+1][z] += (errpixel * 7) >> 4
					pixels[i][j+1][z] += (errpixel * 7) / 16
				}
				if j-1 >= 0 && i+1 < height {
					//pixels[i+1][j-1][z] += (errpixel * 3) >> 4
					pixels[i+1][j-1][z] += (errpixel * 3) / 16
				}
				if i+1 < height && j+1 < width {
					// pixels[i+1][j][z] += (errpixel * 5) >> 4
					// pixels[i+1][j+1][z] += (errpixel) >> 4
					pixels[i+1][j][z] += ((errpixel * 5) / 16)
					pixels[i+1][j+1][z] += ((errpixel) / 16)
				}

			}

		}

	}

	// writeImage(pixels, width, height, "Floydsteinberg")

}

func BayerMulti(pixels [][][]int, noOfColors int, spread int, matrixsize int) {
	width := len(pixels[0])
	height := len(pixels)
	bayermat := make([][][]float64, height)
	for i := range bayermat {
		bayermat[i] = make([][]float64, width)
		for j := range bayermat[i] {
			bayermat[i][j] = make([]float64, 4)
		}
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			for z := 0; z < len(pixels[i][j]); z++ {
				bayermat[i][j][z] = float64(pixels[i][j][z]) / 255
			}
		}
	}
	thresholdmat := createBayerMat(matrixsize, false)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			for z := 0; z < 3; z++ {
				bayermat[i][j][z] +=
					(float64(thresholdmat[i%(len(thresholdmat))][j%(len(thresholdmat))])*
						float64(1/(len(thresholdmat)*len(thresholdmat))) - 0.5) *
						float64(spread)
				bayermat[i][j][z] = math.Floor(bayermat[i][j][z]*float64(noOfColors-1)+0.5) /
					float64(noOfColors-1)
			}
		}
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			for z := 0; z < 3; z++ {
				pixels[i][j][z] = int(bayermat[i][j][z] * 255)
			}
		}
	}

	// writeImage(pixels, width, height, "BayerMulti"+strconv.Itoa(matrixsize))

}
