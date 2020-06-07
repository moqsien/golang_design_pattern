package single

// 单例模式
type singleton struct {
	count int
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
