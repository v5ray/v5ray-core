package filesystemcap

import "github.com/v4fly/v4ray-core/v0/common/platform/filesystem/fsifce"

type FileSystemCapabilitySet interface {
	OpenFileForReadSeek() fsifce.FileSeekerFunc
	OpenFileForRead() fsifce.FileReaderFunc
	OpenFileForWrite() fsifce.FileWriterFunc
	ReadDir() fsifce.FileReadDirFunc
	RemoveFile() fsifce.FileRemoveFunc
}
