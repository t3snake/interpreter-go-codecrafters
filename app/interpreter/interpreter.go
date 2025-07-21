package interpreter

import (
	"fmt"
	"strconv"

	"github.com/codecrafters-io/interpreter-starter-go/app/parser"

	//lint:ignore ST1001 I dont care
	. "github.com/codecrafters-io/interpreter-starter-go/app/token"
)

func EvaluateAst(node *parser.AstNode) (any, error) {
	// TODO typecheck here? but any return type
	switch node.Type {
	case parser.GROUP:
		if len(node.Children) == 0 {
			return nil, fmt.Errorf("interpreter error no children for group node")
		}
		return EvaluateAst(node.Children[0])
	case parser.NUMBERNODE:
		return node.Representation, nil
	case parser.STRINGNODE:
		return node.Representation, nil
	case parser.TERMINAL:
		val, ok := node.Representation.(Token)
		if !ok {
			return nil, fmt.Errorf("interpreter expected Token for Terminal but got %s", node.Representation)
		}
		switch val.Type {
		case TRUE:
			return true, nil
		case FALSE:
			return false, nil
		case NIL:
			return nil, nil
		default:
			return nil, fmt.Errorf("unexpected Token Type for Terminal AST node %s", val.Type)
		}
	default:
		// TODO do other node cases
		return nil, fmt.Errorf("unexpected or not implemented yet")
	}
}

func PrintEvaluation(result any) string {
	if result == nil {
		return "nil"
	}
	switch res := result.(type) {
	case bool:
		if res {
			return "true"
		} else {
			return "false"
		}
	case float64:
		return strconv.FormatFloat(res, 'f', -1, 64)
	case string:
		return res
	default:
		return "error: unknown evaluation"
	}
}
