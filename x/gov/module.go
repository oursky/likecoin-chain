package gov

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	originalgov "github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/cosmos/cosmos-sdk/x/gov/keeper"
	"github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/likecoin/likechain/x/gov/cli"
)

type AppModuleBasic struct {
	originalgov.AppModuleBasic
}

var (
	_ module.AppModule           = AppModule{}
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModuleSimulation = AppModule{}
)

func NewAppModuleBasic(proposalHandlers ...govclient.ProposalHandler) AppModuleBasic {
	return AppModuleBasic{
		AppModuleBasic: originalgov.NewAppModuleBasic(proposalHandlers...),
	}
}

func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

// TODO: RegisterRESTRoutes

// Go only allows "inheritence" from one embedded type
// If both types (in our case originalgov.AppModule and our AppModuleBasic) implements module.AppModuleBasic, then no
// methods will be "inherited" as Go requires developer to explicitly select one implementation (by implementing them).
// Therefore here we simply override one method instead of embedding AppModuleBasic (which actually doesn't reuse the
// code)
type AppModule struct {
	originalgov.AppModule
}

func (AppModule) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

// TODO: RegisterRESTRoutes

func NewAppModule(cdc codec.Codec, keeper keeper.Keeper, ak types.AccountKeeper, bk types.BankKeeper) AppModule {
	return AppModule{
		AppModule: originalgov.NewAppModule(cdc, keeper, ak, bk),
	}
}
