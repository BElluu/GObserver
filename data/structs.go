package data

import "time"

type Servers struct {
	Servers []ServerDetails
}

// type ServerItem struct {
// 	Server ServerDetails
// }

type ServerDetails struct {
	Id             int
	Name           string
	IpAddress      string
	Online         bool
	LastTimeOnline time.Time
	Tags           []string
}

var MyServers = Servers{}

// func (s *Servers) SrvItems() []ServerItem {
// 	return s.Items
// }
