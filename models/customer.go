package models

type Customer struct {
	Name string `json:"name"`
	FavouriteSnack string `json:"favouriteSnack"`
	TotalSnacks int `json:"totalSnacks"`
}