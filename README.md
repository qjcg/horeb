# Horeb [![Build Status][build-logo]][horeb-travis]

![Mt. Horeb][mt-horeb]

Speaking in tongues via stdout.

Somewhat inspired by the [TempleOS](http://templeos.org) [oracle].

## Install

```sh
go get -u github.com/qjcg/horeb
```

Alternatively, you can download the [latest binary release here].

## Usage
### Basic Usage
```
go run main.go blocks.go -h
horeb: Speaking in tongues via stdout
  -c    colorize output
  -d    print all Blocks
  -h    prints usage documentation
  -l    list all Block names and codepoint ranges
  -n int
        number of characters to print (default 30)
```
### Advanced usage
The following line will print a sequence of 30 characters composed of characters from the geometric and emoji sets
```
./horeb -n=30 geometric emoji
```

## Test

Run main unit test suite:

```
go test
```

Run unit and integration tests (after a successful "go install"):

```sh
go test -tags integration
```


## Font Support

For information about fonts supporting specific unicode blocks, see [fileformat.info].

[build-logo]: https://travis-ci.org/qjcg/horeb.svg?branch=master
[horeb-travis]: https://travis-ci.org/qjcg/horeb
[mt-horeb]: http://upload.wikimedia.org/wikipedia/commons/thumb/a/a4/Francis_Frith_%28English_-_Mount_Horeb%2C_Sinai_-_Google_Art_Project_%286787000%29.jpg/306px-Francis_Frith_%28English_-_Mount_Horeb%2C_Sinai_-_Google_Art_Project_%286787000%29.jpg "Mt. Horeb"
[oracle]: https://www.youtube.com/watch?v=jqT-EgUN4y8
[latest binary release here]: https://github.com/qjcg/horeb/releases/latest
[fileformat.info]: http://www.fileformat.info/info/unicode/block/index.htm


## License

MIT.
