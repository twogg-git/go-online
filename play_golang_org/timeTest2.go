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

	fmt.Println("\nCurrent time by Region")
	fmt.Println(getCurrentTimeByRegion("Asia/Singapore"))
	fmt.Println(getCurrentTimeByRegion("Asia/Tokyo"))
	fmt.Println(getCurrentTimeByRegion("Europe/Moscow"))
	fmt.Println(getCurrentTimeByRegion("Europe/Zurich"))
	fmt.Println(getCurrentTimeByRegion("Europe/London"))
	fmt.Println(getCurrentTimeByRegion("America/Caracas"))
	fmt.Println(getCurrentTimeByRegion("America/Bogota"))
	fmt.Println(getCurrentTimeByRegion("America/Denver"))
	fmt.Println(getCurrentTimeByRegion("America/Los_Angeles"))

	fmt.Println("\nCurrent time by Zones ")
	current = time.Now().UTC()
	fmt.Println(getCurrentTimeByZone("CEST"))
	fmt.Println(getCurrentTimeByZone("GMT"))
	fmt.Println(getCurrentTimeByZone("EST"))
	fmt.Println(getCurrentTimeByZone("COT"))
	fmt.Println(getCurrentTimeByZone("PDT",))

	fmt.Println("\nFixedTime")
	fixedTime := getFixedTime(1985, 10, 10, 3, 45, "UTC")
	fmt.Println("? " + fixedTime.Format(FORMAT))
	fmt.Println("> " + getTimeByZone(fixedTime, "PDT",))
	fmt.Println("> " + getTimeByZone(fixedTime, "COT"))
	fmt.Println("> " + getTimeByZone(fixedTime, "CEST"))

}

func getTimeByZone(dateToConvert time.Time, zone string) string {
	return dateToConvert.In(time.FixedZone(zone, getZoneOffset(zone)*60*60)).Format(FORMAT)
}

func getFixedTime(year, month, day, hour, min int, zone string) time.Time {
	loc, _ := time.LoadLocation(zone)
	return time.Date(year, time.Month(month), day, hour, min, 0, 0, loc)
}

func getCurrentTimeByZone(zone string) string {
	current := time.Now().UTC()
	return current.In(time.FixedZone(zone, getZoneOffset(zone)*60*60)).Format(FORMAT)
}

func getCurrentTimeByRegion(region string) string {
	location, _ := time.LoadLocation(region)
	currentDt := time.Now().In(location)
	var output = time.Date(
		currentDt.Year(),
		currentDt.Month(),
		currentDt.Day(),
		currentDt.Hour(),
		currentDt.Minute(),
		0, 0, // ignored secs and nsecs
		location).Format(FORMAT)
	output = output + " (" + region + ")"
	return output
}

func getZoneOffset(zone string) int {
	if zone == "PDT" {
		return -7
	}
	if zone == "COT" {
		return -5
	}
	if zone == "EST" {
		return -4
	}
	if zone == "GMT" {
		return 0
	}
	if zone == "CEST" {
		return +2
	}
	return 0 //GMT
}
