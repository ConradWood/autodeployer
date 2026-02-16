package main

import (
	"context"
	"fmt"
	"path/filepath"

	common "golang.conradwood.net/apis/common"
	"golang.conradwood.net/apis/deploymonkey"
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
	return &common.Void{}, nil
}
