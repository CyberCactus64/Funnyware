# Ransomware written in Go, for educational purposes

## 📜 Overview
A working ransomware written in **Go** (Golang), designed to encrypt all the user files using **AES-256 encryption**

### ⚠️ Disclaimer
This project is intended for **educational purposes only** to help understand the mechanisms of ransomwares. 
Misuse of this code for malicious purposes is illegal and unethical!
Test this code only on a virtual machine or on devices that do not belong to others.

---

## 🚀 Features
1. **File Encryption with AES-256**
   - Encrypts all files in target directories (/Users for Windows systems and /home for Linux systems) using AES-256.

2. **Encryption Key Generation**
   - Generates a unique AES-256 encryption key for each victim.
   - The decryption key is then made available only through a payment, as part of the ransom process.
   - The key is sent to a remote server, ensuring it’s inaccessible without fulfilling the ransom demands.

3. **Ransom Note Generation**
   - Creates a text file with details on:
     - Threats and instructions to install the Tor Browser.
     - A reference to a specific .onion site (for the proof of concept, a normal webpage is used) for payment and decryption software.
   
---

## 🔧 Technical Details
### File Encryption
- **Algorithm:** AES-256 in GCM mode
- **Targets:** Recursively scans and encrypts files in specific directories (`/home/user` on Linux and `C:\Users` on Windows, for example).
- **Exclusions:** Skips not important folders and critical system files to avoid rendering the machine unbootable.

### Written in Go
- **Why Go?**
  - Cross-platform compatibility.
  - Strong support for concurrency and networking.
  - Lightweight binaries for easier distribution.

---

## 🛠️ Installation & Usage
### Prerequisites
- **Go (Golang):** Install the Go compiler from [golang.org](https://golang.org).
- **Dependencies:**
  - Standard Go libraries.
- **External Tools:**
 - To change the software icon in Windows: https://github.com/tc-hib/go-winres (```go install github.com/tc-hib/go-winres@latest```):

### Build Instructions
1. Clone the repository:
   ```bash
   git clone https://github.com/CyberCactus64/funny-ware.git
   ```
2. Build the encryptor:
   ```bash
   cd Encryptor
   go-winres simply --icon icon.png
   go build -ldflags="-H windowsgui" -o Minecraft.exe
   ```
2. Build the decryptor:
   ```bash
   cd Decryptor
   go build -ldflags="-H windowsgui" -o HelpMe.exe
   ```

---

## ✨ FUTURE UPDATES:
 - Function to generate a random 32-byte encryption key, to write in the ransom note.
 - Function to securely transmit the encryption key to an external server.
 - Enhance code obfuscation techniques to bypass Microsoft defender.

---

## 🛡️ Legal and Ethical Use
This project must only be used for **research** and **cybersecurity training**. 
Unauthorized deployment of this program is strictly prohibited and punishable under applicable laws.
I am not responsible for any malicious use of this project. This code is intended solely for educational purposes.

---

### 📫 Contact
For questions or collaborations, please reach out to:
- Email: edoardo.enricomaria.fornasier@gmail.com
- GitHub: [CyberCactus64](https://github.com/CyberCactus64)
