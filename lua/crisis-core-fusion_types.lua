-- Typed models for the CrisisCoreFusion SDK (LuaLS annotations).
--
-- GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
-- params (op.<name>.points[].args.params[]). Field/param types come from the
-- canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
-- @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
-- edit by hand.

---@class Fusion
---@field materia1 string
---@field materia1_mastered boolean
---@field materia2 string
---@field materia2_mastered boolean
---@field result? table

---@class FusionCreateData

---@class Materia
---@field description? string
---@field id number
---@field max_level? number
---@field name string
---@field rarity? string
---@field type string

---@class MateriaLoadMatch
---@field id string

---@class MateriaListMatch

---@class System
---@field status? string

---@class SystemLoadMatch

local M = {}

return M
