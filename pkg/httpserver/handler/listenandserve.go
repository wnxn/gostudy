package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type dollars float32

type database map[string]dollars

var dataList = template.Must(template.New("datalist").Parse(`
<table>
<tr style='text-align: left'>
	<th>item</th>
	<th>price</th>
</tr>
{{range $i,$v:= .}}
<tr>
	<td>{{print $i}}</a></td>
	<td>{{print $v}}</a></td>
</tr>
{{end}}
</table>
`))

func(d dollars)String()string{
	return fmt.Sprintf("$%.2f", d)
}

func main() {
	db := database{"shoes":90, "socks":20}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price",db.price)
	http.HandleFunc("/update",db.update)
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}

func (db database)list(w http.ResponseWriter, req *http.Request){
	err := dataList.Execute(w, db)
	if err != nil{
		log.Fatal(err)
	}
}

func (db database)price(w http.ResponseWriter, req *http.Request){
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item %s\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database)update(w http.ResponseWriter, req *http.Request){
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	fmt.Fprintf(w, "item=%s, price=%s\n",item, price)
}