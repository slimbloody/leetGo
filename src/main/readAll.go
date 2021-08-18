package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./log.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	//content, err := ioutil.ReadAll(file)
	fmt.Println(4<<(^uint(1)>>63))
}
