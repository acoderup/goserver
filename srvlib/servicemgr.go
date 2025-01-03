package srvlib

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/acoderup/goserver/core/logger"
	"github.com/acoderup/goserver/core/netlib"
	"github.com/acoderup/goserver/srvlib/protocol"
)

var (
	SessionAttributeServiceInfo = &serviceMgr{}
	SessionAttributeServiceFlag = &serviceMgr{}
	ServiceMgr                  = &serviceMgr{servicesPool: make(map[int32]map[int32]*protocol.ServiceInfo)}
)

type ServiceRegisteListener interface {
	OnRegiste([]*protocol.ServiceInfo)
	OnUnregiste(*protocol.ServiceInfo)
}

type serviceMgr struct {
	servicesPool map[int32]map[int32]*protocol.ServiceInfo // srvType:srvId:ServiceInfo
	listeners    []ServiceRegisteListener
}

func (this *serviceMgr) AddListener(l ServiceRegisteListener) ServiceRegisteListener {
	this.listeners = append(this.listeners, l)
	return l
}

func (this *serviceMgr) RegisteService(s *netlib.Session, services []*protocol.ServiceInfo) {
	if this == nil || s == nil || len(services) == 0 {
		return
	}

	// 根据对方提供的服务信息获取对方应该和那些其它服务建立连接，然后让其它服务连接这个刚注册的服务（service）
	s.SetAttribute(SessionAttributeServiceInfo, services) // 保存监听服务配置
	for _, service := range services {
		srvid := service.GetSrvId()
		srvtype := service.GetSrvType()
		if _, has := this.servicesPool[srvtype]; !has {
			this.servicesPool[srvtype] = make(map[int32]*protocol.ServiceInfo)
		}
		if _, exist := this.servicesPool[srvtype][srvid]; !exist {
			this.servicesPool[srvtype][srvid] = service
			logger.Logger.Infof("服务注册：%v", service)
			pack := &protocol.SSServiceInfo{}
			pack.Service = service
			sessiontypes := GetCareSessionsByService(service.GetSrvType())
			areaId := service.GetAreaId()
			for _, v1 := range sessiontypes {
				// 地区和服务类型
				ServerSessionMgrSington.Broadcast(int(protocol.SrvlibPacketID_PACKET_SS_SERVICE_INFO), pack, int(areaId), int(v1))
				logger.Logger.Infof("广播服务信息：%v ==> AreaId:%v ServerType:%v", service.GetSrvName(), areaId, v1)
			}

			//if len(this.listeners) != 0 {
			//	for _, l := range this.listeners {
			//		l.OnRegiste(services)
			//	}
			//}
		}
	}
}

// UnregisteService 广播服务关闭
func (this *serviceMgr) UnregisteService(service *protocol.ServiceInfo) {
	if this == nil || service == nil {
		return
	}

	srvid := service.GetSrvId()
	srvtype := service.GetSrvType()
	if v, has := this.servicesPool[srvtype]; has {
		if ss, exist := v[srvid]; exist && ss == service {
			delete(v, srvid)
			logger.Logger.Infof("服务关闭：%v", service)
			pack := &protocol.SSServiceShut{}
			pack.Service = service
			sessiontypes := GetCareSessionsByService(service.GetSrvType())
			areaId := service.GetAreaId()
			for _, v1 := range sessiontypes {
				ServerSessionMgrSington.Broadcast(int(protocol.SrvlibPacketID_PACKET_SS_SERVICE_SHUT), pack, int(areaId), int(v1))
				logger.Logger.Infof("广播服务关闭：%v ==> AreaId:%v ServerType:%v", service.GetSrvName(), areaId, v1)
			}
			//if len(this.listeners) != 0 {
			//	for _, l := range this.listeners {
			//		l.OnUnregiste(service)
			//	}
			//}
		}
	}

}

