package circle

import (
	"math"
	"reflect"
	"testing"
)

func TestCircle_Area(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{name: "check negative value",
			fields:  fields{radius: -10},
			want:    0,
			wantErr: true,
		},
		{name: "check correct value",
			fields:  fields{radius: 8},
			want:    201.06192982974676,
			wantErr: false,
		},
		{name: "check zero",
			fields:  fields{radius: 0},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				radius: tt.fields.radius,
			}
			got, err := c.Area()
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

func TestCircle_Perimeter(t *testing.T) {
	type fields struct {
		radius float64
	}

	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	} {
		{
			name: "check negative value",
			fields: fields{radius: -10},
			want: 0,
			wantErr: true,
		},
		{
			name: "check correct value",
			fields: fields{8},
			want: 50.26548245743669,
			wantErr: false,
		},
		{
			name: "check zero",
			fields: fields{0},
			want: 0,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				radius: tt.fields.radius,
			}
			got, err := c.Perimeter()
			if (err != nil) != tt.wantErr {
				t.Errorf("Perimeter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Perimeter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_String(t *testing.T) {
	type fields struct {
		radius float64
	}

	tests := []struct {
		name string
		fields fields
		want string
	}{
		{
			name: "check zero",
			fields: fields{0},
			want: "Circle: radius 0.00",
		},
		{
			name: "check negative value",
			fields: fields{-10},
			want: "Circle: radius -10.00",
		},
		{
			name: "check positive value",
			fields: fields{8},
			want: "Circle: radius 8.00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				radius: tt.fields.radius,
			}
			if got := c.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name string
		fields fields
		want Circle
		wantErr bool
	}{
		{
			name: "check positive value",
			fields: fields{12},
			want: Circle{12},
			wantErr: false,
		},
		{
			name: "check max positive value",
			fields: fields{math.MaxFloat64},
			want: Circle{math.MaxFloat64},
			wantErr: false,
		},
		{
			name: "check min negative value",
			fields: fields{-math.MaxFloat64},
			want: Circle{0},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.fields.radius)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}