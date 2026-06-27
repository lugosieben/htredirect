# Configuration

htredirect may be configured by htredirect files.
These files should use the `.htredirect` extension.

The main configuration file is `config.htredirect` next to the executable by default. This can be changed with [environment variables](environment.md).


## Syntax

Statements are seperated by semicolons `;`.
Values with spaces in them should be wrapped in single or double quotes.
For readability purposes, it is recommended to always wrap values like URLs in quotes, even if they don't contain spaces.
Newlines and indentation are ignored.

Comments are defined with `--`, ignoring everything after `--` on the same line.

There are 3 types of statements:

### General Configuration

Setting general configuration values, such as the ports to listen on, can be done with `SET`:

```htredirect
SET <KEY> = <VALUE>;
```

Multiple `SET` statements for the same key are not allowed.

More on general configuration can be found in the [general configuration documentation](general_configuration.md).

### Redirects

Redirects are defined with `REDIRECT`:

```htredirect
REDIRECT WHERE
<FIELD> [MODS] <COMPARATOR> <VALUE>,
TO <URL> <METHOD>;
```

Example:

```htredirect
REDIRECT WHERE
HOST LOWER EQUALS 'example.net',
PATH NOT MATCHES '^/api/.*'
TO 'https://example.com/{path}' TEMPORARY;
```

More on the syntax of redirect statements can be found in the [redirects documentation](redirects.md).

### Reads

More config files can be "required" with `READ`:

```htredirect
READ 'path/to/file.htredirect';
```

A file may not be read more than once.

## Order

Redirects earlier in the file take precedence over later redirects.
`READ` statements insert the read file at the position of the `READ` statement.
