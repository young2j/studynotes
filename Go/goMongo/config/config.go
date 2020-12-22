package config

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	credentials = options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		AuthSource:    "anquan",
		Username:      "ysj",
		Password:      "123456",
	}
	// direct                = true
	connectTimeout        = 10 * time.Second
	hosts                 = []string{"localhost:27017"}
	maxPoolSize    uint64 = 20
	minPoolSize    uint64 = 5
	readPreference        = readpref.Primary()
	replicaSet            = "replicaSetDb"

	// ClientOpts mongoClient 连接客户端参数
	clientOpts = &options.ClientOptions{
		Auth:           &credentials,
		ConnectTimeout: &connectTimeout,
		// Direct:         &direct,
		Hosts:          hosts,
		MaxPoolSize:    &maxPoolSize,
		MinPoolSize:    &minPoolSize,
		ReadPreference: readPreference,
		ReplicaSet:     &replicaSet,
	}
)

// ClientOpts mongoClient settings.
var ClientOpts = options.Client().
	// SetAuth(options.Credential{
	// 	AuthMechanism: "SCRAM-SHA-1",
	// 	AuthSource:    "anquan",
	// 	Username:      "ysj",
	// 	Password:      "123456",
	// }).
	SetConnectTimeout(10 * time.Second).
	SetHosts([]string{"localhost:27017"}).
	SetMaxPoolSize(20).
	SetMinPoolSize(5).
	// SetReplicaSet("replicaSetDb").
	SetReadPreference(readpref.Primary())
