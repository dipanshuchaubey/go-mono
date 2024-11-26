package types

type BootcampInfo struct {
	ID          string   `json:"id"`
	Title       string   `json:"name"`
	Description string   `json:"description"`
	Website     string   `json:"website"`
	Email       string   `json:"email"`
	NameSlug    string   `json:"slug"`
	Careers     []string `json:"careers"`
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
	Data []BootcampInfo `json:"data"`
}

type CourseResponse struct {
	Data CourseInfo `json:"data"`
}

type ReviewResponse struct {
	Data []ReviewDetails `json:"data"`
}
