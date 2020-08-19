package main

import "fmt"

func main() {
	/*
		Exercise 1
		Using a COMPOSITE LITERAL: 
			1. Create an ARRAY which holds 5 VALUES of TYPE int and assign VALUES to each index position
			2. Range over the array and print the values out
			3. Using format printing print out the TYPE of the array
	*/
	a := [5]int{1, 2, 3, 4, 42}
   
	for i, v := range a {
		fmt.Printf("index: %v\tvaleu: %v\n", i, v)
	}

	fmt.Printf("type of a: %T\n", a)
	
	/*
		Exercise 2
		Using a COMPOSITE LITERAL: 
			1. Create a SLICE of TYPE int and assign 10 VALUES 
			2. Range over the slice and print the values out
			3. Using format printing print out the TYPE of the slice
	*/
	b := []int{10, 20, 30, 40, 42, 50, 60, 70, 80, 90}

	for i, v := range b {
		fmt.Printf("index: %v\tvalue: %v\n", i, v)
	}

	fmt.Printf("type of b: %T\n", b)

	/*
		Exercise 3
		Using the code from the previous example, use SLICING to create the following new slices which are then printed:
			[10 20 30 40 42]
			[50 60 70 80 90]
			[30 40 42 50 60]
			[20 30 40 42 50]
	*/
	fmt.Println(b[:5])
	fmt.Println(b[5:])
	fmt.Println(b[2:7])
	fmt.Println(b[1:6])


	/*
		Exercise 4
		Follow these steps:
			1. Start with this slice
				[42, 43, 44, 45, 46, 47, 48, 49, 50, 51]
			2. Append to that slice this value
				52
			3. Print out the slice 
			4. In ONE STATEMENT append to that slice these values
				53
				54
				55
			5. print out the slice
			6. Append to the slice this slice
				[56, 57, 58, 59, 60]
			7. print out the slice
	*/
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 53, 54, 55)
	fmt.Println(x)
	
	y :=[]int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

	/*
		Exercise 5
		To DELETE from a slice, we use APPEND along with SLICING. 
		For this hands-on exercise, follow these steps:
			1. start with this slice
				[42, 43, 44, 45, 46, 47, 48, 49, 50, 51]
			2. use APPEND & SLICING to get these values here which you should ASSIGN to a variable “y” and then print:
				[42, 43, 44, 48, 49, 50, 51]
	*/
	x = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y = append(x[:3], x[6:]...)
	fmt.Println(y)

	/*
		Exercise 6
		Create a slice to store the names of all of the states in the United States of America. 
			1. What is the length of your slice? 
			2. What is the capacity? 
			3. Print out all of the values, along with their index position in the slice, without using the range clause. 
		Here is a list of the states:
		` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, 
		` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, 
		` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, 
		` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, 
		` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, 
		` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`
	*/
	z := make([]string, 50, 50)
	z = []string{
		` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, 
		` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, 
		` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, 
		` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, 
		` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, 
		` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,
	}

	fmt.Printf("slice: %v\nlength: %v\tcapacity: %v\n", z, len(z), cap(z))

	for i := 0; i < len(z); i++ {
		fmt.Printf("index: %v\tvalue: %v\n", i, z[i])
	}

	/*
		Exercise 7
		Create a slice of a slice of string ([][]string). 
			1. Store the following data in the multi-dimensional slice:
				"James", "Bond", "Shaken, not stirred!"
				"Miss", "Moneypenny", "Helloooooo, James."
			2. Range over the records, then range over the data in each record.
	*/
	xxs := [][]string{};
	xxs = append(xxs, []string{"James", "Bond", "Shaken, not stirred!"})
	xxs = append(xxs, []string{"Miss", "Moneypenny", "Helloooooo, James."})

	for index, value := range xxs {
		fmt.Printf("index: %v\n", index)
		for i, v := range value {
			fmt.Printf("\tindex: %v\tvalue: %v\n", i, v)
		}
	}


	/*
		Exercise 8
		Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things. 
			1. Store three records in your map. 
				`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
				`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
				`no_dr`, `Being evil`, `Ice cream`, `Sunsets`
			2. Print out all of the values, along with their index position in the slice.
	*/
	m := map[string][]string{
		"bond_james": []string{"Shaken, not stirred!", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr": []string{"Being evil", "Ice cream", "Sunset"},
	}

	for person, things := range m {
		fmt.Printf("person: %v\n", person)
		for k, v := range things {
			fmt.Printf("\tkey: %v\tvalue: %v\n", k, v)
		}
	}

	/*
		Exercise 9
		Using the code from the previous example, add a record to your map. 
		Now print the map out using the “range” loop
	*/
	m["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

	for i, v := range m {
		fmt.Printf("index: %v\tvalue: %v\n", i, v)
	}

	/*
		Exercise 10
		Using the code from the previous example, delete a record from your map. 
		Now print the map out using the “range” loop
	*/
	delete(m, "fleming_ian")
	
	for i, v := range m {
		fmt.Printf("index: %v\tvalue: %v\n", i, v)
	}
}