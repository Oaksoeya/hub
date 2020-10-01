package querier

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func querySubscription(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QuerySubscriptionParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, errors.Wrap(types.ErrorUnmarshal, err.Error())
	}

	plan, found := k.GetSubscription(ctx, params.ID)
	if !found {
		return nil, nil
	}

	res, err := types.ModuleCdc.MarshalJSON(plan)
	if err != nil {
		return nil, errors.Wrap(types.ErrorMarshal, err.Error())
	}

	return res, nil
}

func querySubscriptions(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QuerySubscriptionsParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, errors.Wrap(types.ErrorUnmarshal, err.Error())
	}

	subscriptions := k.GetSubscriptions(ctx)

	start, end := client.Paginate(len(subscriptions), params.Page, params.Limit, len(subscriptions))
	if start < 0 || end < 0 {
		subscriptions = types.Subscriptions{}
	} else {
		subscriptions = subscriptions[start:end]
	}

	res, err := types.ModuleCdc.MarshalJSON(subscriptions)
	if err != nil {
		return nil, errors.Wrap(types.ErrorMarshal, err.Error())
	}

	return res, nil
}

func querySubscriptionsForAddress(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QuerySubscriptionsForAddressParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, errors.Wrap(types.ErrorUnmarshal, err.Error())
	}

	subscriptions := k.GetSubscriptionsForAddress(ctx, params.Address)

	start, end := client.Paginate(len(subscriptions), params.Page, params.Limit, len(subscriptions))
	if start < 0 || end < 0 {
		subscriptions = types.Subscriptions{}
	} else {
		subscriptions = subscriptions[start:end]
	}

	res, err := types.ModuleCdc.MarshalJSON(subscriptions)
	if err != nil {
		return nil, errors.Wrap(types.ErrorMarshal, err.Error())
	}

	return res, nil
}

func querySubscriptionsForPlan(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QuerySubscriptionsForPlanParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, errors.Wrap(types.ErrorUnmarshal, err.Error())
	}

	subscriptions := k.GetSubscriptionsForPlan(ctx, params.ID)

	start, end := client.Paginate(len(subscriptions), params.Page, params.Limit, len(subscriptions))
	if start < 0 || end < 0 {
		subscriptions = types.Subscriptions{}
	} else {
		subscriptions = subscriptions[start:end]
	}

	res, err := types.ModuleCdc.MarshalJSON(subscriptions)
	if err != nil {
		return nil, errors.Wrap(types.ErrorMarshal, err.Error())
	}

	return res, nil
}

func querySubscriptionsForNode(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QuerySubscriptionsForNodeParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, errors.Wrap(types.ErrorUnmarshal, err.Error())
	}

	subscriptions := k.GetSubscriptionsForNode(ctx, params.Address)

	start, end := client.Paginate(len(subscriptions), params.Page, params.Limit, len(subscriptions))
	if start < 0 || end < 0 {
		subscriptions = types.Subscriptions{}
	} else {
		subscriptions = subscriptions[start:end]
	}

	res, err := types.ModuleCdc.MarshalJSON(subscriptions)
	if err != nil {
		return nil, errors.Wrap(types.ErrorMarshal, err.Error())
	}

	return res, nil
}
