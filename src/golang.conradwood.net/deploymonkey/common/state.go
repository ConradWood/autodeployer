package common

import "sync"

var (
	auto_deploying      bool
	auto_deploying_lock sync.Mutex
)

func SetAutoDeploying(b bool) {
	auto_deploying_lock.Lock()
	defer auto_deploying_lock.Unlock()
	auto_deploying = b
}
func IsAutoDeployingCurrently() bool {
	auto_deploying_lock.Lock()
	defer auto_deploying_lock.Unlock()
	return auto_deploying
}
