package keyword

type Service struct{}

func NewService() *Service {
    return &Service{}
}

func (s *Service) List() []Keyword {
    return allEntities
}

func (s *Service) Get(idx int) (*Keyword, error) {
    return &allEntities[idx], nil
}

func (s *Service) Delete(idx int) error {
    allEntities = append(allEntities[:idx], allEntities[idx+1:]...)
    return nil
}
