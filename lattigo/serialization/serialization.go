package serialization

import (
	"encoding"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/google/go-cmp/cmp"
)

var (
	Path = "data/"
)

func CleanFolder(path string) (err error) {
	folder, err := os.Open(path)
	if err != nil {
		return err
	}
	files, err := folder.Readdir(0)
	if err != nil {
		return err
	}
	for i := range files {
		if files[i].Name() != "donotremove.txt" {
			if err := os.Remove(path + files[i].Name()); err != nil {
				log.Println(err)
			}
		}
	}
	return
}

// TestSerialization tests that the serialization to a file of an
// object works properly.
func TestSerialization[T any](object *T, path string) (err error) {

	if err = Serialize[T](object, path); err != nil {
		return fmt.Errorf("Serialize(%T): %w", object, err)
	}

	objectCopy := new(T)

	if err = Deserialize[T](objectCopy, path); err != nil {
		return fmt.Errorf("Deserialize(%T): %w", object, err)
	}

	if err = os.Remove(path); err != nil {
		return fmt.Errorf("os.Remove(%s): %w", path, err)
	}

	if !cmp.Equal(object, objectCopy) {
		return fmt.Errorf("%T != %T", object, objectCopy)
	}

	return
}

// Serialize serializes an object implementing [io.WriterTo]
// or [encoding.BinaryMarshaler] to file specified by the path.
func Serialize[T any](object *T, path string) (err error) {

	var f *os.File
	// #nosec - G304
	if f, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600); err != nil {
		return fmt.Errorf("os.OpenFile(%s): %w", path, err)
	}

	defer f.Close()

	switch object := any(object).(type) {
	case io.WriterTo:

		if _, err = object.WriteTo(f); err != nil {
			return fmt.Errorf("%T.WriteTo: %w", object, err)
		}

	case encoding.BinaryMarshaler:

		var data []byte
		if data, err = object.MarshalBinary(); err != nil {
			return fmt.Errorf("%T.MarshalBinary: %w", object, err)
		}

		if _, err = f.Write(data); err != nil {
			return fmt.Errorf("f.Write: %w", err)
		}

	default:
		return fmt.Errorf("invalid input: should implement io.WriterTo or encoding.BinaryMarshaler")
	}

	return
}

// Deserialize deserializes an [*object] implementing [io.ReaderFrom]
// or [encoding.BinaryUnmarshaler] from file specified by the path.
func Deserialize[T any](object *T, path string) (err error) {

	switch object := any(object).(type) {
	case io.ReaderFrom:

		var f *os.File
		// #nosec - G304
		if f, err = os.OpenFile(path, os.O_RDONLY, 0600); err != nil {
			return fmt.Errorf("os.OpenFile(%s): %w", path, err)
		}

		defer f.Close()

		if _, err = object.ReadFrom(f); err != nil {
			return fmt.Errorf("%T.ReadFrom: %w", object, err)
		}

	case encoding.BinaryUnmarshaler:

		var data []byte
		// #nosec - G304
		if data, err = os.ReadFile(path); err != nil {
			return fmt.Errorf("os.ReadFile(%s): %w", path, err)
		}

		if err = object.UnmarshalBinary(data); err != nil {
			return fmt.Errorf("%T.UnmarshalBinary: %w", object, err)
		}

	default:
		return fmt.Errorf("[%T] does not implement io.ReaderFrom or encoding.BinaryUnmarshaler", object)
	}

	return
}
