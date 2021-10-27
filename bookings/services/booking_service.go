package services

func (s *service) GetBookings() []Booking {
	return s.repository.GetAllBookings()
}
