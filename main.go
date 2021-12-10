package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("test1.xlsx")
	if err != nil {
		fmt.Errorf("error occurred while reading xlsx file: ", err)
		return
	}
	// defer func() {
	// 	if err := f.C err != nil {
	// 		fmt.Errorf(err)
	// 	}
	// }()
	rows, err := f.GetRows("RawData")
    if err != nil {
        fmt.Println(err)
        return
    }
    for _, row := range rows {
        for _, colCell := range row {
            fmt.Print(colCell, "\t")
        }
        fmt.Println()
    }
}