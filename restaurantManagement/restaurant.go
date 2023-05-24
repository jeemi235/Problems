// Restaurant A has T no. of tables, W no. of waiting and C no. of customers for a day.Now when a user comes the manager checks whether the table is
// free or not. If the table is free, then the table will be allocated to the customer. Now if tables are not available then the manager will check
// for waiting. If waiting is available, then the waiting is allocated to the customer.
package main

import (
	"fmt"
)

func main() {
	var table int
	var waiting int
	var customer int
	//Constraints ->    - 1 ≤ T ≤ 20; 1 ≤ W ≤ 10; C > T + W

	fmt.Printf("enter the number of tables between 1 to 20: ")
	fmt.Scanln(&table)

	fmt.Printf("enter the capicity of waiting between 1 to 10: ")
	fmt.Scanln(&waiting)

	// waiting is capicity of waiting tables
	fmt.Printf("enter the number of customers, it must be greater than total capicity of tables and waiting: ")
	fmt.Scanln(&customer)

	// we will ask for customer untill it matches our constraints
	for customer < table+waiting {
		fmt.Printf("Please enter the value greater than %d: ", table+waiting)
		fmt.Scanln(&customer)
	}

	// CustomerList is list of customers
	var CustomerList = make([]string, customer+1)
	for i := 1; i <= customer; i++ {
		fmt.Printf("Enter %dth customer: ", i)
		fmt.Scanln(&CustomerList[i])
	}

	fmt.Println("Your customer list is: ", CustomerList)
	manager(table, waiting, customer, CustomerList)
}

// This function will handle tables and waiting
func manager(table int, waiting int, customer int, CustomerList []string) {
	//count is a variable and will traverse from count to total number of customers
	count := 1

	//countwait will traverse for waiting list
	countwait := 1
	var allocated = make([]string, table+waiting+1)
	var WaitArray = make([]string, waiting+1)

	//WaitingTemp and tableTemp are temp variables to store original values
	WaitingTemp := waiting
	tableTemp := table

	for count <= customer {
		//We will traverse till table capicity becomes 0
		if table > 0 {
			fmt.Printf("%s has been allocated table %d.\n", CustomerList[count], count)
			table -= 1
			allocated[count] = CustomerList[count]

		} else if waiting > 0 {
			//We will traverse till waiting capicity becomes 0
			fmt.Printf("%s has been allocated waiting %d.\n", CustomerList[count], countwait)
			waiting -= 1
			//add values in waitarray and allocated array
			WaitArray[countwait] = CustomerList[count]
			allocated[count] = CustomerList[count]
			countwait += 1

		} else {
			//in this else will relese those customers who couldnt find place
			fmt.Printf("%s leaves the restaurant. \n", CustomerList[count])
		}
		count += 1
	}

	//With countWC we will traverse till allocated table becomes empty
	countWC := 1

	//tableTemp2 is again temp for table's capicity
	tableTemp2 := tableTemp
	tableTemp = tableTemp + WaitingTemp

	//x will store the length of allocated tables
	x := len(allocated) + 1
	for countWC <= x {
		//Traverse till total value of allocated values

		if tableTemp > 0 {
			fmt.Printf("%s has completed the task.\n", allocated[countWC])
			tableTemp -= 1
		}

		if WaitingTemp > 0 {
			fmt.Printf("%s has been allocated table %d.\n", WaitArray[countWC], countWC)
			allocated[tableTemp2+1] = WaitArray[countWC]
			WaitingTemp -= 1
		}

		countWC += 1
	}
}
