package main

import (
	"fmt"
	"net"
	"net/http"
)

var ip_address string

func main() {
	http.HandleFunc("/", myhandler)

	ip := getIP()
	if ip == nil {
		ip_address = "Unknown"
	} else {
		ip_address = ip.To4().String()
	}
	http.ListenAndServe(":80", nil)
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
				fmt.Println(ip.To4().String())
				return ip
			}
			fmt.Println("ignoring : ", ip)
		}
	}
	return nil
}
func myhandler(w http.ResponseWriter, r *http.Request) {
	htmlContent := `
	<!DOCTYPE html>
<html lang="en">
<head>
  <title>Sample HTML Responsive Template` + ip_address + `</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/css/bootstrap.min.css" rel="stylesheet">
  <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/js/bootstrap.bundle.min.js"></script>
</head>
<body>

<div class="container-fluid p-5 bg-primary text-white text-center">
  <h1>Instance IP : ` + ip_address + ` Bootstrap Response Template </h1>
  <p>Resize this responsive page to see the effect!</p> 
</div>
  
<div class="container mt-5">
  <div class="row">
    <div class="col-sm-4">
      <h3>Column 1</h3>
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit...</p>
      <p>Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris...</p>
    </div>
    <div class="col-sm-4">
      <h3>Column 2</h3>
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit...</p>
      <p>Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris...</p>
    </div>
    <div class="col-sm-4">
      <h3>Column 3</h3>        
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit...</p>
      <p>Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris...</p>
    </div>
  </div>
</div>

</body>
</html>
	`
	fmt.Fprintln(w, htmlContent)
}
