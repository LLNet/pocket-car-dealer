package main

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"strconv"

	"github.com/pocketbase/pocketbase"
)

func main() {
	app := pocketbase.New()
	initCommand(app)
	initApi(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func initApi(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/hello",
			Handler: func(c echo.Context) error {
				return handel(c, app)
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
			},
		})

		return nil
	})
}

func handel(c echo.Context, app *pocketbase.PocketBase) error {
	var total int

	err := app.Dao().DB().
		Select("count(*)").
		From("users").
		Row(&total)

	if err != nil {
		log.Fatal(err)
	}

	return c.String(http.StatusOK, "Hello world!"+strconv.Itoa(total))
}

func initCommand(app *pocketbase.PocketBase) {
	app.RootCmd.AddCommand(&cobra.Command{
		Use: "install",
		Run: func(cmd *cobra.Command, args []string) {
			println("installing")
			createCollections(app)
		},
	})
	// make import command to import data from api
	app.RootCmd.AddCommand(&cobra.Command{
		Use: "import",
		Run: func(cmd *cobra.Command, args []string) {
			println("importing")
			importData(app)
		},
	})
}

func importData(app *pocketbase.PocketBase) {
	url := "https://gw.bilinfo.net/listingapi/api/export"
	username := "demo"
	password := "ocfB6XzF73"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
}

func createCollections(app *pocketbase.PocketBase) {
	collection, err := app.Dao().FindCollectionByNameOrId("cars3")

	if err != nil {
	}

	if collection == nil {
		collection := &models.Collection{
			Name:     "cars3",
			Type:     models.CollectionTypeBase,
			ListRule: nil,
			ViewRule: nil,
			/** @TODO change all them after here to be admin only */
			CreateRule: types.Pointer("@request.auth.id != ''"),
			UpdateRule: types.Pointer("@request.auth.id != ''"),
			DeleteRule: types.Pointer("@request.auth.id != ''"),
			Schema: schema.NewSchema(
				&schema.SchemaField{
					Name:     "mileage",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "year",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "make",
					Type:     schema.FieldTypeText,
					Required: true,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "model",
					Type:     schema.FieldTypeText,
					Required: true,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "variant",
					Type:     schema.FieldTypeText,
					Required: true,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "registrationDate",
					Type:     schema.FieldTypeDate,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "serviceBook",
					Type:     schema.FieldTypeBool,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "color",
					Type:     schema.FieldTypeText,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "type",
					Type:     schema.FieldTypeText,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "motor",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "propellant",
					Type:     schema.FieldTypeText,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "numberOfDoors",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "newPrice",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "numberOfGears",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				// gearType
				&schema.SchemaField{
					Name:     "gearType",
					Type:     schema.FieldTypeText,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "motorVolume",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "effect",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "cylinders",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "valvesPerCylinder",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "driveWheels",
					Type:     schema.FieldTypeText,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "trailerWeight",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "gasTankMax",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "litresUsedPer100Km",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "kmPerLiter",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "acceleration0To100",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "topSpeed",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "effectInNm",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				// effectInNmRpm, weight, greenTax, greenTaxPeriod, weightTax, weightTaxPeriod, payload, modelSeries,  numberOfAirbags
				&schema.SchemaField{
					Name:     "effectInNmRpm",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "weight",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "greenTax",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "greenTaxPeriod",
					Type:     schema.FieldTypeText,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "weightTax",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "weightTaxPeriod",
					Type:     schema.FieldTypeText,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "payload",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "modelSeries",
					Type:     schema.FieldTypeText,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "numberOfAirbags",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				// totalWeight, length, width, height, bodyType, comment, pictureCount
				&schema.SchemaField{
					Name:     "totalWeight",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "length",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "width",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "height",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "bodyType",
					Type:     schema.FieldTypeText,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "comment",
					Type:     schema.FieldTypeText,
					Required: false,
					Unique:   false,
				},
				&schema.SchemaField{
					Name:     "pictureCount",
					Type:     schema.FieldTypeNumber,
					Required: false,
					Unique:   false,
				},
			),
		}

		if err := app.Dao().SaveCollection(collection); err != nil {
			println(err.Error())
		}

		println("collection created")
	}
}
