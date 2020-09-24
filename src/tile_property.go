package ttme

import r "github.com/lachee/raylib-goplus/raylib"

type tileProperty struct {
	Name  string  `json:"name"`
	Value string  `json:"value"`
	Color r.Color `json:"color"`
}