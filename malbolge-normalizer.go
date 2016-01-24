package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
)

func normalize(malbolgeFile string) {
	//Operations
	
	op := map[byte]byte{
		4:byte('i'),
		5:byte('<'),
		23:byte('/'),
		39:byte('*'),
		40:byte('j'),
		62:byte('p'),
		68:byte('o'),
		81:byte('v')}
	
	//------------------------------------------------------------------
	
	file, err := os.Open(malbolgeFile) //Open malbolge file
	
	if err != nil {
		fmt.Println("File `"+malbolgeFile+"` not Found")
		log.Fatalln(err)
	}
	
	reader := bufio.NewReader(file) //Create a reader
	
	var content []byte //Catch the content of the file
	
	line,_,err := reader.ReadLine()
	for line != nil {
		if err != nil {
			fmt.Println("Error reading `"+malbolgeFile+"` file")
			log.Fatal(err)
			break
		}
		
		for _, b := range line { //Concat content with new byte array
			content = append(content, b)
		}
		line,_,_ = reader.ReadLine() //Read next line
	}
	
	//------------------------------------------------------------------
	//Content readed let's preset content to normalized version
	for i,b := range content {
		content[i] += byte(i%94)
		
		aux := op[b]
		if aux == 0 { aux = byte('O') } //If it isn't a possible op, replace with uppercas O
		content[i] = aux
	}
	
	fmt.Println(string(content))
}

func extend() {
}

func main() {
	args := os.Args[1:]
	
	normalize(args[0])
}
