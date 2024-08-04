package usecase

import "github.com/enockjeremi/queswer/domain"

type answerUsecase struct {
	answerRepo domain.AnswerRepository
}

func NewAnswerUsecase(qr domain.AnswerRepository) domain.AnswerUsecase {
	return &answerUsecase{answerRepo: qr}
}

func (uc *answerUsecase) GetAllAnswer() ([]domain.Answer, error) {
	return uc.answerRepo.GetAll()
}

func (uc *answerUsecase) CreateAnswer(answer *domain.Answer) error {
	return uc.answerRepo.Create(answer)
}

func (uc *answerUsecase) GetAnswer(answer *domain.Answer, id string) error {
	return uc.answerRepo.GetOne(answer, id)
}

func (uc *answerUsecase) UpdateAnswer(answer *domain.Answer) error {
	return uc.answerRepo.Update(answer)
}

func (uc *answerUsecase) RemoveAnswer(answer *domain.Answer, id string) error {
	return uc.answerRepo.Delete(answer, id)
}

func (uc *answerUsecase) GetQuestion(question *domain.Question, id string) error {
	return uc.answerRepo.GetOneQuestion(question, id)
}
