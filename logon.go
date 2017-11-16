package main

import (
	_ "net/http/pprof"

	"util/logs"

	"core/server"

	"share/pipe"
)

var _ = logs.Debug

//
type Logon struct {
	server.Server
}

//
func NewLogon() *Logon {
	return &Logon{}
}

//
func (this *Logon) Init() bool {
	// config
	if !LoadConfig("conf/") {
		return false
	}

	// monitor// to do
	//go http.ListenAndServe("localhost:6886", nil)

	// recv/send msg among servers
	pipe.Init(Cfg().LanCfg, Cfg().EtcdCfg, handleMsgs)

	return true
}

//
func (this Logon) String() string {
	return "logon"
}
