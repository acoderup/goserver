package handler

import (
	"github.com/acoderup/goserver/core/netlib"
	"github.com/acoderup/goserver/srvlib"
)

/*
 服务注册，一个服务器可以提供多个服务端口
*/

// 依赖于 serversessionregiste
// 需要挂接在serversessionregiste之后
var (
	ServiceHandlerServiceRegisteName = "srv-service-handler"
)

type SessionHandlerServiceRegiste struct {
}

func (this SessionHandlerServiceRegiste) GetName() string {
	return ServiceHandlerServiceRegisteName
}

func (this *SessionHandlerServiceRegiste) GetInterestOps() uint {
	return 1<<netlib.InterestOps_Opened | 1<<netlib.InterestOps_Closed
}

func (this *SessionHandlerServiceRegiste) OnSessionOpened(s *netlib.Session) {
	sc := s.GetSessionConfig()
	if sc.IsClient {
		/*报告自己的监听信息*/
		srvlib.ServiceMgr.ReportService(s)
	} else {
		// 这里标记只有监听端才会触发自动连接
		s.SetAttribute(srvlib.SessionAttributeServiceFlag, 1)
	}
}

func (this *SessionHandlerServiceRegiste) OnSessionClosed(s *netlib.Session) {
	sc := s.GetSessionConfig()
	if !sc.IsClient {
		srvlib.ServiceMgr.ClearServiceBySession(s)
	}
}

func (this *SessionHandlerServiceRegiste) OnSessionIdle(s *netlib.Session) {
}

func (this *SessionHandlerServiceRegiste) OnPacketReceived(s *netlib.Session, packetid int, logicNo uint32, packet interface{}) {
}

func (this *SessionHandlerServiceRegiste) OnPacketSent(s *netlib.Session, packetid int, logicNo uint32, data []byte) {
}

func init() {
	netlib.RegisteSessionHandlerCreator(ServiceHandlerServiceRegisteName, func() netlib.SessionHandler {
		return &SessionHandlerServiceRegiste{}
	})
}
