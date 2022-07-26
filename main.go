package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://ergast.com/api/f1/2021/drivers.json")
	bs := make([]byte, 99999)

	if err == nil {
		resp.Body.Read(bs)
		fmt.Println(string(bs))
	} else {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
