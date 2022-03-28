package model

import (
	"encoding/json"
	"time"
)

type Resume struct {
	Phone         string    `json:"phone,omitempty"`
	Name          string    `json:"name,omitempty"`
	Age           int       `json:"age,omitempty"`
	Job           *string   `json:"job,omitempty"`
	CreateAt      time.Time `json:"create_at,omitempty"`
	UpdateAt      time.Time `json:"update_at,omitempty"`
	ResumeContent *[]byte   `json:"resume_content,omitempty"`
}

func (r Resume) MarshalJSON() ([]byte, error) {
	var temp = map[string]interface{}{
		"phone":     r.Phone,
		"name":      r.Name,
		"age":       r.Age,
		"create_at": r.CreateAt,
		"update_at": r.UpdateAt,
	}
	if r.Job != nil {
		temp["job"] = r.Job
	}
	// if r.CreateAt != nil {
	// 	temp["create_at"] = string(*r.CreateAt)
	// }
	if r.ResumeContent != nil {
		temp["resumeContent"] = string(*r.ResumeContent)
	}
	return json.Marshal(temp)
}
