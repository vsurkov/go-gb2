package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
)

func Output(res *Result) {
	for hash, m := range res.dupl.m {
		var size uint64
		fcount := uint64(len(m))

		fmt.Printf(DebugColor, fmt.Sprintf("HASH %v\n", hash))
		for dir, file := range m {
			fmt.Printf("  %v/%v\n", dir, file.finfo.Name())
			if size == 0 {
				size = uint64(file.finfo.Size())
			}
		}
		fmt.Printf(InfoColor,
			fmt.Sprintf("Files: %v\tSize: %v \tTotal: %v",
				fcount,
				humanize.Bytes(size),
				humanize.Bytes(size*fcount)))
		fmt.Printf(ErrorColor,
			fmt.Sprintf("\t Overspending: %v\n\n", humanize.Bytes(size*(fcount-1))))
	}

	if len(res.dupl.m) == 0 {
		fmt.Printf(InfoColor, "no duplicate files found\n")
	}
}
