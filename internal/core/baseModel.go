package core

type (
	BaseModel struct {
		Id string `json:"_id" bson:"_id,omitempty"`
	}
)
