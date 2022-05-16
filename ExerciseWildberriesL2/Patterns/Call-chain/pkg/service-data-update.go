package pkg

import "fmt"

type UpdateDataService struct {
	Name string
	Next Service
}

func (upd *UpdateDataService) Execute(data *Data) {
	if data.UpdateSource {
		fmt.Printf("Даные в сервисе [%s] уже обработаны \n", upd.Name)
		upd.Next.Execute(data)
		return
	}
	data.UpdateSource = true
	upd.Next.Execute(data)
}

func (upd *UpdateDataService) SetNext(svc Service) {
	upd.Next = svc
}
