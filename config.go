package main

import (
	"path/filepath"

	"github.com/astaxie/beego/config"

	"util/etcd"
	"util/logs"

	"core/net/lan"
)

var _ = logs.Debug

//
type Config struct {
	SrvName string

	// server
	LanCfg  *lan.LanCfg
	EtcdCfg *etcd.SrvCfg
}

func (this *Config) init(fileName string) bool {
	confd, e := config.NewConfig("ini", fileName)
	if e != nil {
		logs.Panicln("load config file failed! file:", fileName, "error:", e)
	}

	//[scribe]
	//open=false
	//addr=localhost:7915

	//[server]
	this.SrvName = confd.String("server::name")
	srvAddr := confd.String("server::addr")
	this.LanCfg = lan.NewLanCfg(this.SrvName, srvAddr)

	//[etcd]
	this.EtcdCfg = &etcd.SrvCfg{}
	this.EtcdCfg.EtcdAddrs = confd.Strings("etcd::addrs")
	this.EtcdCfg.SrvAddr = srvAddr
	this.EtcdCfg.SrvRegPath = confd.String("etcd::reg_path") // to do
	this.EtcdCfg.SrvRegUpTick = confd.DefaultInt64("etcd::reg_uptick", 2000)

	this.EtcdCfg.WatchPaths = confd.Strings("etcd::watch_path")

	//#close client notify
	//close_notify_must=match;data
	//close_notify_cached=battle

	// echo
	logs.Info("gate config:%+v", *this)

	return true
}

//
var g_config = &Config{}

func Cfg() *Config {
	return g_config
}

//
func LoadConfig(confPath string) bool {
	// config
	confFile := filepath.Clean(confPath + "/self.ini")

	return g_config.init(confFile)
}

//
func SrvId() string {
	return Cfg().LanCfg.ServerId()
}

//
func SrvName() string {
	return Cfg().LanCfg.Name
}

// to do add check func
