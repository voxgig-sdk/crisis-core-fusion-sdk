<?php
declare(strict_types=1);

// CrisisCoreFusion SDK utility: result_headers

class CrisisCoreFusionResultHeaders
{
    public static function call(CrisisCoreFusionContext $ctx): ?CrisisCoreFusionResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
