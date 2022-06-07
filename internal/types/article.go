package types

import (
	"net/http"
	"time"
)

// Article is one instance of an article
type Article struct {
	ID    int     `gorm:"type:SERIAL;PRIMARY_KEY" json:"id" example:"1"`
	Name  string  `gorm:"type:varchar;NOT NULL" json:"name" example:"Skittles"`
	Price float64 `gorm:"type:decimal;NOT NULL" json:"price" example:"1.99"`
} // @name Article

// Render implements the github.com/go-chi/render.Renderer interface
func (a *Article) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Bind implements the the github.com/go-chi/render.Binder interface
func (a *Article) Bind(r *http.Request) error {
	return nil
}

// ArticleList contains a list of articles
type ArticleList struct {
	// A list of articles
	Items []*Article `json:"items"`
	// The id to query the next page
	NextPageID int `json:"next_page_id,omitempty" example:"10"`
} // @name ArticleList

// Render implements the github.com/go-chi/render.Renderer interface
func (a *ArticleList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Order is one instance of an order
type Order struct {
	ID       int       `gorm:"type:SERIAL;PRIMARY_KEY" json:"id" example:"1"`
	DateTime time.Time `gorm:"timestamp" json:"lastUpdated,omitempty" example:"0001-01-01 00:00:00+00"`
} // @name Order

// Render implements the github.com/go-chi/render.Renderer interface
func (o *Order) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Bind implements the the github.com/go-chi/render.Binder interface
func (o *Order) Bind(r *http.Request) error {
	return nil
}

// OrderList contains a list of orders
type OrderList struct {
	// A list of orders
	Items []*Order `json:"items"`
	// The id to query the next page
	NextPageID int `json:"next_page_id,omitempty" example:"10"`
} // @name OrderList

// Render implements the github.com/go-chi/render.Renderer interface
func (o *OrderList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
