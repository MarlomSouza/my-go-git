package models

type User struct {
	Login                   string `json:"login"`
	ID                      int    `json:"id"`
	AvatarURL               string `json:"avatar_url"`
	Name                    string `json:"name"`
	Email                   string `json:"email"`
	PublicRepos             int    `json:"public_repos"`
	PublicGists             int    `json:"public_gists"`
	Followers               int    `json:"followers"`
	Following               int    `json:"following"`
	CreatedAt               string `json:"created_at"`
	UpdatedAt               string `json:"updated_at"`
	PrivateGists            int    `json:"private_gists"`
	TotalPrivateRepos       int    `json:"total_private_repos"`
	OwnedPrivateRepos       int    `json:"owned_private_repos"`
	TwoFactorAuthentication bool   `json:"two_factor_authentication"`
}
