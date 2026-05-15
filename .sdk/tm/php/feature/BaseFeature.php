<?php
declare(strict_types=1);

// CrisisCoreFusion SDK base feature

class CrisisCoreFusionBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(CrisisCoreFusionContext $ctx, array $options): void {}
    public function PostConstruct(CrisisCoreFusionContext $ctx): void {}
    public function PostConstructEntity(CrisisCoreFusionContext $ctx): void {}
    public function SetData(CrisisCoreFusionContext $ctx): void {}
    public function GetData(CrisisCoreFusionContext $ctx): void {}
    public function GetMatch(CrisisCoreFusionContext $ctx): void {}
    public function SetMatch(CrisisCoreFusionContext $ctx): void {}
    public function PrePoint(CrisisCoreFusionContext $ctx): void {}
    public function PreSpec(CrisisCoreFusionContext $ctx): void {}
    public function PreRequest(CrisisCoreFusionContext $ctx): void {}
    public function PreResponse(CrisisCoreFusionContext $ctx): void {}
    public function PreResult(CrisisCoreFusionContext $ctx): void {}
    public function PreDone(CrisisCoreFusionContext $ctx): void {}
    public function PreUnexpected(CrisisCoreFusionContext $ctx): void {}
}
