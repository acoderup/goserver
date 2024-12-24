package action

import (
	"github.com/acoderup/goserver/core/logger"
	"github.com/acoderup/goserver/core/netlib"
	"github.com/acoderup/goserver/srvlib"
	"github.com/acoderup/goserver/srvlib/protocol"
)

func init() {
	netlib.Register(int(protocol.SrvlibPacketID_PACKET_SS_BROADCAST), &protocol.SSPacketBroadcast{}, BroadcastHandler)
}

// BroadcastHandler 处理广播消息
func BroadcastHandler(s *netlib.Session, packetId int, data interface{}) error {
	if bp, ok := data.(*protocol.SSPacketBroadcast); ok {
		pd := bp.GetData()
		sp := bp.GetSessParam()
		if v := sp.GetBcss(); v != nil {
			srvlib.ServerSessionMgrSington.Broadcast(int(bp.GetPacketId()), pd, int(v.GetSArea()), int(v.GetSType()))
			return nil
		}
		srvlib.ClientSessionMgrSington.Broadcast(int(bp.GetPacketId()), pd)
	}
	return nil
}

// BroadcastMessage 消息转发
// areaId 区域ID
// serverType 服务器类型
// sp 消息处理的目标服务器或客户端
func BroadcastMessage(areaId, serverType int, packetId int, pack interface{}, sp *protocol.BCSessionUnion) error {
	if packetId == 0 || pack == nil {
		return nil
	}
	data := &protocol.SSPacketBroadcast{
		SessParam: sp,
		PacketId:  int32(packetId),
	}
	if byteData, ok := pack.([]byte); ok {
		data.Data = byteData
	} else {
		byteData, err := netlib.MarshalPacket(packetId, pack)
		if err == nil {
			data.Data = byteData
		} else {
			logger.Logger.Errorf("CreateBroadcastPacket err:%v", err)
			return err
		}
	}
	srvlib.ServerSessionMgrSington.Broadcast(int(protocol.SrvlibPacketID_PACKET_SS_BROADCAST), data, areaId, serverType)
	return nil
}

// BroadcastMessageToServer 消息转发
// s 转发服务器连接
// sp 消息处理的目标服务器或客户端
func BroadcastMessageToServer(s *netlib.Session, packetId int, pack interface{}, sp *protocol.BCSessionUnion) error {
	if packetId == 0 || pack == nil || s == nil {
		return nil
	}
	data := &protocol.SSPacketBroadcast{
		SessParam: sp,
		PacketId:  int32(packetId),
	}
	if byteData, ok := pack.([]byte); ok {
		data.Data = byteData
	} else {
		byteData, err := netlib.MarshalPacket(packetId, pack)
		if err == nil {
			data.Data = byteData
		} else {
			logger.Logger.Errorf("CreateBroadcastPacket err:%v", err)
			return err
		}
	}
	s.Send(int(protocol.SrvlibPacketID_PACKET_SS_BROADCAST), data)
	return nil
}
