package main
import (
	"net/http"
	"log"
)

type CustomServer struct {
	Mux *http.ServeMux
	Port string
	Debug bool
}

func NewCustomServer(port string, debug bool) *CustomServer {
	return &CustomServer{
		Mux: http.NewServeMux(),
		Port: port,
		Debug: debug,
	}
}

func (c *CustomServer) addHandler (url string,handler interface{}) {
	switch handler.(type) {
		case http.Handler:
			c.Mux.Handle(url, handler.(http.Handler))

		case func(http.ResponseWriter, *http.Request):
			c.Mux.HandleFunc(url, handler.(func(http.ResponseWriter, *http.Request)))

		default:
			panic("Handler must be of type http.Handler or func(http.ResponseWriter, *http.Request)")
	}
}

func (c *CustomServer) Start() {
	if c.Debug {
		log.Println("listening on port " + c.Port + ", press ctrl+c to stop")
		log.Println("http://localhost:" + c.Port)
	}
	go func() {
		if err := http.ListenAndServe(":" + c.Port, c.Mux); err != nil {
			log.Fatal(err)
		}
	}()
}

