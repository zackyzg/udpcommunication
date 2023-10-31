/**
 *
 * @package       main
 * @author        YuanZhiGang <zackyuan@yeah.net>
 * @version       1.0.0
 * @copyright (c) 2013-2023, YuanZhiGang
 */

package main

import (
	"fmt"
	"net"
)

func main() {
	// network string, laddr, raddr *UDPAddr
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})

	if err != nil {
		fmt.Println("connect server failed,err:", err)
		return
	}

	defer socket.Close()

	sendData := []byte("hello server")
	_, err = socket.Write(sendData) // 发送数据
	if err != nil {
		fmt.Println("send data failed,err:", err)
		return
	}

	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
	if err != nil {
		fmt.Println("receive data failed,err:", err)
		return
	}

	fmt.Printf("receive:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)

}
