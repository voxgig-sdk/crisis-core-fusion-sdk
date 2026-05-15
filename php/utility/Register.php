<?php
declare(strict_types=1);

// CrisisCoreFusion SDK utility registration

require_once __DIR__ . '/../core/UtilityType.php';
require_once __DIR__ . '/Clean.php';
require_once __DIR__ . '/Done.php';
require_once __DIR__ . '/MakeError.php';
require_once __DIR__ . '/FeatureAdd.php';
require_once __DIR__ . '/FeatureHook.php';
require_once __DIR__ . '/FeatureInit.php';
require_once __DIR__ . '/Fetcher.php';
require_once __DIR__ . '/MakeFetchDef.php';
require_once __DIR__ . '/MakeContext.php';
require_once __DIR__ . '/MakeOptions.php';
require_once __DIR__ . '/MakeRequest.php';
require_once __DIR__ . '/MakeResponse.php';
require_once __DIR__ . '/MakeResult.php';
require_once __DIR__ . '/MakePoint.php';
require_once __DIR__ . '/MakeSpec.php';
require_once __DIR__ . '/MakeUrl.php';
require_once __DIR__ . '/Param.php';
require_once __DIR__ . '/PrepareAuth.php';
require_once __DIR__ . '/PrepareBody.php';
require_once __DIR__ . '/PrepareHeaders.php';
require_once __DIR__ . '/PrepareMethod.php';
require_once __DIR__ . '/PrepareParams.php';
require_once __DIR__ . '/PreparePath.php';
require_once __DIR__ . '/PrepareQuery.php';
require_once __DIR__ . '/ResultBasic.php';
require_once __DIR__ . '/ResultBody.php';
require_once __DIR__ . '/ResultHeaders.php';
require_once __DIR__ . '/TransformRequest.php';
require_once __DIR__ . '/TransformResponse.php';

CrisisCoreFusionUtility::setRegistrar(function (CrisisCoreFusionUtility $u): void {
    $u->clean = [CrisisCoreFusionClean::class, 'call'];
    $u->done = [CrisisCoreFusionDone::class, 'call'];
    $u->make_error = [CrisisCoreFusionMakeError::class, 'call'];
    $u->feature_add = [CrisisCoreFusionFeatureAdd::class, 'call'];
    $u->feature_hook = [CrisisCoreFusionFeatureHook::class, 'call'];
    $u->feature_init = [CrisisCoreFusionFeatureInit::class, 'call'];
    $u->fetcher = [CrisisCoreFusionFetcher::class, 'call'];
    $u->make_fetch_def = [CrisisCoreFusionMakeFetchDef::class, 'call'];
    $u->make_context = [CrisisCoreFusionMakeContext::class, 'call'];
    $u->make_options = [CrisisCoreFusionMakeOptions::class, 'call'];
    $u->make_request = [CrisisCoreFusionMakeRequest::class, 'call'];
    $u->make_response = [CrisisCoreFusionMakeResponse::class, 'call'];
    $u->make_result = [CrisisCoreFusionMakeResult::class, 'call'];
    $u->make_point = [CrisisCoreFusionMakePoint::class, 'call'];
    $u->make_spec = [CrisisCoreFusionMakeSpec::class, 'call'];
    $u->make_url = [CrisisCoreFusionMakeUrl::class, 'call'];
    $u->param = [CrisisCoreFusionParam::class, 'call'];
    $u->prepare_auth = [CrisisCoreFusionPrepareAuth::class, 'call'];
    $u->prepare_body = [CrisisCoreFusionPrepareBody::class, 'call'];
    $u->prepare_headers = [CrisisCoreFusionPrepareHeaders::class, 'call'];
    $u->prepare_method = [CrisisCoreFusionPrepareMethod::class, 'call'];
    $u->prepare_params = [CrisisCoreFusionPrepareParams::class, 'call'];
    $u->prepare_path = [CrisisCoreFusionPreparePath::class, 'call'];
    $u->prepare_query = [CrisisCoreFusionPrepareQuery::class, 'call'];
    $u->result_basic = [CrisisCoreFusionResultBasic::class, 'call'];
    $u->result_body = [CrisisCoreFusionResultBody::class, 'call'];
    $u->result_headers = [CrisisCoreFusionResultHeaders::class, 'call'];
    $u->transform_request = [CrisisCoreFusionTransformRequest::class, 'call'];
    $u->transform_response = [CrisisCoreFusionTransformResponse::class, 'call'];
});
