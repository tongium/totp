# totp
Time-based one-time password demonstration

## URI Format

```
otpauth://totp/{{issuer}}:{{account_name}}?secret={{key_with_base32_format}}&issuer={{issuer}}&algorithm={{algorithm}}&digits={{digits}}&period={{period}}
```

- algorithm: SHA-1
- issuer: Provider
- account_name: Email
- digits: Numer of digit [6,8]. default: 6
- period: Period in second, default: 30


## Reference

- [URI Format](https://github.com/google/google-authenticator/wiki/Key-Uri-Format)
- [TOTP](https://en.wikipedia.org/wiki/Time-based_one-time_password)
- [HOTP](https://datatracker.ietf.org/doc/html/rfc4226)