package poa

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"

	modulev1 "cosmossdk.io/api/cosmos/poa/module/v1"
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/proto/tendermint/crypto"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
)

// ============================================================================
//
// type AppModuleBasic struct
//
// ============================================================================
type AppModuleBasic struct{}

// ----------------------------------------------------------------------------
// AppModuleBasic: module.AppModuleBasic
// ----------------------------------------------------------------------------
var _ module.AppModuleBasic = AppModuleBasic{}

func (AppModuleBasic) Name() string { return ModuleName }

func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *gwruntime.ServeMux) {
	// TODO(zrbecker): Register GRPC Gateway Routes
}

func (AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// TODO(zrbecker): Register Interfaces
}

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
}

// ----------------------------------------------------------------------------
// AppModuleBasic: module.HasGenesisBasics
// ----------------------------------------------------------------------------
var _ module.HasGenesisBasics = AppModuleBasic{}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	// TODO(zrbecker): Setup Default Genesis State
	return []byte("{}")
}

func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	return nil
}

// ============================================================================
//
// type AppModule struct
//
// ============================================================================
type AppModule struct {
	AppModuleBasic
}

func NewAppModule() AppModule {
	return AppModule{}
}

// ----------------------------------------------------------------------------
// AppModule: appmodule.AppModule
// ----------------------------------------------------------------------------
var _ appmodule.AppModule = AppModule{}

func (AppModule) IsAppModule()        {}
func (AppModule) IsOnePerModuleType() {}

// ----------------------------------------------------------------------------
// AppModule: module.AppModuleSimulation
// ----------------------------------------------------------------------------
var _ module.AppModuleSimulation = AppModule{}

func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	// TODO(zrbecker): implement module.AppModuleSimulation
}

func (AppModule) RegisterStoreDecoder(sdr simtypes.StoreDecoderRegistry) {
	// TODO(zrbecker): implement module.AppModuleSimulation
}
func (AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	// TODO(zrbecker): implement module.AppModuleSimulation
	return nil
}

// ----------------------------------------------------------------------------
// AppModule: module.HasServices
// ----------------------------------------------------------------------------
var _ module.HasServices = AppModule{}

func (AppModule) RegisterServices(cfg module.Configurator) {
	// TODO(zrbecker): implement module.HasServices
}

// ----------------------------------------------------------------------------
// AppModule: module.HasABCIGenesis
// ----------------------------------------------------------------------------
var _ module.HasABCIGenesis = AppModule{}

func (AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	// TODO(zrbecker): implement module.HasABCIGenesis
	return []byte("{}")
}

func (AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	fmt.Println("HELLO INIT GENESIS")
	// TODO(zrbecker): implement module.HasABCIGenesis
	hardcodedValidatorPublicKey, err := hex.DecodeString("28FEB0E78F6B3B909C2C72CE02DAEA42C9C509BA7EDD3A8C5821012BFB22CC84")
	if err != nil {
		panic("failed to decode key")
	}
	return []abci.ValidatorUpdate{
		{PubKey: crypto.PublicKey{Sum: &crypto.PublicKey_Ed25519{
			Ed25519: hardcodedValidatorPublicKey,
		}}, Power: 10000000},
	}
}

// ----------------------------------------------------------------------------
// AppModule: module.HasABCIEndBlock
// ----------------------------------------------------------------------------
var _ module.HasABCIEndBlock = AppModule{}

func (AppModule) EndBlock(ctx context.Context) ([]abci.ValidatorUpdate, error) {
	// TODO(zrbecker): implement module.HasABCIEndBlock
	return nil, nil
}

// ----------------------------------------------------------------------------
// AppModule: module.HasBeginBlocker
// ----------------------------------------------------------------------------
var _ appmodule.HasBeginBlocker = AppModule{}

func (am AppModule) BeginBlock(ctx context.Context) error {
	fmt.Println("Hello Authority!!!")
	return nil
}

// ============================================================================
//
// depinject
//
// ============================================================================
type ModuleInputs struct {
	depinject.In

	Config *modulev1.Module
}

type ModuleOutputs struct {
	depinject.Out

	Module appmodule.AppModule
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	m := NewAppModule()

	return ModuleOutputs{Module: m}
}

func init() {
	appmodule.Register(&modulev1.Module{}, appmodule.Provide(ProvideModule))
}
