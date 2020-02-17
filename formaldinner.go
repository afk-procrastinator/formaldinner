package main

// 31 tables, kitchen crew,
// x31 waiters (one alt)
// x10 kitchen crew
// x249 seated

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Person struct declaration:
type Person struct {
	Value     string
	Placement string
	Lastname  string
	Firstname string
	Table     string
	Table2    string
	Table3    string
}

// 32 people:
var waiter []Person

// 10 people:
var kitchen []Person

// 8 people per, 31 total, 249 in all:
var table [][]Person

var value int = 1

var tableNum int = 1

var counter = -1

var kitchenCrewLoc = 1

// Main function:
func main() {

	var people []Person
	var csvFile, _ = os.Open("seating.csv")
	csvFile.Close()
	people = nil
	csvFile, _ = os.Open("seating.csv")
	var reader = csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		people = append(people, Person{
			Firstname: line[0],
			Lastname:  line[1],
		})
	}
	var peopleSlice = people
	var slicedPeople = Shuffle(peopleSlice)
	initFile("first.csv")
	iterateAndChoose(slicedPeople, "first.csv")

}

// Shuffle function taken from https://www.calhoun.io/how-to-shuffle-arrays-and-slices-in-go/
func Shuffle(slice []Person) []Person {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]Person, len(slice))
	n := len(slice)
	for i := 0; i < n; i++ {
		randIndex := r.Intn(len(slice))
		ret[i] = slice[randIndex]
		slice = append(slice[:randIndex], slice[randIndex+1:]...)
	}
	return ret
}

// Chooses the next int people, appends to return crew:
func chooseNext(slice []Person, num int) []Person {
	crew := make([]Person, 0)
	for i := 0; i < num; i++ {
		crew = append(crew, slice[i])
	}
	return crew
}

// Chooses the next int people, removes them from the main slice:
func removeIndex(num int, slice []Person) []Person {

	//returnrearrange(num, 1, slice)

	for i := 0; i < num; i++ {
		slice = append(slice[:0], slice[0+1:]...)
	}
	return slice
}

// File append function, seatType is:
/*
1 = table
2 = KC
3 = waiter
*/
func makeFile(slice []Person, num int, seatType int, title string) string {

	var name string

	f, err := os.OpenFile(title, os.O_APPEND|os.O_WRONLY, 0644)
	d := slice

	// Loops through entirety of my slice
	for _, v := range d {

		valueStr := strconv.Itoa(value)

		// If you're above position (value) 280, you're KC
		switch {
		case value > 280:
			valueStr = "Kitchen Crew number " + strconv.Itoa(kitchenCrewLoc)
			kitchenCrewLoc++
		case value >= 249 && value <= 280:
			var waiterTable = value - 248
			valueStr = "Waiter at table " + strconv.Itoa(waiterTable)
		default:
			// Counter cycles from 0 to 7, picking out 8 people per table
			counter++
			if counter == 8 {
				tableNum++
				counter = 0
			}

			valueStr = strconv.Itoa(tableNum)

		}

		// Full name to print to the csv file. Will be changed for final iterations, currently prints basically everything for clarity's sake.

		nextLoc := nextLocation(tableNum, counter, value, 5, 5)

		name = v.Lastname + "," + v.Firstname + "," + valueStr + "," + nextLoc

		fmt.Fprintln(f, name)

		value++

		if err != nil {
			fmt.Println(err)
			return ""
		}
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return ""
}

// Primary init of the file:
func initFile(title string) {
	f, err := os.Create(title)
	if err != nil {
		f.Close()
		return
	}
}

// Function that iterates through the slice and chooses people to go to certain positions based on index. Also calls in the file creation function.
func iterateAndChoose(slicedPeople []Person, title string) []Person {
	var originalGroup = slicedPeople

	// choose the first 10 to be kitchen crew:
	var nextGroup = chooseNext(slicedPeople, 10)
	makeFile(nextGroup, 10, 2, title)

	// remove the first 10 from the main list:
	removeIndex(10, slicedPeople)

	// choose the next 31 to be waiters:
	nextGroup = chooseNext(slicedPeople, 32)
	makeFile(nextGroup, 32, 3, title)

	// remove the next 31 from the main list:
	removeIndex(32, slicedPeople)

	// append all tables to CSV file:
	for i := 1; i < 32; i++ {
		var seated = chooseNext(slicedPeople, 8)
		makeFile(seated, i, 1, title)
		removeIndex(8, slicedPeople)
	}
	fmt.Println("all completed succesfully!")

	return originalGroup
}

// Calculates the following location. Table is
func nextLocation(table int, place int, location int, waiterTable int, kitchenLoc int) string {
	var newTable int
	var newPlace int
	//var stringToPrint string

	switch place {
	case 0: // 1 -> up a table to 2 (+10)
		newTable = table + 1
		newPlace = place + 1
		location = location + 10

	case 1: // 2 -> up 2 tables to 3 (+19) [first 10 are selected for KC, placed back in after]
		newTable = table + 2
		newPlace = place + 1
		location = location + 19

	case 2: // 3 -> up 3 tables to 4 (+28)
		newTable = table + 3
		newPlace = place + 1
		location = location + 28

	case 3: // 4 -> up 4 tables to 5 (+37)
		newTable = table + 4
		newPlace = place + 1
		location = location + 37

	case 4: // 5 -> up 5 tables to 6 (+46)
		newTable = table + 5
		newPlace = place + 1
		location = location + 46

	case 5: // 6 -> up 6 tables to 7 (+55)
		newTable = table + 6
		newPlace = place + 1
		location = location + 55

	case 6: // 7 -> up 7 tables to 8 (+64)
		newTable = table + 7
		newPlace = place + 1
		location = location + 64

	case 7: // 8 -> up 8 tables to 9 (W) (+73)

		if location >= 249 && location <= 280 { // Waiters -> get placed into position 1 based on their location
			newTable = waiterTable + 1
			newPlace = 1
			location = (newTable * 8) - 7
		} else if location > 280 { // KC
			newTable = kitchenLoc
			newPlace = 2
		} else {
			newTable = table + 8
			newPlace = place + 1
		}
	}
	newString := "New table is: " + strconv.Itoa(newTable) + " and new position is: " + strconv.Itoa(newPlace)
	return newString
}
