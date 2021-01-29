package color

import (
	"github.com/amimof/huego"
	colorful "github.com/lucasb-eyer/go-colorful"
)

func HueToRGB(light *huego.Light) (r, g, b uint8) {
	normalizedX := float64(light.State.Xy[0])
	normalizedY := float64(light.State.Xy[1])

	color := colorful.Xyy(normalizedX, normalizedY, 1.0)
	return color.RGB255()
}
