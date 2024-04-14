package main

import (
	"fmt"
	"os"

	"github.com/dshills/ai-manager/ai"
)

func converse(aimgr *ai.Manager, tData ai.ThreadData, query string) *ai.GeneratorResponse {
	thread, err := aimgr.NewThread(tData)
	if err != nil {
		fmt.Printf("%s %s\n", tData.Model, err)
		os.Exit(1)
	}

	resp, err := thread.Converse(query)
	if err != nil {
		fmt.Printf("%s %s\n", tData.Model, err)
		os.Exit(1)
	}
	return resp
}
