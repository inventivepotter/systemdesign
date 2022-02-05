package smartmusic

import "fmt"

//Playlist implements Playable
type Playlist struct {
	name    string
	current *PlaylistItem
	start   *PlaylistItem
	end     *PlaylistItem
}

func (p *Playlist) Play() {
	cs := *p.current.item
	cs.Play()
	for p.Next() != nil {
		p.Play()
	}
}

func (p Playlist) Name() string {
	return p.name
}

func (p *Playlist) Next() *PlaylistItem {
	if p.current == nil {
		return nil
	}
	pli := *p.current
	p.current = pli.next
	return pli.next
}

func (p *Playlist) Add(pl Playable) {
	if p == pl {
		return
	}
	if p.start == nil {
		//fmt.Println("First item added in Playlist", pl.Name(), p.Name())
		pli := NewPlaylistItem(&pl)
		p.setStart(&pli)
		p.setEnd(&pli)
		p.setCurrent(&pli)
	} else {
		//fmt.Println("Another item added in Playlist", pl.Name(), p.Name())
		pli := NewPlaylistItem(&pl)
		p.end.SetNext(&pli)
		p.setEnd(&pli)
	}
}

func (p *Playlist) setStart(pli *PlaylistItem) {
	p.start = pli
}

func (p *Playlist) setCurrent(pli *PlaylistItem) {
	p.current = pli
}

func (p *Playlist) setEnd(pli *PlaylistItem) {
	p.end = pli
}

func (p *Playlist) Print() {
	fmt.Print(p.Name(), " has \n\t")
	pli := *p.start
	for pli.next != nil {
		fmt.Print((*pli.item).Name(), "\t")
		pli = *pli.next
	}
	fmt.Println((*pli.item).Name(), "\t")
}

func CreatePlaylist(name string) Playlist {
	return Playlist{
		name: name,
	}
}
