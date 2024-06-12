package models

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type SettingType int

const (
	TypeString SettingType = iota
	TypeInt
	TypeFloat
	TypeBool
	TypeJSON
)

// settingTypeMap maps SettingType values to their string representations
var settingTypeMap = map[SettingType]string{
	TypeString: "string",
	TypeInt:    "int",
	TypeFloat:  "float",
	TypeBool:   "bool",
	TypeJSON:   "json",
}

// settingTypeReverseMap maps string representations to their SettingType values
var settingTypeReverseMap = map[string]SettingType{
	"string": TypeString,
	"int":    TypeInt,
	"float":  TypeFloat,
	"bool":   TypeBool,
	"json":   TypeJSON,
}

// Setting represents the settings table in the database
type Setting struct {
	BasicModel

	Key         string      `gorm:"primaryKey"`
	Value       string      `gorm:"not null"`
	Type        SettingType `gorm:"not null"`
	Title       string
	Description string
}

// TODO: Work on Type Safety
// SetValue sets the value of the setting based on the type
// func (s *Setting) SetValue(value interface{}) error {
// 	switch s.Type {
// 	case TypeString:
// 		s.Value = value.(string)
// 	case TypeInt:
// 		s.Value = strconv.Itoa(value.(int))
// 	case TypeFloat:
// 		s.Value = fmt.Sprintf("%f", value.(float64))
// 	case TypeBool:
// 		s.Value = strconv.FormatBool(value.(bool))
// 	case TypeJSON:
// 		jsonValue, err := json.Marshal(value)
// 		if err != nil {
// 			return err
// 		}
// 		s.Value = string(jsonValue)
// 	default:
// 		return fmt.Errorf("unsupported type: %d", s.Type)
// 	}
// 	return nil
// }

// GetValue retrieves the value of the setting based on the type
func (s *Setting) GetValue() (interface{}, error) {
	switch s.Type {
	case TypeString:
		return s.Value, nil
	case TypeInt:
		return strconv.Atoi(s.Value)
	case TypeFloat:
		return strconv.ParseFloat(s.Value, 64)
	case TypeBool:
		return strconv.ParseBool(s.Value)
	case TypeJSON:
		var result interface{}
		if err := json.Unmarshal([]byte(s.Value), &result); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unsupported type: %d", s.Type)
	}
}

func (s *Setting) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":          s.ID,
		"key":         s.Key,
		"value":       s.Value,
		"type":        settingTypeMap[s.Type],
		"title":       s.Title,
		"description": s.Description,
		"created_at":  s.CreatedAt,
		"updated_at":  s.UpdatedAt,
	})
}

// UnmarshalJSON unmarshals the JSON data into the setting
func (s *Setting) UnmarshalJSON(data []byte) error {
	var aux struct {
		Key         string      `json:"key"`
		Value       interface{} `json:"value"`
		Type        string      `json:"type"`
		Title       string      `json:"title"`
		Description string      `json:"description"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	s.Key = aux.Key
	s.Value = aux.Value.(string)
	s.Title = aux.Title
	s.Description = aux.Description

	t, ok := settingTypeReverseMap[aux.Type]
	if !ok {
		return fmt.Errorf("unsupported type: %s", aux.Type)
	}
	s.Type = t
	return nil
}
