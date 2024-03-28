package main
import (
//"github.com/ianlow27/translangopenweb/pkg1"
//"io/ioutil"
//"net/http"
"os"
"bufio"
"fmt"
"strings"
)
func main(){
  strvocab := 
`field|maes{f}
flower|blodeuyn{m}
fence|clawdd{m}`
var pr = fmt.Println
  arr1 := strings.Split(strvocab, "\n");
  gvocab := map[string]string {}
  for _, value := range arr1 {
    //pr(value)
    lstr1 := strings.Split(value, "|");
    //for _, value2 := range lstr1 {
      //pr(value2)
      gvocab[strings.TrimSpace(lstr1[0])] = strings.TrimSpace(lstr1[1])
      pr(lstr1[0] + "__" + gvocab[lstr1[0]])
    //}
  }
  reader := bufio.NewReader(os.Stdin)
  pr("Please enter some text:")
  text, _ := reader.ReadString('\n')
  //pr(text)
  arr2 := strings.Split(strings.TrimSpace(text), " ");
  text = ""
  for _, value := range arr2 {
    text +=  " " + gvocab[value]
  }
  pr("Welsh>" + text)
  /*
  gvocab := map[string]string {
    "field": "maes{f}",
    "flower": "blodeuyn{m}",
  }
  */
  //pr(gvocab["field"])
  //fmt.Println("This is a new project called 'translangopenweb'")
  //pkg1.Pkg1run()
  //-------------------------------
  /*
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
  */
  //-------------------------------
}
