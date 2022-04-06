package keeper_test

import (
	"testing"

	testkeeper "github.com/likecoin/likechain/testutil/keeper"
	"github.com/likecoin/likechain/x/likenft/types"
	"github.com/stretchr/testify/require"
)

func TestParamsValidate(t *testing.T) {
	var err error
	var params types.Params

	params = types.DefaultParams()
	err = params.Validate()
	require.NoError(t, err)

	params = types.Params{
		MintPriceDenom: "nanolike",
	}
	err = params.Validate()
	require.NoError(t, err)

	params = types.Params{
		MintPriceDenom: "",
	}
	err = params.Validate()
	require.Error(t, err, "should not accept empty mint price denom")

	params = types.Params{
		MintPriceDenom: "nanolike123!!!??",
	}
	err = params.Validate()
	require.Error(t, err, "should not accept mint price denom with invalid characters")

	params = types.Params{}
	err = params.Validate()
	require.Error(t, err, "should not accept empty params")
}

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LikenftKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
