package controllers

type IController interface {
	Find()
	FindById(id string)
	Create(entity []string)
	Update(id string)
	Delete(id string)
}
