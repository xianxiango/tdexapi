# Tdex API

## Getting started

```go
var logger log.Logger
logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
logger = level.NewFilter(logger, level.AllowAll())
logger = log.With(logger, "time", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
ctx, _ := context.WithCancel(context.Background())
tdexService := tdex.NewAPIService(
    "URL",
    "api-key",
    "api-secret",
    logger,
    ctx,
)
t := tdex.NewTdex(tdexService)
```

## Golang Tdex API

#### 用户详细信息

```go
td, err := t.UserInfo()
if err != nil {
    panic(err)
}
fmt.Printf("%+v", td)
```

#### 查询有所余额

```go
td, err := t.Balances(tdex.BalancesRequest{
    Type: 1,
})
if err != nil {
    panic(err)
}
fmt.Printf("%+v", td)
```
```
params:
    type uint32
```

#### 查询单个余额

```go
td, err := t.Balance(tdex.BalanceRequest{
    Type: 1,
    Currency: 1,
})
if err != nil {
    panic(err)
}
fmt.Printf("%+v", td)
```
```
params:
    currency uint32
    type uint32
```

#### 提现

```go
td, err := t.Withdraw(tdex.WithdrawRequest{
    Currency: 1,
    Address: "string",
    Amount: 1,
})
if err != nil {
    panic(err)
}
fmt.Printf("%+v", td)
```
```
params:
	currency uint32	 源币种。比如 1 - BTC
	address	string	提现地址
	amount	float64 数量
```

#### 资金划转

```go
td, err := t.Switch(tdex.SwitchRequest{
    Currency: 1,
    Direction: 1,
    Amount: 1,
})
if err != nil {
    panic(err)
}
fmt.Printf("%+v", td)
```
```
params:
    currency uint32
    direction uint32
    amount float64
```

#### 开仓

```go
td, err := t.FuturesOpen(tdex.FuturesOpenRequest{
    Cid:int64, Side: uint32, Scale: float64
})
if err != nil {
    panic(err)
}
fmt.Printf("%+v", td)
```
```
params:
    cid	int64 产品

side	uint32	交易方向。0 - buy 1 - sell。参考

scale	float64 杠杆

volume	uint32	数量

distance	bool 触发时使用价距或价格

price	float64 限价 <=0:市价 singular

timely	uint32	时效性(限价单用) singular。参考

timelyParam	int32	时效性参数

passive	bool 被动性

visible	int32 显示数量 <0:全部可见 >=0隐藏

strategy	uint32	策略。参考

better	bool 以买一卖一价进入订单簿

variable	uint32	策略使用的变量(条件订单用) singular。参考

constant	float64 策略中常量(条件订单用) singular

sl	Object	止损 singular

	-distance	bool 价距|报价 市价单只用用价距

	-param float64 值

tp	Object	止盈 singular

	-distance bool 价距|报价 市价单只用用价距

	-param	float64 值
```


#### 批量平仓

```go
td, err := t.FuturesClose([]tdex.FuturesCloseRequest{
	{}, {}, {},
})
if err != nil {
    panic(err)
}
fmt.Printf("%+v", td)
```
```
params:
     
     list Object[]	持仓列表

		  -cid	int64	产品
		
		  -id	uint64	仓位
		
		  -distance	bool	是否为相对价格
		
		  -price	float64	限价 <=0: 市价
		
		  -timely	uint32	时效性(限价单用) singular。参考
		
		  -timelyParam	int32	时效性参数
		
		  -strategy	uint32	策略。参考
		
		  -variable	uint32	策略使用的变量(条件订单用) singular。参考
		
		  -constant	float64	策略中常量(条件订单用) singular
		
		  -passive	bool	被动性
		
		  -visible	int32	显示数量 <0:全部可见 >=0隐藏
		
		  -better bool	 以买一卖一价进入订单簿
```

#### 全部平仓

```go
td, err := t.FuturesCloseAll([]tdex.FuturesCloseAllRequest{
    1,2,3,
})
if err != nil {
    panic(err)
}
fmt.Printf("%+v", td)
```
```
params:
	list	uint64[]	产品列表 [CID,...]
```

#### 批量取消

```go
td, err := t.FuturesCancel([]tdex.FuturesCancelRequest{
    {Cid:1,ID:1},{Cid:1,ID:1},
})
if err != nil {
    panic(err)
}
fmt.Printf("%+v", td)
```
```
params:
	list	Object[]		委托列表
        cid     int64       产品
        id      uint64      仓位
```

#### 设置止损

