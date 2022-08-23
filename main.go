package main

import ( 
	"fmt"
	"github.com/rohini21m/PortScanner/port"

)
func main(){
	fmt.Println("Building port scanner with Golang")
	open:= port.ScanPort("tcp","localhost",3306)
    fmt.Printf("port open : %t\n",open)
	results:= port.IntialScan("localhost")
	fmt.Println(results)
}