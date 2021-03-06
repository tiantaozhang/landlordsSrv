package msg

const  (
	Normal = 10
)


// Ready to start game
type Ready struct {
	State int `json:"state"`
}

// Landlord for call landlord
type Landlord struct {
	State int `json:"state"`
}

// A PokeU for a user
type PokeU struct {
	State int    `json:"state"`
	Uid   string `json:"Uid"`
	Pokes []Poke `json:"pokes"`
}

// A Poke is just a p
type Poke struct {
	Color  string `json:"color"`
	Number string `json:"number"`
}

// Rest pokes
type Rest struct {
	Pokes []string `json:"pokes"`
}

// Room create
type Room struct {
	Id     string  `json:"id" bson:"_id"`
	Pwd    string  `json:"pwd" bson:"pwd"`
	Owner  string  `json:"owner" bson:"owner"`
	Users  [3]string `json:"users" bson:"users"`
	Status int     `json:"status" bson:"status"`
	Time   int     `json:"time,omitempty" bson:"time"`
}

// RoomS search
type RoomS struct {
	Id string `json:"id"`
}

// RoomP pwd
type RoomP struct {
	Id  string `json:"id"`
	Pwd string `json:"pwd"`
}

// user
type User struct {
	Uid      string `json:"uid" bson:"uid"`
	Pwd      string `json:"pwd" bson:"pwd"`
	NickName string `json:"nickName" bson:"nickName"`
	Status   int    `json:"status" bson:"status"`
}
