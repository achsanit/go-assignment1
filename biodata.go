package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Student struct {
	ID 			string
	FullName 	string
	Address		string
	Job 		string
	JoinReason 	string
}

func main() {
	// get first os argument
	idStudent := os.Args[1]

	// get student by idStudent
	selectedStudent := getStudentById(idStudent)
	
	fmt.Print(
		"ID: ", selectedStudent.ID, "\n",
		"FullName: ", selectedStudent.FullName, "\n",
		"Address: ", selectedStudent.Address, "\n",
		"Job: ", selectedStudent.Job, "\n",
		"Reason: ", selectedStudent.JoinReason, "\n",
	)
}

func getStudentById(id string) Student {
	// new map student with key: string, value: Student
	students := map[string]Student {}

	// read all csv file
	records, err := readCsv("static/biodata.csv")

	if err != nil {
		// return panic when read csv error
		panic("read csv error")
	}

	// add all csv data to students map
	for _, record := range records {
		students[record[0]] = Student{ record[1], record[2], record[5], record[4], record[6]}
	}

	return students[id]
}

// read static csv file return slice of slice string
func readCsv(fileName string) ([][]string, error) {
	f, err := os.Open(fileName) // open file

	if err != nil {
		return [][]string{}, err // if error when open file, return error and empty slice
	}

	defer f.Close() // close file at the end of func

	r := csv.NewReader(f) // read csv

	records, err := r.ReadAll() // read all records of csv

	if err != nil {
		return [][]string{}, err // if error when read all record, return error and empty slice
	}

	return records, err // return recods
}
