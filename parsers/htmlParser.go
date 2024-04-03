package parsers

import (
	"errors"
	"strconv"

	"tikasdimitrios/candystore/models"

	"golang.org/x/net/html"
)

func ProcessHtmlNode(n *html.Node, customers *[]models.Customer) {
	isCutomerTable := false

	if n.Type == html.ElementNode && n.Data == "table"{
		for _, attr := range n.Attr {
			if attr.Key == "class" && attr.Val == "top.customers summary" {
				isCutomerTable = true
			}
		}

		if isCutomerTable{
			for el := n.FirstChild; el != nil; el = el.NextSibling {
				if el.Data == "tbody" {
					for c := el.FirstChild; c != nil; c = c.NextSibling {
						if processCustomerNode(c) != nil{
							*customers = append(*customers, *processCustomerNode(c))
						}
					}
					return
				}
			}
		}
	}

	if n.FirstChild != nil {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if len(*customers) > 0 {
				return
			}
			ProcessHtmlNode(c, customers)
		}
	}

	if n.NextSibling == nil {
		return
	}

	ProcessHtmlNode(n.NextSibling, customers)
}

func processCustomerNode(n *html.Node) *models.Customer {
	if n.Type == html.ElementNode && n.Data == "tr" {
		customer, err := newCustomerFromNode(n)
		if err != nil {
			return nil
		}

		return customer
	}

	return nil
}

func newCustomerFromNode(n *html.Node) (*models.Customer, error) {
	customerData := []string{}
	totalSnacks := 0

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "td" {
			if len(c.Attr) > 0 && c.Attr[0].Key == "x-total-candy" {
				var err error
				totalSnacks, err = strconv.Atoi(c.Attr[0].Val)
				if err != nil {
					return nil, errors.New("total snacks not an integer")
				}
			}
			customerData = append(customerData, c.FirstChild.Data)
		}
	}	

	return &models.Customer{
		Name: customerData[0],
		FavouriteSnack: customerData[1],
		TotalSnacks: totalSnacks,
	}, nil
}