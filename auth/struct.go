package auth

type AuthStruct struct {
	Code         string
	Token        string
	RefreshToken string
	ClientID     string
	TenantID     string
	State        int //Random thing
}
