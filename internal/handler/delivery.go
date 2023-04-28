package handler

type DeliveryHandler interface {
	Create()
	Delete()
	Update()
	Get()
}

type DeliveryHandle struct {
}

func NewDeliveryHandler() DeliveryHandler {
	return &DeliveryHandle{}
}

func (c *DeliveryHandle) Create() {

}

func (c *DeliveryHandle) Delete() {

}

func (c *DeliveryHandle) Update() {

}

func (c *DeliveryHandle) Get() {

}
