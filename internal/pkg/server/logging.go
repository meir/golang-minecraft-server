package server

import Log "log"

func(s *Server) Verbose(log ...interface{}) {
	if s.Settings.Logging.Verbose {
		Log.Println(log...)
	}
}

func(s *Server) Log(log ...interface{}) {
	Log.Println(log...)
}