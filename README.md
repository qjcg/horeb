# Horeb [![Build Status][build-logo]][horeb-travis]

![Mt. Horeb][mt-horeb]

*"Speaking in tongues via stdout."*

Horeb is a CLI tool for generating random sequences of characters from various
[Unicode blocks](https://en.wikipedia.org/wiki/Unicode_block).

One of the sources of inspiration for this tool is the
[TempleOS](http://templeos.org) [oracle].


## Install

```sh
go get -u github.com/qjcg/horeb
```

Alternatively, you can download the [latest binary release here].


## Usage

Print 1000 random dominos:

```sh
$ horeb -n 1000 dominos
```

Print 500 random characters drawn from the `emoji`, `geometric` and `math_alnum`
code blocks:

```sh
$ horeb -n 500 emoji geometric math_alnum
```

List all Unicode block names `horeb` knows about along with their corresponding
[codepoint](https://en.wikipedia.org/wiki/Code_point) ranges:

```sh
$ horeb -l
10100 1013f  aegean_nums
10140 1018f  ancient_greek_nums
 20a0  20cf  currency

[...]
```

Dump all characters from all blocks `horeb` knows about:

```sh
$ horeb -d
10100 1013f  aegean_nums
𐄀 𐄁 𐄂 𐄇 𐄈 𐄉 𐄊 𐄋 𐄌 𐄍 𐄎 𐄏 𐄐 𐄑 𐄒 𐄓 𐄔 𐄕 𐄖 𐄗 𐄘 𐄙 𐄚 𐄛 𐄜 𐄝 𐄞 𐄟 𐄠 𐄡 𐄢 𐄣 𐄤 𐄥 𐄦 𐄧 𐄨 𐄩 𐄪 𐄫 𐄬 𐄭 𐄮 𐄯 𐄰 𐄱 𐄲 𐄳 𐄷 𐄸 𐄹 𐄺 𐄻 𐄼 𐄽 𐄾 𐄿

10140 1018f  ancient_greek_nums
𐅀 𐅁 𐅂 𐅃 𐅄 𐅅 𐅆 𐅇 𐅈 𐅉 𐅊 𐅋 𐅌 𐅍 𐅎 𐅏 𐅐 𐅑 𐅒 𐅓 𐅔 𐅕 𐅖 𐅗 𐅘 𐅙 𐅚 𐅛 𐅜 𐅝 𐅞 𐅟 𐅠 𐅡 𐅢 𐅣 𐅤 𐅥 𐅦 𐅧 𐅨 𐅩 𐅪 𐅫 𐅬 𐅭 𐅮 𐅯 𐅰 𐅱 𐅲 𐅳 𐅴 𐅵 𐅶 𐅷 𐅸 𐅹 𐅺 𐅻 𐅼 𐅽 𐅾 𐅿 𐆀 𐆁 𐆂 𐆃 𐆄 𐆅 𐆆 𐆇 𐆈 𐆉 𐆊 𐆋 𐆌 𐆍 𐆎

 20a0  20cf  currency
₠ ₡ ₢ ₣ ₤ ₥ ₦ ₧ ₨ ₩ ₪ ₫ € ₭ ₮ ₯ ₰ ₱ ₲ ₳ ₴ ₵ ₶ ₷ ₸ ₹ ₺ ₻ ₼ ₽ ₾

[...]
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
