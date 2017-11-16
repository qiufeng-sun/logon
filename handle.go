package main

import (
	"util/logs"

	"core/net/dispatcher/pb"

	"share/handler"
	"share/msg"
	"share/pipe"
)

var _ = logs.Debug

//
func handleMsgs(f *pb.PbFrame) {
	handler.HandleFrame(f)
}

//
func regFunc(msgId msg.EMsg, h func(f *pb.PbFrame)) {
	handler.RegFunc(int32(msgId), h)
}

//
func init() {
	regFunc(msg.EMsg_ID_CSLogin, handleOAuth)
}

//
func handleOAuth(f *pb.PbFrame) {
	// parse
	var m msg.CSLogin
	e := handler.ParseMsgData(f.MsgRaw, &m)
	if e != nil {
		logs.Error("invalid logon msg! fromUrl:%v, error:%v", f.GetSrcUrl(), e)
		// to do
		return
	}
	logs.Info("user logoning: msg=%+v", m)

	// process// to do

	// feedback// to do
	accId := int64(1)
	fb := &msg.SCLogin{}

	pipe.SendMsg(accId, pipe.SrvUrl(), f.GetSrcUrl(), msg.EMsg_ID_SCLogin, fb)
}
