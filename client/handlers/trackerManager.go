package handlers

import (
	"PFS/common"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"
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
	con.Close()
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
	con.Close()
	if rep.Success {
		return rep.List
	}
	return nil
}

//FileRegister registers a given file name to server
func FileRegister(file string, size string) {
	myPort := os.Args[1]
	con, err := net.Dial("tcp", Server)
	var li []common.FileShare
	li = append(li, common.FileShare{Name: file, Size: size})
	if err != nil {
		fmt.Println("Connection to tracker failed, exiting program: " + err.Error())
		return
	}
	msg := &common.MsgReq{}
	msg.MessageType = "Register"
	msg.Reg.List = li
	msg.Reg.IPAddr = strings.Split(con.LocalAddr().String(), ":")[0] + ":" + myPort
	enc := gob.NewEncoder(con)
	enc.Encode(msg)
	rep := &common.MsgRep{}
	dec := gob.NewDecoder(con)
	dec.Decode(rep)
	con.Close()
	if rep.Success {
		fmt.Println("File Registered Sucessfully: ")
		return
	}
	fmt.Println("File Registration failed ...")
	return
}

//GetLocation fetches the list locations of file from server
func GetLocation(file string) ([]common.FileLocation, error) {
	con, err := net.Dial("tcp", Server)
	if err != nil {
		fmt.Println("Connection to tracker failed " + err.Error())
		return nil, err
	}
	msg := &common.MsgReq{}
	msg.MessageType = "FileLocation"
	msg.Loc.Name = file
	enc := gob.NewEncoder(con)
	enc.Encode(msg)
	rep := &common.MsgRep{}
	dec := gob.NewDecoder(con)
	dec.Decode(rep)
	con.Close()
	var list []common.FileLocation
	list = rep.Loc
	return list, nil
}
