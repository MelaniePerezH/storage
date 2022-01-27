package product

import (
	"context"
	"database/sql"
	"log"

	"github.com/extlurosell/meli_bootcamp_go_w3-2/internal/domain"
)

// Repository encapsulates the storage of a Product.
type Repository interface {
	Get(ctx context.Context, id string) (domain.Product, error)
	Store(ctx context.Context, product domain.Product) (domain.Product, error)
	GetAll() ([]domain.Product, error)
	Update(ctx context.Context, p domain.Product) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(ctx context.Context, name string) (domain.Product, error) {
	query := "SELECT id,name,type,count,price FROM products WHERE name=?;"
	row := r.db.QueryRow(query, name)
	p := domain.Product{}
	err := row.Scan(&p.Id, &p.Name, &p.Tipo, &p.Count, &p.Price)
	if err != nil {
		return domain.Product{}, err
	}

	return p, nil
}

func (r *repository) Store(ctx context.Context, product domain.Product) (domain.Product, error) {

	stmt, err := r.db.Prepare("INSERT INTO products(name, type, count, price) VALUES( ?, ?, ?, ? )") // se prepara el SQL
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	var result sql.Result
	result, err = stmt.Exec(product.Name, product.Tipo, product.Count, product.Price) // retorna un sql.Result y un error
	if err != nil {
		return domain.Product{}, err
	}
	insertedId, _ := result.LastInsertId() // del sql.Resul devuelto en la ejecución obtenemos el Id insertado
	product.Id = int(insertedId)
	return product, nil
}

func (r *repository) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	db := r.db
	rows, err := db.Query("SELECT id,name,type,count,price FROM products")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// se recorren todas las filas
	for rows.Next() {
		// por cada fila se obtiene un objeto del tipo Product
		var product domain.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Tipo, &product.Count, &product.Price); err != nil {
			log.Fatal(err)
			return nil, err
		}
		//se añade el objeto obtenido al slice products
		products = append(products, product)
	}
	return products, nil
}
func (r *repository) Update(ctx context.Context, p domain.Product) error {
	query := "UPDATE products SET Tipo=?, Count=?, Price=?  WHERE Name=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, query, p.Name)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
