package e2e_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"

	sdk "github.com/cosmos/cosmos-sdk/types"
	nft "github.com/likecoin/likechain/backport/cosmos-sdk/v0.46.0-alpha2/x/nft"
	nftcli "github.com/likecoin/likechain/backport/cosmos-sdk/v0.46.0-alpha2/x/nft/client/cli"
	"github.com/likecoin/likechain/testutil/network"

	iscncli "github.com/likecoin/likechain/x/iscn/client/cli"
	iscntypes "github.com/likecoin/likechain/x/iscn/types"
	cli "github.com/likecoin/likechain/x/likenft/client/cli"
	types "github.com/likecoin/likechain/x/likenft/types"
	"github.com/stretchr/testify/require"
)

func TestEndToEndIscnBlindBox(t *testing.T) {
	tempDir := t.TempDir() // swap to ioutil for longlived files when debug
	cfg := network.DefaultConfig()

	// Override x/iscn gas fee denom to avoid the need to seed tokens

	// We do not have account addresses until network spawned
	// And do not want to spend time on seeding tester accounts
	// Fix later if using nanolike is a must for testing features

	iscnGenesis := iscntypes.GenesisState{}
	cfg.Codec.MustUnmarshalJSON(cfg.GenesisState[iscntypes.StoreKey], &iscnGenesis)
	iscnGenesis.Params.FeePerByte = sdk.NewDecCoin(
		cfg.BondDenom, sdk.NewInt(iscntypes.DefaultFeePerByteAmount),
	)
	cfg.GenesisState[iscntypes.StoreKey] = cfg.Codec.MustMarshalJSON(&iscnGenesis)

	// Setup network
	net := network.New(t, cfg)
	ctx := net.Validators[0].ClientCtx
	userAddress := net.Validators[0].Address
	txArgs := []string{
		fmt.Sprintf("--from=%s", userAddress.String()),
		"--yes",
		"--output=json",
		fmt.Sprintf("--gas-prices=%s", cfg.MinGasPrices),
		"--broadcast-mode=block",
	}
	queryArgs := []string{
		"--output=json",
	}

	// Seed input files
	createIscnFile, err := os.CreateTemp(tempDir, "create_iscn_*.json")
	require.NoError(t, err)
	require.NotNil(t, createIscnFile)
	_, err = createIscnFile.WriteString(`{
	"recordNotes": "Add IPFS fingerprint",
	"contentFingerprints": [
		"hash://sha256/9564b85669d5e96ac969dd0161b8475bbced9e5999c6ec598da718a3045d6f2e",
		"ipfs://QmNrgEMcUygbKzZeZgYFosdd27VE9KnWbyUD73bKZJ3bGi"
	],
	"stakeholders": [
		{
			"entity": {
				"@id": "did:cosmos:5sy29r37gfxvxz21rh4r0ktpuc46pzjrmz29g45",
				"name": "Chung Wu"
			},
			"rewardProportion": 95,
			"contributionType": "http://schema.org/author"
		},
		{
			"rewardProportion": 5,
			"contributionType": "http://schema.org/citation",
			"footprint": "https://en.wikipedia.org/wiki/Fibonacci_number",
			"description": "The blog post referred the matrix form of computing Fibonacci numbers."
		}
	],
	"contentMetadata": {
		"@context": "http://schema.org/",
		"@type": "Article",
		"name": "使用矩陣計算遞歸關係式",
		"description": "An article on computing recursive function with matrix multiplication.",
		"datePublished": "2019-04-19",
		"version": 1,
		"url": "https://nnkken.github.io/post/recursive-relation/",
		"author": "https://github.com/nnkken",
		"usageInfo": "https://creativecommons.org/licenses/by/4.0",
		"keywords": "matrix,recursion"
	}
}
`)
	require.NoError(t, err)
	err = createIscnFile.Close()
	require.NoError(t, err)

	newClassFile, err := os.CreateTemp(tempDir, "new_class_*.json")
	require.NoError(t, err)
	require.NotNil(t, newClassFile)
	_, err = newClassFile.WriteString(`{
	"name": "New Class",
	"symbol": "CLS",
	"description": "Testing New Class",
	"uri": "ipfs://aabbcc",
	"uri_hash": "aabbcc",
	"metadata": {
		"abc": "def"
	},
	"config": {
		"burnable": false
	}
}
`)
	require.NoError(t, err)
	err = newClassFile.Close()
	require.NoError(t, err)

	// Note use future reveal time so the network does not reveal the class
	// during test

	updateClassFile, err := os.CreateTemp(tempDir, "update_class_*.json")
	require.NoError(t, err)
	require.NotNil(t, updateClassFile)
	_, err = updateClassFile.WriteString(`{
	"name": "Oursky Cat Photos",
	"symbol": "Meowgear",
	"description": "Photos of our beloved bosses.",
	"uri": "https://nft.oursky.com/cats",
	"uri_hash": "",
	"metadata": {
		"name": "Oursky Cat Photos",
		"description": "Photos of our beloved bosses.",
		"image": "ipfs://QmZu3v5qFaTrrkSJC4mz8nLoDbR5kJx1QwMUy9CZhFZjT3",
		"external_link": "https://nft.oursky.com/cats"
	},
	"config": {
		"burnable": true,
		"max_supply": 10,
		"blind_box_config": {
			"mint_periods": [
				{
					"start_time": "2022-01-01T00:00:00Z",
					"allowed_addresses": [],
					"mint_price": 1000000000
				}
			],
			"reveal_time": "2050-01-01T00:00:00Z"
		}
	}
}
`)
	require.NoError(t, err)
	err = updateClassFile.Close()
	require.NoError(t, err)

	createMintable1File, err := os.CreateTemp(tempDir, "create_mintable_1_*.json")
	require.NoError(t, err)
	require.NotNil(t, createMintable1File)
	_, err = createMintable1File.WriteString(`{
	"uri": "https://example.com/1",
	"uri_hash": "111111",
	"metadata": {}
}
`)
	require.NoError(t, err)
	err = createMintable1File.Close()
	require.NoError(t, err)

	createMintable2File, err := os.CreateTemp(tempDir, "create_mintable_2_*.json")
	require.NoError(t, err)
	require.NotNil(t, createMintable2File)
	_, err = createMintable2File.WriteString(`{
		"uri": "https://example.com/2",
		"uri_hash": "222222",
		"metadata": {}
}
`)
	require.NoError(t, err)
	err = createMintable2File.Close()
	require.NoError(t, err)

	createMintable3File, err := os.CreateTemp(tempDir, "create_mintable_3_*.json")
	require.NoError(t, err)
	require.NotNil(t, createMintable3File)
	_, err = createMintable3File.WriteString(`{
		"uri": "https://example.com/3",
		"uri_hash": "333333",
		"metadata": {}
}
`)
	require.NoError(t, err)
	err = createMintable3File.Close()
	require.NoError(t, err)

	updateMintable1File, err := os.CreateTemp(tempDir, "update_mintable_1_*.json")
	require.NoError(t, err)
	require.NotNil(t, updateMintable1File)
	_, err = updateMintable1File.WriteString(`{
	"uri": "ipfs://QmYXq11iygTghZeyxvTZqpDoTomaX7Vd6Cbv1wuyNxq3Fw",
	"uri_hash": "QmYXq11iygTghZeyxvTZqpDoTomaX7Vd6Cbv1wuyNxq3Fw",
	"metadata": {
		"name": "Sleepy Coffee #1",
		"description": "Coffee is very sleepy", 
		"image": "ipfs://QmVhp6V2JdpYftT6LnDPELWCDMkk2aHwQZ1qbWf15KRbaZ",
		"external_url": "ipfs://QmYXq11iygTghZeyxvTZqpDoTomaX7Vd6Cbv1wuyNxq3Fw",
		"attributes": [
			{
				"trait_type": "Cat", 
				"value": "Coffee"
			}
		]
	}
}	
`)
	require.NoError(t, err)
	err = updateMintable1File.Close()
	require.NoError(t, err)

	updateMintable2File, err := os.CreateTemp(tempDir, "update_mintable_2_*.json")
	require.NoError(t, err)
	require.NotNil(t, updateMintable2File)
	_, err = updateMintable2File.WriteString(`{
	"uri": "ipfs://QmZLVbWsLP9EJW5pnGVRNV2eMqrvQ3wBnBWWJVkqDrCZjy",
	"uri_hash": "QmZLVbWsLP9EJW5pnGVRNV2eMqrvQ3wBnBWWJVkqDrCZjy",
	"metadata": {
		"name": "Hungry Chima #1",
		"description": "Chima is very hungry", 
		"image": "ipfs://QmTATCkXDu3u1KLZPuLrSg8RiyWsBdxuYHyT8HSifHcgzV",
		"external_url": "ipfs://QmZLVbWsLP9EJW5pnGVRNV2eMqrvQ3wBnBWWJVkqDrCZjy",
		"attributes": [
			{
				"trait_type": "Cat", 
				"value": "Chima"
			}
		]
	}
}
`)
	require.NoError(t, err)
	err = updateMintable2File.Close()
	require.NoError(t, err)

	updateMintable3File, err := os.CreateTemp(tempDir, "update_mintable_3_*.json")
	require.NoError(t, err)
	require.NotNil(t, updateMintable3File)
	_, err = updateMintable3File.WriteString(`{
	"uri": "ipfs://QmRg1QfbyfVsxvstH85z3FLzBYXEPdmRBYqg1eZwjKxMtK",
	"uri_hash": "QmRg1QfbyfVsxvstH85z3FLzBYXEPdmRBYqg1eZwjKxMtK",
	"metadata": {
		"name": "Smart Chima #2",
		"description": "Chima is very smart", 
		"image": "ipfs://QmPztx7RuugPHgP8yQPLSt6QhP5GZCxGqXf4KspNgffDRt",
		"external_url": "ipfs://QmRg1QfbyfVsxvstH85z3FLzBYXEPdmRBYqg1eZwjKxMtK",
		"attributes": [
			{
				"trait_type": "Cat", 
				"value": "Chima"
			}
		]
	}
}	
`)
	require.NoError(t, err)
	err = updateMintable3File.Close()
	require.NoError(t, err)

	// Create iscn
	out, err := clitestutil.ExecTestCLICmd(
		ctx,
		iscncli.NewCreateIscnTxCmd(),
		append([]string{createIscnFile.Name()}, txArgs...),
	)
	require.NoError(t, err)

	// Get iscn id prefix created
	res := sdk.TxResponse{}
	cfg.Codec.MustUnmarshalJSON(out.Bytes(), &res)
	var iscnIdPrefix string
FindIscnIdPrefix:
	for _, log := range res.Logs {
		for _, event := range log.Events {
			if event.Type == "iscn_record" {
				for _, attr := range event.Attributes {
					if attr.Key == "iscn_id_prefix" {
						iscnIdPrefix = attr.Value
						break FindIscnIdPrefix
					}
				}
			}
		}
	}
	require.NotEmpty(t, iscnIdPrefix)

	// Create class
	out, err = clitestutil.ExecTestCLICmd(
		ctx,
		cli.CmdNewClass(),
		append([]string{fmt.Sprintf("--iscnIdPrefix=%s", iscnIdPrefix), newClassFile.Name()}, txArgs...),
	)
	require.NoError(t, err)

	// Validate event emitted
	res = sdk.TxResponse{}
	cfg.Codec.MustUnmarshalJSON(out.Bytes(), &res)
	actualCreateEvent := parseEventCreateClass(res)
	require.NotEmpty(t, actualCreateEvent.ClassId)
	require.Equal(t, iscnIdPrefix, actualCreateEvent.ParentIscnIdPrefix)
	require.Empty(t, actualCreateEvent.ParentAccount)

	// Query class
	out, err = clitestutil.ExecTestCLICmd(
		ctx,
		cli.CmdShowClassesByISCN(),
		append([]string{iscnIdPrefix}, queryArgs...),
	)
	require.NoError(t, err)

	// Unmarshal and check class data
	classesRes := types.QueryClassesByISCNResponse{}
	cfg.Codec.MustUnmarshalJSON(out.Bytes(), &classesRes)

	require.Len(t, classesRes.Classes, 1)
	class := classesRes.Classes[0]
	require.Equal(t, "New Class", class.Name)
	require.Equal(t, "CLS", class.Symbol)
	require.Equal(t, "Testing New Class", class.Description)
	require.Equal(t, "ipfs://aabbcc", class.Uri)
	require.Equal(t, "aabbcc", class.UriHash)
	classData := types.ClassData{}
	err = classData.Unmarshal(class.Data.Value)
	require.NoError(t, err)
	expectedMetadata, err := types.JsonInput(`{
	"abc": "def"
}`).Normalize()
	require.NoError(t, err)
	actualMetadata, err := classData.Metadata.Normalize()
	require.NoError(t, err)
	require.Equal(t, expectedMetadata, actualMetadata)
	require.Equal(t, types.ClassConfig{
		Burnable: false,
	}, classData.Config)
	require.Equal(t, types.ClassParent{
		Type:              types.ClassParentType_ISCN,
		IscnIdPrefix:      iscnIdPrefix,
		IscnVersionAtMint: 1,
	}, classData.Parent)

	// Update class
	out, err = clitestutil.ExecTestCLICmd(
		ctx,
		cli.CmdUpdateClass(),
		append([]string{class.Id, updateClassFile.Name()}, txArgs...),
	)
	require.NoError(t, err)

	// Validate event emitted
	res = sdk.TxResponse{}
	cfg.Codec.MustUnmarshalJSON(out.Bytes(), &res)
	actualUpdateEvent := parseEventUpdateClass(res)
	require.Equal(t, types.EventUpdateClass{
		ClassId:            class.Id,
		ParentIscnIdPrefix: iscnIdPrefix,
		ParentAccount:      "",
	}, actualUpdateEvent)

	// Query updated class
	out, err = clitestutil.ExecTestCLICmd(
		ctx,
		cli.CmdShowClassesByISCN(),
		append([]string{iscnIdPrefix}, queryArgs...),
	)
	require.NoError(t, err)

	// Unmarshal and check updated class data
	classesRes = types.QueryClassesByISCNResponse{}
	cfg.Codec.MustUnmarshalJSON(out.Bytes(), &classesRes)

	require.Len(t, classesRes.Classes, 1)
	class = classesRes.Classes[0]
	require.Equal(t, "Oursky Cat Photos", class.Name)
	require.Equal(t, "Meowgear", class.Symbol)
	require.Equal(t, "Photos of our beloved bosses.", class.Description)
	require.Equal(t, "https://nft.oursky.com/cats", class.Uri)
	require.Equal(t, "", class.UriHash)
	classData = types.ClassData{}
	err = classData.Unmarshal(class.Data.Value)
	require.NoError(t, err)
	expectedMetadata, err = types.JsonInput(`{
	"name": "Oursky Cat Photos",
	"description": "Photos of our beloved bosses.",
	"image": "ipfs://QmZu3v5qFaTrrkSJC4mz8nLoDbR5kJx1QwMUy9CZhFZjT3",
	"external_link": "https://nft.oursky.com/cats"
}`).Normalize()
	require.NoError(t, err)
	actualMetadata, err = classData.Metadata.Normalize()
	require.NoError(t, err)
	require.Equal(t, expectedMetadata, actualMetadata)
	require.Equal(t, types.ClassConfig{
		Burnable:  true,
		MaxSupply: uint64(10),
		BlindBoxConfig: &types.BlindBoxConfig{
			MintPeriods: []types.MintPeriod{
				{
					StartTime:        time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
					AllowedAddresses: nil,
					MintPrice:        1000000000,
				},
			},
			RevealTime: time.Date(2050, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}, classData.Config)
	require.Equal(t, types.ClassParent{
		Type:              types.ClassParentType_ISCN,
		IscnIdPrefix:      iscnIdPrefix,
		IscnVersionAtMint: 1,
	}, classData.Parent)

	// Create mintable 1
	out, err = clitestutil.ExecTestCLICmd(
		ctx,
		cli.CmdCreateMintableNFT(),
		append([]string{class.Id, "mintable1", createMintable1File.Name()}, txArgs...),
	)
	require.NoError(t, err)

	// Validate event emitted
	res = sdk.TxResponse{}
	cfg.Codec.MustUnmarshalJSON(out.Bytes(), &res)
	createMintable1Event := parseEventCreateMintableNFT(res)
	require.Equal(t, types.EventCreateMintableNFT{
		ClassId:                 class.Id,
		MintableNftId:           "mintable1",
		ClassParentIscnIdPrefix: iscnIdPrefix,
		ClassParentAccount:      "",
	}, createMintable1Event)

	// Create mintable 2
	out, err = clitestutil.ExecTestCLICmd(
		ctx,
		cli.CmdCreateMintableNFT(),
		append([]string{class.Id, "mintable2", createMintable2File.Name()}, txArgs...),
	)
	require.NoError(t, err)

	// Validate event emitted
	res = sdk.TxResponse{}
	cfg.Codec.MustUnmarshalJSON(out.Bytes(), &res)
	createMintable2Event := parseEventCreateMintableNFT(res)
	require.Equal(t, types.EventCreateMintableNFT{
		ClassId:                 class.Id,
		MintableNftId:           "mintable2",
		ClassParentIscnIdPrefix: iscnIdPrefix,
		ClassParentAccount:      "",
	}, createMintable2Event)

	// Create mintable 3
	out, err = clitestutil.ExecTestCLICmd(
		ctx,
		cli.CmdCreateMintableNFT(),
		append([]string{class.Id, "mintable3", createMintable3File.Name()}, txArgs...),
	)
	require.NoError(t, err)

	// Validate event emitted
	res = sdk.TxResponse{}
	cfg.Codec.MustUnmarshalJSON(out.Bytes(), &res)
	createMintable3Event := parseEventCreateMintableNFT(res)
	require.Equal(t, types.EventCreateMintableNFT{
		ClassId:                 class.Id,
		MintableNftId:           "mintable3",
		ClassParentIscnIdPrefix: iscnIdPrefix,
		ClassParentAccount:      "",
	}, createMintable3Event)

	// Query created mintables
	out, err = clitestutil.ExecTestCLICmd(
		ctx,
		cli.CmdMintableNFTs(),
		append([]string{class.Id}, queryArgs...),
	)
	require.NoError(t, err)

	// Unmarshal and check created mintables
	createdMintablesRes := types.QueryMintableNFTsResponse{}
	cfg.Codec.MustUnmarshalJSON(out.Bytes(), &createdMintablesRes)
	require.Equal(t, []types.MintableNFT{
		{
			ClassId: class.Id,
			Id:      "mintable1",
			Input: types.NFTInput{
				Uri:      "https://example.com/1",
				UriHash:  "111111",
				Metadata: types.JsonInput(`{}`),
			},
		},
		{
			ClassId: class.Id,
			Id:      "mintable2",
			Input: types.NFTInput{
				Uri:      "https://example.com/2",
				UriHash:  "222222",
				Metadata: types.JsonInput(`{}`),
			},
		},
		{
			ClassId: class.Id,
			Id:      "mintable3",
			Input: types.NFTInput{
				Uri:      "https://example.com/3",
				UriHash:  "333333",
				Metadata: types.JsonInput(`{}`),
			},
		},
	}, createdMintablesRes.MintableNfts)

	// Update mintable 1
	out, err = clitestutil.ExecTestCLICmd(
		ctx,
		cli.CmdUpdateMintableNFT(),
		append([]string{class.Id, "mintable1", updateMintable1File.Name()}, txArgs...),
	)
	require.NoError(t, err)

	// Validate event emitted
	res = sdk.TxResponse{}
	cfg.Codec.MustUnmarshalJSON(out.Bytes(), &res)
	updateMintable1Event := parseEventUpdateMintableNFT(res)
	require.Equal(t, types.EventUpdateMintableNFT{
		ClassId:                 class.Id,
		MintableNftId:           "mintable1",
		ClassParentIscnIdPrefix: iscnIdPrefix,
		ClassParentAccount:      "",
	}, updateMintable1Event)

	// Update mintable 2
	out, err = clitestutil.ExecTestCLICmd(
		ctx,
		cli.CmdUpdateMintableNFT(),
		append([]string{class.Id, "mintable2", updateMintable2File.Name()}, txArgs...),
	)
	require.NoError(t, err)

	// Validate event emitted
	res = sdk.TxResponse{}
	cfg.Codec.MustUnmarshalJSON(out.Bytes(), &res)
	updateMintable2Event := parseEventUpdateMintableNFT(res)
	require.Equal(t, types.EventUpdateMintableNFT{
		ClassId:                 class.Id,
		MintableNftId:           "mintable2",
		ClassParentIscnIdPrefix: iscnIdPrefix,
		ClassParentAccount:      "",
	}, updateMintable2Event)

	// Delete mintable 3
	out, err = clitestutil.ExecTestCLICmd(
		ctx,
		cli.CmdDeleteMintableNFT(),
		append([]string{class.Id, "mintable3"}, txArgs...),
	)
	require.NoError(t, err)

	// Validate event emitted
	res = sdk.TxResponse{}
	cfg.Codec.MustUnmarshalJSON(out.Bytes(), &res)
	deleteMintable3Event := parseEventDeleteMintableNFT(res)
	require.Equal(t, types.EventDeleteMintableNFT{
		ClassId:                 class.Id,
		MintableNftId:           "mintable3",
		ClassParentIscnIdPrefix: iscnIdPrefix,
		ClassParentAccount:      "",
	}, deleteMintable3Event)

	// Recreate mintable 3
	out, err = clitestutil.ExecTestCLICmd(
		ctx,
		cli.CmdCreateMintableNFT(),
		append([]string{class.Id, "mintable3", updateMintable3File.Name()}, txArgs...),
	)
	require.NoError(t, err)

	// Validate event emitted
	res = sdk.TxResponse{}
	cfg.Codec.MustUnmarshalJSON(out.Bytes(), &res)
	recreateMintable3Event := parseEventCreateMintableNFT(res)
	require.Equal(t, types.EventCreateMintableNFT{
		ClassId:                 class.Id,
		MintableNftId:           "mintable3",
		ClassParentIscnIdPrefix: iscnIdPrefix,
		ClassParentAccount:      "",
	}, recreateMintable3Event)

	// Query created mintables
	out, err = clitestutil.ExecTestCLICmd(
		ctx,
		cli.CmdMintableNFTs(),
		append([]string{class.Id}, queryArgs...),
	)
	require.NoError(t, err)

	// Unmarshal and check updated mintables
	updatedMintablesRes := types.QueryMintableNFTsResponse{}
	cfg.Codec.MustUnmarshalJSON(out.Bytes(), &updatedMintablesRes)
	require.Equal(t, []types.MintableNFT{
		{
			ClassId: class.Id,
			Id:      "mintable1",
			Input: types.NFTInput{
				Uri:      "ipfs://QmYXq11iygTghZeyxvTZqpDoTomaX7Vd6Cbv1wuyNxq3Fw",
				UriHash:  "QmYXq11iygTghZeyxvTZqpDoTomaX7Vd6Cbv1wuyNxq3Fw",
				Metadata: types.JsonInput(`{"name":"Sleepy Coffee #1","description":"Coffee is very sleepy","image":"ipfs://QmVhp6V2JdpYftT6LnDPELWCDMkk2aHwQZ1qbWf15KRbaZ","external_url":"ipfs://QmYXq11iygTghZeyxvTZqpDoTomaX7Vd6Cbv1wuyNxq3Fw","attributes":[{"trait_type":"Cat","value":"Coffee"}]}`),
			},
		},
		{
			ClassId: class.Id,
			Id:      "mintable2",
			Input: types.NFTInput{
				Uri:      "ipfs://QmZLVbWsLP9EJW5pnGVRNV2eMqrvQ3wBnBWWJVkqDrCZjy",
				UriHash:  "QmZLVbWsLP9EJW5pnGVRNV2eMqrvQ3wBnBWWJVkqDrCZjy",
				Metadata: types.JsonInput(`{"name":"Hungry Chima #1","description":"Chima is very hungry","image":"ipfs://QmTATCkXDu3u1KLZPuLrSg8RiyWsBdxuYHyT8HSifHcgzV","external_url":"ipfs://QmZLVbWsLP9EJW5pnGVRNV2eMqrvQ3wBnBWWJVkqDrCZjy","attributes":[{"trait_type":"Cat","value":"Chima"}]}`),
			},
		},
		{
			ClassId: class.Id,
			Id:      "mintable3",
			Input: types.NFTInput{
				Uri:      "ipfs://QmRg1QfbyfVsxvstH85z3FLzBYXEPdmRBYqg1eZwjKxMtK",
				UriHash:  "QmRg1QfbyfVsxvstH85z3FLzBYXEPdmRBYqg1eZwjKxMtK",
				Metadata: types.JsonInput(`{"name":"Smart Chima #2","description":"Chima is very smart","image":"ipfs://QmPztx7RuugPHgP8yQPLSt6QhP5GZCxGqXf4KspNgffDRt","external_url":"ipfs://QmRg1QfbyfVsxvstH85z3FLzBYXEPdmRBYqg1eZwjKxMtK","attributes":[{"trait_type":"Cat","value":"Chima"}]}`),
			},
		},
	}, updatedMintablesRes.MintableNfts)

	// Mint blind nft
	// Note non-owner mint covered by unit test, FIXME improve by seeding another
	// account

	out, err = clitestutil.ExecTestCLICmd(
		ctx,
		cli.CmdMintNFT(),
		append([]string{class.Id}, txArgs...),
	)

	fmt.Printf(out.String())

	// Validate event emitted
	res = sdk.TxResponse{}
	cfg.Codec.MustUnmarshalJSON(out.Bytes(), &res)
	actualMintEvent := parseEventMintNFT(res)
	require.Equal(t, types.EventMintNFT{
		ClassId:                 class.Id,
		NftId:                   "nft1",
		Owner:                   userAddress.String(),
		ClassParentIscnIdPrefix: iscnIdPrefix,
		ClassParentAccount:      "",
	}, actualMintEvent)

	// Query NFT
	out, err = clitestutil.ExecTestCLICmd(
		ctx,
		nftcli.GetCmdQueryNFT(),
		append([]string{class.Id, "nft1"}, queryArgs...),
	)
	require.NoError(t, err)
	nftRes := nft.QueryNFTResponse{}
	cfg.Codec.MustUnmarshalJSON(out.Bytes(), &nftRes)
	require.Equal(t, class.Id, nftRes.Nft.ClassId)
	require.Equal(t, "nft1", nftRes.Nft.Id)
	require.Empty(t, nftRes.Nft.Uri)
	require.Empty(t, nftRes.Nft.UriHash)
	nftData := types.NFTData{}
	err = nftData.Unmarshal(nftRes.Nft.Data.Value)
	require.NoError(t, err)
	require.Equal(t, types.NFTData{
		Metadata: types.JsonInput(`null`),
		ClassParent: types.ClassParent{
			Type:              types.ClassParentType_ISCN,
			IscnIdPrefix:      iscnIdPrefix,
			IscnVersionAtMint: 1,
		},
		ToBeRevealed: true,
	}, nftData)

	// Note: reveal queue & logic tested separately; hard to test with this setup
	// as we cannot override block header time, & don't have access to keepers

	// Burn NFT
	out, err = clitestutil.ExecTestCLICmd(
		ctx,
		cli.CmdBurnNFT(),
		append([]string{class.Id, "nft1"}, txArgs...),
	)
	require.NoError(t, err)

	// Validate event emitted
	res = sdk.TxResponse{}
	cfg.Codec.MustUnmarshalJSON(out.Bytes(), &res)
	actualBurnEvent := parseEventBurnNFT(res)
	require.Equal(t, types.EventBurnNFT{
		ClassId:                 class.Id,
		NftId:                   "nft1",
		Owner:                   userAddress.String(),
		ClassParentIscnIdPrefix: iscnIdPrefix,
	}, actualBurnEvent)

	// Check NFT is burnt
	_, err = clitestutil.ExecTestCLICmd(
		ctx,
		nftcli.GetCmdQueryNFT(),
		append([]string{class.Id, "nft1"}, queryArgs...),
	)
	require.Error(t, err)
	require.Contains(t, err.Error(), "not found")
}