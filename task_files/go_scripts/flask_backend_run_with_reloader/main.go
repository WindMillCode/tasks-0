package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"main/shared"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/windmillcode/go_cli_scripts/v4/utils"
)

var (
	cmd          *exec.Cmd
	restartTimer *time.Timer
	lastMD5Sum   = make(map[string][md5.Size]byte)
	debounceTime = 500 * time.Millisecond // Adjust debounce time as needed
)

func main() {
	flaskAppFolder, err := shared.SetupEnvironmentToRunFlaskApp()
	if err != nil {
		log.Fatalf("Error during setup: %v", err)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	setupFileWatcher(flaskAppFolder, watcher, done)

	restartServer(flaskAppFolder) // Initial start

	// Handle SIGINT and SIGTERM.
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for sig := range signalChan {
			log.Printf("Received %v, exiting...", sig)
			done <- true
		}
	}()

	<-done
}

func setupFileWatcher(root string, watcher *fsnotify.Watcher, done chan bool) {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Failed to traverse Flask app folder: %v", err)
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				handleEvent(event, root)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Printf("Watcher error: %v", err)
			}
		}
	}()
}

func handleEvent(event fsnotify.Event, location string) {

	if event.Op&(fsnotify.Write|fsnotify.Rename|fsnotify.Create) != 0 {
		log.Printf("Detected change: %s", event)
		if shouldRestart(event.Name) {
			if restartTimer != nil {
				restartTimer.Stop()
			}
			restartTimer = time.AfterFunc(debounceTime, func() {
				restartServer(location)
			})
		}
	}
}

func restartServer(location string) {
	utils.CDToLocation(location)

	// Kill the existing process if it's running
	if cmd != nil && cmd.Process != nil {
		// Check if the process has not already exited
		if cmd.ProcessState == nil || !cmd.ProcessState.Exited() {
			err := cmd.Process.Kill()
			if err != nil {
				panic(fmt.Sprintf("Failed to kill server process: %v", err))
			}
		}
	}

	// Start a new process
	cmd = exec.Command("python", "app.py")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		log.Fatalf("Failed to start Flask server: %v", err)
		return
	}

	go func() {
		if err := cmd.Wait(); err != nil {
			log.Printf("Flask server process exited with error: %v", err)
		}
	}()
}

func shouldRestart(filePath string) bool {
	currentMD5, err := fileMD5(filePath)
	if err != nil {
		log.Printf("Error calculating MD5 for %s: %v", filePath, err)
		return false
	}

	if lastMD5, ok := lastMD5Sum[filePath]; ok && currentMD5 == lastMD5 {
		return false // MD5 is the same, no need to restart
	}

	lastMD5Sum[filePath] = currentMD5
	return true
}

func fileMD5(filePath string) ([md5.Size]byte, error) {
	var md5Sum [md5.Size]byte
	file, err := os.ReadFile(filePath)
	if err != nil {
		return md5Sum, err
	}
	md5Sum = md5.Sum(file)
	return md5Sum, nil
}
