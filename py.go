package py

import "fmt"
import "net/http"


func Print(text string) {

  fmt.Println(text)

}

func GetRequest(host string) {
  
  http.Get(host)
  
  
}

func PostRequest(host string) {

  output := http.Post(host)
  
}
