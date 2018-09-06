# Tdex API

## Getting started

```go
var logger log.Logger
logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
logger = level.NewFilter(logger, level.AllowAll())
logger = log.With(logger, "time", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
ctx, _ := context.WithCancel(context.Background())
tdexService := tdex.NewAPIService(
    "https://tl.tdex.com",
    "api-key",
    "api-secret",
    logger,
    ctx,
)
t := tdex.NewTdex(tdexService)
```

## Golang Tdex API

#### ProductCoins

```go
pc, err := b.ProductCoins()
if err != nil {
    panic(err)
}
fmt.Printf("%+v", pc)
```

#### Balances

```go
bl, err := b.Balances(binance.BalancesRequest{
    Type: 1
})
if err != nil {
    panic(err)
}
fmt.Printf("%+v", bl)
```

#### FuturesCloseAll

```go
fc, err := t.FuturesCloseAll([]tdex.FuturesCloseAllRequest{
    111,
    222
})
if err != nil {
    panic(err)
}
fmt.Printf("%+v", fc)
```
 