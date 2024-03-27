package main
import "fmt"
import "github.com/ianlow27/translangopenweb/pkg1"
import (
"io/ioutil"
"net/http"
)
func main(){
  fmt.Println("This is a new project called 'translangopenweb'")
  pkg1.Pkg1run()
  //-------------------------------
  req, err := http.NewRequest(http.MethodGet,
    "http://localhost:8000", nil)
  if(err == nil){
    var response *http.Response
    response, err = http.DefaultClient.Do(req)
    if(err == nil){
      resp, _ := ioutil.ReadAll(response.Body)
      fmt.Println(string(resp))
    }else {
    }
  }
  //-------------------------------
}
