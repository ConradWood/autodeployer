package main

import (
	"context"

	common "golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/deploymonkey"
	cm "golang.conradwood.net/deploymonkey/common"
	"golang.conradwood.net/deploymonkey/deployq"
)

func (dm *DeployMonkey) GetStatus(ctx context.Context, req *common.Void) (*pb.Status, error) {
	b := cm.IsAutoDeployingCurrently() || deployq.IsDeploying()
	res := &pb.Status{
		CurrentlyApplyingSuggestions: b,
	}
	return res, nil

}
