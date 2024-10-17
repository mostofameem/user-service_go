package db

import (
	"base_service/logger"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
)

type CartList struct {
	ProductName string `db:"product_name" json:"product_name"`
	Price       int    `db:"price" json:"price"`
	Quantity    int    `db:"quantity" json:"quantity"`
}

type Cart struct {
	ProductName string `json:"product_name" validate:"required,alpha"`
	Quantity    int    `json:"quantity" validate:"required"`
}
type CartTypeRepo struct {
	table string
}

var cartTypeRepo *CartTypeRepo

func initCartTypeRepo() {
	cartTypeRepo = &CartTypeRepo{
		table: "carts",
	}
}

func GetCartTypeRepo() *CartTypeRepo {
	return cartTypeRepo
}

func (r *CartTypeRepo) GetCart(id int, ch chan []CartList) {

	var AllProduct []CartList

	query, args, err := GetQueryBuilder().
		Select("product_name", "price", "quantity").
		From(r.table).
		Where(sq.Eq{"user_id": id}).
		OrderBy("quantity DESC").
		ToSql()
	if err != nil {
		slog.Error(
			"Failed to create resource types select query getcart",
			logger.Extra(map[string]any{
				"error": err.Error(),
				"query": query,
				"args":  args,
			}),
		)
		ch <- []CartList{}

	}

	err = GetDB().Select(&AllProduct, query, args...)
	if err != nil {
		slog.Error(
			"Failed to get resource types getcart",
			logger.Extra(map[string]any{
				"error": err.Error(),
			}),
		)
		ch <- []CartList{}
	}
	ch <- AllProduct
}
