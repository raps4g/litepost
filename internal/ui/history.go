package ui

func (s *State) saveCurrentRequest() {
    s.History = append(s.History, s.Request) 
}
