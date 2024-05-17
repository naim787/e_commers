package main 

import (
  "fmt"
  "net/http"
  "github.com/julienschmidt/httprouter"
  )
  
  type LogMiddleware struct {
    Hanler http.Handler
  }
  
  func(middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Println(" sebelum exekusi middle whare")
    middleware.Hanler.ServeHTTP(w, r)
    fmt.Println(" sesuda exekusi middle whare")
  }
  
  func main() {
    app := httprouter.New()
    
    app.GET("/", func(w http.ResponseWriter, r *http.Request,_ httprouter.Params){
      fmt.Println("hallo middleware")
      fmt.Fprintf(w, "hallo middleware")
    })
    
    LogMid := new(LogMiddleware)
    LogMid.Hanler = app
    
    http.ListenAndServe(":8080", LogMid)
  }