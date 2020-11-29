package pg_storage

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ilhamrobyana/online-store-evermos-task/entity"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

var orderRequest = &entity.OrderCreateRequest{
	Items: []entity.Item{
		{
			ProductID: 1,
			Quantity:  5,
		},
	},
}

var order = &entity.Order{
	ID:     1,
	UserID: 1,
	Items:  6,
}

var product = &entity.Product{
	ID:            1,
	Name:          "T-Shirt",
	Inventory:     20,
	Price:         30000,
	Discount:      30.0,
	DiscountUntil: time.Now().Add(time.Hour),
}

var users = &entity.User{
	ID:       1,
	Username: "Ilham",
	Password: "12345678",
}

func NewMock() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New() // mock sql.DB
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gdb, err := gorm.Open("postgres", db) // open gorm db
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}
	return gdb, mock
}

func TestGetAll(t *testing.T) {
	client, mock := NewMock()
	orderStorage := &Order{client}
	defer client.Close()

	query := `SELECT \* FROM "orders" WHERE \(user_id\=\$1\)`

	rows := sqlmock.NewRows([]string{"id", "userID", "items"}).
		AddRow(order.ID, order.UserID, order.Items)
	mock.ExpectQuery(query).WithArgs(order.UserID).WillReturnRows(rows)

	orders, err := orderStorage.GetAll(order.UserID)
	assert.NotNil(t, orders)
	assert.NoError(t, err)
}

// Not yet done, it keeps returning Scan error even though it's not present on the acutal run of the code,
// Would implement a go routine to simulate race condition and i'm sure the actual code could handle race conditions
func TestCreate(t *testing.T) {
	client, mock := NewMock()
	orderStorage := &Order{client}
	defer client.Close()

	mock.ExpectBegin()

	query := `INSERT  INTO "orders" \("user_id"\,"items"\) VALUES \(\$1\,\$2\) RETURNING "orders"\."id"`
	rows := sqlmock.NewRows([]string{"id", "userID", "items"}).
		AddRow(order.ID, order.UserID, 0)
	mock.ExpectQuery(query).WithArgs(order.UserID, 0).WillReturnRows(rows)

	createdOrder, err := orderStorage.Create(orderRequest.Items, users.ID)
	fmt.Println(err)
	assert.NotNil(t, createdOrder)
	assert.NoError(t, err)
}
