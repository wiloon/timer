package database

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.etcd.io/bbolt"
)

var db *bbolt.DB

const bucketName = "timer"

func init() {
	var err error
	db, err = bbolt.Open("foo.db", 0666, nil)
	if err != nil {
		log.Println(err)
	}

	_ = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			fmt.Println(err)
		}
		return err
	})
}

func Get(key string) string {
	var value string

	_ = db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))

		valueByte := b.Get([]byte(key))
		if valueByte == nil {
			value = ""
		} else {
			value = string(valueByte)
		}

		return nil
	})
	return value
}

func Set(key, value string) {
	_ = db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			fmt.Println(err)
		}
		err = bucket.Put([]byte(key), []byte(value))
		return err
	})
}
