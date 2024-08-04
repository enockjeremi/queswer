package usecase

import "github.com/enockjeremi/queswer/domain"

type questionUsecase struct {
	questionRepo domain.QuestionRepository
}

func NewQuestionUsecase(qr domain.QuestionRepository) domain.QuestionUsecase {
	return &questionUsecase{questionRepo: qr}
}

func (uc *questionUsecase) GetAllQuestion() ([]domain.Question, error) {
	return uc.questionRepo.GetAll()
}

func (uc *questionUsecase) CreateQuestion(question *domain.Question) error {
	return uc.questionRepo.Create(question)
}

func (uc *questionUsecase) GetQuestion(question *domain.Question, id string) error {
	return uc.questionRepo.GetOne(question, id)
}

func (uc *questionUsecase) UpdateQuestion(question *domain.Question) error {
	return uc.questionRepo.Update(question)
}

func (uc *questionUsecase) RemoveQuestion(question *domain.Question, id string) error {
	return uc.questionRepo.Delete(question, id)
}
