package smartmusic

import "fmt"

//Song implements Playable
type Song struct {
	name string
}

func (s Song) Play() {
	fmt.Println(s.name + " is playing")
}

func (s Song) Name() string {
	return s.name
}

func CreateSong(name string) Song {
	return Song{
		name: name,
	}
}
