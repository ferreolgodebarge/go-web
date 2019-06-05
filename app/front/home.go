package front

import (
	"net/http"
	"text/template"
)

// HomeData is a struct containing data to be injected in the Home template
type HomeData struct {
	PageTitle string
}

// HomeHandler is the main Handler serving template
// which contains 3 functionalities :
//	- contact a database and manage tables, and records
//	- produce and consumme rabbit messages
//	- import and list files in s3 bucket
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	data := HomeData{
		PageTitle: "My title",
	}
	tmpl.Execute(w, data)
}
