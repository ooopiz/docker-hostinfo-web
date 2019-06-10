package main

import (
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

type NetworkInfo struct {
	IfaceName  string
	IfaceAddrs []string
}

func getHostname() string {
	hostname, _ := os.Hostname()
	return hostname
}

func getNetworks() []NetworkInfo {
	var networkInfos []NetworkInfo

	interfaces, _ := net.Interfaces()
	for _, i := range interfaces {
		byNameInterface, _ := net.InterfaceByName(i.Name)
		addresses, _ := byNameInterface.Addrs()
		var addrs []string
		for _, v := range addresses {
			addrs = append(addrs, v.String())
		}
		if addrs != nil {
			networkInfos = append(networkInfos, NetworkInfo{IfaceName: i.Name, IfaceAddrs: addrs})
		}
	}
	return networkInfos
}

func showHostInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("path", r.URL.Path)

	type HostInfo struct {
		Hostname  string
		Timestamp time.Time
		Networks  []NetworkInfo
	}

	items := HostInfo{
		Hostname:  getHostname(),
		Timestamp: time.Now(),
		Networks:  getNetworks(),
	}
	//log.Println(items)

	tmpl, err := template.ParseFiles("./template.html")
	if err != nil {
		log.Fatal("Parse Template Error ...")
	}
	tmpl.Execute(w, items)
}

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/", showHostInfo)
	err := http.ListenAndServe(":9090", handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
