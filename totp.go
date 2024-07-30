// Simple TOTP implementation, based on reference Java code from
// https://datatracker.ietf.org/doc/html/rfc6238

package main

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var numDigits int
	var T, T0, X int64

	tStr := flag.String("t", "", "timestamp (T) as Unix epoch (default current-timestmp)")
	flag.Int64Var(&T0, "t0", 0, "start timestamp (T0) as Unix epoch (default 0)")
	flag.Int64Var(&X, "x", 30, "time step in seconds (X)")
	flag.IntVar(&numDigits, "digits", 6, "number of digits")
	flag.Parse()

	T = time.Now().Unix()
	if *tStr != "" {
		var err error
		T, err = strconv.ParseInt(*tStr, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: timestamp needs to be an integer\n")
			os.Exit(1)
		}
	}
	if numDigits < 1 || numDigits > 8 {
		fmt.Fprintf(os.Stderr, "ERROR: number of digits must be in [1,8] range\n")
		os.Exit(1)
	}
	if X <= 0 {
		fmt.Fprintf(os.Stderr, "ERROR: time step must be a positive integer\n")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	key_encoded := scanner.Text()

	key, err := base32.StdEncoding.DecodeString(key_encoded)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
	//fmt.Printf("key = %v\n", key)

	ct := (T - T0) / X
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

	extract31 := uint32(val[i]&0x7f)<<24 +
		uint32(val[i+1])<<16 +
		uint32(val[i+2])<<8 +
		uint32(val[i+3])

	powers := []uint32{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000}
	res := extract31 % powers[numDigits]

	fmt.Printf("%0"+strconv.Itoa(numDigits)+"d", res)
}
