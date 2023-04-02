package user

type Service struct {
	Repo Repository
}

func (s *Service) GetUserByID(userID int) (*User, error) {
	// Call UserRepository to retrieve user data
	return s.Repo.GetUserByID(userID)
}
