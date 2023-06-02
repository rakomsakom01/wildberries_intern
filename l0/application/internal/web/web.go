package web

import (
    "net/http"
    "fmt"
    "log"
)


type WebServer struct {
    h handler

    mux *http.ServeMux

	infoLog  *log.Logger
	errorLog *log.Logger
}


func (w *WebServer) Start() {
    fmt.Println("http://localhost:8000/")
    http.ListenAndServe("localhost:8000", w.mux)
}


func InitWeb(cacheListRequest chan<- int, cacheListResult <-chan []int, 
        cacheValRequest chan<- int, cacheValResult <-chan map[string]interface{},
        infoLog *log.Logger, errorLog *log.Logger) WebServer {

	mux := http.NewServeMux()

    h := handler{
        cacheListRequest: cacheListRequest,
        cacheListResult : cacheListResult,
        cacheValRequest : cacheValRequest,
        cacheValResult  : cacheValResult,
	    infoLog         : infoLog,
        errorLog        : errorLog,
    }

    mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static"))))
   
    mux.HandleFunc("/", h.home)
    mux.HandleFunc("/query/", h.query)
    mux.HandleFunc("/parse/", h.parse)
    mux.HandleFunc("/json/", h.json)

    return WebServer{
        h: h,
        mux: mux,
 	    infoLog         : infoLog,
        errorLog        : errorLog,       
    }
}
