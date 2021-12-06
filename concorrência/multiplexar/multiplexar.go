package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func titulo(urls ...string) <-chan string {
	c := make(chan string, 2)
	for _, url := range urls {
		go func(url string) { //função anônima
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url) //invocando essa função
	}
	return c
}

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

//multiplexar: misturar (mensagens) num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		titulo("https://www.cod3r.com.br", "https://www.google.com"),
		titulo("https://www.udemy.com", "https://www.youtube.com"),
	)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
