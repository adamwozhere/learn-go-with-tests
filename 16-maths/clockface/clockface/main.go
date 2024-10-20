package main

import (
	"learn-go-with-tests/16-maths/clockface"
	"os"
	"time"
)

// we put a new directory under our main `clockface` directory... also called clockface.
// this is confusing but I'm not sure why - this will be the main package that we'll use
// to actually draw the SVG.

// func main() {
// 	t := time.Now()
// 	sh := clockface.SecondHand(t)
// 	io.WriteString(os.Stdout, svgStart)
// 	io.WriteString(os.Stdout, bezel)
// 	io.WriteString(os.Stdout, secondHandTag(sh))
// 	io.WriteString(os.Stdout, svgEnd)
// }

// func secondHandTag(p clockface.Point) string {
// 	return fmt.Sprintf(`<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
// }

// const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
// <!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
// <svg xmlns="http://www.w3.org/2000/svg"
//      width="100%"
//      height="100%"
//      viewBox="0 0 300 300"
//      version="2.0">`

// const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

// const svgEnd = `</svg>`

// This will build the SVG
// navigate to this directory in the command line then run:
// `go build` to create a `clockface.exe`
// then run `./clockface > clock.svg` to write the file.

// after creating an SVGWriter to create the SVG and testing it by parsing the XML, we can shorten this main function:

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
