package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("myTextFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() { // internally, it advances token based on sperator
		fmt.Println(scanner.Text()) // token in unicode-char

	}
}

//func main(){
//f, err := os.Open(os.Args[1])
//if err != nil{
//fmt.Println("Error:", err)
//os.Exit(1)
//}

//io.Copy(os.Stdout, f)
//}
