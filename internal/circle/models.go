package circle


type Circle struct {
    ID uint
    Sectors []Sector
    UserID uint
}


type Sector struct {
    ID uint
    Name string
    Value float32
    CircleID uint
}
