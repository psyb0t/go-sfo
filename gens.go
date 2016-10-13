package sfo

import (
    "os"
    "path/filepath"
)

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
