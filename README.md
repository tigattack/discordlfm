# Last FM -> Discord status scrobbling

Sets your "game" to the last, or currently playing song played on your last.fm

### !Important!

 - Discord is against custom clients and self-bots like this so while no one that has used this has gotten banned for it yet (as far as i know) and you **probably** wont, i can make no guarantees. **USE ON OWN RISK**. I am not responsible for you getting banned from discord.
 - You will not be able to see the status yourself in the client, you should log into a second account or ask a friend to confirm wether it's working or not.
 - You should probably turn off "Display currently running game as status message" in the client when using this, otherwise it will override this when you launch a game or something.

### To download pre-compiled executables see [releases](https://github.com/jonas747/discordlfm/releases)

### To compile it yourself:
 1. Install and set up go
 2. Run `go get -u github.com/jonas747/discordlfm`
 3. It should now be built in `$GOPATH/bin` unless it threw errors

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
