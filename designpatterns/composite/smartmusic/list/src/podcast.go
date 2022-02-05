package smartmusic

import "fmt"

type Podcast struct {
	name string
}

func (p Podcast) Play() {}

func (p *Podcast) Reference() (Playable, error) {
	var speed string
	fmt.Print("Choose the speed of " + p.name + " to be played 1x, 1.25x, 1.5x : ")
	if _, err := fmt.Scanf("%s", &speed); err != nil {
		fmt.Println(err)
		return nil, err
	}
	wrap := PodcastWrapper{Podcast: p, speed: speed}
	return wrap, nil
}

//Podcast implements Playable
type PodcastWrapper struct {
	Podcast *Podcast
	speed   string
}

func (pw PodcastWrapper) Play() {
	fmt.Println(pw.Podcast.name + " is playing at speed of " + pw.speed)
}

func (pw PodcastWrapper) Reference() (Playable, error) {
	return nil, nil
}

func CreatePodCast(name string) Podcast {
	return Podcast{
		name: name,
	}
}
