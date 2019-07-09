package factoryMethod

type FileSystem interface {
	Open(string) string
}

const (
	default_ = 1 << iota
	main_
	block
)

type DefaultFileSystem struct{}

func (dfs DefaultFileSystem) Open(string) string {
	return "Default File System"
}

type MainFileSystem struct{}

func (mfs MainFileSystem) Open(string) string {
	return "Main File System"
}

type BlockFileSystem struct{}

func (mfs BlockFileSystem) Open(string) string {
	return "Block File System"
}
func NewStore(t int) FileSystem {
	var r FileSystem
	switch t {
	case default_:
		r = new(DefaultFileSystem)
	case main_:
		r = new(MainFileSystem)
	case block:
		r = new(BlockFileSystem)
	}
	return r
}
