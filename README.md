# Zash shell

## Description

This is a small pet project aimed to build minimal version of POSIX compliant shell that's capable of interpreting shell commands, running external programs and builtin commands like cd, pwd, echo etc. Along the way, codebase contains my attempts in Go to implement command parsing, REPL, builtin commands, autocomplete, running executables, etc.

This shell is compatible within linux/macOS. However, it was written using MacOS so there might potentially be some tweaks with syscalls in Linux :D

This project contains the following shell features:

1. Most basic functionality: REPL, builtin commands (echo, exit, type), executables running
2. Navigation (pwd, cd commands)
3. Quotes handling (single quotes, double quotes, backslashes within single/double quotes, execution quoted executables)
4. STD redirection (stdout, stderr) using unix operators (>, 1>, 2>, 1>>, 2>>)
5. Text autocompletion (single match, multiple matches, partial completion with longest common prefix, missing completions)

## P.S.

Purpose of the project was just to get familiar with Go language for myself on practice, so some code snippets might be kinda bad.

## How to run locally:

1. Clone repository locally
2. Run within project directory: `./shell.sh`
