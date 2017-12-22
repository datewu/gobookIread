package eval

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strings"
)

// Env is the environment
type Env map[Var]float64

// Expr is an arithmetic expression
type Expr interface {
	// Eval returns the value of this Expr in the environment env.
	Eval(env Env) float64

	// Check reports errors in this Expr and adds its Vars to the set.
	Check(vars map[Var]bool) error
}

// Var identifies a varoable, e.g. x.
type Var string

// Eval implement Expr interface
func (v Var) Eval(env Env) float64 {
	return env[v]
}

// Check implement Expr interface
func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

// literal is a numberic constant, e.g. 3.1415
type literal float64

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (l literal) Check(_ map[Var]bool) error {
	return nil
}

// unary represents a unary operator expression, e.g. -x.
type unary struct {
	op rune // one of '+' , '-'.
	x  Expr
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupport unary operator: %q", u.op))
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q", u.op)
	}
	return u.x.Check(vars)
}

// binary represents a binary operator expression, e.g. x + y.
type binary struct {
	op   rune // one of '+', '-', '*', '/'
	x, y Expr
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)

	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupport binary operator: %q", b.op))
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected unay op %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

// call represents a function call expression. e.g. sin(x)
type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupport  function call: %q", c.fn))
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}

	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d", c.fn, len(c.args), arity)
	}

	for _, a := range c.args {
		if err := a.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

// ParseAndCheck ready go
func ParseAndCheck(s string) (Expr, error) {
	if s == "" {
		return nil, errors.New("empty epression")
	}

	expr, err := parse(s)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	vars := make(map[Var]bool)
	if err := expr.Check(vars); err != nil {
		log.Println(err)
		return nil, err
	}

	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable: %q", v)
		}
	}
	return expr, nil
}
