package ubereats

import (
	"testing"
)

func TestSeoFeedV1(t *testing.T) {
	pathname := pathname{brand: burgerKing, city: unitedStatesNewYork}
	menu, err := seoFeedV1(pathname)

	if err != nil {
		t.Errorf("expected err to be nil, got %v", err)
	}

	if menu.Restaurant.Name != "Burger King" {
		t.Errorf("expected menu.Restaurant.Name to be Burger King, got %s", menu.Restaurant.Name)
	}

	if len(menu.Items) == 0 {
		t.Errorf("expected menu.Items to have at least 1 item, got %d", len(menu.Items))
	}
}
