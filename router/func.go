package router

import (
	"github.com/guitarpawat/bookshelf/util"
	"html/template"
)

func getFuncMap() template.FuncMap {
	return template.FuncMap{
		"iterateStringList": util.StringSliceToString,
	}
}
