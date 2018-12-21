package sentinel_hub

import (
	"encoding/json"
	"os"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/common"
	tmDb "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	tmTypes "github.com/tendermint/tendermint/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/hub"
	"github.com/ironman0x7b2/sentinel-sdk/x/ibc"
)

const appName = "Sentinel Hub"

var (
	DefaultCLIHome  = os.ExpandEnv("$HOME/.sentinel-hub-cli")
	DefaultNodeHome = os.ExpandEnv("$HOME/.sentinel-hubd")
)

type SentinelHub struct {
	*baseapp.BaseApp
	cdc *codec.Codec

	keyMain       *csdkTypes.KVStoreKey
	keyAccount    *csdkTypes.KVStoreKey
	keyIBC        *csdkTypes.KVStoreKey
	keyCoinLocker *csdkTypes.KVStoreKey

	accountKeeper       auth.AccountKeeper
	feeCollectionKeeper auth.FeeCollectionKeeper
	bankKeeper          bank.Keeper
	ibcKeeper           ibc.Keeper
	hubKeeper           hub.BaseKeeper
}

func NewSentinelHub(logger log.Logger, db tmDb.DB, baseAppOptions ...func(*baseapp.BaseApp)) *SentinelHub {
	cdc := MakeCodec()

	var app = &SentinelHub{
		cdc:           cdc,
		BaseApp:       baseapp.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc), baseAppOptions...),
		keyMain:       csdkTypes.NewKVStoreKey(sdkTypes.KeyMain),
		keyAccount:    csdkTypes.NewKVStoreKey(sdkTypes.KeyAccount),
		keyIBC:        csdkTypes.NewKVStoreKey(sdkTypes.KeyIBC),
		keyCoinLocker: csdkTypes.NewKVStoreKey(sdkTypes.KeyCoinLocker),
	}

	app.accountKeeper = auth.NewAccountKeeper(
		cdc,
		app.keyAccount,
		func() auth.Account {
			return &sdkTypes.AppAccount{}
		},
	)
	app.bankKeeper = bank.NewBaseKeeper(app.accountKeeper)
	app.ibcKeeper = ibc.NewKeeper(app.keyIBC, app.cdc)
	app.hubKeeper = hub.NewBaseKeeper(app.cdc, app.keyCoinLocker, app.bankKeeper)

	app.Router().
		AddRoute(sdkTypes.KeyBank, bank.NewHandler(app.bankKeeper)).
		AddRoute(sdkTypes.KeyIBC, hub.NewIBCHubHandler(app.ibcKeeper, app.hubKeeper))

	app.SetInitChainer(app.initChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)
	app.SetAnteHandler(auth.NewAnteHandler(app.accountKeeper, app.feeCollectionKeeper))

	app.MountStores(app.keyMain, app.keyAccount, app.keyIBC, app.keyCoinLocker)
	err := app.LoadLatestVersion(app.keyMain)
	if err != nil {
		common.Exit(err.Error())
	}

	app.Seal()

	return app
}

func MakeCodec() *codec.Codec {
	cdc := codec.New()

	codec.RegisterCrypto(cdc)
	csdkTypes.RegisterCodec(cdc)
	bank.RegisterCodec(cdc)
	auth.RegisterCodec(cdc)
	sdkTypes.RegisterCodec(cdc)
	ibc.RegisterCodec(cdc)
	hub.RegisterCodec(cdc)

	cdc.Seal()

	return cdc
}

func (app *SentinelHub) BeginBlocker(_ csdkTypes.Context, _ abciTypes.RequestBeginBlock) abciTypes.ResponseBeginBlock {
	return abciTypes.ResponseBeginBlock{}
}

func (app *SentinelHub) EndBlocker(_ csdkTypes.Context, _ abciTypes.RequestEndBlock) abciTypes.ResponseEndBlock {
	return abciTypes.ResponseEndBlock{}
}

func (app *SentinelHub) initChainer(ctx csdkTypes.Context, req abciTypes.RequestInitChain) abciTypes.ResponseInitChain {
	stateJSON := req.AppStateBytes

	genesisState := new(sdkTypes.GenesisState)
	err := app.cdc.UnmarshalJSON(stateJSON, genesisState)
	if err != nil {
		// TODO: https://github.com/cosmos/cosmos-sdk/issues/468
		panic(err)
	}

	for _, gacc := range genesisState.Accounts {
		acc, err := gacc.ToAppAccount()
		if err != nil {
			// TODO: https://github.com/cosmos/cosmos-sdk/issues/468
			panic(err)
		}

		acc.AccountNumber = app.accountKeeper.GetNextAccountNumber(ctx)
		app.accountKeeper.SetAccount(ctx, acc)
	}

	return abciTypes.ResponseInitChain{}
}

func (app *SentinelHub) ExportAppStateAndValidators() (appState json.RawMessage, validators []tmTypes.GenesisValidator, err error) {
	ctx := app.NewContext(true, abciTypes.Header{})
	var accounts []*sdkTypes.GenesisAccount

	appendAccountsFn := func(acc auth.Account) bool {
		account := &sdkTypes.GenesisAccount{
			Address: acc.GetAddress(),
			Coins:   acc.GetCoins(),
		}

		accounts = append(accounts, account)
		return false
	}

	app.accountKeeper.IterateAccounts(ctx, appendAccountsFn)

	genState := sdkTypes.GenesisState{Accounts: accounts}
	appState, err = codec.MarshalJSONIndent(app.cdc, genState)
	if err != nil {
		return nil, nil, err
	}

	return appState, validators, err
}