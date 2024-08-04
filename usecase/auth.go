package usecase

import "github.com/enockjeremi/queswer/domain"

type authUsecase struct {
	authRepo domain.AuthRepository
}

func NewAuthUsecase(ar domain.AuthRepository) domain.AuthUsecase {
	return &authUsecase{authRepo: ar}
}

func (uc *authUsecase) RegisterUser(user *domain.User) error {
	return uc.authRepo.RegisterUserRepo(user)
}

func (uc *authUsecase) VerifyCredentials(user *domain.User) error {
	return uc.authRepo.VerifyCredentialsRepo(user)
}

func (uc *authUsecase) GetUser(user *domain.User, id interface{}) error {
	return uc.authRepo.GetUserRepo(user, id)
}

func (uc *authUsecase) VerifyUsername(user *domain.User, userInput string) error {
	return uc.authRepo.VerifyUsernameRepo(user, userInput)
}

func (uc *authUsecase) CreateProfile(profile *domain.Profile) error {
	return uc.authRepo.CreateProfileRepo(profile)
}

func (uc *authUsecase) UpdateProfile(profile *domain.Profile) error {
	return uc.authRepo.UpdateProfileRepo(profile)
}

func (uc *authUsecase) GetProfile(profile *domain.Profile, id interface{}) error {
	return uc.authRepo.GetProfileRepo(profile, id)
}

func (uc *authUsecase) ChangePassword(user *domain.User, newPassword string) error {
	return uc.authRepo.ChangePasswordRepo(user, newPassword)
}
