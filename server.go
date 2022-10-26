package confs

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"reflect"
	"time"
)
func (c *Container)InitServe(a *reflect.Value) (interface{},error) {
	var b interface{}
	var err error
	switch a.Interface().(type) {
	case gorm.DB:
		b,err = c.InitMysql()
	case redis.Conn:
		b,err = c.InitRedis()
	}
	if err != nil {
		return b,errors.WithStack(err)
	}
	return b,nil
}

func (c *Container)InitMysql() (gorm.DB,error) {
	instanceCfg := c.Datas
	temp := "%s:%s@tcp(%s:%s)/%s?charset=%s"
	username := instanceCfg["username"].(string)
	host := instanceCfg["host"].(string)
	port := instanceCfg["port"].(string)
	password:=instanceCfg["password"].(string)
	database := instanceCfg["database"].(string)
	charset := instanceCfg["charset"].(string)

	dsn := fmt.Sprintf(temp, username, password, host, port, database, charset)
	db, err := gorm.Open("mysql", dsn)
	db.DB().SetMaxIdleConns(int(instanceCfg["max_idle_conn"].(float64)))
	db.DB().SetMaxIdleConns(int(instanceCfg["max_open_conn"].(float64)))
	db.DB().Ping()
	db.DB().SetConnMaxLifetime(10 * time.Minute)
	return *db,err
}

func (c *Container)InitRedis() (redis.Conn,error) {
	r, err := redis.Dial("tcp", ":6379")
	return r,err
}