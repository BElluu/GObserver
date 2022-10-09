package data

type Servers struct {
	Server []ServerDetails
}

type ServerDetails struct {
	Id             string
	Name           string
	IpAddress      string
	Online         bool
	LastTimeOnline string
	Tags           []string
}

var MyServers = Servers{}
