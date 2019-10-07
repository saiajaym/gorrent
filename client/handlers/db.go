package handlers

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

const fileDB = "fileDB"

//DBOpen Opens DB
func DBOpen(file string) (*bolt.DB, error) {
	db, err := bolt.Open(file, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Println("db Open failed " + err.Error())
	}
	db.Update(func(tx *bolt.Tx) error {

		_, err := tx.CreateBucketIfNotExists([]byte(fileDB))
		if err != nil {
			fmt.Println("DBOpen unable to creare filebucket" + err.Error())
			return err
		}
		return err
	})
	err = nil
	return db, err
}

//FileMap maps local file location to file name
func FileMap(name string, path string, db *bolt.DB) {
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(fileDB))
		err := bucket.Put([]byte(name), []byte(path))
		return err
	})
}

//GetPath returns absolute path of a file
func (leech *Leech) GetPath(name string) (string, error) {
	var path string
	err := leech.Db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(fileDB))
		path = string(bucket.Get([]byte(name)))
		return nil
	})
	return path, err
}
