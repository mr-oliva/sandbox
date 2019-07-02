package ip

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/bookun/sandbox/go/judge-clientip/driver/cache"
	"google.golang.org/appengine"
)

type cacher interface {
	Get(context.Context, string) (string, error)
	Add(context.Context, string, string) error
}

type hostResearch struct {
	cache cacher
}

type result struct {
	IP    string `json:"ip"`
	Host  string `json:"host"`
	Kind  string `json:"kind"`
	Error string `json:"error"`
}

func GetIP(w http.ResponseWriter, r *http.Request) {
	//{"ip":"222.145.13.136","host":"p7103136-ipngn32501marunouchi.tokyo.ocn.ne.jp.","kind":"hoge ocn","success":true}

	w.Header().Set("Content-Type", "application/json")

	var resultBuf bytes.Buffer

	clientIP := r.Header.Get("X-Forwarded-For")
	if clientIP == "" {
		clientIP = "222.145.13.136"
	}

	//ctx := context.Background()
	ctx := appengine.NewContext(r)

	cache := &cache.GoogleMemcache{}
	hostResearch := &hostResearch{cache: cache}
	cacheResult, err := hostResearch.cache.Get(ctx, clientIP)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if cacheResult != "" {
		fmt.Fprint(w, cacheResult)
		return
	}

	myDial := func(ctx context.Context, network, address string) (net.Conn, error) {
		d := net.Dialer{}
		return d.DialContext(ctx, network, address)
	}
	resolver := net.Resolver{PreferGo: true, StrictErrors: true, Dial: myDial}
	hosts, err := resolver.LookupAddr(ctx, clientIP)
	if err != nil {
		result := result{IP: clientIP, Host: "-", Kind: "no set", Error: err.Error()}
		if err := json.NewEncoder(&resultBuf).Encode(result); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, resultBuf.String())
		return
	}
	result := result{IP: clientIP, Host: strings.Join(hosts, ","), Kind: "other", Error: ""}
	if err := json.NewEncoder(&resultBuf).Encode(result); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, resultBuf.String())
	if err := hostResearch.cache.Add(ctx, clientIP, resultBuf.String()); err != nil {
		log.Println(err.Error())
	}
}
