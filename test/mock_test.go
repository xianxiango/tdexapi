package tdex_test

import (
	"tdexapi/tdex"

	"github.com/stretchr/testify/mock"
)

type ServiceMock struct {
	mock.Mock
}

func (m *ServiceMock) ProductCoins() (*tdex.ProductCoins, error) {
	args := m.Called()
	return args.Get(0).(*tdex.ProductCoins), args.Error(1)
}
func (m *ServiceMock) UserInfo() (*tdex.UserInfo, error) {
	args := m.Called()
	return args.Get(0).(*tdex.UserInfo), args.Error(1)
}
func (m *ServiceMock) Balances() (*tdex.Balances, error) {
	args := m.Called()
	return args.Get(0).(*tdex.Balances), args.Error(1)
}

func (m *ServiceMock) Balance(obr tdex.BalanceRequest) (*tdex.Balance, error) {
	args := m.Called()
	ob, ok := args.Get(0).(*tdex.Balance)
	if !ok {
		ob = nil
	}
	return ob, args.Error(1)
}
func (m *ServiceMock) Withdraw(obr tdex.WithdrawRequest) (*tdex.Withdraw, error) {
	args := m.Called()
	ob, ok := args.Get(0).(*tdex.Withdraw)
	if !ok {
		ob = nil
	}
	return ob, args.Error(1)
}
func (m *ServiceMock) Switch(obr tdex.SwitchRequest) (*tdex.Switch, error) {
	args := m.Called()
	ob, ok := args.Get(0).(*tdex.Switch)
	if !ok {
		ob = nil
	}
	return ob, args.Error(1)
}
func (m *ServiceMock) FuturesOpen(obr tdex.FuturesOpenRequest) (*tdex.FuturesOpen, error) {
	args := m.Called()
	ob, ok := args.Get(0).(*tdex.FuturesOpen)
	if !ok {
		ob = nil
	}
	return ob, args.Error(1)
}
func (m *ServiceMock) FuturesClose(obr tdex.FuturesCloseRequest) (*tdex.FuturesClose, error) {
	args := m.Called()
	ob, ok := args.Get(0).(*tdex.FuturesClose)
	if !ok {
		ob = nil
	}
	return ob, args.Error(1)
}
func (m *ServiceMock) FuturesCloseAll(obr tdex.FuturesCloseAllRequest) (*tdex.FuturesCloseAll, error) {
	args := m.Called()
	ob, ok := args.Get(0).(*tdex.FuturesCloseAll)
	if !ok {
		ob = nil
	}
	return ob, args.Error(1)
}
func (m *ServiceMock) FururesCancel(obr tdex.FururesCancelRequest) (*tdex.FururesCancel, error) {
	args := m.Called()
	ob, ok := args.Get(0).(*tdex.FururesCancel)
	if !ok {
		ob = nil
	}
	return ob, args.Error(1)
}
func (m *ServiceMock) FururesReplace(obr tdex.FururesReplaceRequest) (*tdex.FururesReplace, error) {
	args := m.Called()
	ob, ok := args.Get(0).(*tdex.FururesReplace)
	if !ok {
		ob = nil
	}
	return ob, args.Error(1)
}
func (m *ServiceMock) Setsl(obr tdex.SetslRequest) (*tdex.Setsl, error) {
	args := m.Called()
	ob, ok := args.Get(0).(*tdex.Setsl)
	if !ok {
		ob = nil
	}
	return ob, args.Error(1)
}
func (m *ServiceMock) Settp(obr tdex.SettpRequest) (*tdex.Settp, error) {
	args := m.Called()
	ob, ok := args.Get(0).(*tdex.Settp)
	if !ok {
		ob = nil
	}
	return ob, args.Error(1)
}
func (m *ServiceMock) Merge(obr tdex.MergeRequest) (*tdex.Merge, error) {
	args := m.Called()
	ob, ok := args.Get(0).(*tdex.Merge)
	if !ok {
		ob = nil
	}
	return ob, args.Error(1)
}
func (m *ServiceMock) Split(obr tdex.SplitRequest) (*tdex.Split, error) {
	args := m.Called()
	ob, ok := args.Get(0).(*tdex.Split)
	if !ok {
		ob = nil
	}
	return ob, args.Error(1)
}
func (m *ServiceMock) Setup(obr tdex.SetupRequest) (*tdex.Setup, error) {
	args := m.Called()
	ob, ok := args.Get(0).(*tdex.Setup)
	if !ok {
		ob = nil
	}
	return ob, args.Error(1)
}
func (m *ServiceMock) SchemeSet(obr tdex.SchemeSetRequest) (*tdex.SchemeSet, error) {
	args := m.Called()
	ob, ok := args.Get(0).(*tdex.SchemeSet)
	if !ok {
		ob = nil
	}
	return ob, args.Error(1)
}
func (m *ServiceMock) SchemeGet(obr tdex.SchemeGetRequest) (*tdex.SchemeGet, error) {
	args := m.Called()
	ob, ok := args.Get(0).(*tdex.SchemeGet)
	if !ok {
		ob = nil
	}
	return ob, args.Error(1)
}
func (m *ServiceMock) GetOrders() (*tdex.GetOrders, error) {
	args := m.Called()
	return args.Get(0).(*tdex.GetOrders), args.Error(1)
}
func (m *ServiceMock) GetPosition() (*tdex.GetPosition, error) {
	args := m.Called()
	return args.Get(0).(*tdex.GetPosition), args.Error(1)
}
func (m *ServiceMock) GetHistory(obr tdex.GetHistoryRequest) (*tdex.GetHistory, error) {
	args := m.Called()
	ob, ok := args.Get(0).(*tdex.GetHistory)
	if !ok {
		ob = nil
	}
	return ob, args.Error(1)
}
func (m *ServiceMock) GetContract(obr tdex.GetContractRequest) (*tdex.GetContract, error) {
	args := m.Called()
	ob, ok := args.Get(0).(*tdex.GetContract)
	if !ok {
		ob = nil
	}
	return ob, args.Error(1)
}
func (m *ServiceMock) SpotBuy(obr tdex.SpotBuyRequest) (*tdex.SpotBuy, error) {
	args := m.Called()
	ob, ok := args.Get(0).(*tdex.SpotBuy)
	if !ok {
		ob = nil
	}
	return ob, args.Error(1)
}
func (m *ServiceMock) SpotSell(obr tdex.SpotSellRequest) (*tdex.SpotSell, error) {
	args := m.Called()
	ob, ok := args.Get(0).(*tdex.SpotSell)
	if !ok {
		ob = nil
	}
	return ob, args.Error(1)
}
func (m *ServiceMock) SpotCancel(obr tdex.SpotCancelRequest) (*tdex.SpotCancel, error) {
	args := m.Called()
	ob, ok := args.Get(0).(*tdex.SpotCancel)
	if !ok {
		ob = nil
	}
	return ob, args.Error(1)
}
func (m *ServiceMock) SpotOrders() (*tdex.SpotOrders, error) {
	args := m.Called()
	return args.Get(0).(*tdex.SpotOrders), args.Error(1)
}
func (m *ServiceMock) SpotHistory(obr tdex.SpotHistoryRequest) (*tdex.SpotHistory, error) {
	args := m.Called()
	ob, ok := args.Get(0).(*tdex.SpotHistory)
	if !ok {
		ob = nil
	}
	return ob, args.Error(1)
}
func (m *ServiceMock) SpotStat(obr tdex.SpotStatRequest) (*tdex.SpotStat, error) {
	args := m.Called()
	ob, ok := args.Get(0).(*tdex.SpotStat)
	if !ok {
		ob = nil
	}
	return ob, args.Error(1)
}
func (m *ServiceMock) RobotClose(obr tdex.RobotCloseRequest) (*tdex.RobotClose, error) {
	args := m.Called()
	ob, ok := args.Get(0).(*tdex.RobotClose)
	if !ok {
		ob = nil
	}
	return ob, args.Error(1)
}
