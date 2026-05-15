# CrisisCoreFusion SDK utility: make_context
require_relative '../core/context'
module CrisisCoreFusionUtilities
  MakeContext = ->(ctxmap, basectx) {
    CrisisCoreFusionContext.new(ctxmap, basectx)
  }
end
