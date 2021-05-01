//(2) Generate Report
package main

import "fmt"

// 2. Generate Shopping List Report
//2a. Ways to Generate Report
func generateReportOption() {
	fmt.Println("Generate Report")
	fmt.Println("1. Total Cost of each Category")
	fmt.Println("2. List of item by category")
	fmt.Println("3. Main Menu.")

}

//2b. generate report by category
func categoryReport() {
	fmt.Println("Total cost by Category")
	categoryCost := make([]float64, len(categories))
	for _, element := range items {
		categoryCost[element.category] += element.unitCost * float64(element.quantity)
	}
	for i, v := range categoryCost {
		fmt.Printf("%v cost: %g \n", categories[i], v)
	}
}

// 2c generate report by categoryList
func categoryList() {
	// var itemsInEachCategoryList [][]string
	itemsInEachCategoryList := make([][]string, len(categories))
	for key, element := range items {
		itemsInEachCategoryList[element.category] = append(itemsInEachCategoryList[element.category], key)
	}
	for cat, val := range itemsInEachCategoryList {
		for _, item := range val {
			fmt.Printf("Category: %v - Item: %s Quantity: %d Unit Cost: %g \n", categories[cat], item, items[item].quantity, items[item].unitCost)
		}
	}
}
