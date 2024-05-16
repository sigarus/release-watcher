package recievers

type Reciever interface {
	SendData(name, release, description, link string) error
	GetName() string
}
