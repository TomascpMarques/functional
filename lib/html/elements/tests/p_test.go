package elements

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/tomascpmarques/frontline/lib/html/elements"
)

func TestP_PushNewElement(t *testing.T) {
	type fields struct {
		Content    []interface{}
		attributes map[string]string
	}
	type args struct {
		e elements.Element
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Paragraph Push",
			fields: fields{
				Content:    elements.ContentList{"Lorem "},
				attributes: nil,
			},
			args: args{
				e: elements.NewP("Yess"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &elements.P{
				Content:    tt.fields.Content,
				Attributes: tt.fields.attributes,
			}
			p.PushNewElement(tt.args.e)
		})
	}
}

func TestP_MarkItUp(t *testing.T) {
	type fields struct {
		Content    []interface{}
		attributes map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Simple Paragraph",
			fields: fields{
				Content:    elements.ContentList{"lorem potato", ", beat"},
				attributes: map[string]string{},
			},
			want: "<p>lorem potato, beat</p>",
		},
		{
			name: "Simple Paragraph With Attributes",
			fields: fields{
				Content:    elements.ContentList{"lorem potato", ", beat"},
				attributes: map[string]string{"name": "one", "yes": "no"},
			},
			want: `<p name="one" yes="no">lorem potato, beat</p>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := elements.P{
				Content:    tt.fields.Content,
				Attributes: tt.fields.attributes,
			}
			if got := p.MarkItUp(); got != tt.want {
				t.Errorf("P.MarkItUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestP_SetAttributes(t *testing.T) {
	type fields struct {
		Content    []interface{}
		attributes map[string]string
	}
	type args struct {
		attr map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "One Attribute",
			fields: fields{
				Content:    []interface{}{},
				attributes: map[string]string{"id": "1"},
			},
			args: args{
				attr: map[string]string{"id": "1"},
			},
		},
		{
			name: "Manny Attribute",
			fields: fields{
				Content:    []interface{}{},
				attributes: map[string]string{"id": "1", "yes": "no", "potato": "2", "boat": "eer"},
			},
			args: args{
				attr: map[string]string{"id": "1", "yes": "no", "potato": "2", "boat": "eer"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &elements.P{
				Content:    tt.fields.Content,
				Attributes: tt.fields.attributes,
			}
			p.SetAttributes(tt.args.attr)
		})
	}
}

func TestP_ReplaceContent(t *testing.T) {
	type fields struct {
		Content    []interface{}
		attributes map[string]string
	}
	type args struct {
		new elements.Element
		pos uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &elements.P{
				Content:    tt.fields.Content,
				Attributes: tt.fields.attributes,
			}
			p.ReplaceContent(tt.args.new, tt.args.pos)
		})
	}
}

func TestNewP(t *testing.T) {
	type args struct {
		content elements.ContentList
	}
	tests := []struct {
		name string
		args args
		want elements.P
	}{
		{
			name: "Only Text",
			args: args{elements.ContentList{"lorem Ipsum"}},
			want: elements.P{Content: elements.ContentList{"lorem Ipsum"}, Attributes: nil},
		},
		{
			name: "Only Text v2",
			args: args{elements.ContentList{"lorem Ipsum <p>YES</p>"}},
			want: elements.P{Content: elements.ContentList{"lorem Ipsum <p>YES</p>"}, Attributes: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := json.Marshal(*elements.NewP(tt.args.content...))
			want, _ := json.Marshal(tt.want)

			if reflect.DeepEqual(got, want) {
				t.Errorf("NewP() = %v,\n-> want %v", got, tt.want)
			}
		})
	}
}
