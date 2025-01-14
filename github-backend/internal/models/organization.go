package models

// "login": "my-org-marlom",
// "id": 194842763,
// "url": "https://api.github.com/orgs/my-org-marlom",
// "avatar_url": "https://avatars.githubusercontent.com/u/194842763?v=4",
// "description": null

type Organization struct {
	Login       string `json:"login"`
	ID          int    `json:"id"`
	URL         string `json:"url"`
	AvatarURL   string `json:"avatar_url"`
	Description string `json:"description"`
}

type OrganizationMember struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
	AvatarURL string `json:"avatar_url"`
	URL       string `json:"html_url"`
}
