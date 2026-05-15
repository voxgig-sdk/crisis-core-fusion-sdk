# CrisisCoreFusion SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module CrisisCoreFusionFeatures
  def self.make_feature(name)
    case name
    when "base"
      CrisisCoreFusionBaseFeature.new
    when "test"
      CrisisCoreFusionTestFeature.new
    else
      CrisisCoreFusionBaseFeature.new
    end
  end
end
