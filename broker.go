package main

import (
	"log"
	"net"
	"fmt"
	"google.golang.org/grpc"
	"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Operaciones"

)



func main() {

	lis, err := net.Listen("tcp", ":9500")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := Operaciones.Server{}

	grpcServer := grpc.NewServer()

	//REGISTRO DE SERVICIOS
	Operaciones.RegisterOrdenServiceServer(grpcServer, &s)
	Operaciones.RegisterSeguimientoServiceServer(grpcServer, &s)
	//FIN REGISTRO DE SERVICIOS


	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
