package vpn

import (
	"reflect"
	"strings"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/hub"
	"github.com/ironman0x7b2/sentinel-sdk/x/ibc"
)

func NewIBCVPNHandler(ibcKeeper ibc.Keeper, vpnKeeper Keeper) csdkTypes.Handler {
	return func(ctx csdkTypes.Context, msg csdkTypes.Msg) csdkTypes.Result {
		switch msg := msg.(type) {
		case ibc.MsgIBCTransaction:
			switch ibcMsg := msg.IBCPacket.Message.(type) {
			case hub.MsgLockerStatus:
				newMsg, _ := msg.IBCPacket.Message.(hub.MsgLockerStatus)
				if strings.HasPrefix(newMsg.LockerID, vpnKeeper.VPNStoreKey.Name()+"/") {
					return handleUpdateNodeStatus(ctx, ibcKeeper, vpnKeeper, msg)
				} else if strings.HasPrefix(newMsg.LockerID, vpnKeeper.SessionStoreKey.Name()+"/") {
					return handleUpdateSessionStatus(ctx, ibcKeeper, vpnKeeper, msg)
				} else {
					errMsg := "Unrecognized locker id: " + newMsg.LockerID
					return csdkTypes.ErrUnknownRequest(errMsg).Result()
				}
			default:
				errMsg := "Unrecognized IBC msg type: " + reflect.TypeOf(ibcMsg).Name()
				return csdkTypes.ErrUnknownRequest(errMsg).Result()
			}
		default:
			errMsg := "Unrecognized msg type: " + reflect.TypeOf(msg).Name()
			return csdkTypes.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleUpdateNodeStatus(ctx csdkTypes.Context, ibcKeeper ibc.Keeper, vpnKeeper Keeper, ibcMsg ibc.MsgIBCTransaction) csdkTypes.Result {
	msg, _ := ibcMsg.IBCPacket.Message.(hub.MsgLockerStatus)
	sequence, err := ibcKeeper.GetIngressLength(ctx, sdkTypes.IngressLengthKey(ibcMsg.IBCPacket.SrcChainID))

	if err != nil {
		return err.Result()
	}

	if ibcMsg.Sequence != sequence {
		return errorInvalidIBCSequence().Result()
	}

	nodeID := msg.LockerID[len(vpnKeeper.VPNStoreKey.Name())+1:]

	if vpnDetails, err := vpnKeeper.GetVPNDetails(ctx, nodeID); true {
		if err != nil {
			return err.Result()
		}

		if vpnDetails == nil {
			return errorVPNNotExists().Result()
		}
	}

	switch msg.Status {
	case sdkTypes.StatusLock:
		if err := vpnKeeper.SetVPNStatus(ctx, nodeID, sdkTypes.StatusActive); err != nil {
			return err.Result()
		}

		if err := vpnKeeper.AddActiveNodeID(ctx, nodeID); err != nil {
			return err.Result()
		}
	case sdkTypes.StatusRelease:
		if err := vpnKeeper.SetVPNStatus(ctx, nodeID, sdkTypes.StatusDeregister); err != nil {
			return err.Result()
		}
	default:
		return errorInvalidLockStatus().Result()
	}

	if err := ibcKeeper.SetIngressLength(ctx, sdkTypes.IngressLengthKey(ibcMsg.IBCPacket.SrcChainID), sequence+1); err != nil {
		return err.Result()
	}

	return csdkTypes.Result{}
}

func handleUpdateSessionStatus(ctx csdkTypes.Context, ibcKeeper ibc.Keeper, vpnKeeper Keeper, ibcMsg ibc.MsgIBCTransaction) csdkTypes.Result {
	msg, _ := ibcMsg.IBCPacket.Message.(hub.MsgLockerStatus)
	sequence, err := ibcKeeper.GetIngressLength(ctx, sdkTypes.IngressLengthKey(ibcMsg.IBCPacket.SrcChainID))

	if err != nil {
		return err.Result()
	}

	if ibcMsg.Sequence != sequence {
		return errorInvalidIBCSequence().Result()
	}

	sessionID := msg.LockerID[len(vpnKeeper.SessionStoreKey.Name())+1:]

	if sessionDetails, err := vpnKeeper.GetSessionDetails(ctx, sessionID); true {
		if err != nil {
			return err.Result()
		}

		if sessionDetails == nil {
			return errorSessionNotExists().Result()
		}
	}

	switch msg.Status {
	case sdkTypes.StatusLock:
		if err := vpnKeeper.SetSessionStatus(ctx, sessionID, sdkTypes.StatusActive); err != nil {
			return err.Result()
		}

		if err := vpnKeeper.AddActiveSessionID(ctx, sessionID); err != nil {
			return err.Result()
		}
	case sdkTypes.StatusRelease:
		if err := vpnKeeper.SetVPNStatus(ctx, sessionID, sdkTypes.StatusEnd); err != nil {
			return err.Result()
		}
	default:
		return errorInvalidLockStatus().Result()
	}

	if err := ibcKeeper.SetIngressLength(ctx, sdkTypes.IngressLengthKey(ibcMsg.IBCPacket.SrcChainID), sequence+1); err != nil {
		return err.Result()
	}

	return csdkTypes.Result{}
}