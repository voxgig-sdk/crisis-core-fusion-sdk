# frozen_string_literal: true

# Typed models for the CrisisCoreFusion SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# Fusion entity data model.
#
# @!attribute [rw] materia1
#   @return [String]
#
# @!attribute [rw] materia1_mastered
#   @return [Boolean]
#
# @!attribute [rw] materia2
#   @return [String]
#
# @!attribute [rw] materia2_mastered
#   @return [Boolean]
#
# @!attribute [rw] result
#   @return [Hash, nil]
Fusion = Struct.new(
  :materia1,
  :materia1_mastered,
  :materia2,
  :materia2_mastered,
  :result,
  keyword_init: true
)

# Request payload for Fusion#create.
#
# @!attribute [rw] materia1
#   @return [String]
#
# @!attribute [rw] materia1_mastered
#   @return [Boolean]
#
# @!attribute [rw] materia2
#   @return [String]
#
# @!attribute [rw] materia2_mastered
#   @return [Boolean]
#
# @!attribute [rw] result
#   @return [Hash, nil]
FusionCreateData = Struct.new(
  :materia1,
  :materia1_mastered,
  :materia2,
  :materia2_mastered,
  :result,
  keyword_init: true
)

# Materia entity data model.
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [Integer]
#
# @!attribute [rw] max_level
#   @return [Integer, nil]
#
# @!attribute [rw] name
#   @return [String]
#
# @!attribute [rw] rarity
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [String]
Materia = Struct.new(
  :description,
  :id,
  :max_level,
  :name,
  :rarity,
  :type,
  keyword_init: true
)

# Request payload for Materia#load.
#
# @!attribute [rw] id
#   @return [String]
MateriaLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Materia#list.
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [Integer, nil]
#
# @!attribute [rw] max_level
#   @return [Integer, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] rarity
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
MateriaListMatch = Struct.new(
  :description,
  :id,
  :max_level,
  :name,
  :rarity,
  :type,
  keyword_init: true
)

# System entity data model.
#
# @!attribute [rw] status
#   @return [String, nil]
System = Struct.new(
  :status,
  keyword_init: true
)

# Request payload for System#load.
#
# @!attribute [rw] status
#   @return [String, nil]
SystemLoadMatch = Struct.new(
  :status,
  keyword_init: true
)

