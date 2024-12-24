package action

import (
	"github.com/acoderup/goserver/core/logger"
	"github.com/acoderup/goserver/core/netlib"
	"github.com/acoderup/goserver/srvlib"
	"github.com/acoderup/goserver/srvlib/protocol"
)

func init() {
	netlib.Register(int(protocol.SrvlibPacketID_PACKET_SS_MULTICAST), &protocol.SSPacketMulticast{}, MulticastHandler)
}

// MulticastHandler 处理广播消息
func MulticastHandler(s *netlib.Session, packetid int, data interface{}) error {
	if mp, ok := data.(*protocol.SSPacketMulticast); ok {
		pd := mp.GetData()
		sis := mp.GetSessions()
		for _, si := range sis {
			ns := getSession(si)
			if ns != nil {
				ns.Send(int(mp.GetPacketId()), pd)
			}
		}
	}
	return nil
}

func getSession(su *protocol.MCSessionUnion) *netlib.Session {
	cs := su.GetMccs()
	if cs != nil {
		return srvlib.ClientSessionMgrSington.GetSession(cs.GetSId())
	}

	ss := su.GetMcss()
	if ss != nil {
		return srvlib.ServerSessionMgrSington.GetSession(int(ss.GetSArea()), int(ss.GetSType()), int(ss.GetSId()))
	}

	return nil
}

// MulticastMessage 消息转发
// areaId 区域ID
// serverType 服务器类型
// sis 消息处理的目标服务器或客户端
func MulticastMessage(areaId, serverType int, packetId int, pack interface{}, sis ...*protocol.MCSessionUnion) error {
	if packetId == 0 || pack == nil {
		return nil
	}
	data := &protocol.SSPacketMulticast{
		Sessions: sis,
		PacketId: int32(packetId),
	}
	if byteData, ok := pack.([]byte); ok {
		data.Data = byteData
	} else {
		byteData, err := netlib.MarshalPacket(packetId, pack)
		if err == nil {
			data.Data = byteData
		} else {
			logger.Logger.Errorf("CreateMulticastPacket err:%v", err)
			return err
		}
	}
	srvlib.ServerSessionMgrSington.Broadcast(int(protocol.SrvlibPacketID_PACKET_SS_MULTICAST), data, areaId, serverType)
	return nil
}

// MulticastMessageToServer 消息转发
// s 转发服务器连接
// sis 消息处理的目标服务器或客户端
func MulticastMessageToServer(s *netlib.Session, packetId int, pack interface{}, sis ...*protocol.MCSessionUnion) error {
	if packetId == 0 || pack == nil || s == nil {
		return nil
	}
	data := &protocol.SSPacketMulticast{
		Sessions: sis,
		PacketId: int32(packetId),
	}
	if byteData, ok := pack.([]byte); ok {
		data.Data = byteData
	} else {
		byteData, err := netlib.MarshalPacket(packetId, pack)
		if err == nil {
			data.Data = byteData
		} else {
			logger.Logger.Errorf("CreateMulticastPacket err:%v", err)
			return err
		}
	}
	s.Send(int(protocol.SrvlibPacketID_PACKET_SS_MULTICAST), data)
	return nil
}
