package autoimport

// func processFile(path string) error {
// 	// Read file contents
// 	file, err := os.OpenFile(path, os.O_RDWR, 0644)

// 	if err != nil {
// 		return err
// 	}

// 	defer file.Close()
// 	code, err := ioutil.ReadAll(file)

// 	if err != nil {
// 		return err
// 	}

// 	// Parse contents
// 	importPaths, err := parse(code, path)

// 	if err != nil {
// 		return err
// 	}

// 	if len(importPaths) == 0 {
// 		return nil
// 	}

// 	// Find package definition
// 	packagePos := bytes.Index(code, []byte("package "))

// 	if packagePos == -1 {
// 		return errors.New("Package definition missing")
// 	}

// 	seekPos := int64(0)

// 	for i := packagePos; i < len(code); i++ {
// 		if code[i] == '\n' {
// 			seekPos = int64(i + 1)
// 			break
// 		}
// 	}

// 	// Seek to the beginning (after the package line)
// 	file.Seek(seekPos, 0)

// 	// importCommand := fmt.Sprintf("\nimport (\n\t\"%s\"\n)\n\n", strings.Join(importPaths, "\"\n\t\""))
// 	// file.WriteString(importCommand)
// 	// file.Write(code[seekPos:])
// 	// file.Sync()
// 	fmt.Println(importPaths)

// 	return nil
// }
