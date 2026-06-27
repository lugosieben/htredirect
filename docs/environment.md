# Environment Variables

Environment Variables are read once at program startup and can be used to set properties.

## Properties

Properties define behavior of the application that is not related to redirects and may also be needed before even parsing the config.

List of changeable properties:

| Name       | Environment Variable  | Description                                     | Default             |
|------------|-----------------------|-------------------------------------------------|---------------------|
| TIMEFORMAT | HTREDIRECT_TIMEFORMAT | Time Format to use, in Go syntax                | 2006-01-02 15:04:05 |
| MAINCONFIG | HTREDIRECT_MAINCONFIG | Path to the main config, relative to the binary | config.htredirect   |
