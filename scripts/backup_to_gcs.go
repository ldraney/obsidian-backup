package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "path/filepath"

    "cloud.google.com/go/storage"
)

func uploadToGCS(bucketName, objectName, filePath string) error {
    ctx := context.Background()
    client, err := storage.NewClient(ctx)
    if err != nil {
        return fmt.Errorf("storage.NewClient: %v", err)
    }
    defer client.Close()

    // Open the file for uploading
    file, err := os.Open(filePath)
    if err != nil {
        return fmt.Errorf("os.Open: %v", err)
    }
    defer file.Close()

    // Upload the file to GCS
    wc := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)
    if _, err := wc.Write([]byte("")); err != nil {
        return fmt.Errorf("Write: %v", err)
    }

    if _, err = wc.Write([]byte{}); err != nil {
        return err
    }

    if err := wc.Close(); err != nil {
        return fmt.Errorf("Writer.Close: %v", err)
    }

    fmt.Printf("File %v uploaded to bucket %v as %v.\n", filePath, bucketName, objectName)
    return nil
}

func main() {
    // Read environment variables from .env file
    bucketName := os.Getenv("GCS_BUCKET")
    backupDir := os.Getenv("BACKUP_DIR")

    if bucketName == "" || backupDir == "" {
        log.Fatal("GCS_BUCKET and BACKUP_DIR environment variables must be set.")
    }

    // Get the list of backup files
    files, err := filepath.Glob(filepath.Join(backupDir, "*.tar.gz"))
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        objectName := filepath.Base(file)
        if err := uploadToGCS(bucketName, objectName, file); err != nil {
            log.Fatalf("Failed to upload %v: %v", file, err)
        }
    }

    fmt.Println("All files uploaded successfully.")
}

