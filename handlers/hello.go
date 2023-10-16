package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"github.com/iMalyk2001/gomicro/handlers"
)
type Hello struct{
	l *log.Logger
}
func NewHello(l *log.Logger) *Hello{
	return &Hello{l}

}


func (h *Hello) ServeHTTP(rw http.ResponseWriter, r*http.Request){
	h.l.Println("Hello World")
	
	d, err:= io.ReadAll(r.Body)
	if err != nil{
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return

	}

	fmt.Fprintf(rw, "Hello %s", d)

}
