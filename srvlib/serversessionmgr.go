package srvlib

import (
	"strings"

	"github.com/acoderup/goserver/core/logger"
	"github.com/acoderup/goserver/core/netlib"
	"github.com/acoderup/goserver/srvlib/protocol"
)

/*
 服务器信息注册，单个服务器可能包含多个子服务端口
*/

var (
	SessionAttributeServerInfo = &ServerSessionMgr{}
	ServerSessionMgrSington    = &ServerSessionMgr{sessions: make(map[int]map[int]map[int]*netlib.Session)}
)

type ServerSessionRegisteListener interface {
	OnRegiste(*netlib.Session)
	OnUnregiste(*netlib.Session)
}

type ServerSessionMgr struct {
	sessions  map[int]map[int]map[int]*netlib.Session //keys=>areaid:type:id
	listeners []ServerSessionRegisteListener
}

// AddListener 添加一个服务器会话注册监听器
func (ssm *ServerSessionMgr) AddListener(l ServerSessionRegisteListener) ServerSessionRegisteListener {
	ssm.listeners = append(ssm.listeners, l)
	return l
}

// RegisteSession 注册一个新的服务器会话
func (ssm *ServerSessionMgr) RegisteSession(s *netlib.Session) bool {
	attr := s.GetAttribute(SessionAttributeServerInfo)
	if attr == nil {
		logger.Logger.Warnf("服务器注册信息为空")
		return false
	}

	srvInfo, ok := attr.(*protocol.SSSrvRegiste)
	if !ok || srvInfo == nil {
		logger.Logger.Warnf("服务器注册信息错误")
		return false
	}

	areaId := int(srvInfo.GetAreaId())
	srvType := int(srvInfo.GetType())
	srvId := int(srvInfo.GetId())

	if _, ok := ssm.sessions[areaId]; !ok {
		ssm.sessions[areaId] = make(map[int]map[int]*netlib.Session)
	}
	if _, ok := ssm.sessions[areaId][srvType]; !ok {
		ssm.sessions[areaId][srvType] = make(map[int]*netlib.Session)
	}

	session, has := ssm.sessions[areaId][srvType][srvId]
	if has && session != nil && session != s {
		logger.Logger.Infof("删除旧服务器注册: %v", srvInfo)
		ssm.UnregisteSession(session)
	}

	logger.Logger.Infof("服务器注册成功：%v", srvInfo)
	ssm.sessions[areaId][srvType][srvId] = s
	if len(ssm.listeners) != 0 {
		for _, l := range ssm.listeners {
			l.OnRegiste(s)
		}
	}

	return true
}

// UnregisteSession 注销一个服务器会话
func (ssm *ServerSessionMgr) UnregisteSession(s *netlib.Session) bool {
	attr := s.GetAttribute(SessionAttributeServerInfo)
	if attr == nil {
		return false
	}

	srvInfo, ok := attr.(*protocol.SSSrvRegiste)
	if !ok || srvInfo == nil {
		return false
	}

	logger.Logger.Infof("尝试删除服务器注册：%v", srvInfo)
	areaId := int(srvInfo.GetAreaId())
	srvType := int(srvInfo.GetType())
	srvId := int(srvInfo.GetId())

	if a, exist := ssm.sessions[areaId]; exist {
		if b, exist := a[srvType]; exist {
			if conn, exist := b[srvId]; exist && s == conn {
				logger.Logger.Infof("删除服务器注册成功 %v", srvInfo)
				delete(b, srvId)
				if len(ssm.listeners) != 0 {
					for _, l := range ssm.listeners {
						l.OnUnregiste(s)
					}
				}
				return true
			} else {
				logger.Logger.Infof("服务器注册信息已经删除")
				return false
			}
		}
	}

	logger.Logger.Infof("服务器注册信息没找到：%v", srvInfo)

	return false
}

// GetSession 根据区域ID、服务器类型和服务器ID获取会话
func (ssm *ServerSessionMgr) GetSession(areaId, srvType, srvId int) *netlib.Session {
	if a, exist := ssm.sessions[areaId]; exist {
		if b, exist := a[srvType]; exist {
			if c, exist := b[srvId]; exist {
				return c
			}
		}
	}
	return nil
}

// GetSessions 获取指定区域和服务器类型的所有会话
func (ssm *ServerSessionMgr) GetSessions(areaId, srvType int) (sessions []*netlib.Session) {
	if a, exist := ssm.sessions[areaId]; exist {
		if b, exist := a[srvType]; exist {
			for _, s := range b {
				sessions = append(sessions, s)
			}
		}
	}
	return
}

// GetServerId 获取指定区域和服务器类型的第一个服务器ID
func (ssm *ServerSessionMgr) GetServerId(areaId, srvType int) int {
	if a, exist := ssm.sessions[areaId]; exist {
		if b, exist := a[srvType]; exist {
			for sid := range b {
				return sid
			}
		}
	}
	return -1
}

// GetServerIdByMaxData 根据最大数据获取指定区域和服务器类型的服务器ID
func (ssm *ServerSessionMgr) GetServerIdByMaxData(areaId, srvType int) int {
	var bestSid int = -1
	var data string
	if a, exist := ssm.sessions[areaId]; exist {
		if b, exist := a[srvType]; exist {
			for sid, s := range b {
				if srvInfo, ok := s.GetAttribute(SessionAttributeServerInfo).(*protocol.SSSrvRegiste); ok && srvInfo != nil {
					if strings.Compare(data, srvInfo.GetData()) <= 0 {
						data = srvInfo.GetData()
						bestSid = sid
					}
				}
			}
		}
	}
	return bestSid
}

// GetServerIds 获取指定区域和服务器类型的所有服务器ID
// 参数:
//   - areaId: 区域ID
//   - srvType: 服务器类型
//
// 返回:
//   - ids: 包含所有匹配服务器ID的切片
func (ssm *ServerSessionMgr) GetServerIds(areaId, srvType int) (ids []int) {
	// 检查指定区域是否存在
	if a, exist := ssm.sessions[areaId]; exist {
		// 检查指定服务器类型是否存在
		if b, exist := a[srvType]; exist {
			// 遍历所有匹配的服务器，收集它们的ID
			for sid := range b {
				ids = append(ids, sid)
			}
		}
	}
	return
}

// Broadcast 向指定区域和服务器类型广播消息
func (ssm *ServerSessionMgr) Broadcast(packetid int, pack interface{}, areaId, srvType int) {
	if areaId >= 0 {
		if srvType >= 0 {
			// 向特定区域和服务器类型广播
			if a, exist := ssm.sessions[areaId]; exist {
				if b, exist := a[srvType]; exist {
					for _, s := range b {
						s.Send(packetid, pack)
					}
				}
			}
		} else {
			// 向特定区域的所有服务器类型广播
			if a, exist := ssm.sessions[areaId]; exist {
				for _, b := range a {
					for _, s := range b {
						s.Send(packetid, pack)
					}
				}
			}
		}
	} else {
		if srvType >= 0 {
			// 向所有区域的特定服务器类型广播
			for _, a := range ssm.sessions {
				if b, exist := a[srvType]; exist {
					for _, s := range b {
						s.Send(packetid, pack)
					}
				}
			}
		} else {
			// 向所有区域的所有服务器类型广播
			for _, a := range ssm.sessions {
				for _, b := range a {
					for _, s := range b {
						s.Send(packetid, pack)
					}
				}
			}
		}
	}
}
