// +build go1.12,!go1.13

package stdlib

// Code generated by 'goexports image/gif'. DO NOT EDIT.

import (
	"image/gif"
	"reflect"
)

func init() {
	Value["image/gif"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Decode":             reflect.ValueOf(gif.Decode),
		"DecodeAll":          reflect.ValueOf(gif.DecodeAll),
		"DecodeConfig":       reflect.ValueOf(gif.DecodeConfig),
		"DisposalBackground": reflect.ValueOf(gif.DisposalBackground),
		"DisposalNone":       reflect.ValueOf(gif.DisposalNone),
		"DisposalPrevious":   reflect.ValueOf(gif.DisposalPrevious),
		"Encode":             reflect.ValueOf(gif.Encode),
		"EncodeAll":          reflect.ValueOf(gif.EncodeAll),

		// type definitions
		"GIF":     reflect.ValueOf((*gif.GIF)(nil)),
		"Options": reflect.ValueOf((*gif.Options)(nil)),
	}
}