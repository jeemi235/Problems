package main

import (
	"fmt"
	"math"
)

func checkPrimeNumber(num int, prime chan int, nonprime chan int) {

	if num == 1 {
		// fmt.Printf("%d is Not Prime\n", num)
		nonprime <- num
		return
	}

	sq_root := int(math.Sqrt(float64(num)))
	for i := 2; i <= sq_root; i++ {
		if num%i == 0 {
			//fmt.Printf("%d is Not Prime\n", num)
			nonprime <- num
			return
		}
	}
	//fmt.Printf("%d is Prime\n", num)
	prime <- num
	return

}

func checkOddEven(num int, odd chan int, even chan int) {
	if num%2 == 0 {
		//fmt.Printf("%d is odd number\n", num)
		even <- num
	} else {
		//fmt.Printf("%d is even number\n", num)
		odd <- num
	}
}

// Constraints -> length of array < 10 ,maximum number: 1000

func main() {

	prime := make(chan int)
	nonprime := make(chan int)
	odd := make(chan int)
	even := make(chan int)
	fmt.Print("please enter the length of array:")
	var size int
	fmt.Scanln(&size)

	for size > 10 {
		fmt.Print("please enter the length of array less than 10:")
		fmt.Scanln(&size)
	}

	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %dth element: ", i+1)
		fmt.Scanln(&arr[i])
		for arr[i] > 1000 || arr[i] < 1 {
			fmt.Print("please enter the value less than 1000:")
			fmt.Scanln(&arr[i])
		}
	}

	for i := 0; i < size; i++ {
		num := arr[i]
		go checkPrimeNumber(num, prime, nonprime)
		select {
		case primenumber := <-prime:
			fmt.Printf("%d is prime.\n", primenumber)
		case nonprimenumber := <-nonprime:
			fmt.Printf("%d is not prime.\n", nonprimenumber)
		}
		go checkOddEven(num, odd, even)
		select {
		case oddnumber := <-odd:
			fmt.Printf("%d is odd.\n", oddnumber)
		case evennumber := <-even:
			fmt.Printf("%d is even.\n", evennumber)
		}
	}
}

/*
// Restaurant A has T no. of tables, W no. of waiting and C no. of customers for a day.Now when a user comes the manager checks whether the table is
// free or not. If the table is free, then the table will be allocated to the customer. Now if tables are not available then the manager will check
// for waiting. If waiting is available, then the waiting is allocated to the customer.
package main

import (
	"fmt"
	"sync"
)

var mut sync.Mutex //pointer

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(4)

	allocate_chan := make(chan string)
	table_chan := make(chan int)
	complete_chan := make(chan string)
	wait_chan := make(chan string)
	leave_chan := make(chan string)

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

	for i := 1; i < table+waiting+1; i++ {
		go allocate(table, waiting, CustomerList, wg, allocate_chan, table_chan)
		select {
		case x := <-allocate_chan:
			fmt.Printf("%s has been allocated table %d.\n", x, <-table_chan)
		}
		go complete(table, waiting, CustomerList, wg, complete_chan)
		select {
		case y := <-complete_chan:
			fmt.Printf("%s has completed the task.\n", y)
		}
		go wait(table, waiting, CustomerList, wg, wait_chan)
		select {
		case z := <-wait_chan:
			fmt.Printf("%s has been allocated waiting %d.\n", z, i)
		}
		go leave(table, waiting, customer, CustomerList, wg, leave_chan)
		select {
		case w := <-leave_chan:
			fmt.Printf("%s leaves the restaurant. \n", w)
		}
	}
	//go allocate(table, waiting, CustomerList, wg,<-allocate_chan)
	// go complete(table, waiting, CustomerList, wg)
	// go leave(table, waiting, customer, CustomerList, wg)
	// go wait(table, waiting, CustomerList, wg)

	wg.Wait()
}

func allocate(table int, waiting int, CustomerList []string, wg *sync.WaitGroup, allocate_chan chan string, table_chan chan int) {
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
		//fmt.Printf("%s has been allocated table %d.\n", CustomerList[custm], Tablecount)
		x := CustomerList[custm]
		allocate_chan <- x
		table_chan <- Tablecount

		allocate -= 1
		Tablecount += 1
		custm += 1
	}
}

func complete(table int, waiting int, CustomerList []string, wg *sync.WaitGroup, complate_chan chan string) {
	defer wg.Done()
	custm := 1
	total := table + waiting
	for custm < total+1 {
		//fmt.Printf("%s has completed the task.\n", CustomerList[custm])
		x := CustomerList[custm]
		complate_chan <- x

		custm += 1
	}
}

func wait(table int, waiting int, CustomerList []string, wg *sync.WaitGroup, wait_chan chan string) {
	defer wg.Done()
	//count will traverse foe wait
	count := table + 1
	//countwait is count for waiting
	countwait := 1
	for waiting > 0 {
		//fmt.Printf("%s has been allocated waiting %d.\n", CustomerList[count], countwait)
		x := CustomerList[count]
		wait_chan <- x

		waiting -= 1
		count += 1
		countwait += 1
	}
}

func leave(table int, waiting int, customer int, CustomerList []string, wg *sync.WaitGroup, leave_chan chan string) {
	defer wg.Done()
	//extra is total count who will be allocated the tables
	extra := table + waiting + 1
	for i := extra; i < customer+1; i++ {
		//fmt.Printf("%s leaves the restaurant. \n", CustomerList[i])
		x := CustomerList[i]
		leave_chan <- x
	}
}

*/

