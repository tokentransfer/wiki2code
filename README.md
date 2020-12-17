# wiki2code

1. clone the wiki from github
```
git clone https://github.com/caivega/ipfslib.wiki.git
```

2. clone this repository from github
```
git clone https://github.com/tokentransfer/wiki2code.git
```

3. build and run the wiki2code
```
go build && ./wiki2code ../ipfslib.wiki/chain错误信息整理.md

go build && ./wiki2code json ../ipfslib.wiki/chain错误信息整理.md > error.json

go build && ./wiki2code golang ../ipfslib.wiki/chain错误信息整理.md | pbcopy
```