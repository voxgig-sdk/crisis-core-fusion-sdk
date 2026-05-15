# CrisisCoreFusion SDK exists test

require "minitest/autorun"
require_relative "../CrisisCoreFusion_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = CrisisCoreFusionSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end
