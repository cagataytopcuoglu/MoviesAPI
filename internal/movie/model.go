package movie

import (
	"MovieAPI/internal/core"
	"github.com/thedevsaddam/govalidator"
)

type (
	Movie struct {
		core.BaseModel `bson:",inline"`
		Name           string `json:"name" bson:"Name"`
		Description    string `json:"description" bson:"Description"`
		Type           string `json:"type" bson:"Type"`
	}
)

func (s *Movie) Validate() map[string][]string {

	rules := govalidator.MapData{
		"name":        []string{"required"},
		"description": []string{"required"},
		"type":        []string{"required"},
	}
	opts := govalidator.Options{
		Data:  s,
		Rules: rules,
	}
	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) > 0 {
		//data, _ := json.MarshalIndent(e, "", "  ")
		//fmt.Println(string(data))
		return e
	}
	return nil
}
