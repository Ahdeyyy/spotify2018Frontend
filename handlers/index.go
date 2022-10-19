package handlers

import (
	"net/http"
	"database/sql"
	"encoding/json"
	"log"
	"html/template"

	"github.com/Ahdeyyy/spotify2018Frontend/models"
	"github.com/Ahdeyyy/spotify2018Frontend/database"
	
)



type Index struct {
	Db *sql.DB
}

type Home struct {
	Db *sql.DB
}

func (i *Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	songs := db.FindAll(i.Db)
	res,err := json.Marshal(songs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	log.Println(r.Method,r.URL,r.Proto, http.StatusOK)
	w.Write(res)


}


func (h *Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	songs := db.FindAll(h.Db)

    allFiles := []string{ "footer.tmpl", "header.tmpl", "page.tmpl"}

    var allPaths []string
    for _, tmpl := range allFiles {
        allPaths = append(allPaths, "./templates/"+tmpl)
    }
	allPaths = append(allPaths, "./templates/home/content.tmpl")

    templates := template.Must(template.New("").Funcs(template.FuncMap{ "msToMin" : msToMin , "addOne" : addOne }).ParseFiles(allPaths...))

	title := "Home"
	data := struct {
		Title string
		Songs []models.Song
	}{
		Title: title,
		Songs: songs,
	}
    
	log.Println(r.Method,r.URL,r.Proto, http.StatusOK)
    templates.ExecuteTemplate(w, "page", data)

}
