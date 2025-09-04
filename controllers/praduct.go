package controllers

import (
	"beshariq_tort_zakaz_bot/models"
	"github.com/beego/beego/v2/server/web"
	"strconv"
	"time"
)

type ProductController struct {
	web.Controller
}

type ProductMalumot struct {
	web.Controller
}

var products []models.Product

func (c *ProductController) Get() {
	c.Data["Products"] = products
	c.TplName = "product.html"
}

func (c *ProductController) Post() {
	price, _ := strconv.Atoi(c.GetString("price"))
	weight, _ := strconv.ParseFloat(c.GetString("weight"), 64)

	product := models.Product{
		ID:          int64(len(products) + 1),
		ImageURL:    c.GetString("image_url"),
		Weight:      weight,
		Price:       price,
		Description: c.GetString("description"),
		CreatedAt:   time.Now(),
	}

	products = append(products, product)
	c.Redirect("product/product", 302)
}

func (c *ProductMalumot) Get() {
	c.Data["Products"] = products
	c.TplName = "malumot.html"
}
