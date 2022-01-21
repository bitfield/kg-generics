#!/bin/bash

for file in ./*; do
  if [ -d "$file" ] ; then
		echo Testing $file...
		cd $file
		go test
		cd ..
	fi
done
