# Design Overview – bw-pwned

This document explains how bw-pwned works under the hood: data flow, architecture, and design choices.

## 1. High-level flow

1. Query local Bitwarden items using the Bitwarden CLI API (`/list/object/items`)
2. For each item:
   - Extract password  
   - SHA-1 hash → uppercase hex  
   - Split hash:
     - prefix = first 5 chars  
     - suffix = rest  
3. Request:

   ```sh
   https://api.pwnedpasswords.com/range/{prefix}
   ```

4. Compare returned suffixes  
5. Record breach count  
6. Return results to CLI layer

## 2. Why k-anonymous?

The HIBP range API sends *only* a hash prefix.  
HIBP returns hundreds of suffixes, making it impossible to determine the original password.  
This keeps the tool privacy-safe.

## 3. Concurrency

A worker pool handles lookups in parallel:

- avoids flooding the HIBP API  
- keeps CPU usage predictable  
- preserves output ordering  

Workers perform:  
hash → prefix/suffix → API fetch → suffix match → result return.

## 4. Caching

bw-pwned caches **prefix responses** from the HIBP range API.

Many different passwords can share the same first 5 SHA‑1 characters, so the tool:

- checks an in-memory cache before making a request  
- reuses the previously-fetched body if the prefix was already seen  
- only queries HIBP when a prefix is new  

This reduces:

- network traffic  
- latency  
- duplicated work for the worker pool  
- the chance of hitting API rate limits  

The cache is in-memory only and lasts for the duration of a single run.  
No data is ever written to disk.

## 5. Error-handling

### Bitwarden errors (blocking)

- BW CLI offline  
- invalid login  
- service unreachable  

→ User must fix the issue before continuing.

### HIBP errors (non-blocking)

- timeouts  
- HTTP errors  
- malformed responses  

→ The specific item is marked as failed, but the program continues safely.

## 6. Data structures

```go
type Result struct {
    Name     string
    Username string
    URL      string
    Count    int
}
```

Separation of concerns:

- `bitwarden/` → BW API  
- `hibp/` → hashing + range lookup  
- `models/` → shared structs and transport types  

## 7. Security Notes

- No passwords ever leave the local machine  
- Only hashed prefixes are sent to HIBP  
- No sensitive data is written to disk  
