package main

import (
	"flag"
	"github.com/whosonfirst/go-whosonfirst-svg"
	"io"
	"log"
	"os"
)

func main() {

	var width = flag.Float64("width", 1024., "...")
	var height = flag.Float64("height", 1024., "...")
	var mercator = flag.Bool("mercator", false, "...")

	flag.Parse()

	o := svg.NewDefaultOptions()

	o.Width = *width
	o.Height = *height
	o.Mercator = *mercator

	for _, path := range flag.Args() {

		r, err := os.Open(path)

		if err != nil {
			log.Fatal(err)
		}

		defer r.Close()

		body, err := io.ReadAll(r)

		if err != nil {
			log.Fatalf("Failed to read %s, %v", path, err)
		}

		err = svg.FeatureToSVG(body, o)

		if err != nil {
			log.Fatal(err)
		}
	}
}
