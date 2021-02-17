package crypto

type Keystore struct {
	namespace string
}

func NewKeystore(namespace string) (k *Keystore) {
	k = &Keystore{
		namespace: namespace,
	}
	return
}

//
//func (k *Keystore) Sign()
