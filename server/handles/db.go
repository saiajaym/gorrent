package handles

import (
	"PFS/common"
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/boltdb/bolt"
)

var (
	fileList      = []byte("fileList")
	fileLoc       = []byte("fielLocation")
	chunkLocation = []byte("chunkLocation")
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
			fmt.Println("Adding to db... " + l.Name)
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
	fmt.Println("At AddIP")
	for i := range list {
		list[i].Name = list[i].Name + "&" + ipaddr
	}
	err := cli.Db.Update(func(tx *bolt.Tx) error {
		err := error(nil)
		bucket := tx.Bucket(fileLoc)
		for _, l := range list {
			//fmt.Printf("cal chunks for %s, %s \n", l.Name, l.Size)
			tot, _ := strconv.Atoi(l.Size)
			//fmt.Printf("cal chunks for %s, %d \n", l.Name, tot)
			var chunks string
			for i := 1; tot >= 0; i++ {
				chunks = chunks + strconv.Itoa(i) + ","
				tot = tot - common.ChunkSize
			}
			fmt.Println("Adding chunks for " + l.Name + chunks)
			err = bucket.Put([]byte(l.Name), []byte(chunks))
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
func (cli *Client) FileLoc(file string) ([]common.FileLocation, error) {
	var list = []common.FileLocation{}
	err := cli.Db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(fileLoc).Cursor()
		prefix := []byte(file)

		for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
			ip := strings.Split(string(k), "&")
			list = append(list, common.FileLocation{IpAddr: ip[1], Chunks: string(v)})
		}

		return nil
	})

	return list, err
}
