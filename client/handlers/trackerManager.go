package handlers

import (
	"PFS/common"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

//Server holds server IP and Port
var Server string

//CheckTracker checks initial tracker status ...if fails, exits client
func CheckTracker() bool {
	con, err := net.Dial("tcp", Server)
	if err != nil {
		fmt.Println("Connection to tracker failed, exiting program: " + err.Error())
		os.Exit(1)
	}
	fmt.Println("Connection to tracker Successful: " + con.RemoteAddr().String())
	return true
}

//GetList fetches file list from download
func GetList() []common.FileShare {
	con, err := net.Dial("tcp", Server)
	if err != nil {
		fmt.Println("Connection to tracker failed, exiting program: " + err.Error())

	}
	msg := &common.MsgReq{}
	msg.MessageType = "FileList"
	enc := gob.NewEncoder(con)
	enc.Encode(msg)
	rep := &common.MsgRep{}
	dec := gob.NewDecoder(con)
	dec.Decode(rep)
	if rep.Success {
		return rep.List
	}
	return nil
}

//FileRegister registers a given file name to server
func FileRegister(file string, size string) {
	con, err := net.Dial("tcp", Server)
	if err != nil {
		fmt.Println("Connection to tracker failed, exiting program: " + err.Error())

	}
	msg := &common.MsgReq{}
	msg.MessageType = "Register"
	msg.Share.Name = file
	msg.Share.Size = size
	enc := gob.NewEncoder(con)
	enc.Encode(msg)
	rep := &common.MsgRep{}
	dec := gob.NewDecoder(con)
	dec.Decode(rep)
	if rep.Success {
		fmt.Println("File Registered Sucessfully: ")
		return
	}
	fmt.Println("File Registration failed ...")
	return
}
