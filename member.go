package clienty

type Member struct {
	client          *Client
	ID              string   `json:"id"`
	Username        string   `json:"username"`
	FullName        string   `json:"fullName"`
	Initials        string   `json:"initials"`
	AvatarHash      string   `json:"avatarHash"`
	Email           string   `json:"email"`
	IDBoards        []string `json:"idBoards"`
	IDOrganizations []string `json:"idOrganizations"`
}
