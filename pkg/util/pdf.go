package util

import (
	"fmt"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func GenerateProductListPDf() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)
}

func generateHeader(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(12, func() {
			err := m.FileImage("", props.Rect{})
			if err != nil {
				fmt.Println("Image file was not loaded 😱 - ", err)
			}
		})
	})
}

func generateBody(m pdf.Maroto) {

}
