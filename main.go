package main

import (
	"log"
	// _ "net/http/pprof"
)

var conf *Config

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf = initConfig()

	startMonitor()

	// leaseCancel("8SfpDMtVcppwwWFNmhMQ2rcwg9Kq6LTFmdJH1pThiDHP")
	// leaseCancel("Ao62m9YcpT2bJJGXc5EXsbxNdLFAFN6HzUN3Gj8izPx3")
	// leaseCancel("2NTHH6ZXe5w9S5A7SwpPGE9jyNs16zTjBZ1euov9z1er")
	// leaseCancel("9qHU5EUzCZkbKte9nNRkU1FS79CMX2rtR5mys4M44ikA")
	// leaseCancel("DcwqqUBHRoAkLB4WbBwWoSjF3DxJi1fcGGDRP7a9QPxL")

	// val := "%s%d%s%s%s%s%s%s__3AM2KY5qur6ms2yP2KeUwFgoKBApnwmGaEn__1371942__3AShXVgRcRis82CwD7o9pz1Ac9vmRYMqELT__A5dxBC57ftdJaBike6iPDnemdShrMA35boUxKikoopqt__5v3NcJroq5JXKqQ9FU5Gvy3Nh1A7PtEwgeqD3eLXQKJu__BJNJVWDP5BmzeiHwGLVNuC3pjnBkczHjoLu5NStU546q__6AgeQHetZSXvvYt8dJ1X3Xas2jeLcAnrjxPM2NfJJU68__4BLqeZg2WVAnCLpVEZAhcbeXEZ32bXjghhs9nYzg343h"
	// dataTransaction("3AGDV6tAeXsuu833PsFshUvmNkviqY2nku1", nil, nil, nil)
}
