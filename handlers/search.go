package handlers

import (
	"net/http"
	"database/sql"
	"html/template"
	"log"
	
	"github.com/Ahdeyyy/spotify2018Frontend/database"
	"github.com/Ahdeyyy/spotify2018Frontend/models"
)

type Search struct {
	Db *sql.DB
}

func (s *Search) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Credentials", "true")
    w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
    w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	query := r.FormValue("artist")
	songs := db.FindArtist(s.Db, query)
	allFiles := []string{ "footer.tmpl", "header.tmpl", "page.tmpl"}

    var allPaths []string
    for _, tmpl := range allFiles {
        allPaths = append(allPaths, "./templates/"+tmpl)
    }
	allPaths = append(allPaths, "./templates/search/content.tmpl")

    templates := template.Must(template.New("").Funcs(template.FuncMap{ "msToMin" : msToMin ,"addOne" : addOne}).ParseFiles(allPaths...))
	title := "Search results for " + query

	data := struct {
		Title string
		Songs []models.Song
		Query string
	}{
		Title: title,
		Songs: songs,
		Query: "Search results for " + query,
	}

    templates.ExecuteTemplate(w, "page", data)
	log.Println(r.Method,r.URL,r.Proto, http.StatusOK)

}
