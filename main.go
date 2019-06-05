package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

func showHostInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	var addrInfo string
	ifaces, _ := net.Interfaces()
	// handle err
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			// process IP address
			addrInfo = addrInfo + fmt.Sprintf("%s", ip) + "\n"
		}
	}

	fmt.Fprint(w, "<div>"+strings.Replace(addrInfo, "\n", "<br>", -1)+"</div>")
}

func main() {
	http.HandleFunc("/", showHostInfo)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
