# lgtm-capture

Capture screen, then compose image for LGTM, then copy it to the clipboard.

Support Mac OS X only :(

Use imagemagick version 6.

```sh
brew unlink imagemagick
brew uninstall imagemagick
brew install imagemagick@6
brew link imagemagick@6 --force
```

# Install

```sh
go get github.com/myamamoto88/lgtm-capture
```

# Usage

```sh
./lgtm-capture
```

or

1. Create application file by Automator
1. Add application to dock
1. Click application icon, or search alfred and execute

