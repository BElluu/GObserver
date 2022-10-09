package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/tatsushid/go-fastping"
	"net"
	"time"
)

func UniqueSlug() string {
	b := make([]byte, 4)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func PingTarget(ipAddress string) bool {
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip", ipAddress)
	if err != nil {
		return false
	}
	//p.MaxRTT = 10000

	p.AddIPAddr(ra)
	found := false

	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		found = true
	}

	p.OnIdle = func() {
		fmt.Println("finish")
	}

	err = p.Run()

	return found
}
