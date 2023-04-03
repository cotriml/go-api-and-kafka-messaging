package repository

import (
	"database/sql"

	"github.com/cotriml/go-api-and-kafka-messaging/internal/entity"
)

type ProductRepositoryMysql struct {
	DB *sql.DB
}

func NewProductRepositoryMySql(db *sql.DB) *ProductRepositoryMysql {
	return &ProductRepositoryMysql{DB: db}
}

func (r *ProductRepositoryMysql) Create(p *entity.Product) error {
	_, err := r.DB.Exec("INSERT INTO products (id, name, price) VALUES (?, ?, ?)", p.ID, p.Name, p.Price)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepositoryMysql) FindAll() ([]*entity.Product, error) {
	rows, err := r.DB.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}
