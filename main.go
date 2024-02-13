package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	var target string = os.Getenv("TARGET")

	for {
		resp, err := http.Get("http://" + target + "/movies")
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("%v\n", strings.ReplaceAll(string(b), "\n", ""))
		}
		time.Sleep(10 * time.Second)
	}

}
