package hipchat

import "strings"

type color string

// yellow, green, red, purple, gray, random (default: 'yellow')

const (
	ColorGreen   = color("green")
	ColorYellow  = color("yellow")
	ColorRed     = color("red")
	ColorPurple  = color("purple")
	ColorGray    = color("gray")
	ColorRandom  = color("random")
	ColorDefault = ColorYellow
)

var colors map[string]color = map[string]color{
	"green":  ColorGreen,
	"yellow": ColorYellow,
	"red":    ColorRed,
	"purple": ColorPurple,
	"gray":   ColorGray,
	"random": ColorRandom,
}

func Color(s string) color {
	if v, ok := colors[strings.ToLower(s)]; ok {
		return v
	}
	return ColorDefault
}
func isColor(s string) bool {
	c := strings.ToLower(s)
	if _, ok := colors[c]; ok {
		return true
	}
	return c == "default"
}
