# BW-pwned

This project interacts with the bitwarden and haveibeenpwned APIs to automatically check passwords that have been pwned

## Use

### Prereqresites

- BW CLI
- Go version go1.24.4 (tested)
  - once installed, navigate to this repo's folder and run `go mod init` && `go mod tidy` to init Go and get required modules.
- Internet access (required for the HIBP API)

### Bitwarden CLI

This application currently uses the local Bitwarden CLI API, so the Bitwarden CLI must be installed.
You can download it from: [bitwarden](https://bitwarden.com/help/cli/#get)

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

### Running Go app

Building is optional — the app can be run directly:

`go run .`

Run this from the directory containing main.go.
Follow the on-screen prompts.
