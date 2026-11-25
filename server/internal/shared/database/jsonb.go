package database

import (
	"encoding/json"
	"fmt"

	"gorm.io/datatypes"
)

// MapToJSONB convierte un mapa de Go en datatypes.JSON de GORM, usando "{}" si es nil.
func MapToJSONB(data map[string]any) (datatypes.JSON, error) {
	if data == nil {
		return datatypes.JSON([]byte("{}")), nil
	}

	raw, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal map to JSONB: %w", err)
	}

	return datatypes.JSON(raw), nil
}

// JSONBToMap convierte datatypes.JSON de GORM en un mapa de Go, retornando un mapa vacio si es nulo o vacio.
func JSONBToMap(raw datatypes.JSON) (map[string]any, error) {
	if len(raw) == 0 || string(raw) == "null" {
		return map[string]any{}, nil
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSONB to map: %w", err)
	}

	if result == nil {
		return map[string]any{}, nil
	}

	return result, nil
}
