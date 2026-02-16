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
	appname := fmt.Sprintf("[appchange %s appdef-id=%d Binary=%s buildid=%d, deplid=%s, artefactid=%d ipc=%s]", req.DeploymentID, appdef.ID, bn, appdef.BuildID, appdef.DeploymentID, appdef.ArtefactID, ipc_s)
	fmt.Printf("Autodeployer reports app %s has changed to %v (undeplyrequest=%v) after %d seconds runtime\n", appname, req.Status, req.UndeployRequestReceived, req.RuntimeSeconds)
	return &common.Void{}, nil
}
