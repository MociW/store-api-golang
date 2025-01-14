package util

import (
	"fmt"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func GenerateProductListPDf() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)

	generateHeader(m)
	generateBody(m)
}

func generateHeader(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(12, func() {
			err := m.FileImage("", props.Rect{
				Center:  false,
				Percent: 75,
			})

			if err != nil {
				fmt.Println("Image file was not loaded 😱 - ", err)
			}
		})
	})

	m.Row(10, func() {
		m.Text("Product List", props.Text{
			Top:   3,
			Style: consts.Bold,
			Align: consts.Left,
			Color: getBlackColor(),
		})
	})
}

func generateBody(m pdf.Maroto) {
	// tableHeadings := []string{"SKU", "Item", "Price", "Quantity"}
}

func getBlackColor() color.Color {
	return color.Color{
		Red:   80,
		Green: 80,
		Blue:  80,
	}
}
