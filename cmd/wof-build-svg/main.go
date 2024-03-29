package main

import (
	"context"
	"flag"
	"github.com/facebookgo/atomicfile"
	"github.com/whosonfirst/go-whosonfirst-iterate/v2/emitter"
	"github.com/whosonfirst/go-whosonfirst-iterate/v2/iterator"
	"github.com/whosonfirst/go-whosonfirst-svg"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	var mode = flag.String("mode", "files", "...")
	var width = flag.Float64("width", 1024., "...")
	var height = flag.Float64("height", 1024., "...")

	flag.Parse()

	o := svg.NewDefaultOptions()

	o.Width = *width
	o.Height = *height

	cb := func(ctx context.Context, path string, fh io.ReadSeeker, args ...interface{}) error {

		if path != emitter.STDIN {

			ext := filepath.Ext(path)

			if ext != ".geojson" {
				return nil
			}
		}

		f, err := io.ReadAll(fh)

		if err != nil {
			// because this: https://github.com/whosonfirst/go-whosonfirst-svg/issues/3
			// log.Println("SKIP", path)
			return nil
		}

		// update to write new file here - take fname and replace ".geojson" with ".svg"
		// unless this is index.STDIN in which case... uh, what ? I guess just STDOUT...

		if path == emitter.STDIN {
			svg.FeatureToSVG(f, o)
			return nil
		}

		root := filepath.Dir(path)
		fname := filepath.Base(path)
		ext := filepath.Ext(path)

		fname = strings.Replace(fname, ext, ".svg", -1)
		svg_path := filepath.Join(root, fname)

		wr, err := atomicfile.New(svg_path, os.FileMode(0644))

		if err != nil {
			wr.Abort()
			return err
		}

		o.Writer = wr
		svg.FeatureToSVG(f, o)

		return wr.Close()
	}

	ctx := context.Background()

	iter, err := iterator.NewIterator(ctx, *mode, cb)

	if err != nil {
		log.Fatal(err)
	}

	sources := flag.Args()
	err = iter.IterateURIs(ctx, sources...)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
