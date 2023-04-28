package handler

type CourierHandler interface {
	Create()
	Delete()
	Update()
	Get()
}

type CourierHandle struct {
}

func NewCourierHandler() CourierHandler {
	return &CourierHandle{}
}

func (c *CourierHandle) Create() {

}

func (c *CourierHandle) Delete() {

}

func (c *CourierHandle) Update() {

}

func (c *CourierHandle) Get() {

}
