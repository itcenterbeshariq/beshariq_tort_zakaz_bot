package controllers

import (
	"beshariq_tort_zakaz_bot/models"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
	"time"
)

type ProductController struct {
	beego.Controller
}

type ProductMalumotController struct {
	beego.Controller
}

// oddiy xotira (DB o‘rnida)
var products []models.Product

// GET /product → mahsulot qo‘shish formasi
func (c *ProductController) Product() {
	c.TplName = "product.html"
}

// POST /product → mahsulotni saqlash
func (c *ProductController) ProductPost() {
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
	c.Redirect("/malumot", 302) // saqlangandan keyin mahsulot ro‘yxatiga yo‘naltiradi
}

// GET /malumot → saqlangan mahsulotlar ro‘yxati
func (c *ProductMalumotController) Malumot() {
	c.Data["Products"] = products
	c.TplName = "malumot.html"
}
