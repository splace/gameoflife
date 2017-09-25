# Conway's Game Of Life

Uses maps/dicts for storage, so basically has an unlimited grid, with speed varying as number of live cells.

for usage see; executables/readme/usage

in/out storage using png files.

very large grids, easily possible with the 'sparse' internal state storage, are very slow to save.

save sequences of images for making movies.  (eg ffmpeg -r 60 -i %d.png test.mp4)

might be useful for generating explanatory images of certain arrangements/evolutions.


