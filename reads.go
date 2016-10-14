package sfo

import (
    "os"
    "io"
    "bytes"
)

func ReadFile(file_path *string) (*File, error) {
    exists, err := PathExists(file_path)

    if err != nil {
        return nil, err
    }

    if !exists {
        return &File{}, nil
    }

    f, err := os.OpenFile(*file_path, os.O_RDONLY, 0644)

    if err != nil {
        return nil, err
    }

    fstat, err := f.Stat()

    if err != nil {
        return nil, err
    }

    file_bytes :=  make([]byte, fstat.Size())
    for {
        _, err = f.Read(file_bytes)

        if err == io.EOF {
            break
        }

        if err != nil {
            return nil, err
        }
    }

    file_bytes = bytes.Trim(file_bytes, "\x00")

    file := &File{
        Bytes: file_bytes,
        Text: string(file_bytes),
    }

    return file, nil
}
