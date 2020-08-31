package interfaces

type PublicInterface struct {
}

func NewPublicInterface() (i *PublicInterface, err error) {
	i = &PublicInterface{}
	return
}

func (i *PublicInterface) Run() (flow <-chan error) {
	errorsFlow := make(chan error)
	return errorsFlow
}

func (i *PublicInterface) Stop() {

}
