package web


import (
    "net/http"
    "html/template"
    "fmt"
    "strconv"
    "sync"
    "encoding/json"
    "log"
)

type handler struct {
    cacheListSync    sync.Mutex
    cacheListRequest chan<- int
    cacheListResult  <-chan []int
    
    cacheValSync     sync.Mutex
    cacheValRequest  chan<- int
    cacheValResult   <-chan map[string]interface{}

	errorLog *log.Logger
	infoLog  *log.Logger
}


func (h *handler) home(w http.ResponseWriter, r *http.Request) {
    tl, err := template.ParseFiles("ui/html/main.html", 
            "ui/html/header.html", "ui/html/footer.html")
        
    if err != nil {
        panic(err)
    }

    h.cacheListSync.Lock()
    h.cacheListRequest <- 0
    cacheList := <-h.cacheListResult
    h.cacheListSync.Unlock()

    
    defer tl.ExecuteTemplate(w, "main", cacheList)
}



func (h *handler) query(w http.ResponseWriter, r *http.Request) {
    tl, err := template.ParseFiles("ui/html/query.html", 
            "ui/html/header.html", "ui/html/footer.html")
        
    if err != nil {
        panic(err)
    }
    
    defer tl.ExecuteTemplate(w, "query", nil)
}


func (h *handler) parse(w http.ResponseWriter, r *http.Request) {
    query := r.FormValue("query")
    
    id, err := strconv.Atoi(query)

    if err != nil {
        id = -1
    }


    http.Redirect(w, r, fmt.Sprintf("/json/?id=%d", id), 301)
}


func (h *handler) json(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("ui/html/json.html", 
            "ui/html/header.html", "ui/html/footer.html")

    if err != nil {
        panic(err)
    }

    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        id = -1
    }

    h.cacheValSync.Lock()
    h.cacheValRequest <- id
    jsonMap := <-h.cacheValResult
    h.cacheValSync.Unlock()

    data, _ := json.MarshalIndent(jsonMap, "  ", "  ")

    tmpl.ExecuteTemplate(w, "json", string(data))
} 

