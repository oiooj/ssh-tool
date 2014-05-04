package main

import (
	"fmt"
	"net"
	"flag"
	//"time"
	"bytes"
	"github.com/flynn/go-crypto-ssh"
	"github.com/mikioh/ipaddr"
)



// An SSH client is represented with a ClientConn. Currently only
// the "password" authentication method is supported.
//
// To authenticate with the remote server you must pass at least one
// implementation of AuthMethod via the Auth field in ClientConfig.


var (
	username = flag.String("u", "root", "ssh user name")
	password = flag.String("p", "password", "ssh user password")
	hosts    = flag.String("h", "192.168.0.1/24", "ssh hosts, use CIDR")
	command  = flag.String("c", "w", "command")
	timeout  = flag.Int("t", 3, "timout duration (s)")
    )


func work_handler(ip string){
	config := &ssh.ClientConfig{
		User: *username,
		Auth: []ssh.AuthMethod{
			ssh.Password(*password),
		},
	}
	client, err := ssh.Dial("tcp", ip+":22", config, *timeout)
	if err != nil {
		fmt.Println("\033[1;31m"+"[Failed]  "+"\033[0m" + ip + "  Failed to dial:" + err.Error())
		return
	}

	session, err := client.NewSession()
	if err != nil {
		fmt.Println("\033[1;31m"+"[Failed]  "+"\033[0m" + ip + "  Failed to create session:" + err.Error())
		return
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(*command); err != nil {
		fmt.Println("\033[1;31m"+"[Failed]  "+"\033[0m" + ip + "  Failed to run command")
		return
	}
	fmt.Println("\033[1;32m"+"[success] "+"\033[0m" + ip)
	fmt.Println("\033[1;34m"+b.String()+"\033[0m")
}


func main(){
	flag.Parse()
	_, ipn, err := net.ParseCIDR(*hosts)
	if err != nil {
		fmt.Println("NOT CIDR IPs")
		return
	}
	nbits, _ := ipn.Mask.Size()
	p, err := ipaddr.NewPrefix(ipn.IP, nbits)
	if err != nil {
		fmt.Println("NewPrefix ERROR")
	}
	//fmt.Println(p.Addr(), p.LastAddr(), p.Len(), p.Netmask(), p.Hostmask())
	hosts := p.Hosts(p.Addr())
	for _, host := range hosts {
		work_handler(host.String())
	}

}
