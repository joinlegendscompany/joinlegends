package stream

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func VideoStream(c *fiber.Ctx) error {
	raw := c.Params("filename")

	filename, err := url.PathUnescape(raw)
	if err != nil {
		return fiber.ErrBadRequest
	}

	path := filepath.Join("/shared/videos", filename)

	file, err := os.Open(path)
	if err != nil {
		return fiber.ErrNotFound
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return fiber.ErrInternalServerError
	}
	size := stat.Size()

	rangeHeader := c.Get("Range")
	if rangeHeader == "" {
		c.Set("Content-Type", "video/mp4")
		c.Set("Content-Length", strconv.FormatInt(size, 10))
		c.Set("Accept-Ranges", "bytes")
		return c.SendFile(path)
	}

	start, end, err := parseRange(rangeHeader, size)
	if err != nil {
		return fiber.ErrRequestedRangeNotSatisfiable
	}

	chunkSize := end - start + 1

	c.Status(fiber.StatusPartialContent)
	c.Set("Content-Type", "video/mp4")
	c.Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, size))
	c.Set("Accept-Ranges", "bytes")
	c.Set("Content-Length", strconv.FormatInt(chunkSize, 10))

	_, err = file.Seek(start, io.SeekStart)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	_, err = io.CopyN(c, file, chunkSize)
	return err
}

func parseRange(header string, size int64) (int64, int64, error) {
	if !strings.HasPrefix(header, "bytes=") {
		return 0, 0, errors.New("invalid range")
	}

	parts := strings.Split(strings.TrimPrefix(header, "bytes="), "-")
	start, _ := strconv.ParseInt(parts[0], 10, 64)

	var end int64
	if parts[1] == "" {
		end = size - 1
	} else {
		end, _ = strconv.ParseInt(parts[1], 10, 64)
	}

	const MAX_CHUNK_SIZE int64 = 1 * 1024 * 1024 // 1 MB

	if start > end || end >= size {
		return 0, 0, errors.New("range out of bounds")
	}

	return start, end, nil
}
