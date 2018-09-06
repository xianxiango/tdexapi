package tdex

import "fmt"

type Tdex interface {
	ProductCoins() (*ProductCoins, error)
	UserInfo() (*UserInfo, error)
	Balances(bqs BalancesRequest) (*Balances, error)
	Balance(bq BalanceRequest) (*Balance, error)
	Withdraw(wq WithdrawRequest) (*Withdraw, error)
	Switch(sq SwitchRequest) (*Switch, error)
	FuturesOpen(fr FuturesOpenRequest) (*FuturesOpen, error)
	FuturesClose(fcr []FuturesCloseRequest) (*FuturesClose, error)
	FuturesCloseAll(fcar []FuturesCloseAllRequest) (*FuturesCloseAll, error)
	FururesCancel(fcr []FururesCancelRequest) (*FururesCancel, error)
	FururesReplace(frr FururesReplaceRequest) (*FururesReplace, error)
	Setsl(sr SetslRequest) (*Setsl, error)
	Settp(sr SettpRequest) (*Settp, error)
	Merge(mr MergeRequest) (*Merge, error)
	Split(sr SplitRequest) (*Split, error)
	Setup(sr SetupRequest) (*Setup, error)
	SchemeSet(ssr SchemeSetRequest) (*SchemeSet, error)
	SchemeGet(sgr SchemeGetRequest) (*SchemeGet, error)
	GetOrders() (*GetOrders, error)
	GetPosition() (*GetPosition, error)
	GetHistory(ghr GetHistoryRequest) (*GetHistory, error)
	GetContract(gcr GetContractRequest) (*GetContract, error)
	SpotBuy(sbr SpotBuyRequest) (*SpotBuy, error)
	SpotSell(ssr SpotSellRequest) (*SpotSell, error)
	SpotCancel(scr SpotCancelRequest) (*SpotCancel, error)
	SpotOrders() (*SpotOrders, error)
	SpotHistory(shr SpotHistoryRequest) (*SpotHistory, error)
	SpotStat(ssr SpotStatRequest) (*SpotStat, error)
	RobotClose(ssr RobotCloseRequest) (*RobotClose, error)
}
type tdex struct {
	Service Service
}

// Error represents Binance error structure with error code and message.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

// Error returns formatted error message.
func (e Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}
func NewTdex(service Service) Tdex {
	return &tdex{
		Service: service,
	}
}

type Data struct {
}
type ProductCoins struct {
	Status uint32
	Data   ProductCoinsData
	Errmsg string
}
type ProductCoinsData struct {
	List []*ProductCoinsList
}
type ProductCoinsList struct {
	ID            int64
	Code          string
	Symbol        string
	Family        string
	MinTrade      float64
	MinCash       float64
	Fee           float64
	FeeType       string
	Digits        uint32
	ShowDigits    uint32
	Confirmations uint32
	Op            uint32
	Types         uint32
	Exchange      ProductCoinsExchange
}
type ProductCoinsExchange struct {
	Max uint32
	Min float64
	To  []uint32
}

func (t *tdex) ProductCoins() (*ProductCoins, error) {
	return t.Service.ProductCoins()
}

type UserInfo struct {
	Status uint32
	Data   UserInfoData
	Errmsg string
}
type UserInfoData struct {
	AccountID     int64
	UID           string
	Vip           uint32
	GgAuth        bool
	Email         string
	Name          string
	Mobile        string
	Lang          string
	LastLoginTime int64
	LastLoginIP   string
}

func (t *tdex) UserInfo() (*UserInfo, error) {
	return t.Service.UserInfo()
}

type Balances struct {
	Status uint32
	Data   BalancesData
	Errmsg string
}
type BalancesData struct {
	List []*BalancesList
}
type BalancesList struct {
	Type      uint32
	Currency  uint32
	Quantity  float64
	Withdraw  float64
	Transfer  float64
	Deposit   float64
	Lock      uint32
	Available uint32
	Status    uint32
}

func (t *tdex) Balances(bq BalancesRequest) (*Balances, error) {
	return t.Service.Balances(bq)
}

type Balance struct {
	Status uint32
	Data   BalanceData
	Errmsg string
}
type BalanceData struct {
	Type      uint32
	Currency  uint32
	Quantity  float64
	Withdraw  float64
	Transfer  float64
	Deposit   float64
	Lock      uint32
	Available uint32
	Status    uint32
}

func (t *tdex) Balance(bq BalanceRequest) (*Balance, error) {
	return t.Service.Balance(bq)
}

