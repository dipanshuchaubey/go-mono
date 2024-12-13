package dto

import (
	pbrq "carthage/protos/bootcamp_service/request"
)

type CreateBootcampBody struct {
	Title       string   `json:"name"`
	Description string   `json:"description"`
	Website     string   `json:"website"`
	Email       string   `json:"email"`
	Phone       string   `json:"phone"`
	Careers     []string `json:"careers"`
	Address     string   `json:"address"`
}

func (b *CreateBootcampBody) FromProto(body *pbrq.CreateBootcampRequest) {
	b.Title = body.Title
	b.Description = body.Description
	b.Website = body.Website
	b.Email = body.Email
	b.Phone = body.Phone
	b.Careers = body.Careers
	b.Address = body.Address
}
