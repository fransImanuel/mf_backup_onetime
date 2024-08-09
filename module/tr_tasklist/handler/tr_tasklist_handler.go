package handler

import (
	"archive/zip"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"mf_backup_onetime/dto"
	"mf_backup_onetime/module/tr_tasklist"
	"os"
	"path/filepath"
	"sync"

	log "github.com/sirupsen/logrus"
)

type TRTasklistHandler struct {
	TRTasklistService tr_tasklist.Service
}

func InitTRTasklistHandler(tr_tasklist tr_tasklist.Service) {
	handler := &TRTasklistHandler{
		TRTasklistService: tr_tasklist,
	}

	handler.ExportPDFTasklistBulk()

}

func (controller *TRTasklistHandler) ExportPDFTasklistBulk() {
	// panic(1)
	IDs := controller.TRTasklistService.BulkExportOneTimeService()
	var wg sync.WaitGroup

	errCh := make(chan error, 1)
	ctxG, cancel := context.WithCancel(context.Background())
	defer cancel()

	// panic(len(IDs))
	fmt.Println("len(IDs): ", len(IDs))
	wg.Add(len(IDs))
	for _, ID := range IDs {
		go func(id string) {

			defer wg.Done()

			select {
			case <-ctxG.Done():
				return
			default:
				Base64, fname, errSrv := controller.TRTasklistService.ExportPDFTasklist(context.Background(), dto.GetTasklistByID{TasklistId: id})

				// fmt.Println("Base64: ", *Base64)
				cancel() // Cancel all other goroutines
				return

				if errSrv.Message != "" {
					errCh <- errors.New(errSrv.Message)
					cancel() // Cancel all other goroutines
					return
				}
				dec, errFile := base64.StdEncoding.DecodeString(*Base64)
				if errFile != nil {
					errCh <- errFile
					cancel() // Cancel all other goroutines
					return
				}

				f, err := os.Create(fname)
				if err != nil {
					errCh <- err
					cancel() // Cancel all other goroutines
					return
				}
				defer f.Close()

				if _, err := f.Write(dec); err != nil {
					errCh <- err
					cancel() // Cancel all other goroutines
					return
				}
				if err := f.Sync(); err != nil {
					errCh <- err
					cancel() // Cancel all other goroutines
					return
				}
			}

		}(ID)

		// Base64, fname, errSrv := controller.TRTasklistService.ExportPDFTasklist(context.Background(), dto.GetTasklistByID{TasklistId: ID})
		// if errSrv.Message != "" {
		// 	fmt.Println("err: ", errSrv.Message)
		// }
		// dec, errFile := base64.StdEncoding.DecodeString(*Base64)
		// if errFile != nil {
		// 	// panic(errFile)
		// 	fmt.Println("err: ", errFile)
		// 	return
		// }

		// f, err := os.Create(fname)
		// if err != nil {
		// 	fmt.Println("err: ", err)
		// 	return
		// }
		// defer f.Close()

		// if _, err := f.Write(dec); err != nil {
		// 	fmt.Println("err: ", err)
		// 	return
		// }
		// if err := f.Sync(); err != nil {
		// 	fmt.Println("err: ", err)
		// 	return
		// }

	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			log.Printf("Encountered error: %v", err)
			return
		}
	}

	log.Println("All tasks completed successfully")
	fmt.Println("len(IDs): ", len(IDs))
	// panic(1)

	// ZIP PROGRESS

	// Define the folder to zip and the output zip file
	folderToZip := "storage"              // The folder you want to zip
	outputZipFile := "MF_TASKLIST-05.zip" // The name of the output zip file
	// outputZipFile := "MF_TASKLIST-06.zip" // The name of the output zip file
	err := ZipFolder(folderToZip, outputZipFile)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Folder successfully zipped to", outputZipFile)
	}

	// ctx.FileAttachment("./"+outputZipFile, outputZipFile)
}

// AddFileToZip adds a file to the zip archive
func AddFileToZip(zipWriter *zip.Writer, basePath, filePath string) error {
	// Open the file to add to the zip archive
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	// Get the relative path of the file
	relPath, err := filepath.Rel(basePath, filePath)
	if err != nil {
		return err
	}

	// Create a new zip file header
	header, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return err
	}

	// Set the relative path for the header
	header.Name = relPath
	header.Method = zip.Deflate

	// Create a writer for the file header
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	// Copy the file content to the zip writer
	_, err = io.Copy(writer, file)
	if err != nil {
		return err
	}

	return nil
}

// ZipFolder zips the specified folder and its contents
func ZipFolder(folderToZip, outputZipFile string) error {
	// Create a new zip file
	zipFile, err := os.Create(outputZipFile)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %w", err)
	}
	defer zipFile.Close()

	// Create a new zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Walk through the directory tree and add files to the zip archive
	err = filepath.WalkDir(folderToZip, func(filePath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			return AddFileToZip(zipWriter, folderToZip, filePath)
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("error walking the path: %w", err)
	}

	return nil
}
