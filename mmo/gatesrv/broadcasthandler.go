package main

import (
	"github.com/acoderup/goserver/core/logger"
	"github.com/acoderup/goserver/core/netlib"
	"github.com/acoderup/goserver/srvlib"
	"github.com/acoderup/goserver/srvlib/protocol"
	"google.golang.org/protobuf/proto"
)

var (
	BroadcastMaker = &BroadcastPacketFactory{}
)

type BroadcastPacketFactory struct {
}

type BroadcastHandler struct {
}

func init() {
	netlib.RegisterHandler(int(protocol.SrvlibPacketID_PACKET_SS_BROADCAST), &BroadcastHandler{})
	netlib.RegisterFactory(int(protocol.SrvlibPacketID_PACKET_SS_BROADCAST), BroadcastMaker)
}

func (this *BroadcastPacketFactory) CreatePacket() interface{} {
	pack := &protocol.SSPacketBroadcast{}
	return pack
}

func (this *BroadcastPacketFactory) CreateBroadcastPacket(sp *protocol.BCSessionUnion, packetid int, data interface{}) (proto.Message, error) {
	pack := &protocol.SSPacketBroadcast{
		SessParam: sp,
		PacketId:  int32(packetid),
	}
	if byteData, ok := data.([]byte); ok {
		pack.Data = byteData
	} else {
		byteData, err := netlib.MarshalPacket(packetid, data)
		if err == nil {
			pack.Data = byteData
		} else {
			logger.Logger.Warn("BroadcastPacketFactory.CreateBroadcastPacket err:", err)
			return nil, err
		}
	}
	return pack, nil
}

func (this *BroadcastHandler) Process(s *netlib.Session, packetid int, data interface{}) error {
	if bp, ok := data.(*protocol.SSPacketBroadcast); ok {
		pd := bp.GetData()
		sp := bp.GetSessParam()
		if bcss := sp.GetBcss(); bcss != nil {
			srvlib.ServerSessionMgrSington.Broadcast(int(bp.GetPacketId()), pd, int(bcss.GetSArea()), int(bcss.GetSType()))
		} else {
			BundleMgrSington.Broadcast(int(bp.GetPacketId()), pd)
		}
	}
	return nil
}
