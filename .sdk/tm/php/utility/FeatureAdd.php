<?php
declare(strict_types=1);

// CrisisCoreFusion SDK utility: feature_add

class CrisisCoreFusionFeatureAdd
{
    public static function call(CrisisCoreFusionContext $ctx, mixed $f): void
    {
        $ctx->client->features[] = $f;
    }
}
