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
	
}
