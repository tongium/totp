# totp
Time-based one-time password demonstration

## URI Format

```
otpauth://totp/$ISSUER:$ACCOUNT_NAME?secret=$BASE32_KEY&issuer=$ISSUER&algorithm=$ALGORITHM&digits=$DIGITS&period=$SECONDS$
```

- ALGORITHM: SHA-1
- ISSUER: Organization
- ACCOUNT_NAME: Email
- DIGITS: 6
- SECONDS: 30


## Reference

- [URI Format](https://github.com/google/google-authenticator/wiki/Key-Uri-Format)
- [TOTP](https://en.wikipedia.org/wiki/Time-based_one-time_password)
- [HOTP](https://datatracker.ietf.org/doc/html/rfc4226)