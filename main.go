package main

import (
	algos "DitheringAround/lib"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

var templ *template.Template
var distance [][]int
var colorpalettes map[string][][]int

type Dataimage struct {
	Success bool
	Image   string
}

func init() {
	templ = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {

	colorpalettes = make(map[string][][]int)
	distance = algos.Caldiscolor()
	algos.ReadPalette(colorpalettes)
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":8080", nil)

}

type inputJSON struct {
	Image        string
	Block        string
	Gray         bool
	Dithering    string
	Customcheck  bool
	Reducecheck  bool
	Palname      string
	Noofcolors   string
	Colorpallete []string `json:"Colorpallete,omitempty"`
}

type outputJSON struct {
	Image        string
	Colorpallete []string `json:"Colorpallete,omitempty"`
}

func handleFunc(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		templ.ExecuteTemplate(w, "index.html", nil)
		return
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var in inputJSON
	var out outputJSON
	err = json.Unmarshal([]byte(b), &in)
	if err != nil {
		fmt.Printf("Unmarshaling Failed = %s", err)
	}

	coI := strings.Index(string(in.Image), ",")
	rawImage := string(in.Image)[coI+1:]
	unbased, _ := base64.StdEncoding.DecodeString(rawImage)
	res := bytes.NewReader(unbased)
	var img image.Image
	switch strings.TrimSuffix(in.Image[5:coI], ";base64") {
	case "image/png":
		img, err = png.Decode(res)
		if err != nil {
			fmt.Println("PNG Decoding  Failed  ")
		}
	case "image/jpeg":
		img, err = jpeg.Decode(res)
		if err != nil {
			fmt.Println("JPEG Decoding  Failed  ")
		}
	}
	imgrgb, err := algos.Image2RBGConv(img)
	if err != nil {
		fmt.Println("Failed to convert Image to RGB ")
	}

	if in.Block != "0" {
		f, _ := strconv.ParseFloat(in.Block, 64)
		imgrgb = algos.NearestNeighbourScaling(imgrgb, 1/f)
		imgrgb = algos.NearestNeighbourScaling(imgrgb, f)
	}
	if in.Gray {
		algos.Grayavg(imgrgb)
	}
	if in.Dithering != "Select" {
		if in.Dithering == "Floyd-Steinberg" {
			algos.Floydsteinberg(imgrgb)
		} else if in.Dithering == "Bayer" {
			algos.Bayer(imgrgb, 2)
		} else if in.Dithering == "Stucki" {
			algos.Stucki(imgrgb)
		}
	}
	if in.Customcheck {

		if in.Palname != "" {

			if len(in.Colorpallete) > 0 {
				var palette [][]int

				for i := 0; i < len(in.Colorpallete); i++ {

					a, _ := strconv.ParseInt(in.Colorpallete[i][1:3], 16, 64)
					b, _ := strconv.ParseInt(in.Colorpallete[i][3:5], 16, 64)
					c, _ := strconv.ParseInt(in.Colorpallete[i][5:7], 16, 64)
					palette = append(palette, []int{int(a), int(b), int(c)})
				}
				algos.Palletize(imgrgb, palette, distance)

			} else {

				algos.Palletize(imgrgb, colorpalettes[in.Palname], distance)
			}

		}

	}

	if in.Reducecheck {
		noc, err := strconv.Atoi(in.Noofcolors)
		if err != nil {
			fmt.Println("Failed to convert string to interger")
		}
		temppal, _ := algos.MedianCut(imgrgb, noc, distance)
		for i := 0; i < len(temppal); i++ {
			out.Colorpallete = append(out.Colorpallete, fmt.Sprintf("#%02x%02x%02x", temppal[i][0], temppal[i][1], temppal[i][2]))
		}
	}
	img, err = algos.RBG2ImageConv(imgrgb, len(imgrgb[0]), len(imgrgb))
	if err != nil {
		fmt.Println("Failed to convert RGB to Image.image")
	}
	var buf bytes.Buffer
	if err = png.Encode(&buf, img); err != nil {
		fmt.Println("Falied to convert image to PNG")
	}

	unbased = buf.Bytes()
	out.Image = base64.StdEncoding.EncodeToString(unbased)
	jsonData, err := json.Marshal(out)
	if err != nil {
		fmt.Println("Failed to Marshal JSON")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
