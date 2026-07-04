# Typed models for the CrisisCoreFusion SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.

from __future__ import annotations

from dataclasses import dataclass
from typing import Optional, Any


@dataclass
class Fusion:
    materia1: str
    materia1_mastered: bool
    materia2: str
    materia2_mastered: bool
    result: Optional[dict] = None


@dataclass
class FusionCreateData:
    materia1: Optional[str] = None
    materia1_mastered: Optional[bool] = None
    materia2: Optional[str] = None
    materia2_mastered: Optional[bool] = None
    result: Optional[dict] = None


@dataclass
class Materia:
    id: int
    name: str
    type: str
    description: Optional[str] = None
    max_level: Optional[int] = None
    rarity: Optional[str] = None


@dataclass
class MateriaLoadMatch:
    id: str


@dataclass
class MateriaListMatch:
    description: Optional[str] = None
    id: Optional[int] = None
    max_level: Optional[int] = None
    name: Optional[str] = None
    rarity: Optional[str] = None
    type: Optional[str] = None


@dataclass
class System:
    status: Optional[str] = None


@dataclass
class SystemLoadMatch:
    status: Optional[str] = None

