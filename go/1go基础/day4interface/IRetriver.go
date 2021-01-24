package day4interface

type IRetriver interface {
	Get(url string) string
}
