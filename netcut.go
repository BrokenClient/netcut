package main

import(
	"encoding/json"
	."fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"net/url"
)

func mod(s string){
	cooked_msg:=url.QueryEscape(s)
	http.Post("https://netcut.cn/api/note/update/","application/x-www-form-urlencoded",strings.NewReader("note_id=xxx&note_content="+cooked_msg))//here replace xxx with your note_id by networking analyse
}
func main(){
	if(len(os.Args)>1){
		mod(os.Args[1])
	}else{
		var cut string
		con,_:=os.Stdin.Stat()
		if con.Mode().String()[0]=='p'{
			//&&con.Mode().Perm()==os.FileMode(0600){
			var buffer=make([]byte,256)
			for {
				n,err:=os.Stdin.Read(buffer[:])
				if err!=nil{
					break
				}
				cut+=string(buffer[:n])
			}
			mod(cut)
			os.Exit(0)
		}
		var v interface{}
		res,_:=http.Get("https://netcut.cn/api/note/data/?note_id=xxx")//here replace xxx with your note_id by networking analyse
		body,_:=ioutil.ReadAll(res.Body)
		json.Unmarshal([]byte(string(body)),&v)
		m := v.(map[string]interface{})
		Println(m["data"].(map[string]interface{})["note_content"])

	}
}

