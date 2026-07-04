<?php
declare(strict_types=1);

// Typed models for the CrisisCoreFusion SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** Fusion entity data model. */
class Fusion
{
    public string $materia1;
    public bool $materia1_mastered;
    public string $materia2;
    public bool $materia2_mastered;
    public ?array $result = null;
}

/** Match filter for Fusion#create (any subset of Fusion fields). */
class FusionCreateData
{
    public ?string $materia1 = null;
    public ?bool $materia1_mastered = null;
    public ?string $materia2 = null;
    public ?bool $materia2_mastered = null;
    public ?array $result = null;
}

/** Materia entity data model. */
class Materia
{
    public ?string $description = null;
    public int $id;
    public ?int $max_level = null;
    public string $name;
    public ?string $rarity = null;
    public string $type;
}

/** Request payload for Materia#load. */
class MateriaLoadMatch
{
    public string $id;
}

/** Match filter for Materia#list (any subset of Materia fields). */
class MateriaListMatch
{
    public ?string $description = null;
    public ?int $id = null;
    public ?int $max_level = null;
    public ?string $name = null;
    public ?string $rarity = null;
    public ?string $type = null;
}

/** System entity data model. */
class System
{
    public ?string $status = null;
}

/** Match filter for System#load (any subset of System fields). */
class SystemLoadMatch
{
    public ?string $status = null;
}

