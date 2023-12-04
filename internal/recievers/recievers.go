package recievers

type Reciever interface {
	SendData(title, description, link string) error
}
