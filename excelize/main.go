package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	// Create a new sheet.
	fileName := "题库.xlsx"
	sheetName := "Sheet1"
	index := f.NewSheet(sheetName)
	err := f.SetColWidth("Sheet1", "A", "D", 20)
	if err != nil {
		panic(0)
	}

	// Set value of a cell.
	f.SetCellValue(sheetName, "A1", "题号")
	f.SetCellValue(sheetName, "B1", "题干")
	f.SetCellValue(sheetName, "C1", "可选项")
	f.SetCellValue(sheetName, "D1", "参考答案")

	// 下面是可变项
	f.SetCellValue(sheetName, "A2", "QUESTION 1")
	f.SetCellValue(sheetName, "B2", "A. the disk is usually nonreducible.")
	f.SetCellValue(sheetName, "C2", "A. the disk is usually nonreducible.\nB. surgical intervention is indicated.\nC. lateral pole disk derangements usually precede medial pole derangements.\nD. during lateral pole derangements, the medial pole is braced just down the eminence in centric relation.")
	f.SetCellValue(sheetName, "D2", "Correct Answer: A")

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)

	// Save spreadsheet by the given path.
	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
	}
}
