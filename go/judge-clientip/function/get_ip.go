package function

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/bookun/sandbox/go/judge-clientip/driver/cache"
	"github.com/bookun/sandbox/go/judge-clientip/entity"
)

var (
	regMap = map[string]*regexp.Regexp{
		"Flet's hikari":  regexp.MustCompile(os.Getenv("fletsRe")),
		"ipoe":           regexp.MustCompile(os.Getenv("ipoeRe")),
		"B flet's":       regexp.MustCompile(os.Getenv("bFletsRe")),
		"Flet's Premium": regexp.MustCompile(os.Getenv("fletsPremiumRe")),
		"Mobile ONE":     regexp.MustCompile(os.Getenv("mobileOne")),
	}
)

type cacher interface {
	Get(context.Context, string) (entity.Result, error)
	Add(context.Context, string, entity.Result) error
}

type hostResearch struct {
	cache cacher
}

func GetIP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resultBuf bytes.Buffer

	clientIP := r.Header.Get("X-Forwarded-For")
	if clientIP == "" {
		clientIP = "222.145.13.136"
	}
	clientIP = strings.Split(clientIP, ",")[0]

	ctx := context.Background()

	cache, err := cache.NewFirebase(ctx, os.Getenv("projectID"))
	if err != nil {
		log.Println(err.Error())
		return
	}
	hostResearch := &hostResearch{cache: cache}
	cacheResult, err := hostResearch.cache.Get(ctx, clientIP)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if !cacheResult.IsEmpty() {
		if err := json.NewEncoder(&resultBuf).Encode(cacheResult); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, resultBuf.String())
		return
	}

	myDial := func(ctx context.Context, network, address string) (net.Conn, error) {
		d := net.Dialer{}
		return d.DialContext(ctx, network, address)
	}
	resolver := net.Resolver{PreferGo: true, StrictErrors: true, Dial: myDial}
	hosts, err := resolver.LookupAddr(ctx, clientIP)
	if err != nil {
		result := entity.Result{IP: clientIP, Host: "-", Kind: "no set", Error: err.Error()}
		if err := json.NewEncoder(&resultBuf).Encode(result); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, resultBuf.String())
		return
	}

	host := strings.Join(hosts, ",")
	kind := getKind(host)

	result := entity.Result{IP: clientIP, Host: host, Kind: kind, Error: ""}
	result.From = "dns"
	if err := json.NewEncoder(&resultBuf).Encode(result); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, resultBuf.String())
	if err := hostResearch.cache.Add(ctx, clientIP, result); err != nil {
		log.Println(err.Error())
	}
}

func getKind(host string) string {
	for provider, hostPattern := range regMap {
		println(provider)
		if hostPattern.MatchString(host) {
			return provider
		}
	}
	return "other"
}
