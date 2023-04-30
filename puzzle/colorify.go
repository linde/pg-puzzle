package puzzle

import (
	"regexp"
)

const (
	// ${1} is replaced with 1,2,3... and makes the colors
	ANSICOLOR_PRE  = "\033[3"
	ANSICOLOR_POST = "m"

	ANSICOLOR_TEMPLATE = ANSICOLOR_PRE + "${1}" + ANSICOLOR_POST
	ANSICOLOR_RESET    = "\033[0m"
)

// TODO consider taking a board instead of a string
// TODO unit test

var COLOR_RE = regexp.MustCompile(`([0-9])`)

// h/t https://twin.sh/articles/35/how-to-add-colors-to-your-console-terminal-output-in-go

func Colorify(str string) string {

	colorizedStr := COLOR_RE.ReplaceAllString(str, ANSICOLOR_TEMPLATE+"${1}"+ANSICOLOR_RESET)
	return colorizedStr

}
