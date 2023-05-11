package app

import "io"

type Clipper interface {
	getImage() (io.ReadCloser, error)
}
