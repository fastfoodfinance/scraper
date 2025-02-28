package internal

type Source string

type Menu struct {
	Source     Source     `json:"source"`
	Restaurant Restaurant `json:"restaurant"`
	Items      []Item     `json:"items"`
}

type Restaurant struct {
	Name    string `json:"name"`
	LogoUrl string `json:"logo_url"`
}

type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	Price       Price  `json:"price"`
}

type Price struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}
