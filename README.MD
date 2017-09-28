# Conway's Game Of Life

Uses maps/dicts for storage, so basically has an unlimited grid, with speed varying as number of live cells.

for usage see; executables/readme/usage

in/out storage using png files.

very large grids, easily possible with the 'sparse' internal state storage, are very slow to save.

save sequences of images for making movies. 

how to: converting to video;

	from snapshots already saved in a directory ( -movie=<<dir>> );
		
		 avconv -i <<dir>>/%d.png -r 25 output.mp4
	
	or (needs "ls -v" to order files in numeric not alpha order.): 
	
		ls -v <<dir>>/* | xargs cat | avconv -i pipe:0 -r 25  output.mp4

			nice to have options on avconv: -pix_fmt yuv420p -loglevel warning
			
	for directly piped snapshots ( uses -o=<???> to stop final image going to stdout ) 
	
		./life -i="glider" -pipeMovie -wrap -size=200 -ticks=500 -o=out | avconv -i pipe:0 -r 25 output.mp4


might be useful for generating explanatory images of certain arrangements/evolutions.


