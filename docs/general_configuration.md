# General Configuration

Configuration fields that are not related to redirects can be set with `SET` statements:

```htredirect
SET <KEY> = <VALUE>;
```

## Keys

| Key       | Description                                               | Type | Default |
|-----------|-----------------------------------------------------------|------|---------|
| `PORT`    | Port to listen on for the Redirect Server (public facing) | int  | 80      |
| `WEBPORT` | Port to listen on for the Web Server (Management)         | int  | 8080    |