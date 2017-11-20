package main

import (
	"context"
	"flag"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/feature"
	"github.com/whosonfirst/go-whosonfirst-index"
	"github.com/whosonfirst/go-whosonfirst-svg"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {

	var mode = flag.String("mode", "files", "...")
	var width = flag.Float64("width", 1024., "...")
	var height = flag.Float64("height", 1024., "...")

	flag.Parse()

	o := svg.NewDefaultOptions()

	o.Width = *width
	o.Height = *height

	cb := func(fh io.Reader, ctx context.Context, args ...interface{}) error {

		path, err := index.PathForContext(ctx)

		if err != nil {
			return err
		}

		if path != index.STDIN {

			ext := filepath.Ext(path)

			if ext != ".geojson" {
				return nil
			}
		}

		f, err := feature.LoadFeatureFromFile(path)

		if err != nil {
			return err
		}

		// update to write new file here - take fname and replace ".geojson" with ".svg"
		// unless this is index.STDIN in which case... uh, what ? I guess just STDOUT...

		svg.FeatureToSVG(f, o)
		return nil
	}

	idx, err := index.NewIndexer(*mode, cb)

	if err != nil {
		log.Fatal(err)
	}

	sources := flag.Args()
	err = idx.IndexPaths(sources)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
