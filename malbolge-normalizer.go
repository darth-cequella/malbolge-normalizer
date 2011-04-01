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
	//Content readed, let's preset content to normalized version
	for i,b := range content {
		content[i] = byte( (int(b)+i)%94 )
		
		aux := op[content[i]]
		if aux == 0 { aux = byte('O') } //If it isn't a possible op, replace with uppercas O
		content[i] = aux
	}
	
	fmt.Println(string(content))
}

func numberNormalize(malbolgeFile string) {
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
	//Content readed, let's preset content to normalized version
	var numberContent []int
	for i,b := range content {
		aux := int( int(b)+i)%94
		numberContent = append(numberContent, aux)
	}
	
	fmt.Println(numberContent)
}
func expand(malbolgeFile string) {
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
	//Content readed, let's extend content from normalized version
	opContent := ""
	for i,b := range content {
		var aux byte

		switch(b){
			case 'i':
				aux = byte(4)
			case '<':
				aux = byte(5)
			case '/':
				aux = byte(23)
			case '*':
				aux = byte(29)
			case 'j':
				aux = byte(40)
			case 'p':
				aux = byte(62)
			case 'o':
				aux = byte(68)
			case 'v':
				aux = byte(81)
		}

		aux -= byte(i)
		if int(aux) < 33 {
			aux += byte(94)
		} else if int(aux) > 126 {
			 aux -= byte(94)
		}

		opContent += string(aux)
	}
	
	fmt.Println(opContent)
}

func main() {
	args := os.Args[1:]
	
	switch args[0]{
		case "-d":
			numberNormalize(args[1])
		case "-n":
			normalize(args[1])
		case "-e":
			expand(args[1])
	}
}
 