package handler

import (
	"github.com/acoderup/goserver/core/logger"
	"github.com/acoderup/goserver/core/netlib"
	"github.com/acoderup/goserver/srvlib"
	"github.com/acoderup/goserver/srvlib/protocol"
)

// 服务器信息注册，将自己的服务器信息发送给对方

var (
	SessionHandlerSrvRegisteName = "session-srv-registe"
)

type SessionHandlerSrvRegiste struct {
}

func (sfl SessionHandlerSrvRegiste) GetName() string {
	return SessionHandlerSrvRegisteName
}

func (sfl *SessionHandlerSrvRegiste) GetInterestOps() uint {
	return 1<<netlib.InterestOps_Opened | 1<<netlib.InterestOps_Closed
}

func (sfl *SessionHandlerSrvRegiste) OnSessionOpened(s *netlib.Session) {
	registerPacket := &protocol.SSSrvRegiste{
		Id:     int32(netlib.Config.SrvInfo.Id),
		Type:   int32(netlib.Config.SrvInfo.Type),
		AreaId: int32(netlib.Config.SrvInfo.AreaID),
		Name:   netlib.Config.SrvInfo.Name,
		Data:   netlib.Config.SrvInfo.Data,
	}
	s.Send(int(protocol.SrvlibPacketID_PACKET_SS_REGISTE), registerPacket)
	if s.GetSessionConfig().IsClient {
		logger.Logger.Infof("目标服务器：%v 发送服务器信息(客户端)：%v", s.GetSessionConfig().Name, registerPacket)
	} else {
		logger.Logger.Infof("发送服务器信息(服务端)：%v", registerPacket)
	}
}

func (sfl *SessionHandlerSrvRegiste) OnSessionClosed(s *netlib.Session) {
	if s.GetSessionConfig().IsClient {
		logger.Logger.Infof("服务器 %v 断开连接", s.GetSessionConfig().Name)
	} else {
		d, ok := s.GetAttribute(srvlib.SessionAttributeServerInfo).(*protocol.SSSrvRegiste)
		if ok {
			logger.Logger.Infof("服务器断开连接：%v", d.GetName())
		} else {
			logger.Logger.Infof("服务器断开连接")
		}
	}
	srvlib.ServerSessionMgrSington.UnregisteSession(s)
}

func (sfl *SessionHandlerSrvRegiste) OnSessionIdle(s *netlib.Session) {
}

func (sfl *SessionHandlerSrvRegiste) OnPacketReceived(s *netlib.Session, packetid int, logicNo uint32, packet interface{}) {
}

func (sfl *SessionHandlerSrvRegiste) OnPacketSent(s *netlib.Session, packetid int, logicNo uint32, data []byte) {
}

func init() {
	// 服务信息注册中间件
	netlib.RegisteSessionHandlerCreator(SessionHandlerSrvRegisteName, func() netlib.SessionHandler {
		return &SessionHandlerSrvRegiste{}
	})

	// 服务信息注册消息
	netlib.RegisterFactory(int(protocol.SrvlibPacketID_PACKET_SS_REGISTE), netlib.PacketFactoryWrapper(func() interface{} {
		return &protocol.SSSrvRegiste{}
	}))
	netlib.RegisterHandler(int(protocol.SrvlibPacketID_PACKET_SS_REGISTE), netlib.HandlerWrapper(func(s *netlib.Session, packetid int, data interface{}) error {
		if registePacket, ok := data.(*protocol.SSSrvRegiste); ok {
			s.SetAttribute(srvlib.SessionAttributeServerInfo, registePacket)
			srvlib.ServerSessionMgrSington.RegisteSession(s)
		}
		return nil
	}))
}
