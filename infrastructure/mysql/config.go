package mysql

import (
	"fmt"
	"time"

	"github.com/woshahua/corona_server/environment"
	"google.golang.org/appengine"
)

type config struct {
	user                    string
	password                string
	protocol                string
	host                    string
	dbName                  string
	maxIdle                 int
	maxOpen                 int
	connMaxLifetime         time.Duration
	connRetryMaxWaitingTime time.Duration
}

func makeConfig() config {
	c := config{
		user:                    environment.GetSharedEnvironments().MysqlUser,
		password:                environment.GetSharedEnvironments().MysqlPassword,
		protocol:                environment.GetSharedEnvironments().MysqlProtocol,
		host:                    environment.GetSharedEnvironments().MysqlHost,
		dbName:                  environment.GetSharedEnvironments().MysqlDbName,
		maxIdle:                 10,
		maxOpen:                 100,
		connMaxLifetime:         time.Second * 5,
		connRetryMaxWaitingTime: time.Second * 5,
	}
	return c
}

func (c config) build() string {
	if appengine.IsAppEngine() {
		// LINK: https://cloud.google.com/appengine/docs/standard/go112/using-cloud-sql?hl=ja#running_the_sample_code
		format := "%s:%s@unix(/cloudsql/%s)/%s"
		return fmt.Sprintf(
			format,
			c.user,
			c.password,
			c.host,
			c.dbName,
		) + "?parseTime=true&loc=Asia%2FTokyo&charset=utf8mb4"
	}

	format := "%s:%s@%s(%s)/%s"
	return fmt.Sprintf(
		format,
		c.user,
		c.password,
		c.protocol,
		c.host,
		c.dbName,
	) + "?parseTime=true&loc=Asia%2FTokyo&charset=utf8mb4"
}