/*
// Restaurant A has T no. of tables, W no. of waiting and C no. of customers for a day.Now when a user comes the manager checks whether the table is
// free or not. If the table is free, then the table will be allocated to the customer. Now if tables are not available then the manager will check
// for waiting. If waiting is available, then the waiting is allocated to the customer.
package main

import (
	"fmt"
)

func main() {
	allocate_chan := make(chan string)
	table_chan := make(chan int)
	complete_chan := make(chan string)
	wait_chan := make(chan string)
	leave_chan := make(chan string)

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

	go allocate(table, waiting, CustomerList, allocate_chan, table_chan)
	go complete(table, waiting, CustomerList, complete_chan)
	go wait(table, waiting, CustomerList, wait_chan)
	go leave(table, waiting, customer, CustomerList, leave_chan)

	for i := 1; i < table+waiting+1; i++ {
		select {
		case x := <-allocate_chan:
			fmt.Printf("%s has been allocated table %d.\n", x, <-table_chan)
		}

		select {
		case y := <-complete_chan:
			fmt.Printf("%s has completed the task.\n", y)
		}

		select {
		case z := <-wait_chan:
			fmt.Printf("%s has been allocated waiting %d.\n", z, i)
		}

		select {
		case w := <-leave_chan:
			fmt.Printf("%s leaves the restaurant. \n", w)
		}
	}
	//go allocate(table, waiting, CustomerList, wg,<-allocate_chan)
	// go complete(table, waiting, CustomerList, wg)
	// go leave(table, waiting, customer, CustomerList, wg)
	// go wait(table, waiting, CustomerList, wg)

}

func allocate(table int, waiting int, CustomerList []string, allocate_chan chan string, table_chan chan int) {

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
		//fmt.Printf("%s has been allocated table %d.\n", CustomerList[custm], Tablecount)
		x := CustomerList[custm]
		allocate_chan <- x
		table_chan <- Tablecount

		allocate -= 1
		Tablecount += 1
		custm += 1
	}
}

func complete(table int, waiting int, CustomerList []string, complate_chan chan string) {

	custm := 1
	total := table + waiting
	for custm < total+1 {
		//fmt.Printf("%s has completed the task.\n", CustomerList[custm])
		x := CustomerList[custm]
		complate_chan <- x

		custm += 1
	}
}

func wait(table int, waiting int, CustomerList []string, wait_chan chan string) {
	//count will traverse foe wait
	count := table + 1
	//countwait is count for waiting
	countwait := 1
	for waiting > 0 {
		//fmt.Printf("%s has been allocated waiting %d.\n", CustomerList[count], countwait)
		x := CustomerList[count]
		wait_chan <- x

		waiting -= 1
		count += 1
		countwait += 1
	}
}

func leave(table int, waiting int, customer int, CustomerList []string, leave_chan chan string) {
	//extra is total count who will be allocated the tables
	extra := table + waiting + 1
	for i := extra; i < customer+1; i++ {
		//fmt.Printf("%s leaves the restaurant. \n", CustomerList[i])
		x := CustomerList[i]
		leave_chan <- x
	}
}
*/
