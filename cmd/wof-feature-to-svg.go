package main

import (
	"flag"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/feature"
	"github.com/whosonfirst/go-whosonfirst-svg"
	"log"
)

func main() {

	var width = flag.Float64("width", 1024., "...")
	var height = flag.Float64("height", 1024., "...")

	flag.Parse()

	o := svg.NewDefaultOptions()

	o.Width = *width
	o.Height = *height

	// TO DO: custom o.Writer goes here...

	for _, path := range flag.Args() {

		f, err := feature.LoadFeatureFromFile(path)

		if err != nil {
			log.Fatal(err)
		}

		svg.FeatureToSVG(f, o)
	}
}
