## hyperscript-go

supersimple [hyperscript](https://github.com/dominictarr/hyperscript)-inspired templating language / html generator for go.

### usage

something like

```go
package main

import (
    "fmt"
    "github.com/fiatjaf/hyperscript-go"
)

func main() {
    fmt.Print(h.Element("div#container", nil, h.HH{
        h.Element("h1.title", nil, h.Text("hello stranger")),
        h.Element(".nonsense", nil, h.HH{
            h.Element("a", h.A{"href": "#"}, h.Text("click here to do nothing")),
            h.HTML("<a href='#'>or here</a>"),
        }),
    }).Render())
}
```

outputs `<div id='container'><h1 class='title'>hello stranger</h1><div class='nonsense'><a href='#'>click here to do nothing</a><a href='#'>or here</a></div></div>`

### be responsible

this library is very naïve and will fail if pushed to awkward situations. for example, when using special characters in attribute names and values.
