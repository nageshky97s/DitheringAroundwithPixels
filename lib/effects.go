package algos

import "math"

// func gamma22(pixels [][]([]int)) [][]([]int) {

// 	return pixels
// }

func SharpenWarpAround(pixels [][]([]int), sharpness int) [][][]int {
	shrp := make([][][]int, len(pixels))
	for i := range shrp {
		shrp[i] = make([][]int, len(pixels[i]))
		for j := range shrp[i] {
			shrp[i][j] = make([]int, len(pixels[i][j]))
		}
	}

	width := len(pixels[0])
	height := len(pixels)

	up := 0
	down := 0
	left := 0
	right := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			for k := 0; k < 3; k++ {
				if i == 0 {
					up = -(pixels[height-1][j][k])
				} else {
					up = -(pixels[i-1][j][k])
				}
				if j == 0 {
					left = -(pixels[i][width-1][k])
				} else {
					left = -(pixels[i][j-1][k])
				}
				if j == width-1 {
					right = -(pixels[i][0][k])
				} else {
					right = -(pixels[i][j+1][k])
				}
				if i == height-1 {
					down = -(pixels[0][j][k])
				} else {
					down = -(pixels[i+1][j][k])
				}

				if pixels[i][j][k]*(sharpness*4+1) +
					+up + left + right + down > 255 {
					shrp[i][j][k] = 255
				} else if pixels[i][j][k]*(sharpness*4+1) +
					+up + left + right + down < 0 {
					shrp[i][j][k] = 0
				} else {
					shrp[i][j][k] = pixels[i][j][k]*(sharpness*4+1) +
						+up + left + right + down
				}

			}
			shrp[i][j][3] = pixels[i][j][3]
		}
	}

	// writeImage(shrp, width, height, "SharpenWrap")
	return shrp
}

func SharpenPadZero(pixels [][]([]int), sharpness int) [][][]int {
	shrp := make([][][]int, len(pixels))
	for i := range shrp {
		shrp[i] = make([][]int, len(pixels[i]))
		for j := range shrp[i] {
			shrp[i][j] = make([]int, len(pixels[i][j]))
		}
	}
	width := len(pixels[0])
	height := len(pixels)
	up := 0
	down := 0
	left := 0
	right := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			for k := 0; k < 3; k++ {
				if i == 0 {
					up = 0
				} else {
					up = -(pixels[i-1][j][k])
				}
				if j == 0 {
					left = 0
				} else {
					left = -(pixels[i][j-1][k])
				}
				if j == width-1 {
					right = 0
				} else {
					right = -(pixels[i][j+1][k])
				}
				if i == height-1 {
					down = 0
				} else {
					down = -(pixels[i+1][j][k])
				}

				if pixels[i][j][k]*(sharpness*4+1) +
					+up + left + right + down > 255 {
					shrp[i][j][k] = 255
				} else if pixels[i][j][k]*(sharpness*4+1) +
					+up + left + right + down < 0 {
					shrp[i][j][k] = 0
				} else {
					shrp[i][j][k] = pixels[i][j][k]*(sharpness*4+1) +
						+up + left + right + down
				}

			}
			shrp[i][j][3] = pixels[i][j][3]
		}
	}

	// writeImage(shrp, width, height, "SharpenZero")
	return shrp
}
func Palletize(src [][][]int, pallete [][]int, distance [][]int) {
	width := len(src[0])
	height := len(src)
	mindis := 0.0
	colind := 0
	dis := 0.0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			mindis = math.MaxFloat64
			colind = 0
			for k := 0; k < len(pallete); k++ {
				dis = colordistancefast(src[i][j], pallete[k], distance)
				if dis < mindis {
					colind = k
					mindis = dis
				}
			}
			src[i][j][0] = pallete[colind][0]
			src[i][j][1] = pallete[colind][1]
			src[i][j][2] = pallete[colind][2]
		}
	}
	// writeImage(src, width, height, "ChangedPallete")
}

func colordistance(c1 []int, c2 []int) float64 {
	return math.Pow(float64(c1[0]-c2[0]), 2) +
		math.Pow(float64(c1[1]-c2[1]), 2) +
		math.Pow(float64(c1[2]-c2[2]), 2)
}

func colordistancefast(c1 []int, c2 []int, distance [][]int) float64 {
	return float64(distance[c1[0]][c2[0]] +
		distance[c1[1]][c2[1]] +
		distance[c1[2]][c2[2]])
}
