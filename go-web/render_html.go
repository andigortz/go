package main 

import "fmt"
import "net/http"
import "html/template"
import "path"

func main() {

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var filepath = path.Join("views", "site.html")
		var tmpl, err = template.ParseFiles(filepath)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data = map[string]interface{} {
			"title" : "Learning Golang",
			"body" : "Learning Go-web",
		}

		err = tmpl.Execute(w, data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)

}