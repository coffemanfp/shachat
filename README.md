# COMMENTS

The comments will be in english and will to have the follow structure:

* All comments will end in a point.
* All comments will be in line.

For the functions:

// Function - Functions things.
//  @param: param string:
//      Thing for thing.
//  @param: param int:
//      Thing for other thing.

For the vars and consts:

Multiples:

// Explication for the vars or consts group.
var (
    Thing1 string
    Thing2 int
)

One:

// Var name - Explication for the var.

For the structs:

// Strucname - Struct explication.
type Struct struct {
    Field string `json:"field"` // (Just for complicated fields) Field explication
}

For the interfaces:

// Interface name - Interface explication.
type I interface{
    Function() // Bla, bla, bla...
}