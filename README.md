# htredirect

htredirect is a lightweight HTTP redirect service written in Go.

## Running

Recommended way to run htredirect is using Docker:

```bash
docker run -d --name htredirect -p 8080:80 -v ./htredirect.yml:/app/htredirect.yml ghcr.io/lugosieben/htredirect:latest
```
> [!NOTE]
> Make sure to replace the port mapping and configuration file path with your own values.

If you want to use compose, you can start with the [`docker-compose.yml`](deployments/docker-compose.yml) in [`/deployments`](deployments)

Alternatively, you can build and run the service directly, but note that the binary requires `htredirect.yml` and `web/templates/*` in its relative path.

## Configuration

```yaml
port: 80
entries:
  - target: "https://example.com/{path}" # URL to redirect to, supports redirecting the request {path}
    method: "permanent" # permanent (301) or temporary (302)
    rules: # All rules need to match for a redirect to happen
      - field: "host" # host or path (trailing slashes are always stripped)
        comparator: "equal" # equal, equal-insensitive, notequal, regex, notregex, prefix, suffix
        value: "example.net" # value to compare against
```

## See it in action

[gh.lugo.ovh/htredirect](https://gh.lugo.ovh/htredirect) redirects here!
```yaml
- target: "https://github.com/lugosieben/{path}"
  method: "temporary"
  rules:
    - field: "host"
      comparator: "equal"
      value: "gh.lugo.ovh"
```

## License

Licensed under the MIT License. See [LICENSE](LICENSE)
