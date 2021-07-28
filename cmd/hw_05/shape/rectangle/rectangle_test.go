package rectangle

import (
	"math"
	"reflect"
	"testing"
)

func TestRectangle_Area(t *testing.T) {
	type fields struct {
		width float64
		height float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{name: "check negative width",
			fields:  fields{width: -10, height: 2},
			want:    0,
			wantErr: true,
		},
		{name: "check correct value",
			fields:  fields{width: 3, height: 4},
			want:    12.00,
			wantErr: false,
		},
		{name: "check negative height",
			fields:  fields{width: 0, height: -2},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				width: tt.fields.width,
				height: tt.fields.height,
			}
			got, err := r.Area()
			if (err != nil) != tt.wantErr {
				t.Errorf("Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Area() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Perimeter(t *testing.T) {
	type fields struct {
		width float64
		height float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{name: "check negative width",
			fields:  fields{width: -10, height: 2},
			want:    0,
			wantErr: true,
		},
		{name: "check correct value",
			fields:  fields{width: 3, height: 4},
			want:    14.00,
			wantErr: false,
		},
		{name: "check negative height",
			fields:  fields{width: 0, height: -3},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				width: tt.fields.width,
				height: tt.fields.height,
			}
			got, err := r.Perimeter()
			if (err != nil) != tt.wantErr {
				t.Errorf("Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Area() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_String(t *testing.T) {
	type fields struct {
		width float64
		height float64
	}

	tests := []struct {
		name string
		fields fields
		want string
	}{
		{
			name: "check zero",
			fields: fields{width: 0, height: 0},
			want: "Rectangle with height 0.00 and width 0.00",
		},
		{
			name: "check negative value",
			fields: fields{width: -10, height: 2},
			want: "Rectangle with height 2.00 and width -10.00",
		},
		{
			name: "check positive value",
			fields: fields{width: 8, height: 3},
			want: "Rectangle with height 3.00 and width 8.00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				width: tt.fields.width,
				height: tt.fields.height,
			}
			if got := r.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type fields struct {
		width float64
		height float64
	}
	tests := []struct {
		name string
		fields fields
		want Rectangle
	}{
		{
			name: "check positive value",
			fields: fields{width: 12, height: 12},
			want: Rectangle{12, 12},
		},
		{
			name: "check max positive value",
			fields: fields{width: math.MaxFloat64, height: math.MaxFloat64},
			want: Rectangle{math.MaxFloat64, math.MaxFloat64},
		},
		{
			name: "check min negative value",
			fields: fields{width: -math.MaxFloat64, height: -math.MaxFloat64},
			want: Rectangle{-math.MaxFloat64, -math.MaxFloat64},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.fields.width, tt.fields.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}