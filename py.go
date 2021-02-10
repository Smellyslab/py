package py

import "fmt"
import "net/http"


func Print(text string) {

  fmt.Println(text)

}

func GetRequest(host string) {
  
  output := http.Get(host)
  fmt.Println(output)
  
}

func PostRequest(host string) {

  output := http.Post(host)
  fmt.Println(output)
  
}
