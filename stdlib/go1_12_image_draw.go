// +build go1.12,!go1.13

package stdlib

// Code generated by 'goexports image/draw'. DO NOT EDIT.

import (
	"image/draw"
	"reflect"
)

func init() {
	Value["image/draw"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Draw":           reflect.ValueOf(draw.Draw),
		"DrawMask":       reflect.ValueOf(draw.DrawMask),
		"FloydSteinberg": reflect.ValueOf(&draw.FloydSteinberg).Elem(),
		"Over":           reflect.ValueOf(draw.Over),
		"Src":            reflect.ValueOf(draw.Src),

		// type definitions
		"Drawer":    reflect.ValueOf((*draw.Drawer)(nil)),
		"Image":     reflect.ValueOf((*draw.Image)(nil)),
		"Op":        reflect.ValueOf((*draw.Op)(nil)),
		"Quantizer": reflect.ValueOf((*draw.Quantizer)(nil)),
	}
}