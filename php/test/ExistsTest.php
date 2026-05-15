<?php
declare(strict_types=1);

// CrisisCoreFusion SDK exists test

require_once __DIR__ . '/../crisiscorefusion_sdk.php';

use PHPUnit\Framework\TestCase;

class ExistsTest extends TestCase
{
    public function test_create_test_sdk(): void
    {
        $testsdk = CrisisCoreFusionSDK::test(null, null);
        $this->assertNotNull($testsdk);
    }
}
