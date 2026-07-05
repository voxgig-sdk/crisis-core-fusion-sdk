# Typed models for the CrisisCoreFusion SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class FusionRequired(TypedDict):
    materia1: str
    materia1_mastered: bool
    materia2: str
    materia2_mastered: bool


class Fusion(FusionRequired, total=False):
    result: dict


class FusionCreateDataRequired(TypedDict):
    materia1: str
    materia1_mastered: bool
    materia2: str
    materia2_mastered: bool


class FusionCreateData(FusionCreateDataRequired, total=False):
    result: dict


class MateriaRequired(TypedDict):
    id: int
    name: str
    type: str


class Materia(MateriaRequired, total=False):
    description: str
    max_level: int
    rarity: str


class MateriaLoadMatch(TypedDict):
    id: str


class MateriaListMatch(TypedDict, total=False):
    description: str
    id: int
    max_level: int
    name: str
    rarity: str
    type: str


class System(TypedDict, total=False):
    status: str


class SystemLoadMatch(TypedDict, total=False):
    status: str
