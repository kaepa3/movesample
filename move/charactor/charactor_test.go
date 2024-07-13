package charactor

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestCharactor_IsNear(t *testing.T) {
	type fields struct {
		PosX int
		PosY int
		Img  *ebiten.Image
		Size int
		Text string
	}
	type args struct {
		o *Charactor
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{"same", fields{0, 0, nil, 20, ""}, args{&Charactor{0, 0, nil, 20, ""}}, true},
		{"10", fields{0, 0, nil, 20, ""}, args{&Charactor{10, 0, nil, 20, ""}}, true},
		{"20", fields{0, 0, nil, 20, ""}, args{&Charactor{20, 0, nil, 20, ""}}, true},
		{"30", fields{0, 0, nil, 20, ""}, args{&Charactor{30, 0, nil, 20, ""}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Charactor{
				PosX: tt.fields.PosX,
				PosY: tt.fields.PosY,
				Img:  tt.fields.Img,
				Size: tt.fields.Size,
				Text: tt.fields.Text,
			}
			if got := c.IsNear(tt.args.o); got != tt.want {
				t.Errorf("Charactor.IsNear() = %v, want %v", got, tt.want)
			}
		})
	}
}
