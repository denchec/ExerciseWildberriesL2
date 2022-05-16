package pkg

import "fmt"

const (
	ServerType   = "server"
	ComputerType = "computer"
	NotebookType = "notebook"
)

type Machine interface {
	GetType() string
	PrintDetails()
}

func FabricMachine(typeName string) Machine {
	switch typeName {
	default:
		fmt.Printf("%s Не существующий тип объекта \n", typeName)
		return nil
	case ServerType:
		return NewServer()
	case ComputerType:
		return NewComputer()
	case NotebookType:
		return NewNotebook()
	}
}
