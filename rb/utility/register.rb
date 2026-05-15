# CrisisCoreFusion SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

CrisisCoreFusionUtility.registrar = ->(u) {
  u.clean = CrisisCoreFusionUtilities::Clean
  u.done = CrisisCoreFusionUtilities::Done
  u.make_error = CrisisCoreFusionUtilities::MakeError
  u.feature_add = CrisisCoreFusionUtilities::FeatureAdd
  u.feature_hook = CrisisCoreFusionUtilities::FeatureHook
  u.feature_init = CrisisCoreFusionUtilities::FeatureInit
  u.fetcher = CrisisCoreFusionUtilities::Fetcher
  u.make_fetch_def = CrisisCoreFusionUtilities::MakeFetchDef
  u.make_context = CrisisCoreFusionUtilities::MakeContext
  u.make_options = CrisisCoreFusionUtilities::MakeOptions
  u.make_request = CrisisCoreFusionUtilities::MakeRequest
  u.make_response = CrisisCoreFusionUtilities::MakeResponse
  u.make_result = CrisisCoreFusionUtilities::MakeResult
  u.make_point = CrisisCoreFusionUtilities::MakePoint
  u.make_spec = CrisisCoreFusionUtilities::MakeSpec
  u.make_url = CrisisCoreFusionUtilities::MakeUrl
  u.param = CrisisCoreFusionUtilities::Param
  u.prepare_auth = CrisisCoreFusionUtilities::PrepareAuth
  u.prepare_body = CrisisCoreFusionUtilities::PrepareBody
  u.prepare_headers = CrisisCoreFusionUtilities::PrepareHeaders
  u.prepare_method = CrisisCoreFusionUtilities::PrepareMethod
  u.prepare_params = CrisisCoreFusionUtilities::PrepareParams
  u.prepare_path = CrisisCoreFusionUtilities::PreparePath
  u.prepare_query = CrisisCoreFusionUtilities::PrepareQuery
  u.result_basic = CrisisCoreFusionUtilities::ResultBasic
  u.result_body = CrisisCoreFusionUtilities::ResultBody
  u.result_headers = CrisisCoreFusionUtilities::ResultHeaders
  u.transform_request = CrisisCoreFusionUtilities::TransformRequest
  u.transform_response = CrisisCoreFusionUtilities::TransformResponse
}
