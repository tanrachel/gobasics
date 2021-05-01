package main

import (
	"fmt"
	"strings"
)

//declare global variables that will serve as core of the shopping list
var categories []string

var items = make(map[string]*itemInfo)

type itemInfo struct {
	category int
	quantity int
	unitCost float64
}

//User input to select which choice they want
var choice string

func main() {
	//add in test data
	addInTestData()

	for {
		//print shopping Menu
		shoppingMenu()
		//ask for user choice
		fmt.Println("Select your choice:")
		fmt.Scanln(&choice)
		choice = strings.ToLower(choice) // in case user types in different variation of quit
		//7 choices on shopping list
		switch choice {
		case "quit":
			fmt.Println("Exiting Shopping List - Bye!")
			return
		//Generate Shopping List View
		case "1":
			fmt.Println("View Shopping List.")
			viewEntireList()
		//Generate Report
		case "2":
			generateReportOption()
			var reportChoice string
			fmt.Println("Choose your report:")
			fmt.Scanln(&reportChoice)
			//Generate SubReport
			switch reportChoice {
			case "1":
				categoryReport()
			case "2":
				categoryList()
			case "3":
				break
			}
		//Add in new item - if item exists, user will be prompted to use (4)Modify
		case "3":
			fmt.Println("Add Items")
			addItem(addItemInfo())
		//Modify - checks that intended item exists before allowing modification
		case "4":
			fmt.Println("Modify Items.")
			intendItem := theIntendedItem()
			if val, ok := items[intendItem]; ok {
				fmt.Printf("Current Category: %v - Current Item: %s Current Quantity: %d Current Unit Cost: %g \n", categories[val.category], intendItem, val.quantity, val.unitCost)
				modifyItem(intendItem, modifyItemInfo())
			} else {
				fmt.Printf("%v does not exist!", intendItem)
			}
		//Delete
		case "5":
			fmt.Println("Delete Item.")
			deleteItem()
		//Print current data
		case "6":
			fmt.Println("Print Current Data.")
			printCurrentDataField()
		//Add in new category
		case "7":
			fmt.Println("Add New Category Name")
			addNewCategory()
		}

	}

	// modifyItem(modifyItemInfo())
	// modifyItemTester()
}

func shoppingMenu() {
	fmt.Println(" ")
	fmt.Println("Shopping List Application")
	fmt.Println("=========================")
	fmt.Println("1. View entire shopping list.")
	fmt.Println("2. Generate Shopping List Report")
	fmt.Println("3. Add Items")
	fmt.Println("4. Modify Items")
	fmt.Println("5. Delete Item")
	fmt.Println("6. Print Current Data")
	fmt.Println("7. Add New Category Name")
	fmt.Println("Quit to Exit.")

}

func addInTestData() {
	categories = append(categories, "Household", "Food", "Drinks")
	items["Fork"] = &itemInfo{category: findCategory("Household"), quantity: 4, unitCost: 3.0}
	items["Plate"] = &itemInfo{category: findCategory("Household"), quantity: 4, unitCost: 3.0}
	items["Bread"] = &itemInfo{category: findCategory("Food"), quantity: 2, unitCost: 2.0}
	items["Coke"] = &itemInfo{category: findCategory("Drinks"), quantity: 5, unitCost: 2.0}
	items["Cups"] = &itemInfo{category: findCategory("Household"), quantity: 5, unitCost: 3.0}
	items["Cake"] = &itemInfo{category: findCategory("Food"), quantity: 3, unitCost: 1.0}
	items["Sprite"] = &itemInfo{category: findCategory("Drinks"), quantity: 5, unitCost: 2.0}

}

func findCategory(findMe string) int {
	for i, val := range categories {
		if val == findMe {
			return i
		}
	}
	return -1
}

// 1. View Entire Shopping List
func viewEntireList() {
	for key, element := range items {
		// fmt.Println("Category:",categories[element.category],"- Item:",string())
		fmt.Printf("Category: %v - Item: %s Quantity: %d Unit Cost: %g \n", categories[element.category], key, element.quantity, element.unitCost)
	}
}

//3. Add Items Information
func addItemInfo() (string, string, int, float64) {
	var addItemName string
	var addCategory string
	var addUnits int
	var addCost float64
	fmt.Println("What is the name of your item?")
	fmt.Scanln(&addItemName)
	fmt.Println("What category does it belong to?")
	fmt.Scanln(&addCategory)
	fmt.Println("How many units are there?")
	fmt.Scanln(&addUnits)
	fmt.Println("How much does it cost per unit?")
	fmt.Scanln(&addCost)
	return addItemName, addCategory, addUnits, addCost

}

// 3a. actual add items function
func addItem(itemName string, category string, units int, cost float64) {
	//check if category exists
	if itemName == "" || category == "" {
		fmt.Println("Input not detected!")
		return
	}
	if findCategory(category) != -1 {
		//check if item already exist
		if _, ok := items[itemName]; ok {
			fmt.Println("This item already exists, please use modify function!")
		} else {
			items[itemName] = &itemInfo{category: findCategory(category), quantity: units, unitCost: cost}
			fmt.Println("Successfully added to the shopping list")
		}
	} else {
		categories = append(categories, category)
		items[itemName] = &itemInfo{category: findCategory(category), quantity: units, unitCost: cost}
		fmt.Println("Successfully added to the shopping list")
	}
}

//5. Delete Item
func deleteItem() {
	var itemToDelete string
	fmt.Println("Enter item name to delete:")
	fmt.Scanln(&itemToDelete)
	if _, ok := items[itemToDelete]; ok {
		delete(items, itemToDelete)
		fmt.Printf("Deleted %v \n", itemToDelete)
	} else {
		fmt.Println("Item not found. Nothing to delete!")
	}
}

//6. print current data field
func printCurrentDataField() {
	if len(items) == 0 {
		fmt.Println("No data found!")
	} else {
		for key, value := range items {
			fmt.Printf("%v - %v\n", key, *value)
		}
	}

}

//7. Add New Category Name
func addNewCategory() {
	var newCategory string
	fmt.Println("What is the New Category Name to add?")
	fmt.Scanln(&newCategory)
	if newCategory == "" {
		fmt.Println("No Input Found!")
	} else {
		if findCategory(newCategory) == -1 {
			categories = append(categories, newCategory)
			fmt.Printf("New category: %v added at index %d \n", newCategory, findCategory(newCategory))
		} else {
			fmt.Printf("Category: %v already exist at index %d \n", newCategory, findCategory(newCategory))
		}
	}
}
