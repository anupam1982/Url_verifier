package main

import (
    "fmt"
    "net/http"
	"net/url"	
	"testing"
	"io/ioutil"
    "os"
	"encoding/json"
	"bytes"
	"log"
)


func Testadd_url(t *testing.T) {
response, err := http.PostForm("http://localhost:9090/add_url",url.Values{"check": {"https://www.mukund.com"}})
   if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))		
		return
	}
	defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
		var prettyJSON bytes.Buffer
		error := json.Indent(&prettyJSON, contents, "", "\t")
		if error != nil {
			log.Println("JSON parse error: ", error)			
			return
		}

    fmt.Printf("%s\n:", string(prettyJSON.Bytes()))   
}


func Testget_url(t *testing.T) {
response,err := http.Get("http://localhost:9090/get_urls")
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", string(contents))
    
	}
}