package dto

import "time"

type CreateUpdateResponse struct {
	CreatedAt time.Time `json:"created_at"`
	CreatedBy *int      `json:"created_by,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	UpdatedBy *int      `json:"updated_by,omitempty"`
}

type DeleteInfoResponse struct {
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	DeletedBy *int      `json:"deleted_by,omitempty"`
}
