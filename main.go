package main

import (
	"log"
	"net/http"
	"os"
	"gomicro/handlers"
	"os/signal"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	handl := handlers.NewHello(l)
	ghandl := handlers.NewGoodbye(l)
	sm := http.NewServeMux()
	
	sm.Handle("/", handl)
	sm.Handle("/goodbye", ghandl)

	s := &http.server{
		addr: ":9090",
		handler: sm,
		idleTimeout: 120*time.Second,
		readTimeout: 1*time.Second,
		writeTimeout: 1*time.Second,
	}
	




go func(){
	err := s.ListenAndServe()
	if err != nil{
		l.Fatal(err)
	}
}()

sigChan := make(chan os.Signal)
signal.notify(sigChan, os.Interrupt)
signal.notify(sigChan, os.Kill)

sig <-sigChan
l.Println("Recieved terminate, graceful shutdown", sig)
tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
s.Shutdown(tc)