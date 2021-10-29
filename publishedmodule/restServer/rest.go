package restServer

import (
	"fmt"
	"log"
	"net/http"
)

type server struct {
}

func (s *server) ServeHTTP(resWriter http.ResponseWriter, req *http.Request) {
	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.WriteHeader(http.StatusOK)
	resWriter.Write([]byte(`{"message":"Hello From Server"}`))

}

func init() {
	fmt.Println("init web api restful")
	//StartServer("127.0.0.1", 80)
}

func StartServer(ip string, port int) {
	server := &server{}
	http.Handle("/", server)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
