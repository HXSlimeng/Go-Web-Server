package sysdept

type Service struct{}

func (s *Service) GetList() []Model {
	testModel := Model{1, 2, "DeptName", "DeptPath", 2}
	return []Model{testModel}
}