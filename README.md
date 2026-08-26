# totp

Simple CLI demonstration for Time-based One-Time Passwords (TOTP).

## Requirements

- Go 1.18+

## Build

```bash
go build ./...
```

## Usage

Generate a setup URI:

```bash
go run . setup --issuer Provider --account user@example.com
```

Generate a TOTP code from a Base32 secret:

```bash
go run . code JBSWY3DPEHPK3PXP
```

## URI format

```text
otpauth://totp/{{issuer}}:{{account_name}}?secret={{base32_secret}}&issuer={{issuer}}&algorithm={{algorithm}}&digits={{digits}}&period={{period}}
```

- algorithm: `sha-1`
- issuer: provider or service name
- account_name: account identifier (for example email)
- digits: number of digits (commonly 6 or 8; default 6)
- period: code validity period in seconds (default 30)

## References

- [Google Authenticator Key URI Format](https://github.com/google/google-authenticator/wiki/Key-Uri-Format)
- [TOTP](https://en.wikipedia.org/wiki/Time-based_one-time_password)
- [HOTP (RFC 4226)](https://datatracker.ietf.org/doc/html/rfc4226)