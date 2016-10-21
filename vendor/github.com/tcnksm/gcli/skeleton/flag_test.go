package skeleton

import "testing"

func TestFix(t *testing.T) {
	tests := []struct {
		in      *Flag
		exp     *Flag
		success bool
	}{
		{
			in: &Flag{
				LongName:    "debug",
				TypeString:  "bool",
				Description: "Run as DEBUG mode",
			},
			exp: &Flag{
				Name:         "debug",
				ShortName:    "d",
				LongName:     "debug",
				VariableName: "debug",
				TypeString:   TypeStringBool,
				Default:      false,
				Description:  "Run as DEBUG mode",
			},
			success: true,
		},

		{
			in: &Flag{
				LongName:    "token",
				TypeString:  "s",
				Description: "",
			},
			exp: &Flag{
				Name:         "token",
				ShortName:    "t",
				LongName:     "token",
				VariableName: "token",
				TypeString:   TypeStringString,
				Default:      "",
				Description:  "",
			},
			success: true,
		},

		{
			in: &Flag{
				LongName:    "ignore-case",
				TypeString:  "bool",
				Description: "",
			},
			exp: &Flag{
				Name:         "ignore-case",
				VariableName: "ignoreCase",
				ShortName:    "i",
				LongName:     "ignore-case",
				TypeString:   TypeStringBool,
				Default:      false,
				Description:  "",
			},
			success: true,
		},

		{
			in: &Flag{
				LongName:    "token",
				TypeString:  "s",
				Description: "",
				Default:     "ABCD1124",
			},
			exp: &Flag{
				Name:         "token",
				ShortName:    "t",
				LongName:     "token",
				VariableName: "token",
				TypeString:   TypeStringString,
				Default:      "ABCD1124",
				Description:  "",
			},
			success: true,
		},
	}

	for i, tt := range tests {

		err := tt.in.Fix()
		if err != nil && !tt.success {
			continue
		}

		if err == nil && !tt.success {
			t.Fatalf("#%d expect Fix to fail", i)
		}

		if err != nil {
			t.Fatalf("#%d expect Fix not to fail but %q", i, err.Error())
		}

		if *tt.in != *tt.exp {
			t.Errorf("#%d expect %v to eq %v", i, tt.in, tt.exp)
		}
	}
	for i, tt := range tests {

		err := tt.in.Fix()
		if err != nil && !tt.success {
			continue
		}

		if err == nil && !tt.success {
			t.Fatalf("#%d expect Fix to fail", i)
		}

		if err != nil {
			t.Fatalf("#%d expect Fix not to fail but %q", i, err.Error())
		}

		if *tt.in != *tt.exp {
			t.Errorf("#%d expect %v to eq %v", i, tt.in, tt.exp)
		}
	}
}

func TestFixTypeString(t *testing.T) {

	tests := []struct {
		in            *Flag
		success       bool
		expTypeString string
		expDefault    interface{}
	}{
		{
			in:            &Flag{TypeString: "int"},
			success:       true,
			expTypeString: TypeStringInt,
			expDefault:    0,
		},

		{
			in:            &Flag{TypeString: "Int"},
			success:       true,
			expTypeString: TypeStringInt,
			expDefault:    0,
		},

		{
			in:            &Flag{TypeString: "i"},
			success:       true,
			expTypeString: TypeStringInt,
			expDefault:    0,
		},

		{
			in:            &Flag{TypeString: "string"},
			success:       true,
			expTypeString: TypeStringString,
			expDefault:    "",
		},

		{
			in:            &Flag{TypeString: "s"},
			success:       true,
			expTypeString: TypeStringString,
			expDefault:    "",
		},

		{
			in:            &Flag{TypeString: "str"},
			success:       true,
			expTypeString: TypeStringString,
			expDefault:    "",
		},

		{
			in:            &Flag{TypeString: "bool"},
			success:       true,
			expTypeString: TypeStringBool,
			expDefault:    false,
		},

		{
			in:            &Flag{TypeString: "b"},
			success:       true,
			expTypeString: TypeStringBool,
			expDefault:    false,
		},

		{
			in:            &Flag{TypeString: "enexpected_type"},
			success:       false,
			expTypeString: TypeStringBool,
			expDefault:    false,
		},
	}

	for i, tt := range tests {

		err := tt.in.fixTypeString()
		if err != nil && !tt.success {
			continue
		}

		if err == nil && !tt.success {
			t.Fatalf("#%d expect fixTypeString to fail", i)
		}

		if err != nil {
			t.Fatalf("#%d expect fixTypeString not to fail but %q", i, err.Error())
		}

		if tt.in.TypeString != tt.expTypeString {
			t.Errorf("#%d expect %q to eq %q", i, tt.in.TypeString, tt.expTypeString)
		}

		if tt.in.Default != tt.expDefault {
			t.Errorf("#%d expect %v to eq %v", i, tt.in.Default, tt.expDefault)
		}
	}
}
