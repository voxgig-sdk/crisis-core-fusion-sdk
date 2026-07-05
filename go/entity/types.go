// Typed models for the CrisisCoreFusion SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import "encoding/json"

// Fusion is the typed data model for the fusion entity.
type Fusion struct {
	Materia1 string `json:"materia1"`
	Materia1Mastered bool `json:"materia1_mastered"`
	Materia2 string `json:"materia2"`
	Materia2Mastered bool `json:"materia2_mastered"`
	Result *map[string]any `json:"result,omitempty"`
}

// FusionCreateData is the typed request payload for Fusion.CreateTyped.
type FusionCreateData struct {
	Materia1 string `json:"materia1"`
	Materia1Mastered bool `json:"materia1_mastered"`
	Materia2 string `json:"materia2"`
	Materia2Mastered bool `json:"materia2_mastered"`
	Result *map[string]any `json:"result,omitempty"`
}

// Materia is the typed data model for the materia entity.
type Materia struct {
	Description *string `json:"description,omitempty"`
	Id int `json:"id"`
	MaxLevel *int `json:"max_level,omitempty"`
	Name string `json:"name"`
	Rarity *string `json:"rarity,omitempty"`
	Type string `json:"type"`
}

// MateriaLoadMatch is the typed request payload for Materia.LoadTyped.
type MateriaLoadMatch struct {
	Id string `json:"id"`
}

// MateriaListMatch is the typed request payload for Materia.ListTyped.
type MateriaListMatch struct {
	Description *string `json:"description,omitempty"`
	Id *int `json:"id,omitempty"`
	MaxLevel *int `json:"max_level,omitempty"`
	Name *string `json:"name,omitempty"`
	Rarity *string `json:"rarity,omitempty"`
	Type *string `json:"type,omitempty"`
}

// System is the typed data model for the system entity.
type System struct {
	Status *string `json:"status,omitempty"`
}

// SystemLoadMatch is the typed request payload for System.LoadTyped.
type SystemLoadMatch struct {
	Status *string `json:"status,omitempty"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedFrom decodes a runtime value (a map[string]any produced by the op
// pipeline) into a typed model T via a JSON round-trip. On any error it
// returns the zero value of T; the op's own (value, error) tuple carries the
// real error.
func typedFrom[T any](v any) T {
	var out T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value ([]any of maps) into a typed
// slice []T via a JSON round-trip, for list ops.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
