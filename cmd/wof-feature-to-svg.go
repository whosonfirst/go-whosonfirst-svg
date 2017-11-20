package main

import (
	"flag"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/feature"
	"github.com/whosonfirst/go-whosonfirst-svg"
	"log"
)

func main() {

	flag.Parse()

	flag.Parse()
	args := flag.Args()

	o := svg.NewDefaultOptions()

	for _, path := range args {

		f, err := feature.LoadFeatureFromFile(path)

		if err != nil {
			log.Fatal(err)
		}

		svg.FeatureToSVG(f, o)
	}
}
