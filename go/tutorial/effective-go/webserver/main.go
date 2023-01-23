package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

var templ = template.Must(template.New("qr").Parse(templateStr))

func main() {
	// fmt.Printf("addr: %v\n", *addr)
	// fmt.Printf("templ: %v\n", templ)
	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
	err := http.ListenAndServe(*addr, nil)
	fmt.Println("http://localhost:1718/?s=Hello+World")
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func QR(w http.ResponseWriter, req *http.Request) {
	// templ.Execute(w, req.FormValue("s"))
	templ.Execute(w, struct {
		Text   string
		Number string
	}{Text: req.FormValue("s"), Number: req.FormValue("n")})
}

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .Text}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.Text}}" />
<br>
{{.Text}}
<br>
<br>
{{end}}
<p>{{.Number}}</p>
<form action="/" name=f method="GET">
    <input name=n value="{{.Number}}" title="Bad number">
    <input maxLength=1024 size=70 name=s value="{{.Text}}" title="Text to QR Encode">
    <input type=submit value="Show QR" name=qr>
</form>
</body>
</html>
`
