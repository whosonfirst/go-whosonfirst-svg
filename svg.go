package svg

import (
	"github.com/fapian/geojson2svg/pkg/geojson2svg"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/properties/geometry"
	"io"
	"os"
)

type Options struct {
	Width  float64
	Height float64
	Writer io.Writer
}

func NewDefaultOptions() *Options {

	opts := Options{
		Width:  1024.0,
		Height: 1024.0,
		Writer: os.Stdout,
	}

	return &opts
}

func FeatureToSVG(f geojson.Feature, opts *Options) error {

	geom, err := geometry.ToString(f)

	if err != nil {
		return err
	}

	s := geojson2svg.New()

	err = s.AddGeometry(geom)

	if err != nil {
		return err
	}

	rsp := s.Draw(opts.Width, opts.Height)
	_, err = opts.Writer.Write([]byte(rsp))

	if err != nil {
		return err
	}

	return nil
}
