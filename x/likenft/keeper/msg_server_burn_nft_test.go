package keeper_test

import (
	"testing"

	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/likecoin/likechain/backport/cosmos-sdk/v0.46.0-alpha2/x/nft"
	"github.com/likecoin/likechain/testutil/keeper"
	iscntypes "github.com/likecoin/likechain/x/iscn/types"
	"github.com/likecoin/likechain/x/likenft/testutil"
	"github.com/likecoin/likechain/x/likenft/types"
	"github.com/stretchr/testify/require"
)

func TestBurnNFTNormal(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	accountKeeper := testutil.NewMockAccountKeeper(ctrl)
	bankKeeper := testutil.NewMockBankKeeper(ctrl)
	iscnKeeper := testutil.NewMockIscnKeeper(ctrl)
	nftKeeper := testutil.NewMockNftKeeper(ctrl)
	msgServer, goCtx, _ := setupMsgServer(t, keeper.LikenftDependedKeepers{
		AccountKeeper: accountKeeper,
		BankKeeper:    bankKeeper,
		IscnKeeper:    iscnKeeper,
		NftKeeper:     nftKeeper,
	})

	// Test Input
	ownerAddressBytes := []byte{0, 1, 0, 1, 0, 1, 0, 1}
	ownerAddress, _ := sdk.Bech32ifyAddressBytes("cosmos", ownerAddressBytes)
	classId := "likenft1aabbccddeeff"
	tokenId := "token1"
	iscnId := iscntypes.NewIscnId("likecoin-chain", "abcdef", 1)

	// Mock keeper calls
	nftKeeper.
		EXPECT().
		HasNFT(gomock.Any(), gomock.Eq(classId), gomock.Eq(tokenId)).
		Return(true)

	nftKeeper.
		EXPECT().
		GetOwner(gomock.Any(), gomock.Eq(classId), gomock.Eq(tokenId)).
		Return(ownerAddressBytes)

	classData := types.ClassData{
		Metadata:     types.JsonInput(`{"aaaa": "bbbb"}`),
		IscnIdPrefix: iscnId.Prefix.String(),
		Config: types.ClassConfig{
			Burnable: true,
		},
	}
	classDataInAny, _ := cdctypes.NewAnyWithValue(&classData)
	nftKeeper.
		EXPECT().
		GetClass(gomock.Any(), gomock.Eq(classId)).
		Return(nft.Class{
			Id:          classId,
			Name:        "Class Name",
			Symbol:      "ABC",
			Description: "Testing Class 123",
			Uri:         "ipfs://abcdef",
			UriHash:     "abcdef",
			Data:        classDataInAny,
		}, true)

	nftKeeper.
		EXPECT().
		Burn(gomock.Any(), gomock.Eq(classId), gomock.Eq(tokenId)).
		Return(nil)

	// Run
	res, err := msgServer.BurnNFT(goCtx, &types.MsgBurnNFT{
		Creator: ownerAddress,
		ClassID: classId,
		NftID:   tokenId,
	})

	// Check output
	require.NoError(t, err)
	require.Equal(t, &types.MsgBurnNFTResponse{}, res)

	// Check mock was called as expected
	ctrl.Finish()
}

func TestBurnNFTTokenNotFound(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	accountKeeper := testutil.NewMockAccountKeeper(ctrl)
	bankKeeper := testutil.NewMockBankKeeper(ctrl)
	iscnKeeper := testutil.NewMockIscnKeeper(ctrl)
	nftKeeper := testutil.NewMockNftKeeper(ctrl)
	msgServer, goCtx, _ := setupMsgServer(t, keeper.LikenftDependedKeepers{
		AccountKeeper: accountKeeper,
		BankKeeper:    bankKeeper,
		IscnKeeper:    iscnKeeper,
		NftKeeper:     nftKeeper,
	})

	// Test Input
	ownerAddressBytes := []byte{0, 1, 0, 1, 0, 1, 0, 1}
	ownerAddress, _ := sdk.Bech32ifyAddressBytes("cosmos", ownerAddressBytes)
	classId := "likenft1aabbccddeeff"
	tokenId := "token1"

	// Mock keeper calls
	nftKeeper.
		EXPECT().
		HasNFT(gomock.Any(), gomock.Eq(classId), gomock.Eq(tokenId)).
		Return(false)

	// Run
	res, err := msgServer.BurnNFT(goCtx, &types.MsgBurnNFT{
		Creator: ownerAddress,
		ClassID: classId,
		NftID:   tokenId,
	})

	// Check output
	require.Error(t, err)
	require.Contains(t, err.Error(), types.ErrNftNotFound.Error())
	require.Nil(t, res)

	// Check mock was called as expected
	ctrl.Finish()
}
