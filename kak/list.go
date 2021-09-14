package kak

// type session struct {
// 	Name    string   `json:"name"`
// 	Clients []string `json:"clients"`
// 	Dir     string   `json:"dir"`
// }

func List() (sessions []Session, err error) {
	return []Session{}, nil

	// TODO probably don't need this func, can just use Sessions() directly

	// s, err := Sessions()
	// if err != nil {
	// 	return s, err
	// }

	// for _, s := range kakSessions {
	// 	session := session{Name: s}

	// 	clients, err := Get("%val{client_list}", "", session.Name, "")
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	if len(clients) > 0 && clients[0] != "" {
	// 		session.Clients = clients
	// 	}

	// 	dir, err := Get("%sh{pwd}", "", session.Name, "")
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	session.Dir = strings.Join(dir, "")

	// 	sessions = append(sessions, session)
	// }

	// return sessions, nil
}
