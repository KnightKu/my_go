package main

import "fmt"
import "flag"
import "bufio"
import "os"
import "io"
import "strconv"
import "time"

import "alg/qsort"
import "alg/bubblesort"

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to store string result")
var alg *string = flag.String("a", "qsort", "Sort algorihtm")

func readValues(infile *string) (values []int, err error) {
	file, err := os.Open(*infile)
	if err != nil {
		fmt.Println("Failed to open input file:", *infile)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("Too long line...")
			return
		}

		str := string(line)
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile *string) error {
	file, err := os.Create(*outfile)
	if err != nil {
		fmt.Println("Failed to create outfile:", *outfile)
		return err
	}
	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}

	return nil
}

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile=", *infile, ",outfile=", *outfile, ",algorihtm=", *alg)
	}

	values, err := readValues(infile)
	if err == nil {
		fmt.Println("Get values:", values)
		t1 := time.Now()
		switch *alg {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Print("Unsupported algorihtm")
		}
		t2 := time.Now()

		fmt.Println("Sorting done, costs:", t2.Sub(t1))

		writeValues(values, outfile)
	} else {
		fmt.Println(err)
	}
}
