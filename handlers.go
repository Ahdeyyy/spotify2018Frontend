package main


import (
	"net/http"
	"database/sql"
	"encoding/json"
	"log"
	"html/template"
	"bytes"
	"os"
	"bufio"
	"strconv"
)


type Search struct {
	db *sql.DB
}

func msToMin(ms int) string {
	min := ms/60000
	sec := (ms%60000)/1000
	return strconv.Itoa(min) + ":" + strconv.Itoa(sec)
}

func addOne(index int) int {
	return index + 1
}

func (s *Search) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Credentials", "true")
    w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
    w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	query := r.FormValue("artist")
	songs := findArtist(s.db, query)
	allFiles := []string{ "footer.tmpl", "header.tmpl", "page.tmpl"}

    var allPaths []string
    for _, tmpl := range allFiles {
        allPaths = append(allPaths, "./templates/"+tmpl)
    }
	allPaths = append(allPaths, "./templates/search/content.tmpl")

    templates := template.Must(template.New("").Funcs(template.FuncMap{ "msToMin" : msToMin ,"addOne" : addOne}).ParseFiles(allPaths...))
    var processed bytes.Buffer
	title := "Search results for " + query

	data := struct {
		Title string
		Songs []Song
		Query string
	}{
		Title: title,
		Songs: songs,
		Query: "Search results for " + query,
	}

    templates.ExecuteTemplate(&processed, "page", data)

    outputPath := "./static/pages/index.html"
    f, _ := os.Create(outputPath)
    fw := bufio.NewWriter(f)
    fw.WriteString(string(processed.Bytes()))
    fw.Flush()
	f.Close()

	log.Println(r.Method,r.URL,r.Proto, http.StatusOK)
	http.ServeFile(w, r, "./static/pages/index.html")

}

type Index struct {
	db *sql.DB
}

func (i *Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	songs := findAll(i.db)
	res,err := json.Marshal(songs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	log.Println(r.Method,r.URL,r.Proto, http.StatusOK)
	w.Write(res)


}

type Home struct {
	db *sql.DB
}

func (h *Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	songs := findAll(h.db)

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
		Songs []Song
	}{
		Title: title,
		Songs: songs,
	}
    var processed bytes.Buffer
    templates.ExecuteTemplate(&processed, "page", data)

    outputPath := "./static/pages/index.html"
    f, _ := os.Create(outputPath)
    fw := bufio.NewWriter(f)
    fw.WriteString(string(processed.Bytes()))
    fw.Flush()
	f.Close()

	log.Println(r.Method,r.URL,r.Proto, http.StatusOK)
	http.ServeFile(w, r, "./static/pages/index.html")

}