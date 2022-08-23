package  port 

import (
	"strconv"
	"net"
	"time"
)

type ScanResult struct {
	Port int
	State string
}


func ScanPort(protocol, hostname string , port int) ScanResult{ //bool{
	result:= ScanResult{Port:port}
	address:= hostname + ":" + strconv.Itoa(port)
	conn,err:= net.DialTimeout(protocol,address,60*time.Second)
	if err!=nil{
	 result.State = "Closed"
		return result//return false
	}
	defer conn.Close()
 result.State = "Open" 
  return result // true
}

func IntialScan(hostname string) []ScanResult{
 var results []ScanResult
	for i:=1;i<=1024;i++{
		results = append(results,ScanPort("tcp",hostname,i))
	}
	return results
}