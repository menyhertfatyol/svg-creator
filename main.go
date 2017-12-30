package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	topColorValues := []int{178, 226, 218, 150}
	bottomColorValues := []int{50, 83, 98, 0}

	var topColors []int
	var bottomColors []int

	rand.Seed(time.Now().Unix())

	for index := 0; index < 3; index++ {
		topColors = append(topColors, topColorValues[rand.Intn(len(topColorValues))])
	}

	for index := 0; index < 3; index++ {
		bottomColors = append(bottomColors, bottomColorValues[rand.Intn(len(bottomColorValues))])
	}

	topColorsString := (strings.Trim(strings.Join(strings.Fields(fmt.Sprint(topColors)), ","), "[]"))
	bottomColorsString := (strings.Trim(strings.Join(strings.Fields(fmt.Sprint(bottomColors)), ","), "[]"))

	svg := `<svg height="640" width="640">
  <defs>
    <linearGradient id="grad1" x1="0%" y1="0%" x2="0%" y2="100%">
      <stop offset="0%" style="stop-color:rgb(` + topColorsString + `);stop-opacity:1" />
      <stop offset="100%" style="stop-color:rgb(` + bottomColorsString + `);stop-opacity:1" />
    </linearGradient>
  </defs>
  <ellipse cx="320" cy="320" rx="300" ry="300" fill="url(#grad1)" />
  <text fill="#ffffff" font-size="80" font-family="Verdana" x="60" y="340">keddestidrog</text>
  Sorry, your browser does not support inline SVG.
</svg>`

	f, err := os.Create("mix_image.svg")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.WriteString(svg)
	if err != nil {
		log.Fatal(err)
	}

}