// OnRegiste 服务器注册
func (this *serviceMgr) OnRegiste(s *netlib.Session) {
	if this == nil || s == nil {
		return
	}

	if s.GetAttribute(SessionAttributeServiceFlag) == nil {
		return
	}
	// 查看已存在的服务信息列表找当前服务器需要建立连接的服务，然后建立连接
	attr := s.GetAttribute(SessionAttributeServerInfo)
	if attr != nil {
		if srvInfo, ok := attr.(*protocol.SSSrvRegiste); ok && srvInfo != nil {
			services := GetCareServicesBySession(srvInfo.GetType())
			for _, v1 := range services {
				if v2, has := this.servicesPool[v1]; has {
					for _, v3 := range v2 {
						func(si *protocol.ServiceInfo, sInfo *protocol.SSSrvRegiste) {
							pack := &protocol.SSServiceInfo{}
							pack.Service = si
							s.Send(int(protocol.SrvlibPacketID_PACKET_SS_SERVICE_INFO), pack)
							logger.Logger.Infof("建立连接： client(%v)==>server(%v)", srvInfo.GetName(), si.GetSrvName())
						}(v3, srvInfo)
					}
				}
			}
		}
	}
}

func (this *serviceMgr) OnUnregiste(s *netlib.Session) {

}

func (this *serviceMgr) ClearServiceBySession(s *netlib.Session) {
	attr := s.GetAttribute(SessionAttributeServiceInfo)
	if attr != nil {
		if services, ok := attr.([]*protocol.ServiceInfo); ok {
			for _, service := range services {
				this.UnregisteService(service)
			}
		}
		s.RemoveAttribute(SessionAttributeServiceInfo)
	}
}

func (this *serviceMgr) ReportService(s *netlib.Session) {
	acceptors := netlib.GetAcceptors()
	cnt := len(acceptors)
	if cnt > 0 {
		pack := &protocol.SSServiceRegiste{
			Services: make([]*protocol.ServiceInfo, 0, cnt),
		}
		for _, v := range acceptors {
			addr := v.Addr()
			if addr == nil {
				continue
			}
			network := addr.Network()
			s := addr.String()
			ipAndPort := strings.Split(s, ":")
			if len(ipAndPort) < 2 {
				continue
			}

			port, err := strconv.Atoi(ipAndPort[len(ipAndPort)-1])
			if err != nil {
				continue
			}

			sc := v.GetSessionConfig()
			si := &protocol.ServiceInfo{
				AreaId:          int32(sc.AreaId),
				SrvId:           int32(sc.Id),
				SrvType:         int32(sc.Type),
				SrvPID:          int32(os.Getpid()),
				SrvName:         sc.Name,
				NetworkType:     network,
				Ip:              sc.Ip,
				Port:            int32(port),
				WriteTimeOut:    int32(sc.WriteTimeout / time.Second),
				ReadTimeOut:     int32(sc.ReadTimeout / time.Second),
				IdleTimeOut:     int32(sc.IdleTimeout / time.Second),
				MaxDone:         int32(sc.MaxDone),
				MaxPend:         int32(sc.MaxPend),
				MaxPacket:       int32(sc.MaxPacket),
				RcvBuff:         int32(sc.RcvBuff),
				SndBuff:         int32(sc.SndBuff),
				SoLinger:        int32(sc.SoLinger),
				KeepAlive:       sc.KeepAlive,
				NoDelay:         sc.NoDelay,
				IsAutoReconn:    sc.IsAutoReconn,
				IsInnerLink:     sc.IsInnerLink,
				SupportFragment: sc.SupportFragment,
				AllowMultiConn:  sc.AllowMultiConn,
				AuthKey:         sc.AuthKey,
				EncoderName:     sc.EncoderName,
				DecoderName:     sc.DecoderName,
				FilterChain:     sc.FilterChain,
				HandlerChain:    sc.HandlerChain,
				Protocol:        sc.Protocol,
				Path:            sc.Path,
				OuterIp:         sc.OuterIp,
			}
			pack.Services = append(pack.Services, si)
		}
		s.Send(int(protocol.SrvlibPacketID_PACKET_SS_SERVICE_REGISTE), pack)
		logger.Logger.Infof("目标服务器：%v 发送服务信息：%v", s.GetSessionConfig().Name, pack)
	}
}

func (this *serviceMgr) GetServices(srvtype int32) map[int32]*protocol.ServiceInfo {
	if v, has := this.servicesPool[srvtype]; has {
		return v
	}
	return nil
}

func (this *serviceMgr) GetService(srvtype, srvid int32) *protocol.ServiceInfo {
	if v, has := this.servicesPool[srvtype]; has {
		if vv, has := v[srvid]; has {
			return vv
		}
	}
	return nil
}

