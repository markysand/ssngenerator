package main

import (
	"reflect"
	"testing"

	"github.com/rickb777/date"
)

func Test_getBirthRange(t *testing.T) {
	now := date.New(2021, 1, 1)

	type args struct {
		c *config
	}
	tests := []struct {
		name     string
		args     args
		wantFrom date.Date
		wantTo   date.Date
	}{
		{
			name: "Child - 0-13 years",
			args: args{
				c: &config{
					child: true,
					now:   now,
				},
			},
			wantFrom: now.AddDate(-13, 0, 0),
			wantTo:   now,
		},
		{
			name: "Teen - 13-18 years",
			args: args{
				c: &config{
					teen: true,
					now:  now,
				},
			},
			wantFrom: now.AddDate(-18, 0, 0),
			wantTo:   now.AddDate(-13, 0, 0),
		},
		{
			name: "Adult - 18-100 years",
			args: args{
				c: &config{
					adult: true,
					now:   now,
				},
			},
			wantFrom: now.AddDate(-100, 0, 0),
			wantTo:   now.AddDate(-18, 0, 0),
		},
		{
			name: "Unknown - 0-100 years",
			args: args{
				c: &config{
					years:  -1,
					months: -1,
					now:    now,
				},
			},
			wantFrom: now.AddDate(-100, 0, 0),
			wantTo:   now,
		},
		{
			name: "Months - interval is a month",
			args: args{
				c: &config{
					years:  -1,
					months: 5,
					now:    now,
				},
			},
			wantFrom: now.AddDate(0, -6, 0),
			wantTo:   now.AddDate(0, -5, 0),
		},
		{
			name: "Years - interval is a year",
			args: args{
				c: &config{
					years:  10,
					months: -1,
					now:    now,
				},
			},
			wantFrom: now.AddDate(-11, 0, 0),
			wantTo:   now.AddDate(-10, 0, 0),
		},
		{
			name: "Years and months - interval is a month",
			args: args{
				c: &config{
					years:  10,
					months: 5,
					now:    now,
				},
			},
			wantFrom: now.AddDate(-10, -6, 0),
			wantTo:   now.AddDate(-10, -5, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFrom, gotTo := getBirthRange(tt.args.c)
			if !reflect.DeepEqual(gotFrom, tt.wantFrom) {
				t.Errorf("getBirthRange() gotFrom = %v, want %v", gotFrom, tt.wantFrom)
			}
			if !reflect.DeepEqual(gotTo, tt.wantTo) {
				t.Errorf("getBirthRange() gotTo = %v, want %v", gotTo, tt.wantTo)
			}
		})
	}
}
