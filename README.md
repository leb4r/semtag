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
# pull the image from Docker Hub
docker pull docker.io/leb4r/semtag

# execute on the current working directory
docker run -it -v $PWD:/src:z -w /src --rm docker.io/leb4r/semtag final
```
