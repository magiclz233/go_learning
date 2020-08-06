package client

import (
	"crypto/sha256"
	"fmt"
	"go_learning/go_base/ch16_package/tempconv"
	"testing"
	"time"
)

func TestTempConvert(t *testing.T) {
	c := tempconv.Celsius(37)
	f := tempconv.Fahrenheit(101)

	f1 := tempconv.CToF(c)
	c1 := tempconv.FToC(f)
	fmt.Println(f1)
	fmt.Println(c1)

	// 大写开头的定义可以在包外被访问
	fmt.Println(tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.BoilingC)
}

func TestConst(t *testing.T) {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)

	fmt.Println(FlagUp, FlagBroadcast, FlagLoopback)
}

type Flags uint

const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // supports broadcast access capability
	FlagLoopback                       // is a loopback interface
	FlagPointToPoint                   // belongs to a point-to-point link
	FlagMulticast                      // supports multicast access capability
)

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
	EB
)

func TestBit(t *testing.T) {
	fmt.Println(B, KB, MB, GB, TB, PB, EB)
}

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func TestIotaExample(t *testing.T) {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB])
}

func TestSha256(t *testing.T) {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}
