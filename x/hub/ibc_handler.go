package hub

import (
	"fmt"
	"reflect"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/ibc"
)

func NewIBCHubHandler(ibcKeeper ibc.Keeper, hubKeeper Keeper) csdkTypes.Handler {
	return func(ctx csdkTypes.Context, msg csdkTypes.Msg) csdkTypes.Result {
		switch msg := msg.(type) {
		case ibc.MsgIBCTransaction:
			switch ibcMsg := msg.IBCPacket.Message.(type) {
			case MsgLockCoins:
				return handleLockCoins(ctx, ibcKeeper, hubKeeper, msg)
			case MsgReleaseCoins:
				return handleReleaseCoins(ctx, ibcKeeper, hubKeeper, msg)
			case MsgReleaseCoinsToMany:
				return handleReleaseCoinsToMany(ctx, ibcKeeper, hubKeeper, msg)
			default:
				errMsg := fmt.Sprintf("Unrecognized IBC msg: %v", reflect.TypeOf(ibcMsg).Name())
				return csdkTypes.ErrUnknownRequest(errMsg).Result()
			}
		default:
			errMsg := fmt.Sprintf("Unrecognized msg type: %v", reflect.TypeOf(msg).Name())
			return csdkTypes.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleLockCoins(ctx csdkTypes.Context, ibcKeeper ibc.Keeper, hubKeeper Keeper, ibcMsg ibc.MsgIBCTransaction) csdkTypes.Result {
	msg, _ := ibcMsg.IBCPacket.Message.(MsgLockCoins)

	if err := msg.ValidateBasic(); err != nil {
		return err.Result()
	}

	sequence, err := ibcKeeper.GetIngressLength(ctx, sdkTypes.IngressLengthKey(ibcMsg.IBCPacket.SrcChainID))

	if err != nil {
		return err.Result()
	}

	if ibcMsg.Sequence != sequence {
		return errorInvalidIBCSequence().Result()
	}

	lockerID := ibcMsg.IBCPacket.SrcChainID + "/" + msg.LockerID
	address := msg.PubKey.Address().Bytes()
	locker, err := hubKeeper.GetLocker(ctx, lockerID)

	if err != nil {
		return err.Result()
	}

	if locker != nil {
		return errorLockerAlreadyExists().Result()
	}

	if err := hubKeeper.LockCoins(ctx, lockerID, address, msg.Coins); err != nil {
		return err.Result()
	}

	packet := sdkTypes.IBCPacket{
		SrcChainID:  ibcMsg.IBCPacket.DestChainID,
		DestChainID: ibcMsg.IBCPacket.SrcChainID,
		Message: MsgLockerStatus{
			LockerID: msg.LockerID,
			Status:   sdkTypes.StatusLock,
		},
	}

	if err := ibcKeeper.PostIBCPacket(ctx, packet); err != nil {
		return err.Result()
	}

	if err := ibcKeeper.SetIngressLength(ctx, sdkTypes.IngressLengthKey(ibcMsg.IBCPacket.SrcChainID), sequence+1); err != nil {
		return err.Result()
	}

	return csdkTypes.Result{}
}

func handleReleaseCoins(ctx csdkTypes.Context, ibcKeeper ibc.Keeper, hubKeeper Keeper, ibcMsg ibc.MsgIBCTransaction) csdkTypes.Result {
	msg, _ := ibcMsg.IBCPacket.Message.(MsgReleaseCoins)

	if err := msg.ValidateBasic(); err != nil {
		return err.Result()
	}

	sequence, err := ibcKeeper.GetIngressLength(ctx, sdkTypes.IngressLengthKey(ibcMsg.IBCPacket.SrcChainID))

	if err != nil {
		return err.Result()
	}

	if ibcMsg.Sequence != sequence {
		return errorInvalidIBCSequence().Result()
	}

	lockerID := ibcMsg.IBCPacket.SrcChainID + "/" + msg.LockerID
	address := msg.PubKey.Address().Bytes()
	locker, err := hubKeeper.GetLocker(ctx, lockerID)

	if err != nil {
		return err.Result()
	}

	if locker == nil {
		return errorLockerNotExists().Result()
	}

	if !locker.Address.Equals(address) {
		return errorInvalidLockerOwnerAddress().Result()
	}

	if locker.Status != sdkTypes.StatusLock {
		return errorInvalidLockerStatus().Result()
	}

	if err := hubKeeper.ReleaseCoins(ctx, lockerID); err != nil {
		return err.Result()
	}

	packet := sdkTypes.IBCPacket{
		SrcChainID:  ibcMsg.IBCPacket.DestChainID,
		DestChainID: ibcMsg.IBCPacket.SrcChainID,
		Message: MsgLockerStatus{
			LockerID: msg.LockerID,
			Status:   sdkTypes.StatusRelease,
		},
	}

	if err := ibcKeeper.PostIBCPacket(ctx, packet); err != nil {
		return err.Result()
	}

	if err := ibcKeeper.SetIngressLength(ctx, sdkTypes.IngressLengthKey(ibcMsg.IBCPacket.SrcChainID), sequence+1); err != nil {
		return err.Result()
	}

	return csdkTypes.Result{}
}

func handleReleaseCoinsToMany(ctx csdkTypes.Context, ibcKeeper ibc.Keeper, hubKeeper Keeper, ibcMsg ibc.MsgIBCTransaction) csdkTypes.Result {
	msg, _ := ibcMsg.IBCPacket.Message.(MsgReleaseCoinsToMany)

	if err := msg.ValidateBasic(); err != nil {
		return err.Result()
	}

	sequence, err := ibcKeeper.GetIngressLength(ctx, sdkTypes.IngressLengthKey(ibcMsg.IBCPacket.SrcChainID))

	if err != nil {
		return err.Result()
	}

	if ibcMsg.Sequence != sequence {
		return errorInvalidIBCSequence().Result()
	}

	lockerID := ibcMsg.IBCPacket.SrcChainID + "/" + msg.LockerID
	address := msg.PubKey.Address().Bytes()
	locker, err := hubKeeper.GetLocker(ctx, lockerID)

	if err != nil {
		return err.Result()
	}

	if locker == nil {
		return errorLockerNotExists().Result()
	}

	if !locker.Address.Equals(address) {
		return errorInvalidLockerOwnerAddress().Result()
	}

	if locker.Status != sdkTypes.StatusLock {
		return errorInvalidLockerStatus().Result()
	}

	if err := hubKeeper.ReleaseCoinsToMany(ctx, lockerID, msg.Addresses, msg.Shares); err != nil {
		return err.Result()
	}

	packet := sdkTypes.IBCPacket{
		SrcChainID:  ibcMsg.IBCPacket.DestChainID,
		DestChainID: ibcMsg.IBCPacket.SrcChainID,
		Message: MsgLockerStatus{
			LockerID: msg.LockerID,
			Status:   sdkTypes.StatusRelease,
		},
	}

	if err := ibcKeeper.PostIBCPacket(ctx, packet); err != nil {
		return err.Result()
	}

	if err := ibcKeeper.SetIngressLength(ctx, sdkTypes.IngressLengthKey(ibcMsg.IBCPacket.SrcChainID), sequence+1); err != nil {
		return err.Result()
	}

	return csdkTypes.Result{}
}