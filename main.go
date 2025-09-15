package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

type Note struct {
	Name    string
	Content string
}

func createNote(path string, note *Note) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(note.Content)
	return err
}
func readNote(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
func listNote(dir string) ([]string, error) {
	files := []string{}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, val := range entries {
		if !val.IsDir() {
			files = append(files, val.Name())
		}
	}
	return files, nil
}
func deleteNote(path string) error {
	return os.Remove(path)
}
func updateNote(path string, extra string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString("\n" + extra)
	return err
}
func main() {
	baseDir := "notes"
	os.MkdirAll(baseDir, 0755)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n--- File Manager Menu ---")
		fmt.Println("1. Create note")
		fmt.Println("2. Read note")
		fmt.Println("3. List notes")
		fmt.Println("4. Update note")
		fmt.Println("5. Delete note")
		fmt.Println("0. Exist")
		fmt.Println("Choose: ")

		scanner.Scan()
		choice := scanner.Text()
		switch choice {
		case "1":
			{
				fmt.Print("Note name: ")
				scanner.Scan()
				name := scanner.Text()
				fmt.Print("Content: ")
				scanner.Scan()
				content := scanner.Text()
				note := &Note{Name: name, Content: content}
				err := createNote(filepath.Join(baseDir, name+".txt"), note)
				if err != nil {
					panic(err)
				} else {
					fmt.Println("Note Created.")
				}
			}
		case "2":
			{

				fmt.Print("Note name to read: ")
				scanner.Scan()
				name := scanner.Text()
				content, err := readNote(filepath.Join(baseDir, name+".txt"))
				if err != nil {
					panic(err)
				} else {
					fmt.Println("Note Content: \n", content)
				}
			}
		case "3":
			{

				files, err := listNote(baseDir)
				if err != nil {
					panic(err)
				} else {
					fmt.Println("Notes: \n", files)
				}
			}

		case "4":
			{
				fmt.Print("Note name to update: ")
				scanner.Scan()
				name := scanner.Text()
				fmt.Print("Content to replace: ")
				scanner.Scan()
				content := scanner.Text()
				err := updateNote(filepath.Join(baseDir, name+".txt"), content)
				if err != nil {
					panic(err)
				} else {
					fmt.Println("Note Updated.")
				}
			}
		case "5":
			{
				fmt.Print("Note name to delete: ")
				scanner.Scan()
				name := scanner.Text()
				err := deleteNote(filepath.Join(baseDir, name+".txt"))
				if err != nil {
					panic(err)
				} else {
					fmt.Println("Note Deleted.")
				}
			}
		case "0":
			{
				fmt.Print("Exiting... ")
				return
			}
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
