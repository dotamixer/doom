package mongo

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func NewClient(opts *Options) (client *mongo.Client, err error) {
	ctx, _ := context.WithTimeout(context.TODO(), 3*time.Second)

	clientOptions := options.Client().
		SetHosts(opts.Hosts).
		SetReplicaSet(opts.ReplicaSet).
		SetMaxPoolSize(opts.MaxPoolSize).
		SetConnectTimeout(time.Second * 5).
		SetMaxConnIdleTime(time.Second)

	if len(opts.Username) == 0 {
		clientOptions.Auth = &options.Credential{
			Username:   opts.Username,
			Password:   opts.Password,
			AuthSource: opts.AuthSource,
		}
	}

	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		logrus.Fatalf("Failed to connect to mongo server. err:[%v]", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logrus.Fatalf("Failed to ping mongo server. err:[%v]", err)
	}
	return client, err
}

func TimeoutCtx() context.Context {
	ctx, _ := context.WithTimeout(context.TODO(), 3*time.Second)
	return ctx
}
