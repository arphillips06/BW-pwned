# BW-pwned
This project interacts with the bitwarden and haveibeenpwned APIs to automatically check passwords that have been pwned

## Use

### Prereqresites
- BW CLI
- Go version go1.24.4 (tested)
- Internet access (required for the HIBP API)

### use 

This application currently uses the local Bitwarden CLI API, so the Bitwarden CLI must be installed.
You can download it from: https://bitwarden.com/help/cli/#get

A quick setup guide for Windows:

Details on how to use this can be found on that site, as a quick run down on Windows

- Download and unzip the Bitwarden CLI.
- Open a terminal in the unzipped folder.
- Run: 
    - `.\bw.exe login`
    - `.\bw.eve serve`
- (Optional) Verify login status
    - `.\bw.eve status`

If items are added to the vault and you wish to check again then you must run `.\bw.exe sync` to update the BW API.


Building is optional — the app can be run directly:

`go run .`

Run this from the directory containing main.go.
Follow the on-screen prompts.
