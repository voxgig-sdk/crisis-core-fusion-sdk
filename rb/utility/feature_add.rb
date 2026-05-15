# CrisisCoreFusion SDK utility: feature_add
module CrisisCoreFusionUtilities
  FeatureAdd = ->(ctx, f) {
    ctx.client.features << f
  }
end
