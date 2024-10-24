package actionlint

// ExprNode is a node of expression syntax tree. To know the syntax, see
// https://docs.github.com/en/actions/learn-github-actions/expressions
type ExprNode interface {
	// Token returns the first token of the node. This method is useful to get position of this node.
	Token() *Token
	// Parent returns the parent node of this node.
	Parent() ExprNode
}

// Variable

// VariableNode is node for variable access.
type VariableNode struct {
	// Name is name of the variable
	Name   string
	tok    *Token
	parent ExprNode
}

// Token returns the first token of the node. This method is useful to get position of this node.
func (n *VariableNode) Token() *Token {
	return n.tok
}

// Parent returns the parent node of this node.
func (n *VariableNode) Parent() ExprNode {
	return n.parent
}

// Literals

// NullNode is node for null literal.
type NullNode struct {
	tok    *Token
	parent ExprNode
}

// Token returns the first token of the node. This method is useful to get position of this node.
func (n *NullNode) Token() *Token {
	return n.tok
}

// Parent returns the parent node of this node.
func (n *NullNode) Parent() ExprNode {
	return n.parent
}

// BoolNode is node for boolean literal, true or false.
type BoolNode struct {
	// Value is value of the boolean literal.
	Value  bool
	tok    *Token
	parent ExprNode
}

// Token returns the first token of the node. This method is useful to get position of this node.
func (n *BoolNode) Token() *Token {
	return n.tok
}

// Parent returns the parent node of this node.
func (n *BoolNode) Parent() ExprNode {
	return n.parent
}

// IntNode is node for integer literal.
type IntNode struct {
	// Value is value of the integer literal.
	Value  int
	tok    *Token
	parent ExprNode
}

// Token returns the first token of the node. This method is useful to get position of this node.
func (n *IntNode) Token() *Token {
	return n.tok
}

// Parent returns the parent node of this node.
func (n *IntNode) Parent() ExprNode {
	return n.parent
}

// FloatNode is node for float literal.
type FloatNode struct {
	// Value is value of the float literal.
	Value  float64
	tok    *Token
	parent ExprNode
}

// Token returns the first token of the node. This method is useful to get position of this node.
func (n *FloatNode) Token() *Token {
	return n.tok
}

// Parent returns the parent node of this node.
func (n *FloatNode) Parent() ExprNode {
	return n.parent
}

// StringNode is node for string literal.
type StringNode struct {
	// Value is value of the string literal. Escapes are resolved and quotes at both edges are
	// removed.
	Value  string
	tok    *Token
	parent ExprNode
}

// Token returns the first token of the node. This method is useful to get position of this node.
func (n *StringNode) Token() *Token {
	return n.tok
}

// Parent returns the parent node of this node.
func (n *StringNode) Parent() ExprNode {
	return n.parent
}

// Operators

// ObjectDerefNode represents property dereference of object like 'foo.bar'.
type ObjectDerefNode struct {
	// Receiver is an expression at receiver of property dereference.
	Receiver ExprNode
	// Property is a name of property to access.
	Property string
	parent   ExprNode
}

// Token returns the first token of the node. This method is useful to get position of this node.
func (n *ObjectDerefNode) Token() *Token {
	return n.Receiver.Token()
}

// Parent returns the parent node of this node.
func (n *ObjectDerefNode) Parent() ExprNode {
	return n.parent
}

// ArrayDerefNode represents elements dereference of arrays like '*' in 'foo.bar.*.piyo'.
type ArrayDerefNode struct {
	// Receiver is an expression at receiver of array element dereference.
	Receiver ExprNode
	parent   ExprNode
}

// Token returns the first token of the node. This method is useful to get position of this node.
func (n *ArrayDerefNode) Token() *Token {
	return n.Receiver.Token()
}

// Parent returns the parent node of this node.
func (n *ArrayDerefNode) Parent() ExprNode {
	return n.parent
}

// IndexAccessNode is node for index access, which represents dynamic object property access or
// array index access.
type IndexAccessNode struct {
	// Operand is an expression at operand of index access, which should be array or object.
	Operand ExprNode
	// Index is an expression at index, which should be integer or string.
	Index  ExprNode
	parent ExprNode
}

// Token returns the first token of the node. This method is useful to get position of this node.
func (n *IndexAccessNode) Token() *Token {
	return n.Operand.Token()
}

// Parent returns the parent node of this node.
func (n *IndexAccessNode) Parent() ExprNode {
	return n.parent
}

// Note: Currently only ! is a logical unary operator

// NotOpNode is node for unary ! operator.
type NotOpNode struct {
	// Operand is an expression at operand of ! operator.
	Operand ExprNode
	tok     *Token
	parent  ExprNode
}

// Token returns the first token of the node. This method is useful to get position of this node.
func (n *NotOpNode) Token() *Token {
	return n.tok
}

// Parent returns the parent node of this node.
func (n *NotOpNode) Parent() ExprNode {
	return n.parent
}

// CompareOpNodeKind is a kind of compare operators; ==, !=, <, <=, >, >=.
type CompareOpNodeKind int

const (
	// CompareOpNodeKindInvalid is invalid and initial value of CompareOpNodeKind values.
	CompareOpNodeKindInvalid CompareOpNodeKind = iota
	// CompareOpNodeKindLess is kind for < operator.
	CompareOpNodeKindLess
	// CompareOpNodeKindLessEq is kind for <= operator.
	CompareOpNodeKindLessEq
	// CompareOpNodeKindGreater is kind for > operator.
	CompareOpNodeKindGreater
	// CompareOpNodeKindGreaterEq is kind for >= operator.
	CompareOpNodeKindGreaterEq
	// CompareOpNodeKindEq is kind for == operator.
	CompareOpNodeKindEq
	// CompareOpNodeKindNotEq is kind for != operator.
	CompareOpNodeKindNotEq
)

