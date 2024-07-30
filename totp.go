package main

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	key_encoded := scanner.Text()

	key, err := base32.StdEncoding.DecodeString(key_encoded)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v", err)
		return
	}
	//fmt.Printf("key = %v\n", key)

	T0 := 0
	X := 30
	ct := (uint64(time.Now().Unix()) - uint64(T0)) / uint64(X)
	//fmt.Printf("ct = %v\n", ct)

	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, uint64(ct))
	//fmt.Printf("ct bytes: %v\n", bs)

	mac := hmac.New(sha1.New, key)
	mac.Write(bs)
	val := mac.Sum(nil)
	//fmt.Println(val)

	i := val[len(val)-1] & 0x0f
	//fmt.Printf("i = %v\n", i)

	extract31 := (uint32(val[i]&0x7f) << 24) |
		(uint32(val[i+1]) << 16) |
		(uint32(val[i+2]) << 8) |
		(uint32(val[i+3]))
	res := extract31 % 1e6
	fmt.Printf("%06d", res)
}
