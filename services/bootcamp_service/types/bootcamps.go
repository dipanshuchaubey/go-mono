package types

import (
	bs "carthage/protos/bootcamp_service"
)

type BootcampInfo struct {
	ID          string   `json:"id"`
	Title       string   `json:"name"`
	Description string   `json:"description"`
	Website     string   `json:"website"`
	Email       string   `json:"email"`
	NameSlug    string   `json:"slug"`
	Careers     []string `json:"careers"`
}

func (b *BootcampInfo) ToProro(proto *bs.BootcampInfo) {
	proto.BootcampId = b.ID
	proto.Title = b.Title
	proto.Description = b.Description
	proto.Email = b.Description
	proto.NameSlug = b.NameSlug
	proto.Website = b.Website
	proto.Careers = b.Careers
}

type CourseInfo struct {
	ID          string `json:"_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ReviewDetails struct {
	ID           string `json:"_id"`
	UserID       string `json:"user"`
	Title        string `json:"title"`
	MessageInfos string `json:"text"`
	Rating       int64  `json:"rating"`
}

type BootcampResponse struct {
	Success bool           `json:"success"`
	Data    []BootcampInfo `json:"data"`
}

type CreateBootcampResponse struct {
	Success bool         `json:"success"`
	Data    BootcampInfo `json:"data"`
	Error   string       `json:"error,omitempty"`
}

type CourseResponse struct {
	Data CourseInfo `json:"data"`
}

type ReviewResponse struct {
	Data []ReviewDetails `json:"data"`
}

type BootcampReviews struct {
	BootcampID string
	Reviews    []ReviewDetails
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}
