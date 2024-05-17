package src

import (
  "fmt"
  "encoding/json"
	"io/ioutil"
  "strconv"
 // "net/http"
  )
  
 type Core struct {
    Data []DataFild`json:"data"`
    Product []Product`json:"Product"`
 }

type DataFild struct {
  	Name     string `json:"Name"`
  	Email    string `json:"Email"`
  	Password string `json:"Password"`
  	Id       int    `json:"Id"`
  	Status   string `json:"Status"`
}

type Product struct {
  	Product_Name string `json:"Product_Name"`
  	Product_Id int `json:"Product_Id"`
  	Product_Bran string `json:"Product_Bran"`
  	Product_Price int `json:"Product_Price"`
  	Product_Descripsi string `json:"Product_Descripsi"`
}
  
func BacaData() Core {
    isiFile, err := ioutil.ReadFile("./data/data.json")
    if err != nil {
        panic("gagal membaca data")
    }
    per := Core{}
    err = json.Unmarshal(isiFile, &per)
    if err != nil {
        panic("error")
    }
    return per
}


func CariEmail(email string) bool {
  data := BacaData()
  for _, dat := range data.Data {
   if dat.Email == email {
     return true
   }
  }
  return false
}

func CariPassword(pass string) (interface{}, bool) {
  data := BacaData()
  for _, dat := range data.Data {
   if dat.Password == pass {
     return dat.Id, true
   }
  }
  return nil, false
}

func CariId(Id int) *DataFild{
  data := BacaData()
  for _, dat := range data.Data {
    if dat.Id == Id {
     result := DataFild(dat)
      return &result
    }
  }
  return nil
}

func ValidasiCookie(cookie string) bool{
  data := BacaData()
  cook, _ := strconv.Atoi(cookie)
  for _, dat := range data.Data {
    if dat.Id == cook {
      fmt.Println("ada cui" + dat.Status)
      if dat.Status == "admin" {
        return true
      }else {
        return false
      }
    }
  }
  return false
}




func TulisDataUser(nama, email, password, id, status string) error {
  
	dd, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	
	
 smuaData := BacaData()
 
	dat := DataFild{
	  Name: nama,
	  Email: email, 
	  Password: password, 
	  Id: dd,
	  Status: status,
	}
	
  smuaData.Data = append(smuaData.Data, dat)

 jsonData, err := json.Marshal(smuaData)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./data/data.json", jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}