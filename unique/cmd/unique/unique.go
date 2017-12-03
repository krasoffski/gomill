package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

var tpl = `package {{.Package}}

func Dedup(s []{{.Type}}) []{{.Type}} {
	lst := make([]{{.Type}}, 0, 0)
	set := make(map[{{.Type}}]struct{})
	for _, i := range s {
		_, ok := set[i]
		if ok {
			continue
		}
		set[i] = struct{}{}
		lst = append(lst, i)
	}
	return lst
}

`

func main() {
	tt := template.Must(template.New("unique").Parse(tpl))
	for i := 1; i < len(os.Args); i++ {
		dest := strings.ToLower(os.Args[i]) + "_unique.go"
		file, err := os.Create(dest)
		if err != nil {
			fmt.Printf("Unable to create %s: %s (skip)\n", dest, err)
			continue
		}

		vals := map[string]string{
			"Type":    os.Args[i],
			"Package": os.Getenv("GOPACKAGE"),
		}
		tt.Execute(file, vals)
		file.Close()
	}
}
