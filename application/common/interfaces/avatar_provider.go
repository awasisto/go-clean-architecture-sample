package interfaces

type AvatarProvider interface {
	GetAvatarUrlByEmail(email string) (avatarUrl string, err error)
}
