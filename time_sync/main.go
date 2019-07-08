package main

import (
	"fmt"
	"time"
)

const oneDay uint = 60 * 60 * 24

const (
	offset    uint = 7200
	createdAt uint = 1550577322
	givenTime uint = createdAt + 11*3600
)

func isOutdated(givenTime, offset, createdAt uint) bool {
	yearNow, monthNow, dayNow := time.Unix(int64(givenTime+offset), 0).UTC().Date()
	year, month, day := time.Unix(int64(createdAt+offset), 0).UTC().Date()
	return !(yearNow == year && monthNow == month && dayNow == day)
}

func main() {
	//
	fmt.Printf("local time: % v \n", time.Now().Local())
	//

	fmt.Printf("=> createdAt UTC: %v\n", time.Unix(int64(createdAt), 0).UTC())
	fmt.Printf("==> createdAt Local: %v\n \n", time.Unix(int64(createdAt), 0))

	fmt.Printf("=> givenTime UTC: %v\n", time.Unix(int64(givenTime), 0).UTC())
	fmt.Printf("==> givenTime Local: %v\n", time.Unix(int64(givenTime), 0))

	fmt.Println(isOutdated(givenTime, offset, createdAt))

}
