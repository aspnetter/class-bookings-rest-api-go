package http_handling

type BookingsResponseBody struct {
	ID         int    `json:"id"`
	MemberName string `json:"name"`
	ClassID    int    `json:"class_id"`
	ClassName  string `json:"class_name"`
	Date       string `json:"date_for"`
}
