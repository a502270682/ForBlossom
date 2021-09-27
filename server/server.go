package server

import (
	"context"
	wechat_account "forBlossem/adapter/account"
	"forBlossem/adapter/log"
	"forBlossem/adapter/mysql"
	"forBlossem/cache"
	"forBlossem/config"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/silenceper/wechat/v2/officialaccount"
	config2 "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/urfave/cli"
)

type Server struct {
	*cli.App
	config      *config.Config
	//Wechat      *officialaccount.OfficialAccount
	//RedisClient *cache.Redis
}

func initWithConfig(ctx context.Context, filePath string) (*config.Config, error) {
	conf, err := config.Load(filePath)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func initMysql(conf *config.Config) error {
	err := mysql.InitializeMainDb(conf.Mysql.Master)
	if err != nil {
		return err
	}
	db := mysql.GetClient()
	mysql.InitEntityDao(db)
	return nil
}

func initRedis(conf *config.Config) (*cache.Redis, error) {
	client := cache.NewRedis(&cache.RedisOpts{
		Host:        conf.Redis.Host,
		Password:    conf.Redis.Password,
		Database:    conf.Redis.Database,
		MaxIdle:     conf.Redis.MaxIdle,
		MaxActive:   conf.Redis.MaxActive,
		IdleTimeout: conf.Redis.IdleTimeout,
	})
	client.SetConn(client.Conn)

	_, err := client.Conn.Get().Do("PING")
	if err != nil {
		return nil, errors.Wrap(err,"redis init failed")
	}
	return client, nil
}

func initWechat(conf *config.Config, cache cache.Cache) *officialaccount.OfficialAccount {
	return officialaccount.NewOfficialAccount(&config2.Config{
		AppID:     conf.WechatConf.AppID,
		AppSecret: conf.WechatConf.AppSecret,
		Cache: cache,
	})
}

func NewServer(ctx context.Context) *Server {
	s := &Server{
		App: cli.NewApp(),
	}
	s.Flags = []cli.Flag{cli.StringFlag{Name: "c", Usage: "Configuration file"}}
	s.Action = func(c *cli.Context) error {
		if c.GlobalString("c") == "" {
			return errors.New("usage: my_go -c configfilepath")
		}

		log.Info(ctx, "start read config: ", c.GlobalString("c"))
		conf, err := initWithConfig(ctx, c.GlobalString("c"))
		if err != nil {
			return errors.Wrap(err, "fail to init conf")
		}

		log.Infof(ctx, "init config success. conf:%+v", conf)
		s.config = conf

		err = initMysql(conf)
		if err != nil {
			return errors.Wrap(err, "fail to init mysql")
		}

		// redis init
		redisClient, err := initRedis(conf)
		if err != nil {
			return errors.Wrap(err, "fail to init redis")
		}
		cache.SetRedisClient(redisClient)

		// wechat init
		account := initWechat(conf, redisClient)
		wechat_account.SetWechatAccount(account)

		r := gin.Default()
		routes(r)
		if err = r.Run(s.config.HTTPPort); err != nil {
			return errors.Wrap(err, "fail to run")
		}

		return nil
	}
	return s
}