type Withdraw struct {
	Status uint32
	Data   WithdrawData
	Errmsg string
}
type WithdrawData struct {
	Currency uint32
	Quantity float64
	Withdraw float64
	Transfer float64
	Deposit  float64
	OrderID  string
}

func (t *tdex) Withdraw(wq WithdrawRequest) (*Withdraw, error) {
	return t.Service.Withdraw(wq)
}

type Switch struct {
	Status uint32
	Data   SwitchData
	Errmsg string
}
type SwitchData struct {
	Currency  uint32
	Quantity  float64
	Unconfirm float64
	Withdraw  float64
	Transfer  float64
	Deposit   float64
	OrderId   string
}

func (t *tdex) Switch(wq SwitchRequest) (*Switch, error) {
	return t.Service.Switch(wq)
}

type FuturesOpen struct {
	Status uint32
	Data   FuturesOpenData
	Errmsg string
}
type FuturesOpenData struct {
}

func (t *tdex) FuturesOpen(wq FuturesOpenRequest) (*FuturesOpen, error) {
	return t.Service.FuturesOpen(wq)
}

type FuturesClose struct {
	Status uint32
	Data   FuturesCloseData
	Errmsg string
}
type FuturesCloseData struct {
	Result FuturesCloseResult
}
type FuturesCloseResult struct {
	Code  uint32
	Error string
}

func (t *tdex) FuturesClose(wq []FuturesCloseRequest) (*FuturesClose, error) {
	return t.Service.FuturesClose(wq)
}

type FuturesCloseAll struct {
	Status uint32
	Data   FuturesCloseAllData
	Errmsg string
}
type FuturesCloseAllData struct {
}

func (t *tdex) FuturesCloseAll(wq []FuturesCloseAllRequest) (*FuturesCloseAll, error) {
	return t.Service.FuturesCloseAll(wq)
}

type FururesCancel struct {
	Status uint32
	Data   FururesCancelData
	Errmsg string
}
type FururesCancelData struct {
}

func (t *tdex) FururesCancel(wq []FururesCancelRequest) (*FururesCancel, error) {
	return t.Service.FururesCancel(wq)
}

type FururesReplace struct {
	Status uint32
	Data   FururesReplaceData
	Errmsg string
}
type FururesReplaceData struct {
}

func (t *tdex) FururesReplace(wq FururesReplaceRequest) (*FururesReplace, error) {
	return t.Service.FururesReplace(wq)
}

type Setsl struct {
	Status uint32
	Data   SetslData
	Errmsg string
}
type SetslData struct {
}

func (t *tdex) Setsl(wq SetslRequest) (*Setsl, error) {
	return t.Service.Setsl(wq)
}

type Settp struct {
	Status uint32
	Data   SettpData
	Errmsg string
}
type SettpData struct {
}

func (t *tdex) Settp(wq SettpRequest) (*Settp, error) {
	return t.Service.Settp(wq)
}

type Merge struct {
	Status uint32
	Data   MergeData
	Errmsg string
}
type MergeData struct {
}

func (t *tdex) Merge(wq MergeRequest) (*Merge, error) {
	return t.Service.Merge(wq)
}

type Split struct {
	Status uint32
	Data   SplitData
	Errmsg string
}
type SplitData struct {
}

func (t *tdex) Split(wq SplitRequest) (*Split, error) {
	return t.Service.Split(wq)
}

type Setup struct {
	Status uint32
	Data   SetupData
	Errmsg string
}
type SetupData struct {
}

func (t *tdex) Setup(wq SetupRequest) (*Setup, error) {
	return t.Service.Setup(wq)
}

type SchemeGet struct {
	Status uint32
	Data   SchemeGetData
	Errmsg string
}
type SchemeGetData struct {
	Shared bool
	Merged bool
}

func (t *tdex) SchemeGet(wq SchemeGetRequest) (*SchemeGet, error) {
	return t.Service.SchemeGet(wq)
}

type SchemeSet struct {
	Status uint32
	Data   SchemeSetData
	Errmsg string
}
type SchemeSetData struct {
}

func (t *tdex) SchemeSet(wq SchemeSetRequest) (*SchemeSet, error) {
	return t.Service.SchemeSet(wq)
}

