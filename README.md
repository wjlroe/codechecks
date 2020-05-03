# codechecks

This is a collection of tools for fixing/linting source code.

## Installation

Check out this repository and run `go install`.

## Zig fmt

Initially this tool exists to wrap the `zig fmt` tool so that it works with various file paths on Windows. It exists as a wrapper tool so that `zig fmt` can be used from tools like [pre-commit](https://pre-commit.com). `pre-commit` passes file paths to tools with UNIX-style file paths, which `zig fmt` can't parse.

Example `.pre-commit-config.yaml` to use this:

```yaml
repos:
-   repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v2.5.0
    hooks:
    -   id: check-yaml
    -   id: end-of-file-fixer
    -   id: trailing-whitespace
    -   id: check-case-conflict
    -   id: check-merge-conflict
    -   id: check-symlinks
-   repo: local
    hooks:
    -   id: zigfmt
        name: zigfmt
        description: Format files with zig fmt.
        entry: codechecks -zig-fmt
        language: system
        files: \.zig$
```