// IsEqualityOp returns true when it represents == or != operator.
func (kind CompareOpNodeKind) IsEqualityOp() bool {
	return kind == CompareOpNodeKindEq || kind == CompareOpNodeKindNotEq
}

func (kind CompareOpNodeKind) String() string {
	switch kind {
	case CompareOpNodeKindLess:
		return "<"
	case CompareOpNodeKindLessEq:
		return "<="
	case CompareOpNodeKindGreater:
		return ">"
	case CompareOpNodeKindGreaterEq:
		return ">="
	case CompareOpNodeKindEq:
		return "=="
	case CompareOpNodeKindNotEq:
		return "!="
	default:
		return ""
	}
}

// CompareOpNode is node for binary expression to compare values; ==, !=, <, <=, > or >=.
type CompareOpNode struct {
	// Kind is a kind of this expression to show which operator is used.
	Kind CompareOpNodeKind
	// Left is an expression for left hand side of the binary operator.
	Left ExprNode
	// Right is an expression for right hand side of the binary operator.
	Right  ExprNode
	parent ExprNode
}

// Token returns the first token of the node. This method is useful to get position of this node.
func (n *CompareOpNode) Token() *Token {
	return n.Left.Token()
}

// Parent returns the parent node of this node.
func (n *CompareOpNode) Parent() ExprNode {
	return n.parent
}

// LogicalOpNodeKind is a kind of logical operators; && and ||.
type LogicalOpNodeKind int

const (
	// LogicalOpNodeKindInvalid is an invalid and initial value of LogicalOpNodeKind.
	LogicalOpNodeKindInvalid LogicalOpNodeKind = iota
	// LogicalOpNodeKindAnd is a kind for && operator.
	LogicalOpNodeKindAnd
	// LogicalOpNodeKindOr is a kind for || operator.
	LogicalOpNodeKindOr
)

func (k LogicalOpNodeKind) String() string {
	switch k {
	case LogicalOpNodeKindAnd:
		return "&&"
	case LogicalOpNodeKindOr:
		return "||"
	default:
		return "INVALID LOGICAL OPERATOR"
	}
}

// LogicalOpNode is node for logical binary operators; && or ||.
type LogicalOpNode struct {
	// Kind is a kind to show which operator is used.
	Kind LogicalOpNodeKind
	// Left is an expression for left hand side of the binary operator.
	Left ExprNode
	// Right is an expression for right hand side of the binary operator.
	Right  ExprNode
	parent ExprNode
}

// Token returns the first token of the node. This method is useful to get position of this node.
func (n *LogicalOpNode) Token() *Token {
	return n.Left.Token()
}

// Parent returns the parent node of this node.
func (n *LogicalOpNode) Parent() ExprNode {
	return n.parent
}

// FuncCallNode represents function call in expression.
// Note that currently only calling builtin functions is supported.
type FuncCallNode struct {
	// Callee is a name of called function. This is string value because currently only built-in
	// functions can be called.
	Callee string
	// Args is arguments of the function call.
	Args   []ExprNode
	tok    *Token
	parent ExprNode
}

// Token returns the first token of the node. This method is useful to get position of this node.
func (n *FuncCallNode) Token() *Token {
	return n.tok
}

// Parent returns the parent node of this node.
func (n *FuncCallNode) Parent() ExprNode {
	return n.parent
}

// VisitExprNodeFunc is a visitor function for VisitExprNode(). The entering argument is set to
// true when it is called before visiting children. It is set to false when it is called after
// visiting children. It means that this function is called twice for the same node. The parent
// argument is the parent of the node. When the node is root, its parent is nil.
type VisitExprNodeFunc func(node, parent ExprNode, entering bool)

func visitExprNode(n, p ExprNode, f VisitExprNodeFunc) {
	f(n, p, true)
	switch n := n.(type) {
	case *ObjectDerefNode:
		visitExprNode(n.Receiver, n, f)
	case *ArrayDerefNode:
		visitExprNode(n.Receiver, n, f)
	case *IndexAccessNode:
		// Index must be visited before Operand to make UntrustedInputChecker work correctly.
		visitExprNode(n.Index, n, f)
		visitExprNode(n.Operand, n, f)
	case *NotOpNode:
		visitExprNode(n.Operand, n, f)
	case *CompareOpNode:
		visitExprNode(n.Left, n, f)
		visitExprNode(n.Right, n, f)
	case *LogicalOpNode:
		visitExprNode(n.Left, n, f)
		visitExprNode(n.Right, n, f)
	case *FuncCallNode:
		for i := range n.Args {
			visitExprNode(n.Args[i], n, f)
		}
	}
	f(n, p, false)
}

// VisitExprNode visits the given expression syntax tree with given function f.
func VisitExprNode(n ExprNode, f VisitExprNodeFunc) {
	visitExprNode(n, nil, f)
}

// FindParent applies predicate to each parent of this node until predicate returns true.
// Then it returns result of predicate. If no parent found, returns nil, false.
func FindParent[T ExprNode](n ExprNode, predicate func(n ExprNode) (T, bool)) (T, bool) {
	parent := n.Parent()
	for parent != nil {
		t, ok := predicate(parent)
		if ok {
			return t, true
		}
		parent = parent.Parent()
	}
	var zero T
	return zero, false
}
