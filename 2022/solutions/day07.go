package solutions

import (
	"aoc22/helpers"
	"fmt"
	"regexp"
	"time"
)

type Dir struct {
	name   string
	size   int
	parent *Dir
	dirs   []*Dir
	files  []File
}

type File struct {
	name string
	size int
}

func day07OneAndTwo() (int, int) {
	lines, _ := helpers.FileToLines("day07.txt")
	// add a fake `$ cd` at end to force recalculation
	lines = append(lines, "$ ls")
	totalSize := 0
	rootDir := Dir{
		name:   "/",
		size:   0,
		parent: nil,
		dirs:   nil,
		files:  nil,
	}
	var curDir *Dir = &rootDir
	currentFiles := []File{}
	for _, line := range lines {
		if string(line[0]) == "$" {
			re := regexp.MustCompile(`\$ (\w+)( (\w+|/|..))?`)
			m := re.FindAllStringSubmatch(line, -1)[0][1:]
			// recalculate sizes on each command
			if len(curDir.files) == 0 {
				curDir.files = currentFiles
			}
			curDir.size = calcSize(curDir)
			updateParent(curDir)
			if m[0] == "cd" {
				curDir = changeDir(m[2], curDir, &rootDir)
			}
			currentFiles = []File{}
		} else {
			re := regexp.MustCompile(`^(\d+|\w+) (\w+.?(\w+)?|\d+)`)
			m := re.FindAllStringSubmatch(line, -1)[0][1:]
			// check for containing directories
			if m[0] == "dir" {
				var newDir *Dir = &Dir{
					name:   m[1],
					parent: curDir,
				}
				curDir.dirs = append(curDir.dirs, newDir)
				continue
			}
			// iterating through files
			currentFiles = append(currentFiles, File{
				size: helpers.StrToInt(m[0]),
				name: m[1],
			})
		}
	}
	// find target directories
	dirs := make(map[string]int)
	iterateDirs(dirs, &rootDir)
	for _, v := range dirs {
		if v <= 100000 {
			totalSize += v
		}
	}

	// part 2
	diskSize := 70000000
	unusedSpaceNeeded := 30000000 - (diskSize - rootDir.size)
	targetSize := diskSize
	for _, size := range dirs {
		if size > unusedSpaceNeeded && size < targetSize {
			targetSize = size
		}
	}
	return totalSize, targetSize
}

func iterateDirs(dirs map[string]int, dir *Dir) {
	for _, dir := range dir.dirs {
		dirs[dir.name+dir.parent.name] = dir.size
		iterateDirs(dirs, dir)
	}
}

func calcSize(curDir *Dir) int {
	fileSize := 0
	dirSize := 0
	for _, file := range curDir.files {
		fileSize += file.size
	}
	for _, dir := range curDir.dirs {
		dirSize += calcSize(dir)
	}
	return fileSize + dirSize
}

func updateParent(curDir *Dir) {
	if curDir.parent == nil {
		return
	}
	size := 0
	for _, file := range curDir.parent.files {
		size += file.size
	}
	for _, dir := range curDir.parent.dirs {
		size += dir.size
	}
	curDir.parent.size = size
	updateParent(curDir.parent)
}

func changeDir(name string, curDir *Dir, rootDir *Dir) *Dir {
	if name == "/" {
		return rootDir
	}
	if name == ".." {
		return curDir.parent
	}
	for _, dir := range curDir.dirs {
		if dir.name == name {
			return dir
		}
	}
	return curDir
}

func Day07() {
	start := time.Now()
	one, two := day07OneAndTwo()
	elapsed := time.Since(start)
	fmt.Printf("Day07, part 1: %d, part 2: %d (%s)\n", one, two, elapsed)
}
