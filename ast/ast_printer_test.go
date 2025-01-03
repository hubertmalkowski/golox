package ast

import "testing"

func TestExprAstPrinter(t *testing.T) {

	test := []struct {
		expr Expr
		want string
	}{
		{
			expr: NewBinary(
				NewUnary(
					NewToken(MinusTT, "-", nil, 1),
					NewLiteral("123"),
				),
				NewToken(StarTT, "*", nil, 1),
				NewLiteral("2"),
			),

			want: "(* (- 123) 2)",
		},
		{
			expr: NewBinary(
				NewLiteral("1"),
				NewToken(PlusTT, "+", nil, 1),
				NewGrouping(
					NewBinary(
						NewLiteral("2"),
						NewToken(StarTT, "*", nil, 1),
						NewLiteral("3"),
					),
				),
			),
			want: "(+ 1 (group (* 2 3)))",
		},
	}

	for _, tt := range test {
		t.Run(tt.want, func(t *testing.T) {
			got, _ := NewAstPrinter().Print(tt.expr)
			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}

}
