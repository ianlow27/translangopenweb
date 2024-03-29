//Golang #15 Adnoddau Cymorth ReadFile len darllen ffeil 240329k 
//Golang #14 arallenw lluosoglinellarnod regexp Split ReplaceAllLiteralString paramedr 240329a 
//------------------------------------------------------------------------
package main
import (
//"github.com/ianlow27/translangopenweb/pkg1"
"io/ioutil"
//"net/http"
"regexp" //240329g-include regexp
"os"
"bufio"
"fmt"
"strings"
)
var pr = fmt.Println
var splt = strings.Split       //240329c - splt for brevity: 'alias'
var mtch = regexp.MatchString  //240329e - mtch for brevity
var rgxc = regexp.Compile      //240329h - rgxc for brevity
func main(){
  //!! Vocabulary Annotation Rules:
  //!! 1) All abbreviations must be followed by a dot
  //!! 2) The abbreviations are as follows: n-noun; f-feminine; m-masculine; adj-adjective; prp-preposition; 2m-2nd mutation;
  //!! 3) Data for language translation pairs are held in subdirectories under the 'langpairs' directory. The English and Welsh bi-directional data directory is 'cym_eng' not 'eng_cym' simply because 'c' precedes 'e' in the alphabet.  (243Tk) 

  strvocab :=  ""
  file1, err := ioutil.ReadFile("./langpairs/cym_eng/cym_eng.txt") //240329k ReadFile
  if(err == nil){
    strvocab =  string(file1)
  }
/* 243Tl
  strvocab := 
`field_n.s.|maes_n.f.s.
flower_n.s.|blodeuyn_n.m.s.
full_adj.|llawn_adj.
of_prp.|o_prp.2m.
fence_n.s.|clawdd_n.m.s.`           
*/ //240329l - moved data to cym_eng.txt ; 240329a-changed {f} to _nfs: 'multiline string'
  arr1 := strings.Split(strvocab, "\n");
  gvocab := map[string]string {}
  for _, value := range arr1 {
    lstr1 := strings.Split(value, "|")
    if(len(lstr1) <= 1){ continue }//240329m len() and continue
    lstr1a := strings.Split(lstr1[0], "_") //240329b-remove right of _: 'Regexp'
    gvocab[strings.TrimSpace(lstr1a[0])] = strings.TrimSpace(lstr1[1])
    pr(lstr1[0] + "__" + gvocab[lstr1a[0]])
  }
  reader := bufio.NewReader(os.Stdin)
  pr("Please enter some text:")
  text, _ := reader.ReadString('\n')
  //pr(text)
  arr2 := strings.Split(strings.TrimSpace(text), " ");
  text = ""
  prevlstr1_1 := ""
  for _, value := range arr2 { 
    lstr1 := splt(gvocab[value], "_")   //240329f-added splt: 'Split'
    bmtch, _ := mtch("2m.",prevlstr1_1);
    if(bmtch){ 
      text +=  " " + Rpl(lstr1[0], "^b", "f") //240329j-splt for brevity: 'ReplaceAllLiteralString'
    }else {
      text +=  " " + lstr1[0] //240329d-splt for brevity
    }
    prevlstr1_1 = lstr1[1]
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
func Rpl(pstr1 string, pstr2 string, pstr3 string)string { //240329i-unable to use alias: 'Parametres'
  var rgx1, _ = regexp.Compile(pstr2);
  return rgx1.ReplaceAllLiteralString(pstr1, pstr3)
} 
