package main

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"time"
)

const (
	k  = "d516e7d9e1a3fc70cb300928603e3ab940270b03"
	pk = "1f4b8f992212e1e3fc00e6b7779e54ed"
)

func addQueryAuth(q url.Values) url.Values {
	// add auth parameters to query
	ts := time.Now().UnixNano()
	hashInput := fmt.Sprintf("%d%s%s", ts, k, pk)
	hash := fmt.Sprintf("%x", md5.Sum([]byte(hashInput)))
	q.Add("ts", fmt.Sprintf("%d", ts))
	q.Add("apikey", pk)
	q.Add("hash", hash)
	return q
}
