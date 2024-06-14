package algos

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"os"
	"strconv"
	"strings"
)

func ImageLoadPixelize(path string) ([][]([]int), error) {

	imgfile, err := os.Open(path)

	if err != nil {
		fmt.Println("Issue while reading in Image")
		os.Exit(1)
	}
	defer imgfile.Close()
	img, _, err := image.Decode(imgfile)
	if err != nil {
		fmt.Println("Issue while decoding Image")
		os.Exit(1)
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)

	var pixels [][]([]int)
	for y := 0; y < height; y++ {
		var row []([]int)
		for x := 0; x < width; x++ {
			index := (y*width + x) * 4
			pix := rgba.Pix[index : index+4]
			row = append(row, []int{int(pix[0]), int(pix[1]), int(pix[2]), int(pix[3])})

		}
		pixels = append(pixels, row)
	}
	return pixels, nil
}

func Image2RBGConv(img image.Image) ([][][]int, error) {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)

	var pixels [][]([]int)
	for y := 0; y < height; y++ {
		var row []([]int)
		for x := 0; x < width; x++ {
			index := (y*width + x) * 4
			pix := rgba.Pix[index : index+4]
			row = append(row, []int{int(pix[0]), int(pix[1]), int(pix[2]), int(pix[3])})

		}
		pixels = append(pixels, row)
	}
	return pixels, nil

}

func RBG2ImageConv(pixels [][][]int, width int, height int) (image.Image, error) {
	outImage := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			outImage.SetRGBA(x, y, color.RGBA{uint8(pixels[y][x][R]), uint8(pixels[y][x][G]), uint8(pixels[y][x][B]), uint8(pixels[y][x][A])})
		}
	}
	return outImage, nil
}

func ReadPalette(colorpalettes map[string][][]int) {

	entries, err := os.ReadDir("./PALfiles")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		file, err := os.Open("./PALfiles/" + e.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		i := 0

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {

			if i > 2 {
				j := strings.Fields(scanner.Text())
				a, _ := strconv.Atoi(j[0])
				b, _ := strconv.Atoi(j[1])
				c, _ := strconv.Atoi(j[2])
				colorpalettes[strings.TrimSuffix(e.Name(), ".pal")] = append(colorpalettes[strings.TrimSuffix(e.Name(), ".pal")], []int{a, b, c})
			}
			i++
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

}
