package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somhwerintheinternet.com/",
		"https://graph.microsoft.com",
	}

	//creando el canal
	ch := make(chan string)

	for _, api := range apis {
		go checkAPI(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Println(<-ch)
	}
	time.Sleep(5 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("¡LISTO! ¡TOMO %v Segundos! \n", elapsed.Seconds())
}

//supervisamos que los enlaces estén habilitados o en funcionamiento

func checkAPI(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		ch <- fmt.Sprintf("ERROR: %s está caído!! \n ", api)
		return
	}

	ch <- fmt.Sprintf("SUCCESS: %s Está en funcionamiento!! \n", api)
}
