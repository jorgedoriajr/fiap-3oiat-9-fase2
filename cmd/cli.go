package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
	postgres2 "hamburgueria/internal/modules/ingredient/infra/database/postgres"
	"hamburgueria/internal/modules/ingredient/service"
	"hamburgueria/internal/modules/product/infra/database/postgres"
	service2 "hamburgueria/internal/modules/product/service"
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
			{
				Name:        "by-product",
				Usage:       "get ingredients by product",
				Description: "get ingredients by product",
				Action: func(c *cli.Context) error {

					ingredientPersistence := postgres2.NewIngredientRepository(
						sql.GetClient("readWrite"),
						sql.GetClient("readOnly"),
						logger.Get(),
					)

					productPersistence := postgres.NewProductRepository(
						sql.GetClient("readWrite"),
						sql.GetClient("readOnly"),
						logger.Get(),
					)

					serviceI := service.NewIngredientFinderService(ingredientPersistence)

					productFinder := service2.NewProductFinderService(productPersistence, *serviceI)

					//command := command2.CreateProductCommand{
					//	Name:        "New Product",
					//	Amount:      10000,
					//	Description: "new product",
					//	Category:    "Dish",
					//	Menu:        true,
					//}

					id := uuid.MustParse("e962a63f-e1bb-4011-9b30-cbdc771ae740")

					createProductResult, err := serviceI.FindIngredientsByProductId(context.TODO(), id)
					if err != nil {
						return err
					}

					result, err := productFinder.FindByIDWithIngredients(context.TODO(), id)
					if err != nil {
						return err
					}

					for _, entity := range createProductResult {
						fmt.Println(JsonPretty(entity))
					}

					fmt.Println(JsonPretty(result))
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Erro ao executar comando - %s", err)
	}
}

func JsonPretty(i interface{}) string {
	r, _ := json.MarshalIndent(i, "", " ")
	return string(r)
}
