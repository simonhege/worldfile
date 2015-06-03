// Copyright 2015 Simon HEGE. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
Package worldfile allows to deals with world file (eg. tfw, pgw or jgw files).

See http://en.wikipedia.org/wiki/World_file for documentation on the file format.
*/
package worldfile

import (
	"bufio"
	"image"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

//WorldFile allows converting between pixels and map coordinates.
//It is not dependent of a specific SRS.
type WorldFile struct {
	A float64 //pixel size in the x-direction in map units/pixel
	D float64 //rotation about y-axis
	B float64 //rotation about x-axis
	E float64 //pixel size in the y-direction in map units, almost always negative
	C float64 //x-coordinate of the center of the upper left pixel
	F float64 //y-coordinate of the center of the upper left pixel
}

//round rounds the float64 to the neareast integer
func round(f float64) int {
	if f < 0 {
		return int(math.Ceil(f - 0.5))
	}
	return int(math.Floor(f + 0.5))
}

//readFloat64 reads and parses a float64 from the buffered reader
func readFloat64(reader *bufio.Reader) (float64, error) {
	str, err := reader.ReadString('\n')
	if err != nil {
		return 0., err
	}
	return strconv.ParseFloat(strings.TrimSpace(str), 64)
}

//Read decodes a world file (6 lines, each one containing a single floating point value)
//See http://en.wikipedia.org/wiki/World_file for details on the accepted file format
func Read(r io.Reader) (WorldFile, error) {
	reader := bufio.NewReader(r)

	var wf WorldFile
	values := []*float64{&wf.A, &wf.D, &wf.B, &wf.E, &wf.C, &wf.F}

	var err error
	for _, v := range values {
		*v, err = readFloat64(reader)
		if err != nil {
			return WorldFile{}, err
		}
	}

	return wf, nil
}

//ReadFile decodes a world file at a given path
func ReadFile(path string) (WorldFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return WorldFile{}, err
	}
	defer file.Close()

	return Read(file)
}

//ToMap transforms an image point into map coordinates
func (wf WorldFile) ToMap(p image.Point) (float64, float64) {
	fX := float64(p.X)
	fY := float64(p.Y)

	x := wf.A*fX + wf.B*fY + wf.C
	y := wf.D*fX + wf.E*fY + wf.F

	return x, y
}

//FromMap transforms map coordinates into an image point
func (wf WorldFile) FromMap(x float64, y float64) image.Point {

	denom := (wf.A*wf.E - wf.D*wf.B)

	fx := (wf.E*x - wf.B*y + wf.B*wf.F - wf.E*wf.C) / denom
	fy := (-wf.D*x + wf.A*y + wf.D*wf.C - wf.A*wf.F) / denom

	return image.Pt(round(fx), round(fy))
}
