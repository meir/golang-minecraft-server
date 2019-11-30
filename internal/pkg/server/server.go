package server

import (
	"gomc/internal/pkg/config"
	"gomc/internal/pkg/packets"
	"gomc/internal/pkg/packets/serverbound"
	"io/ioutil"
	"net"
	"strconv"
)

type Server struct {
	Entities []interface{}
	Worlds []interface{}
	Settings config.ServerProperties
}

func NewServer() *Server {
	return &Server{
		[]interface{}{},
		[]interface{}{},
		config.GetServerSettings(),
	}
}

func(s *Server) Start() {
	var ipstr = s.Settings.Server.ServerIP + ":" + strconv.Itoa(s.Settings.Server.ServerPort)
	println("Starting GoMC server on", ipstr)
	host, err := net.Listen("tcp", ipstr)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := host.Accept()
		s.Verbose("Connection incoming from", conn.RemoteAddr().String())
		if err != nil {
			s.Verbose("Failed connection")
			s.Verbose(err.Error())
			continue
		}
		//conn.Write()
		msg, err := ioutil.ReadAll(conn)
		if err != nil {
			s.Verbose("Could not receive message correctly")
			continue
		}
		s.Verbose("message includes: \"" + string(msg) + "\"")
		packet := packets.Parse(msg)
		if packet.PacketID == 0 {
			handshake := serverbound.NewHandshakePacket(msg)
			handshake.Print(handshake)
		}
	}
}
