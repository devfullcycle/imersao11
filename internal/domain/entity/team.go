package entity

type Team struct {
    ID string
    Name string
    Players []*Player
}

func NewTeam(id, name string) *Team {
    return &Team{
        ID: id,
        Name: name,
    }
}

func (t *Team) AddPlayer(player *Player) {
    t.Players = append(t.Players, player)
}

func (t *Team) RemovePlayer(player *Player) {
    for i, p := range t.Players {
        if p.ID == player.ID {
            t.Players = append(t.Players[:i], t.Players[i+1:]...)
            return
        }
    }
}