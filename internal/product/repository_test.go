package product

import (
	"context"
	"database/sql"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/extlurosell/meli_bootcamp_go_w3-2/internal/domain"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	txdb.Register("txdb", "mysql", "meli_sprint_user:Meli_Sprint#123@/storage")
}

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("txdb", uuid.New().String())
	if err == nil {
		return db, db.Ping()
	}
	return db, err
}
func TestCreateOK(t *testing.T) {
	db, _ := sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/storage")
	newProduct := domain.Product{
		Name:  "melii",
		Tipo:  "apto",
		Count: 5,
		Price: 141,
	}

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	repo := NewRepository(db)
	res, _ := repo.Store(ctx, newProduct)

	assert.Equal(t, newProduct.Name, res.Name)
	assert.Equal(t, newProduct.Tipo, res.Tipo)
	assert.Equal(t, newProduct.Count, res.Count)
	assert.Equal(t, newProduct.Price, res.Price)
}

func TestGetByName(t *testing.T) {
	db, _ := sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/storage")
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	nombre := "melii"

	repo := NewRepository(db)
	res, _ := repo.Get(ctx, nombre)

	assert.Equal(t, nombre, res.Name)
	assert.IsType(t, domain.Product{}, res)
}

func TestGetAll(t *testing.T) {
	db, _ := sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/storage")
	data := []domain.Product{
		{Id: 1,
			Name:  "meli",
			Tipo:  "casa",
			Count: 50,
			Price: 50,
		},
		{
			Id:    4,
			Name:  "santi",
			Tipo:  "apto",
			Count: 5,
			Price: 141,
		}, {
			Id:    5,
			Name:  "melii",
			Tipo:  "apto",
			Count: 5,
			Price: 141,
		},
	}
	repo := NewRepository(db)
	res, _ := repo.GetAll()

	assert.Equal(t, data, res)
}

func Test_sqlRepository_Store(t *testing.T) {
	db, err := InitDb()
	assert.NoError(t, err)

	repository := NewRepository(db)
	ctx := context.TODO()
	data := domain.Product{
		Name:  "melii",
		Tipo:  "apto",
		Count: 5,
		Price: 4561,
	}
	res, err := repository.Store(ctx, data)
	assert.NoError(t, err)

	assert.Equal(t, data.Name, res.Name)
	assert.Equal(t, data.Tipo, res.Tipo)
	assert.Equal(t, data.Count, res.Count)
	assert.Equal(t, data.Price, res.Price)
}

func TestGetByNamemySQL(t *testing.T) {
	db, err := InitDb()
	assert.NoError(t, err)
	ctx := context.TODO()
	nombre := "melii"

	repo := NewRepository(db)
	res, _ := repo.Get(ctx, nombre)
	assert.Equal(t, nombre, res.Name)
	assert.IsType(t, domain.Product{}, res)
}
