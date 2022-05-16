package pkg

import "fmt"

type Device struct {
	Name string
	Next Service
}

func (device *Device) Execute(data *Data) {
	if data.GetSource {
		fmt.Printf("Даные [%s] уже получены \n", device.Name)
		device.Next.Execute(data)
		return
	}
	data.GetSource = true
	device.Next.Execute(data)
}

func (device *Device) SetNext(svc Service) {
	device.Next = svc
}
