package database

import (
	"fmt"
	"go.etcd.io/bbolt"
	"log"
)

var db *bbolt.DB

const bucketName = "timer"

func init() {
	var err error
	db, err = bbolt.Open("/tmp/foo.db", 0666, nil)
	if err != nil {
		log.Println(err)
	}
}

func Get(key string) string {
	var value string
	_ = db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		value = string(b.Get([]byte(key)))
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
