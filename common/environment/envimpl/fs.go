package envimpl

import (
	"github.com/v4fly/v4ray-core/v0/common/environment"
	"github.com/v4fly/v4ray-core/v0/common/platform/filesystem"
	"github.com/v4fly/v4ray-core/v0/common/platform/filesystem/fsifce"
)

type fileSystemDefaultImpl struct{}

func (f fileSystemDefaultImpl) ReadDir() fsifce.FileReadDirFunc {
	return filesystem.NewFileReadDir
}

func (f fileSystemDefaultImpl) RemoveFile() fsifce.FileRemoveFunc {
	return filesystem.NewFileRemover
}

func (f fileSystemDefaultImpl) OpenFileForReadSeek() fsifce.FileSeekerFunc {
	return filesystem.NewFileSeeker
}

func (f fileSystemDefaultImpl) OpenFileForRead() fsifce.FileReaderFunc {
	return filesystem.NewFileReader
}

func (f fileSystemDefaultImpl) OpenFileForWrite() fsifce.FileWriterFunc {
	return filesystem.NewFileWriter
}

func NewDefaultFileSystemDefaultImpl() environment.FileSystemCapabilitySet {
	return fileSystemDefaultImpl{}
}
