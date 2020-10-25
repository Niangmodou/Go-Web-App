package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Welcome struct {
	Name string
	Time string
}

func main() {
	//Instantiate Welcome struct object
	welcome := Welcome{"Hello", time.Now().Format(time.Stamp)}

	//Locating html file
	templates := template.Must(template.ParseFiles("templates/index.html"))

	//Handling the / server route
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		if name := req.FormValue("name"); name != "" {
			welcome.Name = name
		}

		if err := templates.ExecuteTemplate(w, "index.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	//Starting the webserver to listen to port 8080
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))

}
