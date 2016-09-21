package h

import "testing"

func TestManyThings(t *testing.T) {
	for i, test := range tests {
		t.Logf("test %d", i)
		if test.source.Render() != test.expected {
			t.Errorf("'%s' should be '%s'.", test.source.Render(), test.expected)
		}
	}
}

var tests = []x{
	x{Text("some text"), "some text"},
	x{Element("main", A{"data-key": "value"}, nil), "<main data-key='value'></main>"},
	x{Element("div", nil, HH{Text("hello"), Text(" "), Text("world")}), "<div>hello world</div>"},
	x{Element("ul", A{"  style  ": "color: #333333   "}, HH{
		Element("li ", nil, HH{Text("1")}),
	}), "<ul   style  ='color: #333333   '><li>1</li></ul>"},
	x{Element("tree", nil, HH{
		Element("leaf", A{"color": "green"}, Text("~")),
		Element("leaf", A{"color": "green"}, Text("^")),
		Element("leaf", A{"color": "green"}, Text("#")),
	}), "<tree><leaf color='green'>~</leaf><leaf color='green'>^</leaf><leaf color='green'>#</leaf></tree>"},
	x{HH{
		Element("dt", nil, Text("a title")),
		Element("dd", nil, Text("a description")),
	}, "<dt>a title</dt><dd>a description</dd>"},
	x{Element("b.prelude#that", A{"author": "Bach"}, nil), "<b author='Bach' class='prelude' id='that'></b>"},
	x{Element(".chaconne", A{"class": "music"}, nil), "<div class='music chaconne'></div>"},
	x{Element("b#nothing", A{"class": "blank", "id": ""}, nil), "<b class='blank' id=''></b>"},
}

type x struct {
	source   H
	expected string
}
