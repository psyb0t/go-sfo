// Golang Simplified File Operations
package sfo

import (
    "os"
    "io"
)

type File struct {
    Bytes []byte
    Text string
}

func ReadFile(file_path *string) (*File, error) {
    f, err := os.OpenFile(*file_path, os.O_RDONLY, 0644)
    defer f.Close()

    if err != nil {
        return nil, err
    }

    bytes := make([]byte, 1024)
    for {
        part, err := f.Read(bytes)

        if err != io.EOF {
            return nil, err
        }

        if part == 0 {
            break
        }
    }

    file := &File{
        Bytes: bytes,
        Text: string(bytes),
    }

    return file, nil
}
