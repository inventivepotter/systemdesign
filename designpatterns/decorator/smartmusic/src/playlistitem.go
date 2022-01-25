package smartmusic

type PlaylistItem struct {
	item *Playable
	next *PlaylistItem
}

func (pli *PlaylistItem) SetNext(apli *PlaylistItem) {
	pli.next = apli
}

func NewPlaylistItem(pl *Playable) PlaylistItem {
	return PlaylistItem{
		item: pl,
		next: nil,
	}
}
