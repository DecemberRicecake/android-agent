package main


import (
	"fmt"
	"html"
	"net/http"
	"net"
	"github.com/sevlyar/go-daemon"
	"log"
)

func runDaemon() (cntxt *daemon.Context) {
	cntxt = &daemon.Context{
		PidFilePerm: 0644,
		LogFileName: "/sdcard/android-agent.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       022,
	}
	child, err := cntxt.Reborn()
	if err != nil {
		log.Print("Unale to run: ", err)
	}
	if child != nil {
		return nil
	}
	return cntxt
}

func main() {
	ip := getOutboundIP()
	listenip := net.JoinHostPort(net.IP.String(ip), "9999")

	cntxt := runDaemon()
	if cntxt == nil {
		log.Printf("android-agent listening on %s", listenip)
		return
	}
	defer cntxt.Release()
	log.Print("- - - - - - - - - - - - - - -")
	log.Print("daemon started")

	ServeHTTP(listenip)
}


func getOutboundIP() (ip net.IP) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

//func typeof(v interface{}) string {
//	return fmt.Sprintf("%T", v)
//}

func ServeHTTP(ip string){
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(ip, nil)
}


func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("request from %s: %s %q", r.RemoteAddr, r.Method, r.URL)
	fmt.Fprintf(w, "go-daemon: %q", html.EscapeString(r.URL.Path))
	fmt.Fprintln(w, "hello world")
}