package main

import (
	"context"
	"fmt"
	"path/filepath"

	common "golang.conradwood.net/apis/common"
	"golang.conradwood.net/apis/deploymonkey"
	"golang.conradwood.net/deploymonkey/appstats"
	"golang.yacloud.eu/apis/autodeployercommon"
)

func (e *DeployMonkey) AppChanged(ctx context.Context, req *deploymonkey.AppChangeRequest) (*common.Void, error) {
	appdef := req.AppRef.AppDef
	ipc_s := "ipc not supported"
	if req.IPCSupported {
		ipc_s = "IPCName=" + req.IPCName
	}
	bn := filepath.Base(appdef.Binary)
	prefix := fmt.Sprintf("[appchange binary=%s deplid=%s appdef-deplid=%s] ", bn, req.DeploymentID, appdef.DeploymentID)
	fmt.Printf(prefix+"Autodeployer reports status has changed to %v after %d seconds runtime\n", req.Status, req.RuntimeSeconds)
	fmt.Printf(prefix+"   appdef-id: %d, ipc=%s, buildid=%d, health=%v\n", appdef.ID, ipc_s, appdef.BuildID, req.Health)
	fmt.Printf(prefix+"   undeployrequest received=%v\n", req.UndeployRequestReceived)
	err := appstats.ChangeAppStats(ctx, appdef.ID, func(as *deploymonkey.AppStats) error {
		if req.IPCSupported {
			as.IPCTrackStatus = deploymonkey.IPCTrackStatus_Supported
		}
		if as.IPCTrackStatus != deploymonkey.IPCTrackStatus_Supported {
			if req.Status == autodeployercommon.DeploymentStatus_TERMINATED {
				if req.IPCSupported {
					as.IPCTrackStatus = deploymonkey.IPCTrackStatus_Supported
				} else {
					as.IPCTrackStatus = deploymonkey.IPCTrackStatus_NotSupported
				}
			}
		}
		if req.Status == autodeployercommon.DeploymentStatus_TERMINATED {
			if !req.UndeployRequestReceived {
				as.UnexpectedRestarts = as.UnexpectedRestarts + 1
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &common.Void{}, nil
}
