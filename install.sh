#!/bin/bash
function cls {
	# clear screen and move cursor to 0,0
	echo -e "\033[2J\033[0;0H"
}
cls
echo "Building application..."
go build .
cls
echo "Preparing executable..."
chmod +x ./plezk
cls
echo "Moving executable to /usr/local/bin..."
mv ./plezk /usr/local/bin/plezk
cls
echo "Finished installing plezk!"
