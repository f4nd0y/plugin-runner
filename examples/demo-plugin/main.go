//go:build tinygo.wasm

package main

import (
	"context"
	"fmt"

	"github.com/f4nd0y/plugin-runner/proto/common"
)

// main is required for TinyGo to compile to Wasm.
func main() {
	common.RegisterCommon(DemoPlugin{})
}

type DemoPlugin struct{}

var _ common.Common = (*DemoPlugin)(nil)

func (m DemoPlugin) Prepare(_ context.Context, request *common.Request) (*common.Reply, error) {
	return &common.Reply{
		Code:    0,
		Message: fmt.Sprintf("PrepareRequest Done"),
	}, nil
}

func (m DemoPlugin) Run(_ context.Context, request *common.Request) (*common.Reply, error) {
	return &common.Reply{
		Code:    0,
		Message: fmt.Sprintf("RunRequest Done"),
	}, nil
}

func (m DemoPlugin) Rollback(_ context.Context, request *common.Request) (*common.Reply, error) {
	return &common.Reply{
		Code:    0,
		Message: fmt.Sprintf("RollbackRequest Done"),
	}, nil
}

func (m DemoPlugin) Check(_ context.Context, request *common.Request) (*common.Reply, error) {
	return &common.Reply{
		Code:    0,
		Message: fmt.Sprintf("CheckRequest Done"),
	}, nil
}
