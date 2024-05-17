package main 
import (
  "fmt"
  "math/rand"
	"time"
  "log"
  "net/http"
  "github.com/julienschmidt/httprouter"
  "html/template"
  "e_commers/src"
  "e_commers/massage"
  "strconv"
	//"embed"
  )
/*
//go:embed views/*.html
var tmpFS embed.FS
var myTmplates = template.Must(template.ParseFS(tmpFS, "views/*.html")) */


  // batas proklamasi
  var GlobalCodeRandom int
  
  func randomNumber() int {
    return rand.Intn(900000) + 100000
  }
  
  func buatCookie(w http.ResponseWriter, id interface{}) {
      cookie := http.Cookie{
        Name:     "X-NM",
        Value:    strconv.Itoa(id.(int)),
        Path:     "/",
        HttpOnly: true, // Setel cookie sebagai HttpOnly
      }
     http.SetCookie(w, &cookie)
  }
  
  func htmlParsig(w http.ResponseWriter, nmFile string, nilai interface{}) {
    myTmplates := template.Must(template.ParseFiles("views/layouts/header.html","views/layouts/footer.html","views/layouts/nav.html","views/" + nmFile)) 
    myTmplates.ExecuteTemplate(w, nmFile, nilai)
  }
  
  func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    
    cookie, err := r.Cookie("X-NM")
    if err != nil {
      fmt.Println("tidak ada cookie")
      htmlParsig(w, "home.html", nil)
      return
    }
    
      hasil := src.ValidasiCookie(cookie.Value)
    if hasil {
      data := src.BacaData()
      htmlParsig(w, "Admin.html", data)
    }else {
      htmlParsig(w, "home.html",nil)
    }
  }
  
  func NotFound(w http.ResponseWriter, r *http.Request) {
   // fmt.Fprintln(w, "hallo", nil)
   http.Error(w, "404 Not Found", http.StatusNotFound)
  }
  
  func Search(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    
	nama := r.URL.Query().Get("srch")
	
	data := map[string]interface{}{
	  "Nama" : nama,
	}
	htmlParsig(w, "Search.html", data);
//	fmt.Fprint(w, nama, "typedata")
  }
  
  func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    htmlParsig(w, "Login.html", nil)
  }
  
  func LoginAkun(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
   email := r.URL.Query().Get("email")
   hasil := src.CariEmail(email)
   
   var text string
   
   if hasil == true {
     
    password := r.URL.Query().Get("password")
    id, hsl := src.CariPassword(password)
    if hsl == true {
      buatCookie(w, id)
      http.Redirect(w, r, "/", http.StatusSeeOther)
    } else {
      text = "password salah"
      htmlParsig(w, "Login.html", text)
    }
    
   } else {
     text = "email salah pasword salah"
     htmlParsig(w, "Login.html", text)
   }
  }
  
  // form akun
  func CreateAccount(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    
	rand.Seed(time.Now().UnixNano())
	random := randomNumber()
    
  data := map[string]interface{}{
	  "Id" : random,
	  "Status" : "user",
	}
	htmlParsig(w, "create.html", data);
  }
  
  // pembuatan akun
  func CreateAkun(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    // Mendapatkan data dari body request
        err := r.ParseForm()
        if err != nil {
            http.Error(w, "Gagal memproses form", http.StatusBadRequest)
            return
        }
        
    Name := r.Form.Get("Name")
    Email := r.Form.Get("Email")
    Password := r.Form.Get("Password")
    Id := r.Form.Get("Id")
    Status := r.Form.Get("Status")
     
     
    htmlParsig(w, "Ferivikasi.html", nil)
  
 go func() {
    err = src.TulisDataUser(Name, Email, Password, Id, Status)
     if err != nil {
          http.Error(w, "Gagal menulis data", http.StatusInternalServerError)
          return
     }
  
     GlobalCodeRandom = randomNumber()
     fmt.Println("sorkat bro :" , GlobalCodeRandom)
    
     err = massage.EmailOTP(Email, GlobalCodeRandom)
     if err != nil {
        http.Error(w, "Gagal mengirim email", http.StatusInternalServerError)
        return
      }
  }()
  
  }
  
  func Ferifikasi(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    
    cod :=  r.URL.Query().Get("Code")
  
  // Konversi string menjadi int
    codInt, err := strconv.Atoi(cod)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
  if codInt == GlobalCodeRandom {
    
    hasil := src.BacaData()
  //  fmt.Println("hasil mag", hasil)
    htmlParsig(w,"Admin.html", hasil)
    
    //fmt.Println(hasil)
    // htmlParsig(w,"Admin.html", hasil)
  } else {
     dat := map[string] interface{} {
      "Error" : "Kamu Memasuka Code Ferifikasi Yang Salah",
    }
    htmlParsig(w, "Ferivikasi.html", dat)
  }
  
  fmt.Println(codInt, GlobalCodeRandom)
  }
  
  func DetailData(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    err := r.ParseForm()
        if err != nil {
            http.Error(w, "Gagal memproses form", http.StatusBadRequest)
            return
        }
        
    id := r.Form.Get("Detail")
    iD, err := strconv.Atoi(id)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
   hasil := src.CariId(iD)
   fmt.Println(hasil)
   fmt.Println(hasil.Email)
   htmlParsig(w, "detail.html", hasil)
  }
 
  
  func main() {
    app := httprouter.New()
    
    staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("../Public")));
	app.GET("/static/*filepath", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		staticHandler.ServeHTTP(w, r)
	})
    
    
    
    app.GET("/", Index)
    app.GET("/search", Search)
    app.GET("/login", Login)
    app.GET("/loginakun", LoginAkun)
    app.GET("/create", CreateAccount)
    app.POST("/crtaccount", CreateAkun)
    app.GET("/verify", Ferifikasi)
    app.POST("/detail", DetailData)
    app.NotFound = http.HandlerFunc(NotFound)
    log.Fatal(http.ListenAndServe(":8080", app))
  }