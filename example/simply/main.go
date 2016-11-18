package main

import (
	"fmt"

	"github.com/xboston/go-admitad-yml"
)

func main() {
	// Create YML
	admitadymlCat := admitadyml.NewYML("BestShop", "Best online seller Inc.", "http://best.seller.ru/")

	// Additional info
	admitadymlCat.Shop.Platform = "CMS"
	admitadymlCat.Shop.Version = "2.3"
	admitadymlCat.Shop.Agency = "Agency"
	admitadymlCat.Shop.Email = "CMS@CMS.ru"

	// id, rate, plus
	admitadymlCat.AddCurrency("RUR", "1", 0)

	// Categories
	admitadymlCat.AddCategory(1, 0, "Книги")
	admitadymlCat.AddCategory(2, 1, "Детективы")
	admitadymlCat.AddCategory(3, 1, "Боевики")
	admitadymlCat.AddCategory(4, 0, "Видео")
	admitadymlCat.AddCategory(5, 4, "Комедии")
	admitadymlCat.AddCategory(6, 0, "Принтеры")
	admitadymlCat.AddCategory(7, 0, "Оргтехника")

	// Simple Offer
	offer := admitadyml.Offer{
		ID:                   "123",
		Available:            true,
		Bid:                  21,
		URL:                  "http://best.seller.ru/product_/12348",
		Price:                600,
		OldPrice:             800,
		CurrencyID:           "USD",
		CategoryID:           6,
		Picture:              []string{"http://best.seller.ru/img/device12345.jpg", "http://best.seller.ru/img/device12.jpg"},
		Store:                false,
		Pickup:               true,
		Delivery:             false,
		Name:                 "Наручные часы Casio A1234567B",
		Vendor:               "Casio",
		VendorCode:           "A1234567B",
		Description:          "Изящные наручные часы.",
		SalesNotes:           "Необходима предоплата.",
		ManufacturerWarranty: true,
		CountryOfOrigin:      "Япония",
		Cpa:                  1,
		Age:                  &admitadyml.Age{Unit: "year", Value: "6"},
		TopSeller:            true,
	}
	offer.AddBarcode("0123456789012")
	offer.AddAge("year", "18")

	// Offer vendor.model
	offer2 := admitadyml.Offer{
		ID:                   "12341",
		Available:            true,
		Type:                 admitadyml.TypeVendorModel,
		Bid:                  13,
		URL:                  "http://best.seller.ru/product/12344",
		Price:                16800,
		OldPrice:             17000,
		CurrencyID:           "USD",
		CategoryID:           6,
		Picture:              []string{"http://best.seller.ru/img/device12345.jpg"},
		Store:                false,
		Pickup:               false,
		Delivery:             true,
		TypePrefix:           "Принтер",
		Vendor:               "HP",
		Model:                "Deskjet D2663",
		Description:          "Серия принтеров для людей, которым нужен надежный, простой в использовании цветной принтер для повседневной печати...",
		SalesNotes:           "Необходима предоплата.",
		ManufacturerWarranty: true,
		CountryOfOrigin:      "Япония",
		Cpa:                  1,
		Rec:                  "123123,1214,243",
		Expiry:               "P5Y",
		Weight:               2.07,
		Dimensions:           "100/25.45/11.112",
	}
	offer2.AddBarcode("1234567890120")
	offer2.AddParam("Максимальный формат", "", "А4")
	offer2.AddParam("Технология печати", "", "термическая струйная")
	offer2.AddParam("Тип печати", "", "Цветная")
	offer2.AddParam("Количество страниц в месяц", "стр", "1000")
	offer2.AddParam("Потребляемая мощность", "Вт", "20")
	offer2.AddParam("Вес", "кг", "2.73")

	err := offer.Validate()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = offer2.Validate()
	if err != nil {
		fmt.Println(err.Error())
	}

	admitadymlCat.AddOffer(offer)
	admitadymlCat.AddOffer(offer2)

	admitadyml.ExportToFile(admitadymlCat, "./admitad.xml", true)
}
