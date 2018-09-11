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
    Type uint32
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
    Currency uint32
    Type uint32
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
	Currency uint32	 源币种。比如 1 - BTC
	Address	string	提现地址
	Amount	float64 数量
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
    Currency uint32
    Direction uint32
    Amount float64
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
    Cid			int64 产品

	Side		uint32	交易方向。0 - buy 1 - sell。参考

	Scale		float64 杠杆

	Volume		uint32	数量

	Distance	bool 触发时使用价距或价格

	Price		float64 限价 <=0:市价 singular

	Timely		uint32	时效性(限价单用) singular。参考

	TimelyParam	int32	时效性参数

	Passive		bool 被动性

	Visible		int32 显示数量 <0:全部可见 >=0隐藏

	Strategy	uint32	策略。参考

	Better		bool 以买一卖一价进入订单簿

	Variable	uint32	策略使用的变量(条件订单用) singular。参考

	Constant	float64 策略中常量(条件订单用) singular

	Sl			Object	止损 singular

		-Distance	bool 价距|报价 市价单只用用价距

		-Param float64 值

	Tp	Object	止盈 singular

		-Distance bool 价距|报价 市价单只用用价距

		-Param	float64 值
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
     
     List Object[]	持仓列表

		  -Cid			int64	产品
		
		  -ID			uint64	仓位
		
		  -Distance		bool	是否为相对价格
		
		  -Price		float64	限价 <=0: 市价
		
		  -Timely		uint32	时效性(限价单用) singular。参考
		
		  -TimelyParam	int32	时效性参数
		
		  -Strategy		uint32	策略。参考
		
		  -Variable		uint32	策略使用的变量(条件订单用) singular。参考
		
		  -Constant		float64	策略中常量(条件订单用) singular
		
		  -Passive		bool	被动性
		
		  -Visible		int32	显示数量 <0:全部可见 >=0隐藏
		
		  -Better 		bool	 以买一卖一价进入订单簿
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
	List	uint64[]	产品列表 [CID,...]
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
	List	Object[]		委托列表
        Cid     int64       产品
        ID      uint64      仓位
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

	Cid			int64	产品

	ID			uint64	仓位
	
	Distance	bool	是否为相对价格
	
	Price		float64	限价 <=0: 市价
	
	Timely		uint32	时效性(限价单用) singular。参考
	
	TimelyParam	int32	时效性参数
	
	Strategy	uint32	策略。参考
	
	Variable	uint32	策略使用的变量(条件订单用) singular。参考
	
	Constant	float64 策略中常量(条件订单用) singular
	
	Passive		bool	被动性
	
	Visible		int32	显示数量 <0:全部可见 >=0隐藏
	
	Better		bool	以买一卖一价进入订单簿
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

	Cid	 		int64	产品

	ID			uint64	仓位
	
	Distance	bool	是否为相对价格
	
	Price		float64	限价 <=0: 市价
	
	Timely		uint32	时效性(限价单用) singular。参考
	
	TimelyParam	int32	时效性参数
	
	Strategy	uint32	策略。参考
	
	Variable	uint32	策略使用的变量(条件订单用) singular。参考
	
	Constant	float64 策略中常量(条件订单用) singular
	
	Passive		bool	被动性
	
	Visible		int32	显示数量 <0:全部可见 >=0隐藏
	
	Better		bool	以买一卖一价进入订单簿
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

	Cid		int64	产品

	List	uint64[]	要合仓的仓位列表
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

	Cid		int64	产品

	ID		uint64	仓位

	Volume	uint64	数量
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

	Cid	uint32	产品
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

	Cid	uint32	产品
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

	PageSize 可选	int32	页大小
	Page 可选		int32	当前页码
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

	Symbol	string	产品符号。目前只有 BTCUSD
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

	Amount		float64	数量

	Price 可选	float64	价格。如果市价为 0，限价不为 0
	
	Symbol		string	交易对。如 TDUSDT
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

	Amount	float64	数量

	Price 可选	float64	价格。如果市价为 0，限价不为 0
	
	Symbol	string	交易对。如 TDUSDT
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

	BeginTime	string	开始时间。2017-01-01

	EndTime	string	结束时间。2017-09-13
	
	PageSize 可选	int32	页大小
	
	Page 可选	int32	当前页码
```

