# Conway's Game Of Life

Uses maps/dicts for storage, so basically has an unlimited grid, with speed varying as number of live cells.

for usage see; executables/readme/usage

in/out storage using png files.

very large grids, easily possible with the 'sparse' internal state storage, are very slow to save.

can save sequences of images for making movies. (might be useful for generating explanatory images of certain arrangements/evolutions.)

# making video

from snapshots already saved in a directory; (made using option; -movie=\<dir>)

	 ffmpeg -i <dir>/%d.png -r 25 output.mp4

alternatively; (requires using "ls -v" so images are sequenced in numeric not alpha order.) 

	ls -v <dir>/* | xargs cat | ffmpeg -i pipe:0 -r 25  output.mp4

	## some other nice to have options on ffmpeg: -pix_fmt yuv420p -loglevel warning

and to create directly from piped snapshots; (uses -o to stop final image going to stdout and so movie.) 

	./life -i="glider" -pipeMovie -wrap -size=200 -ticks=500 -o=/dev/null | ffmpeg -i pipe:0 -r 25 glider.mp4
