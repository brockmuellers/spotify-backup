# spotify-library-backup

WIP

I love the convenience of Spotify, but hate not being able to manage my own music library. There are a few annoying things about the Spotify "Your Library": most notably, there's a 10k song limit. There are also reports of albums mysteriously disappearing from people's libraries, or libraries being lost entirely. Spotify may not be around forever. I may need to change accounts. Some bad actor may destroy all of Spotify's data centers. In any case, I've spent hundreds of hours on the music discovery process, so I have no interest in losing any of that history.

So, here's a tool to back up a Spotify library. I'd like to eventually be able to combine this with my old iTunes library, restore the library, etc. All in good time.

### Notes to self

running executable: from repo root, assuming $HOME/go/bin is added to $PATH,
```
$ go install
$ spotify-backup
```

compile and test package, from package root:
```
$ go build
$ go test
```
