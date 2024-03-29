package ptag_test

import (
	"reflect"
	"testing"

	"github.com/1sl4nds/moses/pkg/tag"
	"github.com/1sl4nds/moses/pkg/tag/ptag"
)

func Test_New(t *testing.T) {
	type args struct {
		pubKeyStr string
		relayURL  string
		petname   string
	}
	tests := []struct {
		name string
		args args
		want tag.Tag
	}{
		{
			name: "SHOULD construct petname variant of Tag",
			args: args{
				pubKeyStr: "pub-key-str",
				relayURL:  "relay-url",
				petname:   "petname",
			},
			want: tag.Tag{ptag.Type},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.want.Push(tt.args.pubKeyStr)
			tt.want.Push(tt.args.relayURL)
			tt.want.Push(tt.args.petname)
			got := ptag.New(tt.args.pubKeyStr, tt.args.relayURL, tt.args.petname)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("wanted %v, got %v", tt.want, got)
			}
			t.Logf("got %v", got)
		})
	}
}