func init() {
	// service registe
	netlib.RegisterFactory(int(protocol.SrvlibPacketID_PACKET_SS_SERVICE_REGISTE), netlib.PacketFactoryWrapper(func() interface{} {
		return &protocol.SSServiceRegiste{}
	}))
	netlib.RegisterHandler(int(protocol.SrvlibPacketID_PACKET_SS_SERVICE_REGISTE), netlib.HandlerWrapper(func(s *netlib.Session, packetid int, pack interface{}) error {
		if sr, ok := pack.(*protocol.SSServiceRegiste); ok {
			ServiceMgr.RegisteService(s, sr.GetServices())
		}
		return nil
	}))

	// service info
	netlib.RegisterFactory(int(protocol.SrvlibPacketID_PACKET_SS_SERVICE_INFO), netlib.PacketFactoryWrapper(func() interface{} {
		return &protocol.SSServiceInfo{}
	}))
	netlib.RegisterHandler(int(protocol.SrvlibPacketID_PACKET_SS_SERVICE_INFO), netlib.HandlerWrapper(func(s *netlib.Session, packetid int, pack interface{}) error {
		if sr, ok := pack.(*protocol.SSServiceInfo); ok {
			service := sr.GetService()
			if service != nil {
				sc := &netlib.SessionConfig{
					Id:              int(service.GetSrvId()),
					Type:            int(service.GetSrvType()),
					AreaId:          int(service.GetAreaId()),
					Name:            service.GetSrvName(),
					Ip:              service.GetIp(),
					OuterIp:         service.GetOuterIp(),
					Port:            int(service.GetPort()),
					WriteTimeout:    time.Duration(service.GetWriteTimeOut()),
					ReadTimeout:     time.Duration(service.GetReadTimeOut()),
					IdleTimeout:     time.Duration(service.GetIdleTimeOut()),
					MaxDone:         int(service.GetMaxDone()),
					MaxPend:         int(service.GetMaxPend()),
					MaxPacket:       int(service.GetMaxPacket()),
					RcvBuff:         int(service.GetRcvBuff()),
					SndBuff:         int(service.GetSndBuff()),
					IsClient:        true,
					IsAutoReconn:    true,
					AuthKey:         service.GetAuthKey(),
					SoLinger:        int(service.GetSoLinger()),
					KeepAlive:       service.GetKeepAlive(),
					NoDelay:         service.GetNoDelay(),
					IsInnerLink:     service.GetIsInnerLink(),
					SupportFragment: service.GetSupportFragment(),
					AllowMultiConn:  service.GetAllowMultiConn(),
					EncoderName:     service.GetEncoderName(),
					DecoderName:     service.GetDecoderName(),
					FilterChain:     service.GetFilterChain(),
					HandlerChain:    service.GetHandlerChain(),
					Protocol:        service.GetProtocol(),
					Path:            service.GetPath(),
				}
				if !sc.AllowMultiConn && netlib.ConnectorMgr.IsConnecting(sc) {
					logger.Logger.Warnf("%v:%v %v:%v had connected, not allow multiple connects", sc.Id, sc.Name, sc.Ip, sc.Port)
					return nil
				}
				sc.Init()
				err := netlib.Connect(sc)
				if err != nil {
					logger.Logger.Warn("connect server failed err:", err)
				}
			}
		}
		return nil
	}))

	// service shutdown
	netlib.RegisterFactory(int(protocol.SrvlibPacketID_PACKET_SS_SERVICE_SHUT), netlib.PacketFactoryWrapper(func() interface{} {
		return &protocol.SSServiceShut{}
	}))
	netlib.RegisterHandler(int(protocol.SrvlibPacketID_PACKET_SS_SERVICE_SHUT), netlib.HandlerWrapper(func(s *netlib.Session, packetid int, pack interface{}) error {
		if sr, ok := pack.(*protocol.SSServiceShut); ok {
			service := sr.GetService()
			if service != nil {
				netlib.ShutConnector(service.GetIp(), int(service.GetPort()))
			}
		}
		return nil
	}))

	ServerSessionMgrSington.AddListener(ServiceMgr)
}
