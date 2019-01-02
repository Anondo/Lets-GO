package main

import (
	"fmt"
	"log"

	"github.com/mediocregopher/radix.v2/redis"
)

type Album struct {
	Title  string
	Artist string
}

func (a *Album) populate(amap map[string]string) {
	a.Title = amap["title"]
	a.Artist = amap["artist"]
}
func (a *Album) showInfo() {
	fmt.Println("Title : ", a.Title)
	fmt.Println("Artist : ", a.Artist)
}

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	//resp , err := conn.Cmd("HMSET", "album:1", "title", "Those whom the God detests", "artist", "Nile")
	//if resp.Err != nil {
	//	log.Fatal(resp.Err.Error())
	//}
	//fmt.Println("New Album Added")
	reply, err := conn.Cmd("HGETALL", "album:1").Map()
	if err != nil {
		log.Fatal(err)
	}
	album := Album{}
	album.populate(reply)

	album.showInfo()

}
