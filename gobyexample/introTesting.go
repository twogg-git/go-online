package main

import (
	"fmt"
	"time"
)

// yyyy-MM-dd HH:mm TimeZone
const FORMAT = "2006-01-02 15:04 MST"

func main() {

	fmt.Println("\nRise and shine Gopher Clock!")

	fmt.Println("\nServer Current time")
	current := time.Now().Local()
	fmt.Println(current.Format(FORMAT))

	fmt.Println("Some samples from intro book")

	var inta float64 = 5
	var floatb float64 = 1.2
	fmt.Println("Testing aritmetic ")
	fmt.Println(inta, " + ", floatb, " = ", inta+floatb)
	fmt.Println(inta, " - ", floatb, " = ", inta-floatb)
	fmt.Println(inta, " * ", floatb, " = ", inta*floatb)
	fmt.Println(inta, " / ", floatb, " = ", inta/floatb)

	var stringc string = "Testing"
	fmt.Println("Testing strings len", len(stringc), " ", stringc[0])

	stringd := "String no asignment "
	stringe := "What is the diference?"
	fmt.Println("Testing no asigments ", stringd, " ", stringe)

	var stringf = stringd + stringe
	fmt.Println("Concat ", stringf)

	fmt.Println("Testing string validations ", (stringd == stringe))
	fmt.Println("Testing string concat ", (stringf == (stringd + stringe)))

	var inputa float64
	fmt.Scanf("%f", &inputa)
	outputa := inputa * 2
	fmt.Println("Input ", inputa, " Output ", outputa)

	months := [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	fmt.Println("Testing Months")
	for i := 0; i < 12; i++ {
		fmt.Println(months[i])
	}

	fmt.Println("Now only even months!")
	for i := 0; i < 12; i++ {
		if i%2 == 1 {
			fmt.Println(months[i])
		}
	}

	fmt.Println("Birthdays!")
	for i := 1; i <= 12; i++ {
		var monthLimit int = 31
		if i%2 == 1 {
			if i == 2 {
				monthLimit = 28
			}
			monthLimit = 30
		}
		for j := 1; j <= monthLimit; j++ {
			if i == 6 && j == 11 {
				fmt.Println(i, "/", j, " Juani")
			}
			if i == 5 && j == 15 {
				fmt.Println(i, "/", j, " Sis")
			}
			if i == 7 && j == 20 {
				fmt.Println(i, "/", j, " Chiki")
			}
			if i == 12 && j == 21 {
				fmt.Println(i, "/", j, " Cuñis")
			}
			if i == 10 && j == 11 {
				fmt.Println(i, "/", j, " Nani")
			}
			if i == 10 && j == 10 {
				fmt.Println(i, "/", j, " MUA")
			}
			if i == 12 && j == 31 {
				fmt.Println(i, "/", j, "Sun/Earth")
			}
		}
	}

	fmt.Println("Birthdays with switch")
	for m := 1; m <= 12; m++ {
		mLimit := 31
		if m == 2 {
			mLimit = 28
		} else if m%2 == 1 {
			mLimit = 30
		}
		for d := 1; d <= mLimit; d++ {
			switch {
			case m == 6 && d == 11:
				fmt.Println(months[m-1], "/", d, "Juani")
			case m == 10 && d == 10:
				fmt.Println(months[m-1], "/", d, "Mua!")
			case m == 5 && d == 15:
				fmt.Println(months[m-1], "/", d, "Sis")
			case m == 7 && d == 20:
				fmt.Println(months[m-1], "/", d, "Chikis")
			case m == 10 && d == 11:
				fmt.Println(months[m-1], "/", d, "Nani")
			case m == 12 && d == 25:
				fmt.Println(months[m-1], "/", d, "Newton")
			}
		}
	}

	var mMonths = make(map[int]string)
	mapKey := 0
	for _, month := range months {
		mMonths[mapKey] = month
		mapKey++
	}
	fmt.Println("How many months ", len(mMonths))
	fmt.Println("Comming one Chikis! ", mMonths[6])
	fmt.Println("My BD ", mMonths[9])
	fmt.Println("Mons ", mMonths[5], " Sis ", mMonths[4], " Nani ", mMonths[9])

	datesMD := [8][2]int{
		{6, 11},
		{10, 10},
		{10, 11},
		{5, 15},
		{12, 21},
		{7, 20},
		{1, 31},
		{12, 25}}
	for date := range datesMD {
		fmt.Println("Dates ", date)
	}

}
