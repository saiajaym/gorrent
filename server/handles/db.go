package handles

import (
	"PFS/common"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/boltdb/bolt"
)

var (
	fileList = []byte("fileList")
	fileLoc  = []byte("fielLocation")
)

//DBOpen opens and return db handle
func DBOpen(file string) (*bolt.DB, error) {
	db, err := bolt.Open(file, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Println("db Open failed " + err.Error())
	}
	db.Update(func(tx *bolt.Tx) error {

		_, err := tx.CreateBucketIfNotExists(fileList)
		if err != nil {
			fmt.Println("DBOpen unable to creare filebucket" + err.Error())
			return err
		}

		_, err = tx.CreateBucketIfNotExists(fileLoc)
		if err != nil {
			fmt.Println("DBOpen unable to creare filebucket" + err.Error())
			return err
		}
		return err
	})
	err = nil
	return db, err
}

//AddList Adds list of files to tracker
func (cli *Client) AddList(list []common.FileShare) error {
	//fmt.Println("At addList:" + )
	err := cli.Db.Update(func(tx *bolt.Tx) error {
		err := error(nil)
		bucket := tx.Bucket(fileList)

		for _, l := range list {
			fmt.Println("Adding to bd:.." + l.Name)
			err = bucket.Put([]byte(l.Name), []byte(l.Size))
		}
		return err
	})

	if err != nil {
		fmt.Println("DBUpdate of file list failed" + err.Error())
	}
	return err
}

//AddIPAdd Adds list of files and corresponding IPaddr to db
func (cli *Client) AddIPAdd(list []common.FileShare, ipaddr string) error {
	err := cli.Db.Update(func(tx *bolt.Tx) error {
		err := error(nil)
		bucket := tx.Bucket(fileLoc)
		for _, l := range list {
			ex := bucket.Get([]byte(l.Name))
			if ex == nil {
				err = bucket.Put([]byte(l.Name), []byte(ipaddr))
			} else {
				err = bucket.Put([]byte(l.Name), []byte(string(ex)+"#"+ipaddr))
			}
		}

		return err
	})
	if err != nil {
		fmt.Println("DBUpdate of file loc failed" + err.Error())
	}

	return err
}

//GetList pulls file list from DB and returns back lisr
func (cli *Client) GetList() ([]common.FileShare, error) {
	var list = []common.FileShare{}
	err := cli.Db.View(func(tx *bolt.Tx) error {

		bucket := tx.Bucket(fileList)

		c := bucket.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			list = append(list, common.FileShare{Name: string(k), Size: string(v)})
		}
		if len(list) == 0 {
			return errors.New("No files found")
		}
		return nil

	})
	if err != nil {
		fmt.Println("GetList failed " + err.Error())
	}
	return list, err
}

//FileLoc returns IP address of files
//returns nil if file not found
func (cli *Client) FileLoc(file common.FileShare) ([]common.FileShare, error) {
	var list = []common.FileShare{}
	err := cli.Db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(fileLoc)
		li := bucket.Get([]byte(file.Name))
		if li == nil {
			return errors.New("File Not Present")
		}
		sli := strings.Split(string(li), "#")
		for _, filName := range sli {
			list = append(list, common.FileShare{Name: filName, Size: ""})

		}
		return nil
	})

	return list, err
}
