package main

import (
	"context"
	"fmt"

	common "golang.conradwood.net/apis/common"
	"golang.conradwood.net/apis/deploymonkey"
)

func (e *DeployMonkey) AppChanged(ctx context.Context, req *deploymonkey.AppChangeRequest) (*common.Void, error) {
	appdef := req.AppRef.AppDef
	appname := fmt.Sprintf("[%s appdef-id=%d Binary=%s buildid=%d, deplid=%s, artefactid=%d]", req.DeploymentID, appdef.ID, appdef.Binary, appdef.BuildID, appdef.DeploymentID, appdef.ArtefactID)
	fmt.Printf("Autodeployer reports app %s has changed to %v (undeplyrequest=%v) after %d seconds runtime\n", appname, req.Status, req.UndeployRequestReceived, req.RuntimeSeconds)
	return &common.Void{}, nil
}
