# Contributing to bw-pwned

Thanks for your interest in contributing! Even small improvements — fixing typos, cleaning up error messages, simplifying functions — are welcome.

## How to Contribute

### 1. Fork the repository

```sh
git clone https://github.com/yourname/bw-pwned.git
cd bw-pwned
```

### 2. Create a feature branch

```sh
git checkout -b feature/my-change
```

### 3. Make your changes

Keep contributions small and focused when possible.  
Avoid committing secrets or any Bitwarden data.

### 4. Run the app locally

Ensure Bitwarden CLI is logged in and serving:

```sh
bw login
bw serve
go run .
```

### 5. Submit a Pull Request

Please:

- describe the motivation for your change  
- keep PRs focused on a single topic  
- include expected behaviour if you modify output  
- add small tests if introducing new logic  

Maintainers may suggest adjustments.

---

## Coding Guidelines

### Go Style

- Follow standard Go idioms  
- Favour clarity over cleverness  
- Avoid unnecessary global state  
- Handle errors explicitly  
- Keep functions small and testable  

### Security

Because this tool interacts with password data:

- do not log passwords  
- do not store passwords  
- treat all vault responses carefully  

If you discover a security issue, contact the maintainer privately rather than opening a public issue.

---

## Issues and Feature Requests

If you encounter a bug or have a feature idea:

- open an issue  
- include reproduction steps  
- include your Go version and OS  
- include any error messages  

Feature requests should describe the problem or use case.

---

## Thank You

bw-pwned is a small project, and contributions genuinely help keep it improving. Thank you for taking the time to contribute!
