package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/dannyt66/discordgo"
	"github.com/shkh/lastfm-go/lastfm"
)

const (
	VERSION_MAJOR = 2
	VERSION_MINOR = 1
	VERSIN_PATCH  = 0
)

var (
	VersionString = fmt.Sprintf("%d.%d.%d", VERSION_MAJOR, VERSION_MINOR, VERSIN_PATCH)
)

var (
	flagDiscordToken string
	flagLFMAPIKey    string
	flagLFMUsername  string
	flagNoSong       string
	flagStatus       string
)

func init() {
	flag.StringVar(&flagDiscordToken, "t", "", "Discord token")
	flag.StringVar(&flagLFMAPIKey, "l", "", "Last.fm api key")
	flag.StringVar(&flagLFMUsername, "u", "", "Last.fm username")
	flag.StringVar(&flagNoSong, "g", "", "Game to set to if there hasn't been a new song for a while")
	flag.StringVar(&flagStatus, "s", "", "Status the bot should set the user to")
	flag.Parse()
}

func main() {
	log.Println("Starting up... v" + VersionString)

	if flagDiscordToken == "" {
		fatal("No discord token specified")
	}

	if flagLFMAPIKey == "" {
		fatal("No lastfm api key specified")
	}

	if flagLFMUsername == "" {
		fatal("No last.fm username specified")
	}

	// Setup lastfm
	lfm := lastfm.New(flagLFMAPIKey, "")

	// Setup discord
	dsession, err := discordgo.New(flagDiscordToken)
	if err != nil {
		fatal("Error creating discord session:", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	dsession.AddHandlerOnce(func(s *discordgo.Session, r *discordgo.Ready) {
		wg.Done()
	})

	err = dsession.Open()
	if err != nil {
		fatal("Error opening discord ws conn:", err)
	}

	wg.Wait()

	log.Println("Ready received! Ctrl-c to stop.")
	run(dsession, lfm)
}

func run(s *discordgo.Session, lfm *lastfm.Api) {
	// Run continously untill somethign catches fire or an std
	ticker := time.NewTicker(time.Second * 15)
	lastPlaying := ""
	for {
		<-ticker.C
		playing, err, isPlaying := check(lfm)
		if playing != lastPlaying {
			if isPlaying == false {
				if flagNoSong == "" {
					s.UpdateListeningStatus("")
				} else {
					s.UpdateListeningStatus(flagNoSong)
				}
				log.Println("Not currently playing anything.")
			} else if err != nil {
				log.Println("Error checking:", err)
				continue
			} else {
				s.UpdateListeningStatus(playing)
				log.Println("Updated status to:", playing)
			}
		}
		lastPlaying = playing
	}
}

func check(lfm *lastfm.Api) (string, error, bool) {
	recent, err := lfm.User.GetRecentTracks(map[string]interface{}{"user": flagLFMUsername})
	if err != nil {
		return "", err, true
	}

	if len(recent.Tracks) < 1 {
		return "", errors.New("No tracks"), true
	}
	track := recent.Tracks[0]
	if track.NowPlaying == "" {
		return "", nil, false
	} else {
		return track.Artist.Name + " - " + track.Name, nil, true
	}
}

func fatal(args ...interface{}) {
	log.Println(args...)
	log.Print("Press enter to quit...")

	var input rune
	fmt.Scanf("%c", &input)

	os.Exit(1)
}
