package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	websites := map[string]string{
		"Google": "https://google.com",
		"AWS":    "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["AWS"])

}
