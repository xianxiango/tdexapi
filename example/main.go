package main

import (
	"context"
	"fmt"
	"os"
	"tdexapi"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
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

	// pc, err := t.FuturesCloseAll([]tdex.FuturesCloseAllRequest{111, 222})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v", pc)

	// dc, err := t.FuturesClose([]tdex.FuturesCloseRequest{
	// 	{}, {}, {},
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v", dc)

	cc, err := t.Balances(tdex.BalancesRequest{Type: 1})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", cc)

	// ac, err := t.UserInfo()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v", ac)
}
