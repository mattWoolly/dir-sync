package main

import (
    "fmt"
    "os"
    "path/filepath"
    "io/ioutil"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: sync_dirs <dir1> <dir2>")
        os.Exit(1)
    }

    dir1 := os.Args[1]
    dir2 := os.Args[2]

    if err := syncDirectories(dir1, dir2); err != nil {
        fmt.Printf("Error syncing directories: %v\n", err)
        os.Exit(1)
    }
}

func syncDirectories(dir1, dir2 string) error {
    entriesDir1, err := readDirectory(dir1)
    if err != nil {
        return err
    }

    entriesDir2, err := readDirectory(dir2)
    if err != nil {
        return err
    }

    return syncUniqueFiles(entriesDir1, entriesDir2, dir1, dir2)
}

func syncUniqueFiles(entriesDir1, entriesDir2 map[string]struct{}, dir1, dir2 string) error {
    for file := range entriesDir1 {
        if _, exists := entriesDir2[file]; !exists {
            src := filepath.Join(dir1, file)
            dest := filepath.Join(dir2, file)
            if err := copyFile(src, dest); err != nil {
                return err
            }
        }
    }

    for file := range entriesDir2 {
        if _, exists := entriesDir1[file]; !exists {
            src := filepath.Join(dir2, file)
            dest := filepath.Join(dir1, file)
            if err := copyFile(src, dest); err != nil {
                return err
            }
        }
    }

    return nil
}

func readDirectory(dir string) (map[string]struct{}, error) {
    entries := make(map[string]struct{})
    err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !info.IsDir() {
            relPath, err := filepath.Rel(dir, path)
            if err != nil {
                return err
            }
            entries[relPath] = struct{}{}
        }
        return nil
    })
    return entries, err
}

func copyFile(src, dest string) error {
    input, err := ioutil.ReadFile(src)
    if err != nil {
        return err
    }

    destDir := filepath.Dir(dest)
    if err := os.MkdirAll(destDir, 0755); err != nil {
        return err
    }

    if err := ioutil.WriteFile(dest, input, 0644); err != nil {
        return err
    }

    return nil
}
