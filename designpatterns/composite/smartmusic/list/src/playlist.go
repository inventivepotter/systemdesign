package smartmusic

//Playlist implements Playable
type Playlist struct {
	name  string
	items []*Playable
}

func (p Playlist) Play() {
	for _, item := range p.items {
		(*item).Play()
	}
}

func (p *Playlist) Reference() Playable {
	return p
}

func (p *Playlist) Add(pl Playable) {
	tada := pl.Reference()
	p.items = append(p.items, &tada)
}

func CreatePlaylist(name string) Playlist {
	return Playlist{
		name: name,
	}
}
