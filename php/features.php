<?php
declare(strict_types=1);

// CrisisCoreFusion SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class CrisisCoreFusionFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new CrisisCoreFusionBaseFeature();
            case "test":
                return new CrisisCoreFusionTestFeature();
            default:
                return new CrisisCoreFusionBaseFeature();
        }
    }
}
