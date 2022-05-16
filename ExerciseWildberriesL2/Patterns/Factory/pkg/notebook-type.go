package pkg

import "fmt"

type Notebook struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewNotebook() Notebook {
	return Notebook{
		Type:    NotebookType,
		Core:    4,
		Memory:  32,
		Monitor: true,
	}
}

func (note Notebook) GetType() string {
	return note.Type
}

func (note Notebook) PrintDetails() {
	fmt.Printf("%s Core: [%d], Memory: [%d], Monitor: [%v] \n", note.Type, note.Core, note.Memory, note.Monitor)
}
