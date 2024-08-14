package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

type HostDetails struct {
	IPAddress string
	Hostname  string
}

var host HostDetails = HostDetails{}

func main() {
	var port string
	flag.StringVar(&port, "port", "80", "http port for the application")
	flag.Parse()
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
	if _, err := strconv.Atoi(port); err != nil {
		port = "80"
	}
	http.ListenAndServe(":"+port, nil)
}

func heathCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Success")
}

func listFiles(directory string) ([]fs.DirEntry, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func htmlTemplate(w http.ResponseWriter) {
	storyindex := rand.Intn(5)
	directory := "./templates" // Replace with the directory you want to list files from
	files, err := listFiles(directory)
	if err != nil {
		fmt.Fprintln(w, "opening directory error")
		return
	}
	//fmt.Println(files)
	file, err := os.Open(directory + string(os.PathSeparator) + files[storyindex].Name())
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
