package main

import(
	
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)
func check(e error){
	if e != nil {
		log.Fatalln("err")
	}
}
func main(){
	csvfile, err := os.Open("mycsv.csv")
	check(err)

	r:= csv.NewReader(csvfile)

	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Question: %s Answer %s\n", record[0], record[1])
	}
}