package data

type Servers struct {
	Servers []ServerDetails
}

// type ServerItem struct {
// 	Server ServerDetails
// }

type ServerDetails struct {
	Id             string
	Name           string
	IpAddress      string
	Online         bool
	LastTimeOnline string
	Tags           []string
}

var MyServers = Servers{}

// func (s *Servers) SrvItems() []ServerItem {
// 	return s.Items
// }
