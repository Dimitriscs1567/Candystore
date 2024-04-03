package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"tikasdimitrios/candystore/models"
	"tikasdimitrios/candystore/parsers"

	"golang.org/x/net/html"
)

func main(){
	resp, err := http.Get("https://candystore.zimpler.net/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
        panic(err)
    }

	var customers []models.Customer
	parsers.ProcessHtmlNode(doc, &customers)
	sort.Slice(customers, func(i, j int) bool {
		return customers[i].TotalSnacks > customers[j].TotalSnacks
	})

	jsonResult, err := json.MarshalIndent(customers, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonResult))
}
