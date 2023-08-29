package pizzashop

import (
	"assg/pizzashop/pkg/common/models"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cucumber/godog"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type WorldContext struct {
	// featureVariable {[a:string] : string}
	pizza []models.Pizza
}

type pizzaShopFeatures struct {
	resp *httptest.ResponseRecorder
	*WorldContext
}

func (w *WorldContext) setWorldContext(responseJSON []models.Pizza) {

	fmt.Println(responseJSON, "responseJSON")
	w.pizza = responseJSON

	fmt.Println(w, "setWorldContext")
}

func (w *WorldContext) getWorldContext() []models.Pizza {
	fmt.Println(w)
	return w.pizza
}

var db *gorm.DB
var server *gin.Engine
var response *http.Response

func Init(ctx *godog.TestSuiteContext) {
	var err error
	var url = "postgres://postgres:Manushya@123@localhost:5433/pizza_shop_test"

	db, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		fmt.Println("Unable to create test database")
		return
	}

	// db.Close()

	db.AutoMigrate(models.Pizza{})
	db.AutoMigrate(models.Topping{})
}

func (p *pizzaShopFeatures) resetResponse(*godog.Scenario) {
	p.resp = httptest.NewRecorder()

	//db insertion operations
	var Pizza = models.Pizza{
		Name:  "Margherita",
		Price: 3000,
	}

	var Topping = models.Topping{
		Name:       "Pepperoni",
		Price:      300,
		IsInternal: false,
	}

	db.Create(&Topping)
	db.Create(&Pizza)
}

func (p *pizzaShopFeatures) thereIsARunningServer() error {
	server = gin.Default()

	handler := &handler{Db: db}
	server.GET("/pizzas", handler.getPizzas)
	go server.Run()
	return nil
}

func (p *pizzaShopFeatures) iMakeAGETRequestTo(endpoint string) error {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080%s", endpoint))
	if err != nil {
		return fmt.Errorf("failed to make GET request: %v", err)
	}
	response = resp

	return nil
}

func (p *pizzaShopFeatures) theResponseStatusCodeShouldBe(expectedStatusCode int) error {
	if response.StatusCode != expectedStatusCode {
		return fmt.Errorf("expected status code %d, but got %d", expectedStatusCode, response.StatusCode)
	}
	return nil
}

func (p *pizzaShopFeatures) theResponseShouldBeAJSONArray(body *godog.DocString) error {
	// var pizzaList []models.Pizza
	var actual, expected []models.Pizza
	w := &WorldContext{}
	if err := json.Unmarshal([]byte(body.Content), &expected); err != nil {
		return err
	}

	actualBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	if err := json.Unmarshal(actualBody, &actual); err != nil {
		return fmt.Errorf("failed to unmarshal actual JSON: %v", err)
	}

	fmt.Println(actual, expected, "actual")

	if expected[0].Name != actual[0].Name && expected[0].Price != actual[0].Price {
		return fmt.Errorf("failed to decode JSON response: %v", err)
	}

	w.setWorldContext(expected)
	return nil
}

func (p *pizzaShopFeatures) theWorldContextResponse() {
	pizzaData := p.getWorldContext()

	fmt.Println(pizzaData)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		TestSuiteInitializer: Init,
		ScenarioInitializer:  InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	p := pizzaShopFeatures{}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		p.resetResponse(sc)
		return ctx, nil
	})
	// call main.parser()

	ctx.Step(`a running server`, p.thereIsARunningServer)
	ctx.Step(`I send a "GET" request to "([^"]*)"$`, p.iMakeAGETRequestTo)
	ctx.Step(`the response code should be (\d+)$`, p.theResponseStatusCodeShouldBe)
	ctx.Step(`the response should match json:$`, p.theResponseShouldBeAJSONArray)
	ctx.Step(`get the above data`, p.theWorldContextResponse)
}
