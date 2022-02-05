package main

import (
	sm "github.com/inventivepotter/systemdesign/designpatterns/smartmusic"
)

func main() {
	p1 := sm.CreatePlaylist("Play List 1")
	p2 := sm.CreatePlaylist("Play List 2")
	p3 := sm.CreatePlaylist("Play List 3")
	s0 := sm.CreateSong("Song 0")
	s1 := sm.CreateSong("Song 1")
	s2 := sm.CreateSong("Song 2")
	s3 := sm.CreateSong("Song 3")
	s4 := sm.CreateSong("Song 4")
	s5 := sm.CreateSong("Song 5")
	s6 := sm.CreateSong("Song 6")
	s7 := sm.CreateSong("Song 7")
	s8 := sm.CreateSong("Song 8")
	s9 := sm.CreateSong("Song 9")
	c1 := sm.CreatePodCast("Podcast 1")
	c2 := sm.CreatePodCast("Podcast 2")
	p1.Add(&s1)
	p1.Add(&s2)
	p1.Add(&s3)
	p1.Add(&s4)
	p1.Add(&s5)
	p2.Add(&s0)
	p2.Add(&p1)
	p2.Add(&s6)
	p2.Add(&c2)
	p2.Add(&p3)
	p3.Add(&s7)
	p3.Add(&s8)
	p3.Add(&s9)
	p3.Add(&c1)
	p3.Add(&c2)
	p2.Play()
}
