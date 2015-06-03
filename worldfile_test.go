// Copyright 2015 Simon HEGE. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package worldfile

import (
	"image"
	"testing"
)

var dataToMap = []struct {
	wf WorldFile
	p  image.Point
	x  float64
	y  float64
}{
	{WorldFile{32.0, 0.0, 0.0, -32.0, 691200.0, 4576000.0}, image.Pt(171, 343), 696672.0, 4565024.0},
}

func TestToMap(t *testing.T) {

	for _, tt := range dataToMap {

		outX, outY := tt.wf.ToMap(tt.p)

		if outX != tt.x {
			t.Errorf("ToMap().x => %f, want %f", outX, tt.x)
		}
		if outY != tt.y {
			t.Errorf("ToMap().y => %f, want %f", outY, tt.y)
		}
	}
}

var dataFromMap = []struct {
	wf WorldFile
	x  float64
	y  float64
	p  image.Point
}{
	{WorldFile{32.0, 0.0, 0.0, -32.0, 691200.0, 4576000.0}, 696672.0, 4565024.0, image.Pt(171, 343)},
	{WorldFile{32.0, 0.0, 0.0, -32.0, 691200.0, 4576000.0}, 696687.0, 4565039.0, image.Pt(171, 343)},
	{WorldFile{32.0, 0.0, 0.0, -32.0, 691200.0, 4576000.0}, 696657.0, 4565009.0, image.Pt(171, 343)},
}

func TestFromMap(t *testing.T) {

	for _, tt := range dataFromMap {

		outP := tt.wf.FromMap(tt.x, tt.y)

		if outP != tt.p {
			t.Errorf("ToMap() => %v, want %v", outP, tt.p)
		}
	}
}
