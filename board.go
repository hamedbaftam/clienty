package clienty

import "fmt"

type Board struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Desc           string `json:"desc"`
	Closed         bool   `json:"closed"`
	IDOrganization string `json:"idOrganization"`
	Pinned         bool   `json:"pinned"`
	Starred        bool   `json:"starred"`
	URL            string `json:"url"`
	ShortURL       string `json:"shortUrl"`
}

func (m *Member) GetBoards() (err error) {
	//path := fmt.Sprintf("members/%s/boards", m.ID)
	//err = m.client.Get(path, args, &boards)
	fmt.Println("Erfg")
	//for i := range boards {
	//	fmt.Println(i)
	//}
	//return
	return
}
