package entity

type Player struct {
    ID string
    Name string
    Price float64
}

func NewPlayer(id, name string, price float64) *Player {
    return &Player{
        ID: id,
        Name: name,
        Price: price,
    }
}