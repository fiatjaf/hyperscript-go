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
        h.Element("h1", nil, h.Text("hello stranger")),
        h.Element("p", nil,
            h.Element("a", h.A{"href": "#"}, h.Text("click here to do nothing")),
        ),
    }).Render())
}
```

outputs `<div id='container'><h1>hello stranger</h1><p><a href='#'>click here to do nothing</a></p></div>`

### be responsible

this library is very naïve and will fail if pushed to awkward situations. for example, when using special characters in attribute names and values, or if you pass valid html to text nodes (if you want to do that, escape the html yourself first).
