package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"economy_system/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo, NewCoinRepo, NewGiftRepo, NewKafkaProducer)

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: c.Database.Driver,
		DSN:        c.Database.Source,
	}), &gorm.Config{Logger: gormLogger.Default.LogMode(gormLogger.Info)})
	if err != nil {
		return nil, nil, err
	}
	d := &Data{
		db:  db,
		rdb: nil,
	}

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return d, cleanup, nil
}
