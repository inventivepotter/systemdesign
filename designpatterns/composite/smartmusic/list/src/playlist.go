package smartmusic

import "fmt"

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

func (p *Playlist) Reference() (Playable, error) {
	return p, nil
}

func (p *Playlist) Add(pl Playable) {
	item, err := pl.Reference()
	if err != nil {
		fmt.Println(err)
	}
	p.items = append(p.items, &item)
}

func CreatePlaylist(name string) Playlist {
	return Playlist{
		name: name,
	}
}
