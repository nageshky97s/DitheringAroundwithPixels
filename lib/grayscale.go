package algos

func Grayavg(pixels [][]([]int)) {
	width := len(pixels[0])
	height := len(pixels)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {

			avg := (pixels[i][j][R] + pixels[i][j][G] + pixels[i][j][B]) / 3
			pixels[i][j][R] = avg
			pixels[i][j][G] = avg
			pixels[i][j][B] = avg

		}
	}

}

func Lumagray(pixels [][]([]int)) {
	width := len(pixels[0])
	height := len(pixels)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {

			gry := int(0.2126*float32(pixels[i][j][R]) + 0.7152*float32(pixels[i][j][G]) + 0.0722*float32(pixels[i][j][B]))
			pixels[i][j][R] = gry
			pixels[i][j][G] = gry
			pixels[i][j][B] = gry

		}
	}

}

func Desaturate(pixels [][]([]int)) {
	width := len(pixels[0])
	height := len(pixels)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {

			maxi := 0
			mini := 0
			if pixels[i][j][R] >= pixels[i][j][G] {
				if pixels[i][j][R] >= pixels[i][j][B] {
					maxi = pixels[i][j][R]
				} else {
					maxi = pixels[i][j][B]
				}

			} else if pixels[i][j][G] >= pixels[i][j][B] {
				maxi = pixels[i][j][G]

			} else {
				maxi = pixels[i][j][B]
			}

			if pixels[i][j][R] <= pixels[i][j][G] {
				if pixels[i][j][R] <= pixels[i][j][B] {
					mini = pixels[i][j][R]
				} else {
					mini = pixels[i][j][B]
				}

			} else if pixels[i][j][G] <= pixels[i][j][B] {
				mini = pixels[i][j][G]

			} else {
				mini = pixels[i][j][B]
			}
			gry := (maxi + mini) / 2
			pixels[i][j][R] = gry
			pixels[i][j][G] = gry
			pixels[i][j][B] = gry

		}
	}

}
