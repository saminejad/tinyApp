package main

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	_, _, s := time.Now().Clock() //get time second for seed

	hash := md5.New() // create new hash func

	rand.Seed(int64(s)) // add second for seed

	b := make([]byte, 5) // create new binnary array

	binary.LittleEndian.PutUint32(b, uint32(rand.Int())) // convert int to byte array and use rand for input
	// liitle_endian => left to right high value to low value
	hash.Write(b)
	pass := hash.Sum(nil)

	fmt.Println(hex.EncodeToString(pass)[0:10]) // convert array of byte to string
}
