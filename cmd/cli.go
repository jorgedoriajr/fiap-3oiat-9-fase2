package main

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v2"
	"hamburgueria/internal/modules/product/infra/database/postgres"
	"hamburgueria/internal/modules/product/usecase"
	command2 "hamburgueria/internal/modules/product/usecase/command"
	"hamburgueria/pkg/logger"
	"hamburgueria/pkg/sql"
	"hamburgueria/pkg/starter"
	"log"
	"os"
)

func main() {

	_ = context.TODO()
	starter.Initialize()
	sql.Initialize()

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:        "create-product",
				Usage:       "create new product",
				Description: "create new product",
				Action: func(c *cli.Context) error {

					productPersistence := postgres.NewProductRepository(
						sql.GetClient("readWrite"),
						sql.GetClient("readOnly"),
						logger.Get(),
					)

					productUseCase := usecase.NewCreateProductUseCase(productPersistence)

					command := command2.CreateProductCommand{
						Name:        "New Product",
						Amount:      10000,
						Description: "new product",
						Category:    "DISH",
						Menu:        true,
					}

					createProductResult, err := productUseCase.AddProduct(c.Context, command)
					if err != nil {
						return err
					}

					fmt.Println(createProductResult)
					return nil
				},
			},

			{
				Name:        "create-product-category",
				Usage:       "create new product category",
				Description: "create new product category",
				Action: func(c *cli.Context) error {

					productPersistence := postgres.NewProductRepository(
						sql.GetClient("readWrite"),
						sql.GetClient("readOnly"),
						logger.Get(),
					)

					productUseCase := usecase.NewCreateProductUseCase(productPersistence)

					command := command2.CreateProductCommand{
						Name:        "New Product",
						Amount:      10000,
						Description: "new product",
						Category:    "Dish",
						Menu:        true,
					}

					createProductResult, err := productUseCase.AddProduct(c.Context, command)
					if err != nil {
						return err
					}

					fmt.Println(createProductResult)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Erro ao executar comando - %s", err)
	}
}
