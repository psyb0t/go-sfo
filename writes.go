package sfo

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
