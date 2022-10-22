package auth

type InfoStruct struct {
	Code  string
	State int //Random thing
	Token TokenStruct
}

type TokenStruct struct {
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	ExtExpiresIn int    `json:"ext_expires_in"`
	AccessToken  string `json:"access_token"`
}
