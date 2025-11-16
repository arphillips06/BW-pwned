# BW-pwned

BW-pwned checks the passwords in your Bitwarden vault against the **Have I Been Pwned** (HIBP) *Pwned Passwords* database using the anonymous **k-anonymous range API**.  
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

- **Bitwarden CLI**  
  Download: [Bitwarden CLI](https://bitwarden.com/help/cli/#get)
- **Go 1.24.x** (tested)
After cloning this repo:  

  ```sh
  go mod init bw-hibp-check
  go mod tidy
  ```

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

## Running the app

No need to build — just run it:

```sh
go run .
```

You should see:

```sh
Bitwarden → HIBP checker starting...
1. Check status
2. Get single item
3. List all items
Choose an option [1-3]:
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
