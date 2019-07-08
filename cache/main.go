package main

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/powerman/rpc-codec/jsonrpc2"
)

var (
	inviterIDstr = "Roman"
	cacheMap     = map[string]int{
		"AAA@gmail.com": 33,
	}
	inviteCache     = cache.New(24*time.Hour, time.Hour)
	ErrInviteQuota  = jsonrpc2.Error{211, "DAILY_INVITE_QUOTA_EXCEEDED", nil}
	inviterMailsMap = make(map[string]int)
)

// func init() {
// 	cacheExpirationTime := time.Until(now.EndOfDay())
// 	inviteCache.Set(inviterIDstr, cacheMap, cacheExpirationTime)
// }

func main() {
	// cacheExpirationTime := time.Until(now.EndOfDay())
	// fmt.Println(cacheExpirationTime)

	mapa, found := inviteCache.Get(inviterIDstr)
	if found {
		inviterMailsMap = mapa.(map[string]int)
	}
	fmt.Println(inviterMailsMap)

	// ErrInviteQuota.Data = []string{"aaa", "bbb"}
	// fmt.Println(ErrInviteQuota)

}
