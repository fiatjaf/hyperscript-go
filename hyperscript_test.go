package h

import "testing"

func TestManyThings(t *testing.T) {
	for i, test := range tests {
		if test.source.Render() != test.expected {
			t.Errorf("failed at %d. '%s' should be '%s'.", i, test.source.Render(), test.expected)
		}
	}
}

var tests = []x{
	x{Text("some text"), "some text"},
	x{Element("main", A{"data-key": "value"}, nil), "<main data-key='value'></main>"},
	x{Element("div", nil, []H{Text("hello"), Text(" "), Text("world")}), "<div>hello world</div>"},
	x{Element("ul", A{"  style  ": "color: #333333   "}, []H{
		Element("li", nil, []H{Text("1")}),
	}), "<ul style='color: #333333'><li>1</li></ul>"},
}

type x struct {
	source   H
	expected string
}
