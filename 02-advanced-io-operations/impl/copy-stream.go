package impl

import "io"

func CopyNOffset(dst io.Writer, src io.ReadSeeker, offset, length int64) (int64, error) {
	if _, err := src.Seek(offset, io.SeekStart); err != nil {
		return 0, err
	}

	return io.CopyN(dst, src, length)
}
