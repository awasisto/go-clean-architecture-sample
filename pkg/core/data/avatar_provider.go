package data

type AvatarProvider interface {
	GetAvatarUrlByEmail(email string) (avatarUrl string, err error)
}