type GetOrders struct {
	Status uint32
	Data   GetOrdersData
	Errmsg string
}
type GetOrdersData struct {
	List []GetOrdersList
}
type GetOrdersList struct {
	ID          int64
	Pid         int64
	Cid         uint32
	Rid         int64
	Side        uint32
	Kind        uint32
	Attempt     uint32
	Scale       float64
	Volume      uint32
	Visible     uint32
	Filled      uint32
	Distance    uint32
	Price       float64
	Shared      uint32
	Timely      uint32
	TimelyParam uint32
	Deadline    uint32
	Passive     uint32
	Better      uint32
	Strategy    uint32
	Variable    uint32
	Constant    float64
	Relative    uint32
	Vertex      int64
	Activated   int64
	Triggered   int64
	State       uint32
	SerialID    uint32
	Notes       string
	CreatedAt   string
	UpdatedAt   string
}

func (t *tdex) GetOrders() (*GetOrders, error) {
	return t.Service.GetOrders()
}

type GetPosition struct {
	Status uint32
	Data   GetPositionData
	Errmsg string
}
type GetPositionData struct {
	List []GetPositionList
}
type GetPositionList struct {
	ID        int64
	Cid       uint32
	Shared    uint32
	Side      uint32
	Price     float64
	Volume    uint32
	Scale     float64
	Margin    float64
	Repay     float64
	Force     float64
	Fee       float64
	Notes     string
	Completed uint32
	Bankrupt  uint32
	SerialID  uint32
	Index     uint32
	CreatedAt string
	UpdatedAt string
}

func (t *tdex) GetPosition() (*GetPosition, error) {
	return t.Service.GetPosition()
}

type GetHistory struct {
	Status uint32
	Data   GetHistoryData
	Errmsg string
}
type GetHistoryData struct {
	List      []GetHistoryList
	Total     int32
	PageCount int32
	PageSize  int32
	Page      int32
}
type GetHistoryList struct {
	ID          int64
	Lid         int64
	Cid         uint32
	Pid         int64
	Rid         int64
	Strategy    uint32
	Variable    uint32
	Constant    float64
	Relative    uint32
	Vertex      int64
	Attempt     uint32
	Side        uint32
	Kind        uint32
	Scale       float64
	Volume      uint32
	Filled      uint32
	Distance    uint32
	Price       float64
	Shared      uint32
	Timely      uint32
	TimelyParam uint32
	Deadline    uint32
	Passive     uint32
	Visible     uint32
	Better      uint32
	Activated   int64
	Triggered   int64
	State       uint32
	SerialID    uint32
	Origin      int64
	Reason      int32
	Fee         float64
	Rote        uint32
	FinalPrice  float64
	XVolume     int32
	Action      uint32
	Notes       string
	CreatedAt   string
}

func (t *tdex) GetHistory(wq GetHistoryRequest) (*GetHistory, error) {
	return t.Service.GetHistory(wq)
}

type GetContract struct {
	Status uint32
	Data   GetContractData
	Errmsg string
}
type GetContractData struct {
	MarkPrice             string
	FairBasis             string
	FairBasisRate         string
	OpenInterest          string
	Turnover24h           string
	FundingRate           string
	FundingTimestamp      string
	IndicativeFundingRate string
}

func (t *tdex) GetContract(wq GetContractRequest) (*GetContract, error) {
	return t.Service.GetContract(wq)
}

type SpotBuy struct {
	Status uint32
	Data   SpotBuyData
	Errmsg string
}
type SpotBuyData struct {
	ID int
}

func (t *tdex) SpotBuy(sq SpotBuyRequest) (*SpotBuy, error) {
	return t.Service.SpotBuy(sq)
}

type SpotSell struct {
	Status uint32
	Data   SpotSellData
	Errmsg string
}
type SpotSellData struct {
	ID int
}

func (t *tdex) SpotSell(sq SpotSellRequest) (*SpotSell, error) {
	return t.Service.SpotSell(sq)
}

type SpotCancel struct {
	Status uint32
	Data   SpotCancelData
	Errmsg string
}
type SpotCancelData struct {
}

func (t *tdex) SpotCancel(sq SpotCancelRequest) (*SpotCancel, error) {
	return t.Service.SpotCancel(sq)
}

type SpotOrders struct {
	Status uint32
	Data   SpotOrdersData
	Errmsg string
}
type SpotOrdersData struct {
	List []SpotOrdersList
}
type SpotOrdersList struct {
	ID         int
	Symbol     string
	Currency   int
	Rid        int
	Direction  int
	Quantity   float64
	Completed  float64
	Price      float64
	Reality    float64
	Commission float64
	State      int
	CreateTime string
}

func (t *tdex) SpotOrders() (*SpotOrders, error) {
	return t.Service.SpotOrders()
}

