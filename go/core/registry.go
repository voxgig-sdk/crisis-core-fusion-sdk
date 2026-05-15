package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewFusionEntityFunc func(client *CrisisCoreFusionSDK, entopts map[string]any) CrisisCoreFusionEntity

var NewMateriaEntityFunc func(client *CrisisCoreFusionSDK, entopts map[string]any) CrisisCoreFusionEntity

var NewSystemEntityFunc func(client *CrisisCoreFusionSDK, entopts map[string]any) CrisisCoreFusionEntity

