package voxgigcrisiscorefusionsdk

import (
	"github.com/voxgig-sdk/crisis-core-fusion-sdk/go/core"
	"github.com/voxgig-sdk/crisis-core-fusion-sdk/go/entity"
	"github.com/voxgig-sdk/crisis-core-fusion-sdk/go/feature"
	_ "github.com/voxgig-sdk/crisis-core-fusion-sdk/go/utility"
)

// Type aliases preserve external API.
type CrisisCoreFusionSDK = core.CrisisCoreFusionSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type CrisisCoreFusionEntity = core.CrisisCoreFusionEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type CrisisCoreFusionError = core.CrisisCoreFusionError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewFusionEntityFunc = func(client *core.CrisisCoreFusionSDK, entopts map[string]any) core.CrisisCoreFusionEntity {
		return entity.NewFusionEntity(client, entopts)
	}
	core.NewMateriaEntityFunc = func(client *core.CrisisCoreFusionSDK, entopts map[string]any) core.CrisisCoreFusionEntity {
		return entity.NewMateriaEntity(client, entopts)
	}
	core.NewSystemEntityFunc = func(client *core.CrisisCoreFusionSDK, entopts map[string]any) core.CrisisCoreFusionEntity {
		return entity.NewSystemEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewCrisisCoreFusionSDK = core.NewCrisisCoreFusionSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewCrisisCoreFusionSDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *CrisisCoreFusionSDK  { return NewCrisisCoreFusionSDK(nil) }
func Test() *CrisisCoreFusionSDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
