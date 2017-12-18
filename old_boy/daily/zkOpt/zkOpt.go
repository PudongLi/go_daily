package main

import (
	zk "old_boy/day_01/daily/lib/go-zookeeper-master/zk"
	"fmt"
	"time"
)

var (
	ZKList = []string{"10.255.224.96:2181", "10.255.224.97:2181", "10.255.224.98:2181"}

)

func main(){

}
/*
	get a connection
	parameters:
		zkList: zookeepr node list
	return:
		zk.Conn
 */
func getConnect(zkList []string) (conn *zk.Conn){
	conn, _, err := zk.Connect(zkList, 10*time.Second)
	if err != nil{
		fmt.Println(err)
	}
	return
}

func createNode(node string) bool{


}

func deleteNode(node string) bool{

}

func deleteAll(node string) bool{

}



