package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"text/template"
)

type HostDetails struct {
	IPAddress string
	Hostname  string
}

var host HostDetails = HostDetails{}

func main() {
	http.HandleFunc("/", myhandler)
	http.HandleFunc("/health", heathCheck)
	hostname, err := os.Hostname()
	if err != nil {
		host.Hostname = "Error while retriving hostname"
	} else {
		host.Hostname = hostname
	}
	ip := getIP()
	if ip == nil {
		host.IPAddress = "Unknown"
	} else {
		host.IPAddress = ip.To4().String()
	}
	http.ListenAndServe(":80", nil)
}

func heathCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Success")
}

func htmlTemplate(w http.ResponseWriter) {
	file, err := os.Open("templates/homepage.html")
	if err != nil {
		fmt.Fprintln(w, "template opening error")
		return
	}
	defer file.Close()
	fileContent, err := io.ReadAll(file)
	if err != nil {
		fmt.Fprintln(w, "template reading error")
		return
	}
	template.New("test").Parse(string(fileContent))
	s := template.Must(template.New("test").Parse(string(fileContent)))
	s.Execute(w, host)
}

func getIP() net.IP {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Println(err)
			return nil
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip.IsPrivate() && ip.To4() != nil {
				return ip
			}
		}
	}
	return nil
}
func myhandler(w http.ResponseWriter, r *http.Request) {
	htmlTemplate(w)
}
