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
	Placement string
	Lastname  string
	Firstname string
	Table     string
	Table2    string
	Table3    string
}

// Main function:
func main() {

	// // 32 people:
	// var waiter []Person
	// // 10 people:
	// var kitchen []Person
	// // 8 people per, 31 total, 249 in all:
	// var table [][]Person

	// Reading and parsing the original CSV file:

	// runBody("seating.csv", "first.csv", 1)
	// runBody("first.csv", "second.csv", 2)
	// runBody("second.csv", "third.csv", 3)

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
			Placement: line[0],
			Firstname: line[1],
			Lastname:  line[2],
			//Table:     line[3],
		})
	}
	var peopleSlice = people
	var slicedPeople = Shuffle(peopleSlice)
	initFile("first.csv")
	iterateAndChoose(slicedPeople, "first.csv", 1)

	csvFile, _ = os.Open("first.csv")
	csvFile.Close()
	people = nil
	csvFile, _ = os.Open("first.csv")
	reader = csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		people = append(people, Person{
			Placement: line[0],
			Firstname: line[1],
			Lastname:  line[2],
			// Table:     line[3],
			// Table2:    line[4],
		})
	}
	peopleSlice = people
	slicedPeople = Shuffle(peopleSlice)
	initFile("second.csv")
	iterateAndChoose(slicedPeople, "second.csv", 2)

	csvFile, _ = os.Open("second.csv")
	csvFile.Close()
	people = nil
	csvFile, _ = os.Open("second.csv")
	reader = csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		people = append(people, Person{
			Placement: line[0],
			Firstname: line[1],
			Lastname:  line[2],
			// Table:     line[3],
			// Table2:    line[5],
			// Table3:    line[6],
		})
	}
	peopleSlice = people
	slicedPeople = Shuffle(peopleSlice)
	initFile("third.csv")
	iterateAndChoose(slicedPeople, "third.csv", 3)

}

// Function to read the previous file, create the second, and do the iteration thingamajig.
func runBody(firstFile string, secondFile string, iteration int) {
	var people []Person
	var csvFile, _ = os.Open(firstFile)
	csvFile.Close()
	people = nil
	csvFile, _ = os.Open(firstFile)
	var reader = csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		people = append(people, Person{
			Placement: line[0],
			Firstname: line[1],
			Lastname:  line[2],
		})
	}
	var peopleSlice = people
	var slicedPeople = Shuffle(peopleSlice)
	initFile(secondFile)
	iterateAndChoose(slicedPeople, secondFile, iteration)

	csvFile, _ = os.Open(secondFile)
	csvFile.Close()
	people = nil
	csvFile, _ = os.Open(secondFile)
	reader = csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		people = append(people, Person{
			Placement: line[0],
			Firstname: line[1],
			Lastname:  line[2],
		})
	}
	peopleSlice = people

	cleanUp(peopleSlice)

	fmt.Println(peopleSlice)

}

// Function to "clean-up" the final file. Places all in numberic/alphabetical order. Will add in an ability to re-pars the files to combine them all together.
func cleanUp(slice []Person) []Person {
	initFile("final.csv")
	var secSlice []Person
	d := slice
	fmt.Println(d)

	for i, v := range d {
		var (
			toGo, kind = strconv.Atoi(v.Placement)
		)
		if kind != nil {
			fmt.Println(kind)
		}
		fmt.Println(v)
		secSlice = append(rearrange(int(i), toGo, d))
	}
	fmt.Println(secSlice)
	return secSlice
}

// Function to move a slice item:
func rearrange(remove int, place int, input []Person) []Person {
	slice := input
	val := slice[remove]
	slice = append(slice[:remove], slice[remove+1:]...)
	newSlice := make([]Person, place+1)
	copy(newSlice, slice[:place])
	newSlice[place] = val
	slice = append(newSlice, slice[place:]...)
	fmt.Println("SLICE:", slice)
	return slice
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
func makeFile(slice []Person, num int, seatType int, title string, iteration int) {

	var name string

	f, err := os.OpenFile(title, os.O_APPEND|os.O_WRONLY, 0644)
	d := slice

	for _, v := range d {
		if seatType == 3 {
			v.Table = "Waiter"
			name = v.Placement + "," + v.Lastname + "," + v.Firstname
		} else if seatType == 2 {
			v.Table = "KC"
			name = v.Placement + "," + v.Lastname + "," + v.Firstname
		} else {
			v.Table = strconv.Itoa(num)
			name = v.Placement + "," + v.Lastname + "," + v.Firstname
		}

		fmt.Fprintln(f, name+","+v.Table+","+v.Table2)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Primary init of the file:
func initFile(title string) {
	f, err := os.Create(title)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
}

// Function that iterates through the slice and chooses people to go to certain positions based on index.
// Also calls in the file creation function.
func iterateAndChoose(slicedPeople []Person, title string, iteration int) []Person {
	var originalGroup = slicedPeople

	// choose the first 10 to be kitchen crew:
	var nextGroup = chooseNext(slicedPeople, 10)
	makeFile(nextGroup, 10, 2, title, iteration)

	// remove the first 10 from the main list:
	removeIndex(10, slicedPeople)

	// choose the next 31 to be waiters:
	nextGroup = chooseNext(slicedPeople, 32)
	makeFile(nextGroup, 32, 3, title, iteration)

	// remove the next 31 from the main list:
	removeIndex(32, slicedPeople)

	// append all tables to CSV file:
	for i := 1; i < 32; i++ {
		var table = chooseNext(slicedPeople, 8)
		makeFile(table, i, 1, title, iteration)
		removeIndex(8, slicedPeople)
	}
	fmt.Println("all completed succesfully!")

	return originalGroup
}
