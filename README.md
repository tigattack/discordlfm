# Last FM -> Discord status scrobbling

Sets your "game" to the last, or currently playing song played on your last.fm

### To download pre-compiled executables see [releases](https://github.com/jonas747/discordlfm/releases)

### To compile it yourself:
 1. Install and set up go
 2. Run `go get -u github.com/jonas747/discordlfm`
 3. Switch to the develop branch of discordgo (master statuses currently broken) (`cd $GOPATH/src/github.com/bwmarrin/discordgo && git checkout develop`
 4. Re-compile discordlfm (`go install -a -v github.com/jonas747/discordlfm`)
 5. It should now be built in `$GOPATH/bin` unless it threw errors

Usage:
```
  -g string
        Game to set to if there hasn't been a new song for a while (default "", clears playing status)
  -l string
        Last.fm api key
  -n int
        Number of seconds without a new song for it to be considered nothing. (default 600)
  -t string
        Discord token
  -u string
        Last.fm username
```

Example: `./discordlfm -t "my secret discord token" -l "my secret last.fm api key" -u "jonasr747"`

Guide on retrieving the lfm api key and discord token: https://github.com/jonas747/discordlfm/issues/2#issuecomment-278950579
