package ubereats

import (
	"com/github/fastfoodfinance/scraper/internal"
	internalHttp "com/github/fastfoodfinance/scraper/internal/http"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type brand string

const (
	burgerKing brand = "burger-king"
)

type city string

const (
	unitedStatesNewYork city = "new-york-city"
)

type pathname struct {
	brand brand
	city  city
}

type seoFeedV1Response struct {
	Data struct {
		Elements []struct {
			Type          string `json:"type"`
			Title         string `json:"title"`
			LogoUrl       string `json:"logoUrl"`
			CurrencyCode  string `json:"currencyCode"`
			DishCarousels []struct {
				Dishes []struct {
					CatalogItem struct {
						Title           string `json:"title"`
						ItemDescription string `json:"itemDescription"`
						ImageUrl        string `json:"imageUrl"`
						Price           int    `json:"price"`
					} `json:"catalogItem"`
				} `json:"dishes"`
			} `json:"dishCarousels"`
		} `json:"elements"`
	} `json:"data"`
}

func seoFeedV1Request(pathname pathname) *http.Request {
	url := "https://www.ubereats.com/_p/api/getSeoFeedV1"

	body := fmt.Sprintf(`{"pathname": "/brand-city/%s/%s"}`, pathname.city, pathname.brand)
	bodyReader := strings.NewReader(body)

	req, err := http.NewRequest("POST", url, bodyReader)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-csrf-token", "x")

	return req
}

func seoFeedV1ResponseToMenu(response *seoFeedV1Response) (*internal.Menu, error) {
	restaurant := internal.Restaurant{}
	for _, element := range response.Data.Elements {
		if element.Type == "customEnterAddress" {
			restaurant.Name = strings.Replace(element.Title, " delivered to your door", "", 1)
			restaurant.LogoUrl = element.LogoUrl
		}
	}

	items := []internal.Item{}
	for _, element := range response.Data.Elements {
		if element.Type == "dishCarouselList" {
			for _, dishCarousel := range element.DishCarousels {
				for _, dish := range dishCarousel.Dishes {
					item := internal.Item{
						Name:        dish.CatalogItem.Title,
						Description: dish.CatalogItem.ItemDescription,
						ImageUrl:    dish.CatalogItem.ImageUrl,
						Price: internal.Price{
							Amount:   dish.CatalogItem.Price,
							Currency: element.CurrencyCode,
						},
					}
					items = append(items, item)
				}
			}
		}
	}

	menu := &internal.Menu{
		Restaurant: restaurant,
		Items:      items,
	}

	return menu, nil
}

func seoFeedV1(pathname pathname) (*internal.Menu, error) {
	request := seoFeedV1Request(pathname)
	response := &seoFeedV1Response{}
	err := internalHttp.DoJson(request, response)

	if err != nil {
		return nil, err
	}

	return seoFeedV1ResponseToMenu(response)
}
