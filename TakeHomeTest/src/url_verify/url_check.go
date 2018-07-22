package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
	"net/url"    
	"time"
	"os"
	"bufio"
	"encoding/json"
	"sync"
)

var mutex = &sync.Mutex{}

type Urldetail struct {
	id       int
	URL      string
	Status   string	
}

type PageVariables struct {  
  Urldetails       []Urldetail
}

func isValidUrl(toTest string) string {
    _, err := url.ParseRequestURI(toTest)
    if err != nil {
        return "false"
    } else {
        return "true"
    }
}


func SiteStatus(toTest string) string {	
	timeout := time.Duration(800 * time.Millisecond)
	client := &http.Client{
		Timeout: timeout,
	}	
	resp, err := client.Get(toTest)	
    if (err == nil) && (resp.StatusCode == 200) {
        return "UP"
    } else {
        return "DOWN"
    }    
}

func WriteToFile(site_name string) {  
    mutex.Lock()   
    file, err := os.OpenFile("sites.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
	    mutex.Unlock();
        log.Fatalf("failed opening file: %s", err)
    }
    defer file.Close()
 
    len, err := file.WriteString(site_name)
    if err != nil {
	    mutex.Unlock();
        log.Fatalf("failed writing to file: %s", err)
    }
	mutex.Unlock();
    fmt.Printf("\nLength: %d bytes", len)
    fmt.Printf("\nFile Name: %s", file.Name())
}

func readLines(path string) ([]string, error) {
  mutex.Lock();		
  file, err := os.Open(path)
  if err != nil {
    mutex.Unlock();
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())	
  } 
  mutex.Unlock(); 
  return lines, scanner.Err()
}

func delete_url(w http.ResponseWriter, r *http.Request){
	line_to_del := r.FormValue("del_line")
	//fmt.Println("delurl:", line_to_del)
	lines, err := readLines("sites.txt")
	mutex.Lock();
	err = os.Remove("sites.txt")	
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		mutex.Unlock();
		return
	}
	mutex.Unlock();
	var msg string
	for _,line := range lines {
		 if(line_to_del != line){
			nn:= line + "\n"
			WriteToFile(nn)
		 }
		}
	msg = "Successfully deleted: " + line_to_del
	addmsgBytes, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}	
	w.Write(addmsgBytes)
}

func get_urls(w http.ResponseWriter, r *http.Request){

   MyUrlDetails:= []Urldetail{}   
   lines, err := readLines("sites.txt")
	for i, line := range lines {
		var kk = i;
		online_status := SiteStatus(line)
		MyUrlDetails = append(MyUrlDetails,Urldetail{			
			id:   kk,
			URL: line,
			Status: online_status,
			})
		}		
	urlListBytes, err := json.Marshal(MyUrlDetails)
	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}	
	w.Write(urlListBytes)	  
  }
  
  func check_duplicate(url_to_check string) string {	
	lines, err := readLines("sites.txt")    
	for _,line := range lines {
	     //fmt.Printf("Compare: |%s|%s|\n",url_to_check,line)
		 if(url_to_check == line){
			return "true"
		 }		
    }	
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		return "false"
	}  
	return "false"	
}

func add_url(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
	var msg string
    if r.Method == "GET" {	    
        t, _ := template.ParseFiles("url_check.html")
        t.Execute(w, nil)		
    } else {
        r.ParseForm()
		//fmt.Println("check:", r.FormValue("check")) //get request method
		url_exists := check_duplicate(r.FormValue("check"))		
		if(url_exists == "true"){
		  msg = "Error: URL exists"
		} else {
			status := isValidUrl(r.FormValue("check"))
			if(status == "true"){					
				nn:= r.FormValue("check") + "\n"
				WriteToFile(nn)
				msg = "Valid URL"			
			} else{
				msg = "InValid URL"
			}
	    }			
    }	
	addmsgBytes, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}	
	w.Write(addmsgBytes)	
}

func main() {
    http.HandleFunc("/", get_urls) // setting router rule	
    http.HandleFunc("/add_url", add_url)
	http.HandleFunc("/delete_url", delete_url)
    err := http.ListenAndServe(":9090", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}