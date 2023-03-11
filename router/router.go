package router

import "net/http"

type Router interface {
	Get(uri string, f func(writer http.ResponseWriter, request *http.Request))
	Put(uri string, f func(writer http.ResponseWriter, request *http.Request))
	Post(uri string, f func(writer http.ResponseWriter, request *http.Request))
	Delete(uri string, f func(writer http.ResponseWriter, request *http.Request))
	Serve(port string)
}
