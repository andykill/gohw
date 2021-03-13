package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	out, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	fileInfo, err := out.Stat()
	if err != nil {
		return err
	}
	if fileSize := fileInfo.Size(); offset > fileSize {
		return ErrOffsetExceedsFileSize
	}

	in, err := os.Create(toPath)
	if err != nil {
		return err
	}

	if limit > fileInfo.Size() {
		limit = fileInfo.Size()
	}
	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(out)

	_, err = io.Copy(in, barReader)
	if err != nil {
		return err
	}
	bar.Finish()
	return nil
}
