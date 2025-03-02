package ubereats

import (
	"log"

	"github.com/fastfoodfinance/scraper/internal"
)

const source internal.Source = "ubereats.com"

func Menus() []*internal.Menu {
	pathnames := []pathname{
		{brand: burgerKing, city: unitedStatesNewYork},
	}

	menus := []*internal.Menu{}
	for _, pathname := range pathnames {
		menu, err := seoFeedV1(pathname)
		if err != nil {
			log.Println(err)
			continue
		}

		menus = append(menus, menu)
	}

	return menus
}