```go
td, err := t.Setsl(tdex.SetslRequest{Cid: int64, ID: uint64,...})
if err != nil {
    panic(err)
}
fmt.Printf("%+v", td)
```
```
parmas:

	cid	int64	产品

	id	uint64	仓位
	
	distance	bool	是否为相对价格
	
	price	float64	限价 <=0: 市价
	
	timely	uint32	时效性(限价单用) singular。参考
	
	timelyParam	int32	时效性参数
	
	strategy	uint32	策略。参考
	
	variable	uint32	策略使用的变量(条件订单用) singular。参考
	
	constant	float64 策略中常量(条件订单用) singular
	
	passive	bool	被动性
	
	visible	int32	显示数量 <0:全部可见 >=0隐藏
	
	better	bool	以买一卖一价进入订单簿
```

#### 设置止盈

```go
td, err := t.Settp(tdex.SettpRequest{Cid: int64, ID: uint64,...})
if err != nil {
    panic(err)
}
fmt.Printf("%+v", td)
```
```
parmas:

	cid	int64	产品

	id	uint64	仓位
	
	distance	bool	是否为相对价格
	
	price	float64	限价 <=0: 市价
	
	timely	uint32	时效性(限价单用) singular。参考
	
	timelyParam	int32	时效性参数
	
	strategy	uint32	策略。参考
	
	variable	uint32	策略使用的变量(条件订单用) singular。参考
	
	constant	float64 策略中常量(条件订单用) singular
	
	passive	bool	被动性
	
	visible	int32	显示数量 <0:全部可见 >=0隐藏
	
	better	bool	以买一卖一价进入订单簿
```

#### 合仓

```go
td, err := t.Merge(tdex.MergeRequest{Cid: int64, List: []uint64{}})
if err != nil {
    panic(err)
}
fmt.Printf("%+v", td)
```
```
params:

	cid	int64	产品

	list	uint64[]	要合仓的仓位列表
```

#### 分仓

```go
td, err := t.Split(tdex.SplitRequest{Cid: int64, ID: uint64, Volume: uint64})
if err != nil {
    panic(err)
}
fmt.Printf("%+v", td)
```
```
params:

	cid	int64	产品

	id	uint64	仓位

	volume	uint64	数量
```

#### 获取 用户选项

```go
td, err := t.SchemeGet(tdex.SchemeGetRequest{Cid: int64})
if err != nil {
    panic(err)
}
fmt.Printf("%+v", td)
```
```
params:

	cid	uint32	产品
```
#### 设置 用户选项

```go
td, err := t.SchemeSet(tdex.SchemeSetRequest{Cid: int64})
if err != nil {
    panic(err)
}
fmt.Printf("%+v", td)
```
```
params:

	cid	uint32	产品
```

#### 获取订单

```go
td, err := t.GetOrders()
if err != nil {
    panic(err)
}
fmt.Printf("%+v", td)
```

#### 获取持仓

```go
td, err := t.GetPosition()
if err != nil {
    panic(err)
}
fmt.Printf("%+v", td)
```
#### 获取历史信息

```go
td, err := t.GetHistory(tdex.GetHistoryRequest{PageSize: int32, Page: int32})
if err != nil {
    panic(err)
}
fmt.Printf("%+v", td)
```
```
params:

	pageSize 可选	int32	页大小
	page 可选	int32	当前页码
```

#### 获取合约信息

```go
td, err := t.GetContract(tdex.GetContractRequest{Symbol: string})
if err != nil {
    panic(err)
}
fmt.Printf("%+v", td)
```
```
params:

	symbol	string	产品符号。目前只有 BTCUSD
```

#### 现货买入

```go
td, err := t.SpotBuy(tdex.SpotBuyRequest{Amount: float64, Price: float64, Symbol: string})
if err != nil {
    panic(err)
}
fmt.Printf("%+v", td)
```
```
params:

	amount	float64	数量

	price 可选	float64	价格。如果市价为 0，限价不为 0
	
	symbol	string	交易对。如 TDUSDT
```

#### 现货卖出

```go
td, err := t.SpotSell(tdex.SpotSellRequest{Amount: float64, Price: float64, Symbol: string})
if err != nil {
    panic(err)
}
fmt.Printf("%+v", td)
```
```
params:

	amount	float64	数量

	price 可选	float64	价格。如果市价为 0，限价不为 0
	
	symbol	string	交易对。如 TDUSDT
```

#### 现货订单历史

```go
td, err := t.SpotHistory(tdex.SpotHistoryRequest{BeginTime: string, EndTime: string, PageSize: int32, Page: int32})
if err != nil {
    panic(err)
}
fmt.Printf("%+v", td)
```
```
params:

	beginTime	string	开始时间。2017-01-01

	endTime	string	结束时间。2017-09-13
	
	pageSize 可选	int32	页大小
	
	page 可选	int32	当前页码
```

