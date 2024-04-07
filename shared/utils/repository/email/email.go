package emailRepo

type IEmailRepository interface {
	SendVerification(code string) error
}

type EmailRepository struct{}
// SendVerification implements IEmailRepository.
func (*EmailRepository) SendVerification(code string) error {
	panic("unimplemented")
}

func NewEmailRepository() IEmailRepository {
	return &EmailRepository{}
}
