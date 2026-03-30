package main

import (
	"flag"
	"fmt"
	"sync"
)

var (
	lock_autodeployers = flag.Bool("lock_autodeployers", true, "if false do not lock individual autodeployers")
	ad_glob_lock       sync.Mutex
	ad_host_lock       = make(map[string]*adlock)
)

type adlock struct {
	host string
	lock sync.Mutex
}

// wait for a lock on an autodeployer host
func lockAutodeployerHost(host string) *adlock {
	if !*lock_autodeployers {
		return &adlock{}
	}
	ad_glob_lock.Lock()
	ad := ad_host_lock[host]
	if ad == nil {
		ad = &adlock{host: host}
		ad_host_lock[host] = ad
	}
	ad_glob_lock.Unlock()
	fmt.Printf("locking autodeployer host \"%s\"\n", host)
	ad.lock.Lock()
	fmt.Printf("autodeployer host \"%s\" locked\n", host)
	return ad
}

func (a *adlock) Unlock() {
	if !*lock_autodeployers {
		return
	}
	a.lock.Unlock()
	fmt.Printf("autodeployer host \"%s\" unlocked\n", a.host)
}
