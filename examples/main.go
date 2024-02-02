package main

import (
	"context"
	"fmt"
	"log"

	"github.com/f4nd0y/plugin-runner/proto/common"
)

//go:generate tinygo build -o demo-plugin/demo.wasm -scheduler=none -target=wasi --no-debug demo-plugin/main.go

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()
	p, err := common.NewCommonPlugin(ctx)
	if err != nil {
		return err
	}

	demoPlugin, err := p.Load(ctx, "demo-plugin/demo.wasm")
	if err != nil {
		return err
	}
	defer demoPlugin.Close(ctx)

	reply, err := demoPlugin.Prepare(ctx, &common.Request{})
	if err != nil {
		return err
	}
	fmt.Println(reply.GetMessage())

	reply, err = demoPlugin.Run(ctx, &common.Request{})
	if err != nil {
		return err
	}
	fmt.Println(reply.GetMessage())

	reply, err = demoPlugin.Check(ctx, &common.Request{})
	if err != nil {
		return err
	}
	fmt.Println(reply.GetMessage())

	reply, err = demoPlugin.Rollback(ctx, &common.Request{})
	if err != nil {
		return err
	}
	fmt.Println(reply.GetMessage())
	return nil
}
