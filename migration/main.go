package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	"github.com/ilhamrobyana/online-store-evermos-task/entity"
	"github.com/ilhamrobyana/online-store-evermos-task/pg_storage"
)

func main() {
	client, err := pg_storage.GetPGClient()
	if err != nil {
		fmt.Println(err)
	}
	migrateScheme(client)
}

func migrateScheme(DB *gorm.DB) {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("can't load .env : %v", err))
	}

	isTableDropped := os.Getenv("DROP_TABLE")
	if isTableDropped == "true" {
		DB.Model(&entity.OrderProduct{}).RemoveForeignKey("order_id", "orders(id)")
		DB.Model(&entity.OrderProduct{}).RemoveForeignKey("product_id", "products(id)")
		DB.Model(&entity.Order{}).RemoveForeignKey("user_id", "users(id)")
		DB.DropTableIfExists(
			&entity.OrderProduct{},
			&entity.User{},
			&entity.Order{},
			&entity.Product{},
		)
	}

	DB.AutoMigrate(
		&entity.OrderProduct{},
		&entity.User{},
		&entity.Order{},
		&entity.Product{},
	)
	DB.Model(&entity.OrderProduct{}).AddForeignKey("order_id", "orders(id)", "RESTRICT", "RESTRICT")
	DB.Model(&entity.OrderProduct{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")
	DB.Model(&entity.Order{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}
