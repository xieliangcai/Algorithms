package fileSystemInterface

type FileSystemNode interface {
	CountNumOfFiles() int
	CountSizeOfFiles() int
	GetPath() string
	RemoveSubNode(node FileSystemNode)
	AddSubNode(node FileSystemNode)
}

