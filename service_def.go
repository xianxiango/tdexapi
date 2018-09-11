package tdex

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
)

type Service interface {
	ProductCoins() (*ProductCoins, error)
	UserInfo() (*UserInfo, error)
	Balances(bqs BalancesRequest) (*Balances, error)
	Balance(bq BalanceRequest) (*Balance, error)
	Withdraw(wq WithdrawRequest) (*Withdraw, error)
	Switch(sq SwitchRequest) (*Switch, error)
	FuturesOpen(fr FuturesOpenRequest) (*FuturesOpen, error)
	FuturesClose(fcr []FuturesCloseRequest) (*FuturesClose, error)
	FuturesCloseAll(fcar []FuturesCloseAllRequest) (*FuturesCloseAll, error)
	FuturesCancel(fcr []FuturesCancelRequest) (*FuturesCancel, error)
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
}
type apiService struct {
	URL    string
	APIKey string
	Signer Signer
	Logger log.Logger
	Ctx    context.Context
}

func NewAPIService(url, apiKey string, signer string, logger log.Logger, ctx context.Context) Service {
	if logger == nil {
		logger = log.NewNopLogger()
	}
	if ctx == nil {
		ctx = context.Background()
	}
	TSigner := &HmacSigner{
		Key: []byte(signer),
	}
	return &apiService{
		URL:    url,
		APIKey: apiKey,
		Signer: TSigner,
		Logger: logger,
		Ctx:    ctx,
	}
}

func (as *apiService) request(method string, endpoint string, params map[string]interface{},
	apiKey bool, sign bool) (*http.Response, error) {
	transport := &http.Transport{}
	client := &http.Client{
		Transport: transport,
	}

	buf, err := json.Marshal(params)
	body := bytes.NewBuffer(buf)

	url := fmt.Sprintf("%s/openapi/v1%s", as.URL, endpoint)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create request")
	}
	req.WithContext(as.Ctx)
	if apiKey {
		t := time.Now()
		expires := strconv.FormatInt(t.UTC().UnixNano(), 10)
		// expires = expires[:10]
		sign := as.Signer.Sign([]byte(endpoint+expires), buf)
		req.Header.Set("api-expires", expires)
		req.Header.Set("api-key", as.APIKey)
		req.Header.Set("api-signature", sign)
		req.Header.Set("Content-Type", "application/json")

	}
	// q := req.URL.Query()
	// for key, val := range params {
	// 	q.Add(key, val)
	// }

	// if sign {
	// 	t := time.Now()
	// 	expires := strconv.FormatInt(t.UTC().UnixNano(), 10)
	// 	expires = expires[:10]
	// 	sign := as.Signer.Sign([]byte(endpoint + expires + "{}"))
	// 	level.Debug(as.Logger).Log("queryString", q.Encode())
	// 	q.Add("api-signature", sign)
	// 	q.Add("api-expires", expires)
	// 	level.Debug(as.Logger).Log("api-signature", sign)
	// 	level.Debug(as.Logger).Log("api-expires", expires)
	// }
	// req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
