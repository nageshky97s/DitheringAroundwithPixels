package algos

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

func Kmeans(pixels [][][]int, noOfColors int, maxIter int) {

	width := len(pixels[0])
	height := len(pixels)
	clusterHeads := make([][]int, noOfColors)
	clusterGroup := make([][][]int, noOfColors)
	colorind := make([][][]int, noOfColors)

	totalpixels := len(pixels) * len(pixels[0])
	fmt.Println(totalpixels)

	//Finding unique colors
	uniquecolors := make(map[[3]int]int)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {

			uniquecolors[[3]int{pixels[i][j][0], pixels[i][j][1], pixels[i][j][2]}]++
		}
	}
	totalColors := len(uniquecolors)
	fmt.Println(totalColors)
	if noOfColors > totalColors {
		fmt.Println("The No. of Colors Entered is more than the available colors in the palette")
		return
	}
	//No need to do the Fisher Yates Shuffling to get a random numbers since the order of a map is not guaranteed ie it is randomized
	count := 0
	for i := range uniquecolors {
		clusterHeads[count] = i[:]
		count++
		if count == noOfColors {
			break
		}
	}
	// totalIter := 0
	// if noOfColors == totalColors {
	// 	totalIter = totalColors - 2
	// } else {
	// 	totalIter = noOfColors
	// }
	// //FisherYates shuffle
	// ind := make([]int, totalpixels)
	// for i := 0; i < totalpixels; i++ {
	// 	ind[i] = i
	// }
	// for i := 0; i < totalIter; i++ {
	// 	j := randRange(i, totalpixels-1)
	// 	ind[i], ind[j] = ind[j], ind[i]
	// }
	// for i := 0; i < noOfColors; i++ {
	// 	clusterHeads[i] = pixels[ind[i]/len(pixels[0])][ind[i]%len(pixels[0])]
	// }

	for a := 0; a < maxIter; a++ {
		fmt.Println(a)
		clusterGroup = make([][][]int, noOfColors)
		colorind = make([][][]int, noOfColors)
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				mindis := math.MaxFloat64
				minind := 0
				dis := 0.0
				for z := 0; z < noOfColors; z++ {
					dis = colordistance(clusterHeads[z], pixels[i][j])

					if dis < mindis {
						mindis = dis
						minind = z
					}

				}
				clusterGroup[minind] = append(clusterGroup[minind], pixels[i][j])
				colorind[minind] = append(colorind[minind], []int{i, j})
			}
		}
		for i := 0; i < noOfColors; i++ {
			meanR := 0
			meanG := 0
			meanB := 0

			for j := 0; j < len(clusterGroup[i]); j++ {
				meanR += clusterGroup[i][j][0]
				meanG += clusterGroup[i][j][1]
				meanB += clusterGroup[i][j][2]
			}
			clusterHeads[i][0] = meanR / len(clusterGroup[i])
			clusterHeads[i][1] = meanG / len(clusterGroup[i])
			clusterHeads[i][2] = meanB / len(clusterGroup[i])

		}

	}
	for i := 0; i < noOfColors; i++ {
		for j := 0; j < len(colorind[i]); j++ {
			pixels[colorind[i][j][0]][colorind[i][j][1]][0] = clusterHeads[i][0]
			pixels[colorind[i][j][0]][colorind[i][j][1]][1] = clusterHeads[i][1]
			pixels[colorind[i][j][0]][colorind[i][j][1]][2] = clusterHeads[i][2]
		}
	}
	// writeImage(pixels, width, height, "KmeansColorReduced")
	// fmt.Println("Done")
}

func randRange(min, max int) int {
	return rand.Intn(max+1-min) + min
}

