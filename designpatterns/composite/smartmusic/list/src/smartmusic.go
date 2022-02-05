package smartmusic

//Playable interface allows us to use playlist or song to be within playlist item
type Playable interface {
	Play()
	Reference() Playable
}
