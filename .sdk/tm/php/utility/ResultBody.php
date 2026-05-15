<?php
declare(strict_types=1);

// CrisisCoreFusion SDK utility: result_body

class CrisisCoreFusionResultBody
{
    public static function call(CrisisCoreFusionContext $ctx): ?CrisisCoreFusionResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
