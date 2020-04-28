package router

import "html/template"

func getFuncMap() template.FuncMap {
	return template.FuncMap{
		"iterateStringList": iterateStringList,
	}
}

func iterateStringList(in []string) string {
	res := ""
	for k, v := range in {
		if k != 0 {
			res += ", "
		}
		res += v
	}
	return res
}
