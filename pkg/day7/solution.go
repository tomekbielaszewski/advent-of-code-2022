package day1

import (
	"advent-of-code-2022/pkg"
	"strings"
)

func solution(lines []string) int {
	root := loadFileSystem(lines)
	calcSizeOfDirectories(root)
	return sumAllDirectoriesWithSizeLessThan(100000, root)
}

func solution2(lines []string) int {
	root := loadFileSystem(lines)
	calcSizeOfDirectories(root)

	spaceRequired := root.size - (70000000 - 30000000)
	return findTheSmallestEnoughDirectoryToDelete(spaceRequired, root)
}

func loadFileSystem(lines []string) *file {
	root := &file{
		isDir:    true,
		size:     0,
		children: make(map[string]*file),
		parent:   nil,
	}
	currentFile := root
	for i, line := range lines {
		if i == 0 {
			continue
		}
		if []rune(line)[0] == '$' {
			split := strings.Split(line, " ")
			switch split[1] {
			case "cd":
				target := split[2]
				if target == ".." {
					currentFile = currentFile.parent
				} else {
					currentFile = currentFile.children[target]
				}
				break
			case "ls":
				currentFile.ls(lines[i+1:])
			}
		}
	}
	return root
}

func calcSizeOfDirectories(root *file) {
	size := 0
	for _, f := range root.children {
		if f.isDir {
			calcSizeOfDirectories(f)
		}
		size += f.size
	}
	root.size = size
}

func sumAllDirectoriesWithSizeLessThan(maxSize int, root *file) int {
	sum := 0
	for _, f := range root.children {
		if f.isDir {
			if f.size <= maxSize {
				sum += f.size
			}
			sum += sumAllDirectoriesWithSizeLessThan(maxSize, f)
		}
	}
	return sum
}

func findTheSmallestEnoughDirectoryToDelete(spaceRequired int, root *file) int {
	minEnough := root.size
	for _, f := range root.children {
		if f.isDir {
			if f.size < minEnough && f.size >= spaceRequired {
				minEnough = f.size
			}
			childMinEnough := findTheSmallestEnoughDirectoryToDelete(spaceRequired, f)
			if childMinEnough < minEnough && f.size >= spaceRequired {
				minEnough = childMinEnough
			}
		}
	}
	return minEnough
}

type file struct {
	isDir    bool
	size     int
	children map[string]*file
	parent   *file
}

func (f *file) cd(fileName string) *file {
	return f.children[fileName]
}

func (f *file) ls(lines []string) {
	for _, line := range lines {
		if []rune(line)[0] == '$' {
			return
		}
		split := strings.Split(line, " ")
		if split[0] == "dir" {
			dir := &file{
				isDir:    true,
				size:     0,
				children: make(map[string]*file),
				parent:   f,
			}
			f.children[split[1]] = dir
			continue
		}

		childFile := &file{
			isDir:    false,
			size:     pkg.StringToIntNoErr(split[0]),
			children: make(map[string]*file),
			parent:   f,
		}
		f.children[split[1]] = childFile
	}
}
