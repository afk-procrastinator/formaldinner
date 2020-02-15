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

	counter := -1
	f, err := os.OpenFile(title, os.O_APPEND|os.O_WRONLY, 0644)
	d := slice

	for _, v := range d {

		valueStr := strconv.Itoa(value)

		switch {
		case value > 280:
			valueStr = valueStr + " kc"
		default:

			counter++
			if counter > 7 {
				tableNum++
				counter = 0
				fmt.Println("counted to ", tableNum)
			}

			valueStr = strconv.Itoa(tableNum)

		}

		name = v.Lastname + "," + v.Firstname + "," + valueStr + "," + strconv.Itoa(counter)

		//fmt.Println("value", name)

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

// Function that iterates through the slice and chooses people to go to certain positions based on index.
// Also calls in the file creation function.
// slicedPeople is the slice to iterate through
// title is the string of the txt file
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

// PLACEMENT STUFF:

// 9 people per table, 9 is waiter
// 1 -> up a table to 2 (+10)
// 2 -> up 2 tables to 3 (+19) [first 10 are selected for KC, placed back in after]
// 3 -> up 3 tables to 4 (+28)
// 4 -> up 4 tables to 5 (+37)
// 5 -> up 5 tables to 6 (+46)
// 6 -> up 6 tables to 7 (+55)
// 7 -> up 7 tables to 8 (+64)
// 8 -> up 8 tables to 9 (W) (+73)
// 9 (W) -> up 9 tables to 1 (+74)

func newPlacement(location int) {
	// Calculate the new placement:
	switch {
	case location%4 == 0 && location%8 != 0:
		fmt.Println("at position 4")
	case location%2 == 0 && location%4 != 0 && location%2 != 0:
		fmt.Println("at position 8")
	case location%3 == 0 && location%6 != 0:
		fmt.Println("at position 3")
	case location%6 == 0:
		fmt.Println("at position 6")
	case location%4 != 0 && location%6 != 0 && location%2 == 0:
		fmt.Println("at position 2")
	case location%5 == 0:
		fmt.Println("at position 5")
	default:
		fmt.Println("number 1")
	}
}

func rearrange(remove int, place int, input []Person) []Person {
	slice := input
	val := slice[remove]
	slice = append(slice[:remove], slice[remove+1:]...)
	newSlice := make([]Person, place+1)
	copy(newSlice, slice[:place])
	newSlice[place] = val
	slice = append(newSlice, slice[place:]...)
	return slice
}
