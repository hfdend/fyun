package main

import (
	_ "github.com/hfdend/fyun/init"
	"github.com/hfdend/fyun/object"
	"fmt"
	"net/http"
	"log"
	"text/template"
	"github.com/hfdend/fyun/g"
)

var HTML = fmt.Sprintf(`<!DOCTYPE html>
<html lang="zh">
<head>
<meta charset="utf-8">
<style>
td {
    background: #FFF;
    align-content: center;
    text-align: center;
}

tr {
    margin: 0;
    border: 0;
    padding: 0;
}

table {
    width: 100%%;
}

div {
    width: 800px;
    margin: auto;
    background: #ccc;
}
</style>
</head>
<body>
<div>
<table>
	<tr>
		<td>文件</td>
		<td>类型</td>
		<td>大小</td>
	</tr>
	{{range .}}
	<tr>
		<td><a href="{{if .List}}{{.Name}}{{else}}%s{{.Path}}{{end}}">{{.Name}}</a></td>
		<td>{{if .List}}目录{{else}}文件{{end}}</td>
		<td>{{.Size}}</td>
	</tr>
	{{end}}
</table>
<div>
</body>
</html>
`,
	g.CdnUrl,
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestPath := r.URL.Path
		tree, err := object.GetTree()
		if err != nil {
			log.Println(err)
			http.NotFound(w, r)
			return
		}
		var list []*object.Object
		if requestPath == "/" {
			list = tree.List
		} else {
			var ok bool
			var obj *object.Object
			if obj, ok = tree.GetObject(requestPath); !ok {
				http.NotFound(w, r)
				return
			} else if len(obj.List) == 0 {
				return
			} else {
				list = obj.List
			}
		}
		tpl := template.Must(template.New("").Parse(HTML))
		tpl.Execute(w, list)
	})
	log.Fatalln(http.ListenAndServe(g.Addr, nil))
}
