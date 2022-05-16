package pkg

import "fmt"

type Computer struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewComputer() Computer {
	return Computer{
		Type:    ComputerType,
		Core:    8,
		Memory:  64,
		Monitor: true,
	}
}

func (comp Computer) GetType() string {
	return comp.Type
}

func (comp Computer) PrintDetails() {
	fmt.Printf("%s Core: [%d], Memory: [%d], Monitor: [%v] \n", comp.Type, comp.Core, comp.Memory, comp.Monitor)
}