func MedianCut(pixels [][][]int, noOfColors int, distance [][]int) ([][]int, [][]int) {
	// start := time.Now()
	width := len(pixels[0])
	height := len(pixels)
	// totalpixels := len(pixels) * len(pixels[0])
	// fmt.Println(totalpixels)

	//Finding unique colors
	uniquecolors := make(map[[3]int]int)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {

			uniquecolors[[3]int{pixels[i][j][0], pixels[i][j][1], pixels[i][j][2]}]++
		}
	}
	totalColors := len(uniquecolors)
	fmt.Println(totalColors)
	fmt.Println(noOfColors)
	if noOfColors > totalColors {
		fmt.Println("The No. of Colors Entered is more than the available colors in the palette")
		return nil, nil
	}
	count := 0
	colorslice := make([][]int, totalColors)
	var buckets [][][]int
	for i := range uniquecolors {
		colorslice[count] = i[:]
		count++
	}
	buckets = append(buckets, colorslice)
	// fmt.Println(len(buckets[0]))

	for a := 0; a < noOfColors-1; a++ {
		minmax := []int{math.MaxInt, math.MaxInt, math.MaxInt, math.MinInt, math.MinInt, math.MinInt}
		for i := 0; i < len(buckets[0]); i++ {
			for j := 0; j < 3; j++ {
				if buckets[0][i][j] < minmax[j] {
					minmax[j] = buckets[0][i][j]
				}
				if buckets[0][i][j] > minmax[j+3] {
					minmax[j+3] = buckets[0][i][j]
				}
			}

		}
		ind := randomizeselectChannel(minmax)
		sort.Slice(buckets[0], func(i, j int) bool {
			// edge cases
			if len(buckets[0]) == 0 {
				return false // when bucket to be cut is empty
			}
			// both slices len() > 0, so can test this now:
			return buckets[0][i][ind] < buckets[0][j][ind]
		})
		buckets = append(buckets, buckets[0][0:len(buckets[0])/2])
		buckets = append(buckets, buckets[0][len(buckets[0])/2:])
		buckets = buckets[1:]

	}
	allcolors := make([][]int, len(buckets))
	colorpostions := make([][]int, len(buckets))
	for i := 0; i < len(buckets); i++ {
		sumR := 0
		sumG := 0
		sumB := 0
		for j := 0; j < len(buckets[i]); j++ {
			sumR += buckets[i][j][0]
			sumG += buckets[i][j][1]
			sumB += buckets[i][j][2]
		}
		sumR /= len(buckets[i])
		sumG /= len(buckets[i])
		sumB /= len(buckets[i])
		allcolors[i] = []int{sumR, sumG, sumB}
	}

	dis := 0.0
	colind := 0
	mindis := 0.0

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			mindis = math.MaxFloat64
			colind = 0
			for z := 0; z < len(allcolors); z++ {
				dis = colordistancefast(pixels[i][j], allcolors[z], distance)
				// dis = colordistance(pixels[i][j], allcolors[z])
				if dis < mindis {
					colind = z
					mindis = dis
				}
			}
			pixels[i][j][0] = allcolors[colind][0]
			pixels[i][j][1] = allcolors[colind][1]
			pixels[i][j][2] = allcolors[colind][2]
			colorpostions[colind] = []int{i, j}
		}
	}
	// elapsed := time.Since(start)
	// log.Printf("Binomial took %s", elapsed)
	// writeImage(pixels, width, height, "MedianColors"+strconv.Itoa(noOfColors))
	return allcolors, colorpostions
}

func randomizeselectChannel(rgb []int) int {

	maxval := math.MaxInt
	if rgb[3]-rgb[0] >= rgb[4]-rgb[1] && rgb[3]-rgb[0] >= rgb[5]-rgb[2] {
		maxval = rgb[3] - rgb[0]
	} else if rgb[4]-rgb[1] >= rgb[3]-rgb[0] && rgb[4]-rgb[1] >= rgb[5]-rgb[2] {
		maxval = rgb[4] - rgb[1]
	} else {
		maxval = rgb[5] - rgb[2]
	}
	maxsl := []int{}
	for i := 0; i < 3; i++ {
		if rgb[i+3]-rgb[i] == maxval {
			maxsl = append(maxsl, i)
		}
	}
	if len(maxsl) > 1 {
		maxval = rand.Intn(len(maxsl))
	} else {
		maxval = maxsl[0]
	}

	return maxval

}

func Caldiscolor() [][]int {
	distances := make([][]int, 256)
	for i := 0; i < 256; i++ {
		distances[i] = make([]int, 256)
	}
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			distances[i][j] = int(math.Pow(float64(i-j), 2))
		}
	}
	return distances
}
