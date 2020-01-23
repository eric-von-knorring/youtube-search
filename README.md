# Youtube search

A command line tool to search for youtube videos and output the tile and link on standard out.

## Usage

This can be used to get youtube search result and use them in scripts. 

As an example the following script lets you select one of the search results and open the video link with mpv.

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

## Examples
