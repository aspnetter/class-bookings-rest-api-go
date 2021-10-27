package http_handling

type BookingsRequestBody struct {
	MemberName string `json:"member_name"`
	Date       string `json:"date_for"`
}
