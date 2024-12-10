package day9

import (
	"aoc/input"
	"fmt"
)

const day int = 9

type File struct {
	fileId int
	size   int
	blocks []*Block
}

type Block struct {
	blockId int
	file    *File
}

func (b Block) String() string {
	if b.file == nil {
		return fmt.Sprintf("{blockId: %d, fileId: nil}", b.blockId)
	}
	return fmt.Sprintf("{blockId %d:, fileId: %d}", b.blockId, b.file.fileId)
}

func (f File) GetBlockIds() []int {
	blockIds := make([]int, len(f.blocks))
	for i, block := range f.blocks {
		blockIds[i] = block.blockId
	}
	return blockIds
}

func (f File) PrintBlockIds() {
	for _, block := range f.blocks {
		fmt.Printf("%d ", block.blockId)
	}
	fmt.Printf("\n")
}

func (f File) String() string {
	return fmt.Sprintf("{fileId: %d, size: %d }", f.fileId, f.size)
}

func GetCheckSum(blocks []*Block) int {
	total := 0
	for pos, block := range blocks {
		if block.file != nil {
			total += pos * block.file.fileId
		}
	}
	return total
}

func PrintFileSystem(blocks []*Block) {
	for _, b := range blocks {
		if b.file != nil {
			fmt.Printf("%d", b.file.fileId)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Printf("\n")
}

func Part1() int {
	diskContents := input.GetFileContents(day)
	diskSize := len(diskContents)

	files := make([]File, 0)
	blocks := make([]*Block, 0)
	freeBlocks := make([]*Block, 0)

	id := 0
	i := 0
	blockCounter := 0

	for {
		fileSize := input.GetInt(string(diskContents[i]))

		newFile := File{
			fileId: id,
			size:   fileSize,
			blocks: make([]*Block, 0),
		}

		for j := 0; j < newFile.size; j++ {
			fileBlock := Block{
				blockId: blockCounter,
				file:    &newFile,
			}
			blocks = append(blocks, &fileBlock)
			newFile.blocks = append(newFile.blocks, &fileBlock)
			blockCounter += 1
		}

		files = append(files, newFile)

		if i+1 < diskSize {
			for j := 0; j < input.GetInt(string(diskContents[i+1])); j++ {
				freeBlock := Block{
					blockId: blockCounter,
					file:    nil,
				}
				blocks = append(blocks, &freeBlock)
				freeBlocks = append(freeBlocks, &freeBlock)
				blockCounter += 1
			}
		}

		id += 1
		i += 2
		if i > diskSize-1 {
			break
		}
	}

	//PrintFileSystem(blocks)

	// now compact files right to left
	nextFreeBlock := 0

	for i := len(files) - 1; i >= 0; i-- {
		for j := len(files[i].blocks) - 1; j >= 0; j-- {
			if nextFreeBlock == len(freeBlocks) {
				break
			}

			oldBlockId := files[i].blocks[j].blockId

			// don't swap if current block id < free block id
			if oldBlockId < freeBlocks[nextFreeBlock].blockId {
				continue
			}

			freeBlocks[nextFreeBlock].file = &files[i]
			files[i].blocks[j] = freeBlocks[nextFreeBlock]

			blocks[oldBlockId].file = nil

			nextFreeBlock += 1

		}
	}

	//PrintFileSystem(blocks)

	// calculate check sum
	checksum := GetCheckSum(blocks)
	fmt.Printf("Part 1 answer: %d\n", checksum)

	return checksum
}

func Part2() int {
	diskContents := input.GetFileContents(day)
	diskSize := len(diskContents)

	files := make([]File, 0)
	blocks := make([]*Block, 0)

	id := 0
	i := 0
	blockCounter := 0

	for {
		fileSize := input.GetInt(string(diskContents[i]))

		newFile := File{
			fileId: id,
			size:   fileSize,
			blocks: make([]*Block, 0),
		}

		for j := 0; j < newFile.size; j++ {
			fileBlock := Block{
				blockId: blockCounter,
				file:    &newFile,
			}
			blocks = append(blocks, &fileBlock)
			newFile.blocks = append(newFile.blocks, &fileBlock)
			blockCounter += 1
		}

		files = append(files, newFile)

		if i+1 < diskSize {
			emptyBlocks := input.GetInt(string(diskContents[i+1]))
			if emptyBlocks > 0 {
				for j := 0; j < emptyBlocks; j++ {
					freeBlock := Block{
						blockId: blockCounter,
						file:    nil,
					}
					blocks = append(blocks, &freeBlock)
					blockCounter += 1
				}
			}
		}

		id += 1
		i += 2
		if i > diskSize-1 {
			break
		}
	}

	//PrintFileSystem(blocks)

	for i := len(files) - 1; i >= 0; i-- {
		fileSize := files[i].size

		firstEmptyBlock := 0
		spaceFound := 0
		for blockId, block := range blocks {
			if block.file == nil {
				// empty block
				spaceFound += 1
			} else {
				spaceFound = 0
			}

			if spaceFound == fileSize {
				firstEmptyBlock = blockId - fileSize + 1
				break
			}
		}

		// have we found space and is the space before the file?
		if spaceFound == fileSize && firstEmptyBlock < files[i].blocks[0].blockId {
			nextFreeBlock := firstEmptyBlock
			for j := 0; j < fileSize; j++ {
				oldBlockId := files[i].blocks[j].blockId
				blocks[oldBlockId].file = nil
				blocks[nextFreeBlock].file = &files[i]
				files[i].blocks[j] = blocks[nextFreeBlock]
				nextFreeBlock += 1
			}
		}
		//PrintFileSystem(blocks)
	}

	//PrintFileSystem(blocks)

	// calculate check sum
	checksum := GetCheckSum(blocks)
	fmt.Printf("Part 2 answer: %d\n", checksum)

	return checksum
}
