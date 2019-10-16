package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/wiloon/pingd-log/logconfig"
	"go.etcd.io/bbolt"
	"net/http"
	"time"
)

const bucketTimer = "timer"

func main() {
	logconfig.Init()
	// get system time
	systemTime := time.Now()
	systemTimeStr := systemTime.Format("2006-01-02T15:04:05Z07:00")
	systemTimeStrDay := systemTime.Format("2006-01-02")
	// get net time via http request
	resp, _ := http.Get("http://www.ntsc.ac.cn")
	netTime := resp.Header.Get("Date")

	log.Infof("system time: %v", "net time: %v", systemTimeStr, netTime)

	//save to db //基于文件的数据库

	db, err := bbolt.Open("/tmp/foo.db", 0666, nil)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	_ = db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucketTimer))
		if err != nil {
			fmt.Println(err)
		}
		err = b.Put([]byte(systemTimeStrDay), []byte(systemTimeStr))
		return err
	})

	db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketTimer))
		v := b.Get([]byte(systemTimeStrDay))
		fmt.Printf("value: %s\n", v)
		return nil
	})

	//search
	//查询某一天最早的时间

	// 历史数据整理
	// 从实时表统计  ，结果 插入历史表
}
