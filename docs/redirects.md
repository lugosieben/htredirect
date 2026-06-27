# Redirects

Redirects are defined with `REDIRECT` in `.htredirect` files:

```htredirect
-- Comment
REDIRECT WHERE
[<FIELD> [MODS] <COMPARATOR> <VALUE>] -- Multiple rules can be seperated by commas
TO <URL> <METHOD>;
```

## Fields

Request-determined fields that will be compared against.

| Field  | Description                                                   |
|--------|---------------------------------------------------------------|
| `HOST` | The host of the request                                       |
| `PATH` | The path of the request. Trailing slashes are always stripped |

## Modifiers

Modify the field before comparison or the result after comparison. Any number of modifiers can be used. They are space separated.

| Modifier | Description                                      |
|----------|--------------------------------------------------|
| `LOWER`  | Convert the field to lowercase before comparison |
| `NOT`    | Invert the comparison result                     |

## Comparators

Compare the field against the value, possibly with modifiers applied.

| Comparator | Description                                                   |
|------------|---------------------------------------------------------------|
| `EQUALS`   | True if the field equals the value                            |
| `MATCHES`  | True if the field matches the regex value (ECMAScript syntax) |
| `PREFIX`   | True if the field starts with the value                       |
| `SUFFIX`   | True if the field ends with the value                         |


## Value

Value to compare the field against. If the value contains spaces, it must be wrapped in single or double quotes.
It is still recommended to wrap values like URLs in quotes, even if they don't contain spaces, for readability reasons.

## URL

The URL to redirect to. It supports placeholders:

| Placeholder | Description                                                  |
|-------------|--------------------------------------------------------------|
| `{path}`    | The path of the request, without leading or trailing slashes |

## Methods

How to send the redirect:

| Method      | Description                          |
|-------------|--------------------------------------|
| `PERMANENT` | Send a permanent redirect (HTTP 301) |
| `TEMPORARY` | Send a temporary redirect (HTTP 302) |