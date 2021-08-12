# semtag

Tag your repository according to Semantic Versioning.

[pnikosis/semtag](https://github.com/pnikosis/semtag), but using Go.

## Installation

Using Homebrew

```bash
brew tap leb4r/tap/semtag
brew install leb4r/tap/semtag
```

## Usage

## Container

```bash
# pull the image from DockerHub
docker pull leb4r/semtag

# run on a git repository
docker run -it -v $PWD:/src -w /src --rm leb4r/semtag final
```
