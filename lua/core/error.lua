-- CrisisCoreFusion SDK error

local CrisisCoreFusionError = {}
CrisisCoreFusionError.__index = CrisisCoreFusionError


function CrisisCoreFusionError.new(code, msg, ctx)
  local self = setmetatable({}, CrisisCoreFusionError)
  self.is_sdk_error = true
  self.sdk = "CrisisCoreFusion"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function CrisisCoreFusionError:error()
  return self.msg
end


function CrisisCoreFusionError:__tostring()
  return self.msg
end


return CrisisCoreFusionError
