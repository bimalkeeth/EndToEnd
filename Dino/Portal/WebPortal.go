package Portal

import (
	"fmt"
	"net/http"
)

func RunWebPortal(addr string )error{
	http.HandleFunc("/",rootHandler)
	err:=http.ListenAndServe(addr,nil)
	return err
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {

   fmt.Fprintf(writer,"Welcome to web portal",request.RemoteAddr)

}