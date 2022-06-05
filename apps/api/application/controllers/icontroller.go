package controllers

type IController interface {
	FindAll()
	FindById(id string)
	Create(entity []string)
	Update(id string)
	Delete(id string)
}
