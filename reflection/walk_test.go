package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Levi"},
			[]string{"Levi"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Levi", "Trost"},
			[]string{"Levi", "Trost"},
		},
		{
			"Struct with non string fields",
			struct {
				Name string
				Age  int
			}{"Levi", 33},
			[]string{"Levi"},
		},
		{
			"Nested struct with string field",
			struct {
				Name    string
				Profile struct {
					Age  int
					City string
				}
			}{"Eren Jaeger", struct {
				Age  int
				City string
			}{15, "Shiganshina"}},
			[]string{"Eren Jaeger", "Shiganshina"},
		},
		{
			"Pointer to things",
			&Person{
				"Eren Jaeger",
				Profile{15, "Shiganshina"},
			},
			[]string{"Eren Jaeger", "Shiganshina"},
		},
		{
			"Slices",
			[]Profile{
				{33, "Shiganshina"},
				{34, "Trost"},
			},
			[]string{"Shiganshina", "Trost"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "Shiganshina"},
				{34, "Trost"},
			},
			[]string{"Shiganshina", "Trost"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v but wanted %v", got, test.ExpectedCalls)
			}
		})
	}

	// Since the order of keys are not guaranteed in maps, we test it separately
	t.Run("Maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "SHINZOU",
			"Baz": "WO SASAGEYO",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "SHINZOU")
		assertContains(t, got, "WO SASAGEYO")
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
