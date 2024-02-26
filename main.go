package main

import (
	"fmt"
	"os"
)

type Interpreter struct {
	Memory [100]int 
	InstPointer int // Instruction Pointer
	MemPointer int // Memory Pointer
	AddStack []int // Address Stack
	Program string
	Input string
	Output string
}

func NewInterpreter(program string) *Interpreter {
	return &Interpreter{
		Program: program,
	}
}

func (i *Interpreter) Reset() {
	i.Memory = [100]int{}
	i.InstPointer = 0
	i.MemPointer = 0
	i.AddStack = []int{}
	i.Output = ""
	i.Input = ""
}

func (i *Interpreter) GetInput() int{
	val := 0
	if(len(i.Input) > 0) {
		val = int(i.Input[0])
		i.Input = i.Input[1:]
	}
	return val
}

func (i *Interpreter) SetOutput() {
	i.Output += string(rune(i.Memory[i.MemPointer]))
}

func (i *Interpreter) Interpret() {
	eof := false

	for !eof {
		
		if i.InstPointer >= len(i.Program) || i.InstPointer < 0 {
			eof = true
			break
		}

		switch i.Program[i.InstPointer] {
			case '>':
				i.MemPointer++
			case '<':
				if i.MemPointer > 0 {
					i.MemPointer--
				}
			case '+':
				i.Memory[i.MemPointer]++
			case '-':
				i.Memory[i.MemPointer]--
			case '.':
				i.SetOutput()
			case ',':
				i.Memory[i.MemPointer] = i.GetInput()
			case '[':
				if i.Memory[i.MemPointer] != 0 {
					i.AddStack = append(i.AddStack, i.InstPointer)
				} else {
					count := 1
					for count > 0 {
						i.InstPointer++
						if i.Program[i.InstPointer] == '[' {
							count++
						} else if i.Program[i.InstPointer] == ']' {
							count--
						}
					}
				}
			case ']':
				i.InstPointer, i.AddStack = i.AddStack[len(i.AddStack)-1] - 1, i.AddStack[:len(i.AddStack)-1]
		}

		i.InstPointer++
	}

	if len(i.Output) > 0 {
		fmt.Printf("%v\n", i.Output)
	}
}

func main() {

	args := os.Args

	if len(args) > 1 {
		fileArg := args[1]
		if len(args) > 2 {
			fileArg = args[2]
		}
		i := NewInterpreter(GetProgStringFromFile(fileArg))
		i.Interpret()
		return
	} else {
		fmt.Println("File not found. Please provide a file path.")
	}

	progString := GetProgStringFromFile(args[1])
	fmt.Println(progString)

	i := NewInterpreter(progString);
	i.Interpret()
}

func GetProgStringFromFile(path string) string {

	fmt.Println("Reading file: ", path)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error: ", err)
		return ""
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Error: ", err)
		return ""
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println("Error: ", err)
		return ""
	}

	return string(bs)
}