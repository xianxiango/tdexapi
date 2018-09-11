package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	tdex "github.com/tdex-exchange/golang-tdex-api"
)

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowAll())
	logger = log.With(logger, "time", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	ctx, _ := context.WithCancel(context.Background())
	tdexService := tdex.NewAPIService(
		"https://tl.tdex.com",
		"7AGFdayY1ePgJbCDsa1gVivym9UhUSG3oazY3fwAy1dpX5UxsmHa7EzuQVy1XzoJ",
		"7AGFdbCfPWCB2wh4W2Ln5roMRhnhQVvMvWCsNFpEnc4MLyE28TZX38v92uZ48NXf",
		logger,
		ctx,
	)
	t := tdex.NewTdex(tdexService)

	td, err := t.FuturesCancel([]tdex.FuturesCancelRequest{
		{Cid: 1, ID: 1}, {Cid: 1, ID: 1},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", td)
}
