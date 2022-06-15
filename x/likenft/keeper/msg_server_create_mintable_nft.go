package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/likecoin/likechain/x/likenft/types"
)

func (k msgServer) CreateMintableNFT(goCtx context.Context, msg *types.MsgCreateMintableNFT) (*types.MsgCreateMintableNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	class, classData, err := k.GetClass(ctx, msg.ClassId)
	if err != nil {
		return nil, err
	}
	parent, err := k.ValidateAndRefreshClassParent(ctx, msg.ClassId, classData.Parent)
	if err != nil {
		return nil, err
	}
	if err := k.validateReqToMutateMintableNFT(ctx, msg.Creator, class, classData, parent, true); err != nil {
		return nil, err
	}

	// check id not already exist
	if _, exists := k.GetMintableNFT(ctx, msg.ClassId, msg.Id); exists {
		return nil, types.ErrMintableNftAlreadyExists
	}

	mintableNFT := types.MintableNFT{
		ClassId: msg.ClassId,
		Id:      msg.Id,
		Input:   msg.Input,
	}

	// Deduct minting fee
	userAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf(err.Error())
	}
	err = k.DeductFeeForMintingNFT(ctx, userAddress, mintableNFT.Size())
	if err != nil {
		return nil, err
	}

	// set record
	k.SetMintableNFT(ctx, mintableNFT)

	// Emit event
	ctx.EventManager().EmitTypedEvent(&types.EventCreateMintableNFT{
		ClassId:                 msg.ClassId,
		MintableNftId:           msg.Id,
		ClassParentIscnIdPrefix: parent.IscnIdPrefix,
		ClassParentAccount:      parent.Account,
	})

	return &types.MsgCreateMintableNFTResponse{
		MintableNft: mintableNFT,
	}, nil
}
