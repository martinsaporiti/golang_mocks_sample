package services

type User interface {
	Save(name string) (int, error)
}
