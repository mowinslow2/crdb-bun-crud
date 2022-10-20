package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	model "main.go/models"
)

func main() {
	ctx := context.Background()

	// Connection String
	addr := "your/connection/string/here" //TODO: UPDATE THIS CONNECTION STRING
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(addr)))

	// Bun!
	db := bun.NewDB(sqldb, pgdialect.New())

	// To see executed queries in stdout
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))

	createCustomerTable(ctx, db)
	insertCustomer(ctx, db)
	updateCustomer(ctx, db)
	selectCustomer(ctx, db)
	deleteCustomer(ctx, db)
	dropCustomerTable(ctx, db)

	defer db.Close()
}

func createCustomerTable(ctx context.Context, db *bun.DB) {
	_, err := db.NewCreateTable().Model((*model.Customer)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully created customers table")
	}
}

func insertCustomer(ctx context.Context, db *bun.DB) {
	customer := &model.Customer{Name: "Morgan"}

	_, err := db.NewInsert().Model(customer).Exec(ctx)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Sucessfully inserted to the customers table")
	}
}

func updateCustomer(ctx context.Context, db *bun.DB) {
	customer := &model.Customer{Name: "Winslow"}
	_, err := db.NewUpdate().Model(customer).Where("name = ?", "Morgan").Exec(ctx)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Sucessfully updated the customers table")
	}
}

func selectCustomer(ctx context.Context, db *bun.DB) {
	var customers []model.Customer
	err := db.NewSelect().Model(&customers).Scan(ctx)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(customers[0].Name)
	}
}

func deleteCustomer(ctx context.Context, db *bun.DB) {
	_, err := db.NewDelete().Model(&model.Customer{}).Where("name = ?", "Winslow").Exec(ctx)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully deleted from the customers table")
	}
}

func dropCustomerTable(ctx context.Context, db *bun.DB) {
	_, err := db.NewDropTable().Model((*model.Customer)(nil)).IfExists().Exec(ctx)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully dropped the customers table")
	}
}
