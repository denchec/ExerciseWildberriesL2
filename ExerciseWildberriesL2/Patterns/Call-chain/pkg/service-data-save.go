package pkg

import "fmt"

type SaveDataService struct {
	Next Service
}

func (save *SaveDataService) Execute(data *Data) {
	if data.GetSource {
		fmt.Println("Данные не обработаны")
		return
	}
	fmt.Println("Данные сохранены")
}

func (save *SaveDataService) SetNext(svc Service) {
	save.Next = svc
}
