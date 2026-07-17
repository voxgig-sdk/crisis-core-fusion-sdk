-- CrisisCoreFusion SDK exists test

local sdk = require("crisis-core-fusion_sdk")

describe("CrisisCoreFusionSDK", function()
  it("should create test SDK", function()
    local testsdk = sdk.test(nil, nil)
    assert.is_not_nil(testsdk)
  end)
end)
