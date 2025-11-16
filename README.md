# bw-pwned

bw-pwned checks the passwords in your Bitwarden vault against the **Have I Been Pwned** (HIBP) *Pwned Passwords* database using the anonymous **k-anonymous range API**.  
No passwords are ever sent over the network — only hashed prefixes.

## Features

- Reads login items from the local Bitwarden CLI API  
- Hashes each password (SHA-1 → upper case)  
- Queries the HIBP `/range/{prefix}` endpoint  
- Compares suffixes and shows breach counts  
- Supports:
  - checking a single item  
  - listing all items  
  - viewing Bitwarden status  
- Safe/unbreached items are hidden to keep output clean  
- No password data is stored

## Prerequisites

### **Bitwarden CLI**

Download: [Bitwarden CLI](https://bitwarden.com/help/cli/#get)

### **Go 1.24.x or newer**

Required for module installation and compiling from source.
[Go Install](https://go.dev/doc/install)

### **Internet access**

Needed for HIBP lookups.

- **Internet access** (for HIBP lookups)

## Bitwarden CLI Setup

This app relies on the *local* Bitwarden CLI API.  
You must log in and start the local service.

On Windows:

1. Download + unzip the Bitwarden CLI  
2. In the extracted folder:

   ```powershell
   .\bw.exe login
   .\bw.exe serve
   ```

3. (Optional) Check login status:

   ```powershell
   .\bw.exe status
   ```

4. If you’ve added items and want to re-check:

   ```powershell
   .\bw.exe sync
   ```

## Installation Options

You can use **either**:

### **1. Install directly**

Requires **Go 1.24.x** or newer:

``` sh
go install github.com/arphillips06/bw-pwned@latest
```

This places the binary in:

- **Windows:** `%USERPROFILE%\go\bin`\

Then run:

``` sh
bw-pwned
```

------------------------------------------------------------------------

### **2. Run from source**

``` sh
git clone https://github.com/arphillips06/bw-pwned
cd bw-pwned
go run .
```

Follow the on-screen prompts.

## Project Layout

```sh
/bitwarden       BW CLI API client logic  
/hibp            SHA-1 hashing + HIBP range lookup  
/models          Shared data structures  
main.go          CLI menu + program flow
```

## License

MIT — see LICENSE file.
