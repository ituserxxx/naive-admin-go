package utils

import (
	"bytes"
	"math/rand"
	"strconv"
	"time"

	svg "github.com/ajstarks/svgo"
)

func GenerateSVG(width, height int) ([]byte,string) {
	rand.Seed(time.Now().UnixNano())

	var svgContent bytes.Buffer
	canvas := svg.New(&svgContent)
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height, "fill:white")
	canvas.Circle(width/2, height/2, width/2-5, "fill:#EEE")

	code := strconv.Itoa(rand.Intn(9999))
	canvas.Text(width/2, height/2, code, "text-anchor:middle; font-size:40px; fill:black;")

	canvas.End()

	return svgContent.Bytes(),code
}

