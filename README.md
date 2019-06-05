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


## Examples
