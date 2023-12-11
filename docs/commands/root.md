# Root Command
Parent command the gets executed with all other commands. The flags are inherited. 

## Flags
- `log-file` or `l` is the name of the file to output to
    - optional: defaults to `STDERR`
    - Log open-time flags: `os.O_CREATE|os.O_WRONLY|os.O_APPEND`
    - Validation failures are logged in `JSON` format. 