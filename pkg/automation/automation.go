package automation

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/chromedp/chromedp"
	"github.com/omaressameldin/water-bot/internal/env"
)

func navigateToProduct(productId string) chromedp.Action {
	return chromedp.Navigate(fmt.Sprintf("%s/%s", PRODUCTS_LINK, productId))
}

func setPostCode() []chromedp.Action {
	postcode, err := env.GetPostCode()
	if err != nil {
		log.Fatal(err)
	}

	return []chromedp.Action{
		chromedp.Navigate(MAIN_URL),
		chromedp.SetValue(`//input[@id="input_plz"]`, postcode),
		chromedp.Click(`//button[@id="button_plz"]`),
		chromedp.WaitVisible(`//div[@class="products-list-default products-list-sale"]`),
	}
}

func login() []chromedp.Action {
	email, err := env.GetEmail()
	password, err := env.GetPassword()
	if err != nil {
		log.Fatal(err)
	}

	return []chromedp.Action{
		chromedp.Stop(),
		chromedp.Navigate(LOGIN_URL),
		chromedp.WaitVisible(`//input[@type="email"]`),
		chromedp.SetValue(`//input[@type="email"]`, email),
		chromedp.SetValue(`//input[@type="password"]`, password),
		chromedp.Submit(`//input[@type="email"]`),
		chromedp.WaitVisible(`//nav[contains(@class, "account-nav")]`),
	}
}

func addStillWaterToCart(stillWaterBoxes int) []chromedp.Action {
	stillWaterId, err := env.GetStillWaterId()
	if err != nil {
		log.Fatal(err)
	}

	return addToCart(stillWaterBoxes, stillWaterId)
}

func addSparklingWaterToCart(sparklingWaterBoxes int) []chromedp.Action {
	sparklingWaterId, err := env.GetSparklingWaterId()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(sparklingWaterId)
	return addToCart(sparklingWaterBoxes, sparklingWaterId)
}

func addToCart(qty int, productId string) []chromedp.Action {
	productFormInput := `//input[@name="product"]`

	return []chromedp.Action{
		chromedp.Stop(),
		navigateToProduct(productId),
		chromedp.WaitReady(productFormInput),
		chromedp.SetValue(`//input[@id="qty"]`, strconv.Itoa(qty)),
		chromedp.Submit(productFormInput),
		chromedp.WaitVisible(`//div[contains(@class, "page messages")]`),
	}
}

func checkout() []chromedp.Action {
	return []chromedp.Action{
		chromedp.Navigate(CHECKOUT_URL),
	}
}

func startTasks(tasks []chromedp.Action) {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
	defer cancel()
	err := chromedp.Run(ctx, tasks...)
	if err != nil {
		log.Fatal(err)
	}
}

func OrderWater(stillWaterBoxes, sparklingWaterBoxes int) {
	var buf []byte
	tasks := setPostCode()
	// tasks = append(tasks, login()...)
	tasks = append(tasks, addStillWaterToCart(stillWaterBoxes)...)
	tasks = append(tasks, addSparklingWaterToCart(sparklingWaterBoxes)...)
	tasks = append(tasks, checkout()...)
	tasks = append(tasks, chromedp.WaitVisible(`//div[@class="page-wrapper"]`), chromedp.Screenshot(`//div[@class="page-wrapper"]`, &buf))
	startTasks(tasks)

	if err := ioutil.WriteFile("order.png", buf, 0644); err != nil {
		log.Fatal("error")
	}
	log.Printf("ordering %d still, and %d sparkling", stillWaterBoxes, sparklingWaterBoxes)
}
