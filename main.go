package main

import (
	"flag"
	"log"
	"net"
)

var (
	policy = []byte(`<?xml version="1.0"?>
<!DOCTYPE cross-domain-policy SYSTEM "http://www.macromedia.com/xml/dtds/cross-domain-policy.dtd">
<cross-domain-policy>
<allow-access-from domain="*" to-ports="*" secure="false" />
</cross-domain-policy>`)
	addr = flag.String("addr", ":843", "host:port for bind, default :843")
)

func Listener(addr string) (listener net.Listener) {
	var err error
	if listener, err = net.Listen("tcp", addr); err != nil {
		log.Fatal(err)
	}
	return
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	conn.Write(policy)
}

func Handler(listener net.Listener) {
	for {
		if conn, err := listener.Accept(); err == nil {
			go HandleConnection(conn)
		} else {
			log.Println(err)
		}
	}
}

func main() {
	flag.Parse()
	Handler(Listener(*addr))
}
