//(4) Modification
package main

import (
	"fmt"
	"strconv"
)

//4. Modify Existing Items
func theIntendedItem() string {
	var itemToModify string
	//ask for the existing item
	fmt.Println("What is the name of your item?")
	fmt.Scanln(&itemToModify)
	return itemToModify
}
func modifyItemInfo() []string {
	//initialize variables
	modifyInfo := []string{}
	fields := []string{"item name", "Category", "Quantity", "Unit cost"}

	for i := 0; i < 4; i++ {
		var tempValue string
		fmt.Printf("Enter new %v. Enter for no change.\n", fields[i])
		fmt.Scanln(&tempValue)
		modifyInfo = append(modifyInfo, tempValue)
	}
	return modifyInfo

}
func modifyItem(toModify string, modifications []string) {
	fields := []string{"item name", "Category", "Quantity", "Unit cost"}
	//check if we category exists or no change
	if findCategory(modifications[1]) == -1 && modifications[1] != "" {
		categories = append(categories, modifications[1])
	}
	//make all the modifications
	for i := 1; i < 4; i++ {
		if modifications[i] != "" {
			switch i {
			//make changes to category
			case 1:
				items[toModify].category = findCategory(modifications[i])
			//make changes to quantity only if the input is an integer
			case 2:
				num, err := strconv.Atoi(modifications[i])
				if err == nil {
					items[toModify].quantity = num
				} else {
					modifications[i] = ""
				}
				// items[toModify].quantity, _ = strconv.Atoi(modifications[i])
			//make changes to unitCost only if ParseFloat is true
			case 3:
				num, err := strconv.ParseFloat(modifications[i], 64)
				if err == nil {
					items[toModify].unitCost = num
				} else {
					modifications[i] = ""
				}
				// items[toModify].unitCost, _ = strconv.ParseFloat(modifications[i], 64)
			}
		}
	}
	// if it requires to change the key
	if modifications[0] != "" {
		//change &item to new key created and then delete the old key
		items[modifications[0]] = items[toModify]
		delete(items, toModify)
	}
	//print out what was changed
	for index, val := range modifications {
		if val == "" {
			fmt.Println("No changes to", fields[index], "made")
		}
	}

}

//tester for modifying items
func modifyItemTester() {
	//item name x category x quantity 100 unitcost 100
	viewEntireList()
	modifyItem("Coke", []string{"", "", "100", "100"})
	viewEntireList()
	//item name x category Household quantity 200 unit cost 200
	modifyItem("Coke", []string{"", "Household", "200", "200"})
	viewEntireList()
	// item name Cola category Favorites 1 1000
	modifyItem("Coke", []string{"Cola", "Favorite", "1", "1000"})
	viewEntireList()
	fmt.Println("Printing all data")
	fmt.Println(categories)
	viewEntireList()

}
