package main

import (
	"context"

	common "golang.conradwood.net/apis/common"
	"golang.conradwood.net/apis/deploymonkey"
)

func (e *DeployMonkey) AppChanged(ctx context.Context, req *deploymonkey.AppChangeRequest) (*common.Void, error) {
	return &common.Void{}, nil
}
