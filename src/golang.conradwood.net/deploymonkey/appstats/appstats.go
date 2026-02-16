package appstats

import (
	"context"
	"sync"

	"golang.conradwood.net/apis/deploymonkey"
	"golang.conradwood.net/deploymonkey/db"
)

var (
	appstats_lock sync.Mutex
)

func ChangeAppStats(ctx context.Context, appdefid uint64, changer func(*deploymonkey.AppStats) error) error {
	appstats_lock.Lock()
	defer appstats_lock.Unlock()
	appdef := &deploymonkey.ApplicationDefinition{ID: appdefid}
	appstats, err := db.DefaultDBAppStats().ByAppDef(ctx, appdef.ID)
	if err != nil {
		return err
	}
	var appstat *deploymonkey.AppStats
	if len(appstats) != 0 {
		appstat = appstats[0]
	} else {
		appstat = &deploymonkey.AppStats{AppDef: appdef}
		_, err = db.DefaultDBAppStats().Save(ctx, appstat)
		if err != nil {
			return err
		}
	}
	err = changer(appstat)
	if err != nil {
		return err
	}
	err = db.DefaultDBAppStats().Update(ctx, appstat)
	if err != nil {
		return err
	}
	return nil
}
