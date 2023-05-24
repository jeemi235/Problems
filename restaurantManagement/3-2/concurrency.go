// Restaurant A has T no. of tables, W no. of waiting and C no. of customers for a day.Now when a user comes the manager checks whether the table is
// free or not. If the table is free, then the table will be allocated to the customer. Now if tables are not available then the manager will check
// for waiting. If waiting is available, then the waiting is allocated to the customer.
package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(4)

	var table int
	var waiting int
	var customer int
	//Constraints ->    - 1 ≤ T ≤ 20; 1 ≤ W ≤ 10; C > T + W

	fmt.Printf("Enter the number of tables between 1 to 20: ")
	fmt.Scanln(&table)

	fmt.Printf("Enter the capicity of waiting between 1 to 10: ")
	fmt.Scanln(&waiting)

	// waiting is capicity of waiting tables
	fmt.Printf("Enter the number of customers, it must be greater than total capicity of tables and waiting: ")
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
		for CustomerList[i] == "" {
			fmt.Printf("please enter valid value of %dth customer: ", i)
			fmt.Scanln(&CustomerList[i])
		}
	}
	fmt.Println("Your customer list is: ", CustomerList)

	go allocate(table, waiting, CustomerList, wg)
	go complete(table, waiting, CustomerList, wg)
	go leave(table, waiting, customer, CustomerList, wg)
	go wait(table, waiting, CustomerList, wg)

	wg.Wait()
}

func allocate(table int, waiting int, CustomerList []string, wg *sync.WaitGroup) {
	defer wg.Done()
	//Tablecount will traverse for tables
	Tablecount := 1
	//custm will travers for customer in list
	custm := 1
	//allocate is total number of customer who will be allocated tables
	allocate := table + waiting
	for allocate > 0 {
		if Tablecount > table {
			Tablecount = 1
		}
		fmt.Printf("%s has been allocated table %d.\n", CustomerList[custm], Tablecount)
		allocate -= 1
		Tablecount += 1
		custm += 1
	}
}

func complete(table int, waiting int, CustomerList []string, wg *sync.WaitGroup) {
	defer wg.Done()
	custm := 1
	total := table + waiting
	for custm < total+1 {
		fmt.Printf("%s has completed the task.\n", CustomerList[custm])
		custm += 1
	}
}

func wait(table int, waiting int, CustomerList []string, wg *sync.WaitGroup) {
	defer wg.Done()
	//count will traverse foe wait
	count := table + 1
	//countwait is count for waiting
	countwait := 1
	for waiting > 0 {
		fmt.Printf("%s has been allocated waiting %d.\n", CustomerList[count], countwait)
		waiting -= 1
		count += 1
		countwait += 1
	}
}

func leave(table int, waiting int, customer int, CustomerList []string, wg *sync.WaitGroup) {
	defer wg.Done()
	//extra is total count who will be allocated the tables
	extra := table + waiting + 1
	for i := extra; i < customer+1; i++ {
		fmt.Printf("%s leaves the restaurant. \n", CustomerList[i])
	}
}
