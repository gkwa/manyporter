# manyporter

Purpose:
Use goldmark to fetch markdown frontmatter as json.

Motivation:
Feed my markdown notes into [quickwit](https://quickwit.io/)

## 


```bash

./manyporter test3

```

this gives error:

```
MarshalIndent failed json: unsupported type: map[interface {}]interface {}
```

It is for that reason why we need more. test1 works:


```

./manyporter test1

```

works:


```log

[mtm@ttorres-t14s:manyporter(master)]$ ./manyporter test1
{
  "filetype": "product",
  "x": {
    "z": {
      "n": {
        "p": 10,
        "z": [
          "test",
          2
        ]
      }
    }
  }
}

```



## example usage

```bash
make && ./manyporter run

```

## install manyporter


on macos/linux:
```bash

brew install gkwa/homebrew-tools/manyporter

```


on windows:

```powershell

TBD

```


## Links

- https://quickwit.io/
- https://github.com/yuin/goldmark?tab=readme-ov-file#goldmark