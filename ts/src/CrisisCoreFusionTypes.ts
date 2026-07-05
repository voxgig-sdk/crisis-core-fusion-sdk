// Typed models for the CrisisCoreFusion SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface Fusion {
  materia1: string
  materia1_mastered: boolean
  materia2: string
  materia2_mastered: boolean
  result?: Record<string, any>
}

export interface FusionCreateData {
  materia1: string
  materia1_mastered: boolean
  materia2: string
  materia2_mastered: boolean
  result?: Record<string, any>
}

export interface Materia {
  description?: string
  id: number
  max_level?: number
  name: string
  rarity?: string
  type: string
}

export interface MateriaLoadMatch {
  id: string
}

export interface MateriaListMatch {
  description?: string
  id?: number
  max_level?: number
  name?: string
  rarity?: string
  type?: string
}

export interface System {
  status?: string
}

export interface SystemLoadMatch {
  status?: string
}

