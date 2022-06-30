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

func TestMintNFTNormal(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	accountKeeper := testutil.NewMockAccountKeeper(ctrl)
	bankKeeper := testutil.NewMockBankKeeper(ctrl)
	iscnKeeper := testutil.NewMockIscnKeeper(ctrl)
	nftKeeper := testutil.NewMockNftKeeper(ctrl)
	msgServer, goCtx, keeper := setupMsgServer(t, keeper.LikenftDependedKeepers{
		AccountKeeper: accountKeeper,
		BankKeeper:    bankKeeper,
		IscnKeeper:    iscnKeeper,
		NftKeeper:     nftKeeper,
	})
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Test Input
	ownerAddressBytes := []byte{0, 1, 0, 1, 0, 1, 0, 1}
	ownerAddress, _ := sdk.Bech32ifyAddressBytes("cosmos", ownerAddressBytes)
	iscnId := iscntypes.NewIscnId("likecoin-chain", "abcdef", 1)
	classId := "likenft1aabbccddeeff"
	tokenId := "token1"
	uri := "ipfs://a1b2c3"
	uriHash := "a1b2c3"
	metadata := types.JsonInput(
		`{
	"abc": "def",
	"qwerty": 1234,
	"bool": false,
	"null": null,
	"nested": {
		"object": {
			"abc": "def"
		}
	}
}`)

	// Mock keeper calls
	classData := types.ClassData{
		Metadata:     types.JsonInput(`{"aaaa": "bbbb"}`),
		IscnIdPrefix: iscnId.Prefix.String(),
		Config: types.ClassConfig{
			Burnable: false,
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

	keeper.SetClassesByISCN(ctx, types.ClassesByISCN{
		IscnIdPrefix: iscnId.Prefix.String(),
		ClassIds:     []string{classId},
	})

	iscnKeeper.
		EXPECT().
		GetContentIdRecord(gomock.Any(), gomock.Eq(iscnId.Prefix)).
		Return(&iscntypes.ContentIdRecord{
			OwnerAddressBytes: ownerAddressBytes,
			LatestVersion:     1,
		})

	wrappedOwnerAddress, _ := sdk.AccAddressFromBech32(ownerAddress)
	nftKeeper.
		EXPECT().
		Mint(gomock.Any(), gomock.Any(), wrappedOwnerAddress)

	// Run
	res, err := msgServer.MintNFT(goCtx, &types.MsgMintNFT{
		Creator:  ownerAddress,
		ClassId:  classId,
		Id:       tokenId,
		Uri:      uri,
		UriHash:  uriHash,
		Metadata: metadata,
	})

	// Check output
	require.NoError(t, err)
	require.Equal(t, classId, res.Nft.ClassId)
	require.Equal(t, uri, res.Nft.Uri)
	require.Equal(t, uriHash, res.Nft.UriHash)

	var nftData types.NFTData
	err = nftData.Unmarshal(res.Nft.Data.Value)
	require.NoErrorf(t, err, "Error unmarshal class data")
	require.Equal(t, metadata, nftData.Metadata)
	require.Equal(t, iscnId.Prefix.String(), nftData.IscnIdPrefix)

	// Check mock was called as expected
	ctrl.Finish()
}

func TestMintNFTInvalidTokenID(t *testing.T) {
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
	uri := "ipfs://a1b2c3"
	uriHash := "a1b2c3"
	metadata := types.JsonInput(
		`{
	"abc": "def",
	"qwerty": 1234,
	"bool": false,
	"null": null,
	"nested": {
		"object": {
			"abc": "def"
		}
	}
}`)

	// Run
	res, err := msgServer.MintNFT(goCtx, &types.MsgMintNFT{
		Creator:  ownerAddress,
		ClassId:  classId,
		Id:       "123456", // x/nft requires token id to start with letter
		Uri:      uri,
		UriHash:  uriHash,
		Metadata: metadata,
	})

	// Check output
	require.Error(t, err)
	require.Contains(t, err.Error(), types.ErrInvalidTokenId.Error())
	require.Nil(t, res)

	// Check mock was called as expected
	ctrl.Finish()
}