type SpotHistory struct {
	Status uint64
	Data   SpotHistoryData
	Errmsg string
}
type SpotHistoryData struct {
	List      []SpotHistoryList
	Total     int32
	PageCount int32
	PageSize  int32
	Page      int32
}
type SpotHistoryList struct {
	ID         int
	Symbol     string
	Currency   int
	Rid        int
	Direction  int
	Quantity   float64
	Completed  float64
	Price      float64
	Reality    float64
	Commission float64
	State      int
	CreateTime string
}

func (t *tdex) SpotHistory(sq SpotHistoryRequest) (*SpotHistory, error) {
	return t.Service.SpotHistory(sq)
}

type SpotStat struct {
	Status uint64
	Data   SpotStatData
	Errmsg string
}
type SpotStatData struct {
	Quantity float64
	Worth    float64
}

func (t *tdex) SpotStat(sq SpotStatRequest) (*SpotStat, error) {
	return t.Service.SpotStat(sq)
}

type RobotClose struct {
	Status uint64
	Data   RobotCloseData
	Errmsg string
}
type RobotCloseData struct {
}

func (t *tdex) RobotClose(sq RobotCloseRequest) (*RobotClose, error) {
	return t.Service.RobotClose(sq)
}

type BalancesRequest struct {
	Type uint32
}
type BalanceRequest struct {
	Currency uint32
	Type     uint32
}
type SwitchRequest struct {
	Currency  uint32
	Direction uint32
	Amount    float64
}

type WithdrawRequest struct {
	Currency uint32
	Address  string
	Amount   float64
}

type FuturesOpenRequest struct {
	Cid         int64
	Side        uint32
	Scale       float64
	Volume      uint32
	Distance    bool
	Price       float64
	Timely      uint32
	TimelyParam int32
	Passive     bool
	Visible     int32
	Strategy    uint32
	Better      bool
	Variable    uint32
	Constant    float64
	Sl          FuturesOpenRequestSl
	Tp          FuturesOpenRequestSTp
}
type FuturesOpenRequestSl struct {
	Distance bool
	Param    float64
}
type FuturesOpenRequestSTp struct {
	Distance bool
	Param    float64
}

type FuturesCloseRequest struct {
	Cid         int64
	ID          uint64
	Distance    bool
	Price       float64
	Timely      uint32
	TimelyParam int32
	Strategy    uint32
	Variable    uint32
	Constant    float64
	Passive     bool
	Visible     int32
	Better      bool
}
type FuturesCloseAllRequest uint64

type FururesCancelRequest struct {
	Cid int64
	Id  uint64
}

type FururesReplaceRequest struct {
	ID          uint64
	Relative    bool
	Cid         int64
	Side        uint32
	Scale       float64
	Volume      uint32
	Distance    bool
	Price       float64
	Timely      uint32
	TimelyParam int32
	Passive     bool
	Visible     int32
	Strategy    uint32
	Better      bool
	Variable    uint32
	Constant    float64
}
type SetslRequest struct {
	Cid         int64
	ID          uint64
	Distance    bool
	Price       float64
	Timely      uint32
	TimelyParam int32
	Strategy    uint32
	Variable    uint32
	Constant    float64
	Passive     bool
	Visible     int32
	Better      bool
}
type SettpRequest struct {
	Cid         int64
	ID          uint64
	Distance    bool
	Price       float64
	Timely      uint32
	TimelyParam int32
	Strategy    uint32
	Variable    uint32
	Constant    float64
	Passive     bool
	Visible     int32
	Better      bool
}
type MergeRequest struct {
	Cid  int64
	List []uint64
}
type SplitRequest struct {
	Cid    int64
	ID     uint64
	Volume uint64
}
type SetupRequest struct {
	Cid    int64
	ID     uint64
	Target uint32
}
type SchemeSetRequest struct {
	Cid     uint32
	Options SchemeSetRequestOptions
}
type SchemeSetRequestOptions struct {
	Shared bool
	Merged bool
}
type SchemeGetRequest struct {
	Cid uint32
}
type GetHistoryRequest struct {
	PageSize int32
	Page     int32
}
type GetContractRequest struct {
	Symbol string
}
type SpotBuyRequest struct {
	Amount float64
	Price  float64
	Symbol string
}
type SpotSellRequest struct {
	Amount float64
	Price  float64
	Symbol string
}
type SpotCancelRequest struct {
	ID     int
	Symbol string
}
type SpotHistoryRequest struct {
	BeginTime string
	EndTime   string
	PageSize  int32
	Page      int32
}
type SpotStatRequest struct {
	Symbol    string
	BeginTime int64
	EndTime   int64
}
type RobotCloseRequest struct {
	Type string
}
