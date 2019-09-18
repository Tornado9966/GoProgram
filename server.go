package main

import ( 
	"fmt" 
	"net/http"
	"time" 
	"encoding/json"
	"log"
)

func handleStart(resWr http.ResponseWriter, req *http.Request) {
	
}

func handleCurrentTime(resWr http.ResponseWriter, req *http.Request) {
	jsonCur, err := json.Marshal(map[string]string{ "time": time.Now().Format(time.RFC3339) })
	if err != nil { // Error
		log.Println(err)
		return
	}
	fmt.Println("Time success")
	fmt.Println(time.Now().Format(time.RFC3339))
	resWr.Write(jsonCur) // Sending time
}

func main() {
	
}
