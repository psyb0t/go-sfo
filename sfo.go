// Golang Simplified File Operations
package sfo

import (
    "os"
    "io"
    "bytes"
    "path/filepath"
)

type File struct {
    Bytes []byte
    Text string
}

func ReadFile(file_path *string) (*File, error) {
    exists, err := PathExists(file_path)

    if err != nil {
        return nil, err
    }

    if !exists {
        return &File{}, nil
    }

    f, err := os.OpenFile(*file_path, os.O_RDONLY, 0644)
    defer f.Close()

    if err != nil {
        return nil, err
    }

    file_bytes := make([]byte, 1024)
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

func CreateFile(file_path *string) error {
    exists, err := PathExists(file_path)

    if err != nil {
        return err
    }

    if exists {
        return nil
    }

    dir := filepath.Dir(*file_path)

    exists, err = PathExists(&dir)

    if err != nil {
        return err
    }

    if !exists {
        err = os.MkdirAll(dir, 0644)

        if err != nil {
            return err
        }
    }

    file, err := os.Create(*file_path)
    defer file.Close()

    if err != nil {
        return err
    }

    return nil
}

func WriteStringToFile(file_path *string, text *string) error {
    file, err := ReadyFile(file_path)

    if err != nil {
        return err
    }

    _, err = file.WriteString(*text)

    if err != nil {
        return err
    }

    err = file.Sync()

    if err != nil {
        return err
    }

    return nil
}

func WriteBytesToFile(file_path *string, bytes *[]byte) error {
    file, err := ReadyFile(file_path)

    if err != nil {
        return err
    }

    _, err = file.Write(*bytes)

    if err != nil {
        return err
    }

    err = file.Sync()

    if err != nil {
        return err
    }

    return nil
}

func PathExists(path *string) (bool, error) {
    _, err := os.Stat(*path)

    if os.IsNotExist(err) {
        return false, nil
    }

    if err != nil {
        return false, err
    }

    return true, nil
}

func ReadyFile(file_path *string) (*os.File, error) {
    err := CreateFile(file_path)

    if err != nil {
        return nil, err
    }

    file, err := os.OpenFile(*file_path, os.O_WRONLY, 0644)

    if err != nil {
        return nil, err
    }

    return file, nil
}
