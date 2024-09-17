package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/gonum/mat"
)

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}
func main() {
	// Read file
	f, err := os.Open("./data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// File to Dataframe
	df := dataframe.ReadCSV(f)
	fmt.Println(df)
	fmt.Println(df.Describe())
}
