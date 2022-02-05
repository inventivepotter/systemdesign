package smartmusic

import "fmt"

//Song implements Playable
type Song struct {
	name string
}

func (s Song) Play() {
	fmt.Println(s.name + " is playing")
}

func (s *Song) Reference() (Playable, error) {
	return s, nil
}

func CreateSong(name string) Song {
	return Song{
		name: name,
	}
}
