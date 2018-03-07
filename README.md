# Last FM -> Discord status scrobbling

Sets your "listening to" status on Discord to the currently playing song according to your Last.fm profile.

For easy starting of this script, grab the latest release and the `run.bat.sample` file.
Make a copy of it and rename to `run.bat`, then input your details and save. You can now run this to start the script.

### !Important!

 - Discord is against custom clients and self-bots like this so while no one that has used this has been banned for it yet (as far as I know) and you **probably** wont, I can make no guarantees. **USE AT YOUR OWN RISK**. I am not responsible for you getting banned from Discord.
 - You will not be able to see the status yourself in the client, you should log into a second account or ask a friend to confirm wether it's working or not.
 - You should probably turn off "Display currently running game as status message" in the client when using this, otherwise it will override this when you launch a game.


### To compile it yourself:
 1. Install and set up [Go](https://golang.org/doc/install) and [Git](https://git-scm.com/download).
 2. Run `go get -u <repo>` where `<repo>` is the URL of the repo, beginning `github.com/`.
 3. It should now be built in `$GOPATH/bin` unless it threw errors

Usage:
```
  -g string
        Game to set to if there hasn't been a new song for a while (if unset, the playing status will be cleared when nothing is playing)
  -l string
        Last.fm api key
  -t string
        Discord token
  -u string
        Last.fm username
  -s string
        Status the bot should set the user to (if unset, the user's client status is respected. If set, bot will over-ride every client)
```

Example: `./discordlfm -t "my secret discord token" -l "my secret last.fm api key" -u "jonasr747"`

Guide on retrieving the lfm api key and discord token: https://github.com/jonas747/discordlfm/issues/2#issuecomment-278950579
