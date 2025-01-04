package subject

import "github.com/IlhamRanggaKurniawan/Teamers.git/internal/database/entity"

type SubjectService interface {
	CreateSubject(name string, classId uint64) (*entity.Subject, error)
	FindAllSubjects(classId uint64) (*[]entity.Subject, error)
	DeleteSubject(id uint64) error
}

type subjectService struct {
	subjectRepository SubjectRepository
}

func NewService(subjectRepository SubjectRepository) SubjectService {
	return &subjectService{subjectRepository: subjectRepository}
}

func (s *subjectService) CreateSubject(name string, classId uint64) (*entity.Subject, error) {
	subject, err := s.subjectRepository.Create(name, classId)

	if err != nil {
		return nil, err
	}

	return subject, nil
}

func (s *subjectService) FindAllSubjects(classId uint64) (*[]entity.Subject, error) {
	subjects, err := s.subjectRepository.FindAllByClassId(classId)

	if err != nil {
		return nil, err
	}

	return subjects, nil
}

func (s *subjectService) DeleteSubject(id uint64) error {
	err := s.subjectRepository.Delete(id)

	return err
}
