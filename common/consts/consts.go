package consts

import (
	"fmt"
)

const (
	IdServiceAddress  = "service.id.address"
	IdServiceStartIdx = "service.id.startidx"
	IdServiceStep     = "service.id.step"

	MongoDbName   = "mongodb.dbname"
	MongoDbAddr   = "mongodb.addr"
	MongoDbUser   = "mongodb.user"
	MongoDbPsw    = "mongodb.psw"
	RedisAddress  = "redis.address"
	RedisPassword = "redis.password"
	RedisDb       = "redis.db"
	Neo4JDbUrl    = "neo4j.url"

	ApiServiceAddress    = "service.api.address"
	ApiServiceStaticPath = "service.api.staticpath"

	LeafMsgServiceAddress = "service.message.leaf.addrs"
	MsgServiceAddress     = "service.message.address"
	ParentServiceAddress  = "service.parent.address"
	IsLeafServer          = "service.message.isleaf"
	ErrorLogPath          = "log.error.path"
	AccessLogPath         = "log.access.path"

	TlsEnable = "security.tls.enable"

	SessionCookieName = "session.cookie"
	SessionPrefix     = "session.prefix"

	PrivateToken = "security.token"
)

func ErrGetConfigFailed(configName string, err error) string {
	return fmt.Sprintf("get config(%s) failed!err:=%s\n", configName, err.Error())
}

func ErrDialRemoteServiceFailed(addr string, err error) string {
	return fmt.Sprintf("dial to server(%s) failed!err:=%s\n", addr, err.Error())
}

func ErrRPCCallFailed(service string, function string, err error) string {
	return fmt.Sprintf("rpc call(%s,%s) failed!err:=%s\n", service, function, err.Error())
}
