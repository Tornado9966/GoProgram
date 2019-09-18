package main

import ( 
	
)

func handleStart(resWr http.ResponseWriter, req *http.Request) {
	resWr.Write([]byte("Hello from server!"))
	fmt.Println("Page success")
}

func handleCurrentTime(resWr http.ResponseWriter, req *http.Request) {
	
}

func main() {
	// Request handles
	http.HandleFunc("/", handleStart)
	http.HandleFunc("/time", handleCurrentTime)
	// Deffered call
	defer http.ListenAndServe(":8795", nil)
	// Server status message
	fmt.Println("Server started successfully and listen on 8795 port"
}
