# Youtube search

A command line client to search for youtube videos and output the tile and link on standard out.

## Usage

This can be used to get youtube search result and use them in scripts.

As an example the following script lets you select one of the search with [fuzzy find](https://github.com/junegunn/fzf) results and open the video link with [mpv](https://mpv.io/).

```bash
#!/bin/sh

youtube-search $* | fzf | awk ' {print $NF}' | xargs mpv
```

## Installation
Requires:
 - Go version 1.13 or later.
 - Make

Clone the git repository then enter the git project with your faviorite terminal and run:
```
sudo make install
```

## Sample output
```
$ youtube-search asdf
asdfmovie 1-12 (Complete Collection) - https://www.youtube.com/watch?v=CWBnMBz-zZU
asdfmovie12 - https://www.youtube.com/watch?v=UsVWkAq1s0U
THE MUFFIN SONG (asdfmovie feat. Schmoyoho) - https://www.youtube.com/watch?v=LACbVhgtx9I
asdfmovie 1-11 (Complete Collection) - https://www.youtube.com/watch?v=2CeurT8vjU0
asdfmovie - https://www.youtube.com/watch?v=UsVWkAq1s0U&amp;list=PL3A5849BDE0581B19
asdfmovie11 - https://www.youtube.com/watch?v=LJnpwL-2Drc
asdfmovie10 - https://www.youtube.com/watch?v=foFKXS6Nyho
Asdf Movie Alle Teile Deutsch - https://www.youtube.com/watch?v=fhS9Y3SfuRQ
The Best of Asdf - https://www.youtube.com/watch?v=K8KHbrKeQEI
all asdfmovie songs in order - https://www.youtube.com/watch?v=ovn3rfBa79A
asdfmovie - https://www.youtube.com/watch?v=IYnsfV5N2n8
{YTP} ~ ultrasdf - https://www.youtube.com/watch?v=b4dhm_CVm_0
SFM FNaF: asdfmovie 1-11 (five nights at freddy&#39;s animation) - https://www.youtube.com/watch?v=eyjQFTiJluU
(50k and up subs) Asdfmovie10  - Roblox, lego and original - https://www.youtube.com/watch?v=obcR0RVxqXM
Asdfmovie 9 (Original and Minecraft version) - https://www.youtube.com/watch?v=V173Y4uM5dE
asdfmovie8 RealLife/Original Comparison - https://www.youtube.com/watch?v=NXrtLVyToJs
Cyanide &amp; Happiness Compilation - #14 - https://www.youtube.com/watch?v=5k4VP0C4byY
Eddsworld - Fun Dead - https://www.youtube.com/watch?v=3w1pFW44xkM
Llamas with Hats 1-12: The Complete Series - https://www.youtube.com/watch?v=jJOwdrTA8Gw
asdfmovie9 - https://www.youtube.com/watch?v=8l6T3fwxAyw
STOCKTALES - https://www.youtube.com/watch?v=WXIoWp92ux8
ASDF Movie Dublado - https://www.youtube.com/watch?v=g4FFrfTAMkQ&amp;list=PLOLH3sPZwKEJAb1cvG-BUHxDZJg6fzyUQ
asdfmovie3 - https://www.youtube.com/watch?v=oY6tCnu-1Do
EVERYBODY DO THE FLOP (asdfmovie song) - https://www.youtube.com/watch?v=L5inD4XWz4U
asdfmovie12: deleted scenes - https://www.youtube.com/watch?v=FJLpNSjE1J4
asdfmovie5 - https://www.youtube.com/watch?v=tCnj-uiRCn8
ASDF movie 1, 2 und 3  Österreichisch - https://www.youtube.com/watch?v=yN_kvY_5SWo
ASDF Movie 1,2,3,4,5,6 - https://www.youtube.com/watch?v=PXGljpsidd4
asdfmovie7 - https://www.youtube.com/watch?v=5xniR1GN69U
```

