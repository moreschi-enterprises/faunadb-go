package faunadb

// Event's action types. Usually used as a parameter for Insert or Remove functions.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/events
const (
	ActionCreate = "create"
	ActionUpdate = "update"
	ActionDelete = "delete"
	ActionAdd    = "add"
	ActionRemove = "remove"
)

// Time unit. Usually used as a parameter for Time functions.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/epoch
const (
	TimeUnitDay         = "day"
	TimeUnitHalfDay     = "half day"
	TimeUnitHour        = "hour"
	TimeUnitMinute      = "minute"
	TimeUnitSecond      = "second"
	TimeUnitMillisecond = "millisecond"
	TimeUnitMicrosecond = "microsecond"
	TimeUnitNanosecond  = "nanosecond"
)

// Normalizers for Casefold
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/casefold
const (
	NormalizerNFKCCaseFold = "NFKCCaseFold"
	NormalizerNFC          = "NFC"
	NormalizerNFD          = "NFD"
	NormalizerNFKC         = "NFKC"
	NormalizerNFKD         = "NFKD"
)

// Helper functions

func varargs(expr ...interface{}) interface{} {
	if len(expr) == 1 {
		return expr[0]
	}

	return expr
}

// Optional parameters

// EventsOpt is an boolean optional parameter that describes if the query should include historical events.
// For more information about events, check https://docs.fauna.com/fauna/current/api/fql/types.
//
// Functions that accept this optional parameter are: Paginate.
//
// Deprecated: The Events function was renamed to EventsOpt to support the new history API.
// EventsOpt is provided here for backwards compatibility. Instead of using Paginate with the EventsOpt parameter,
// you should use the new Events function.
func EventsOpt(events interface{}) OptionalParameter {
	return func(fn unescapedObj) {
		fn["events"] = wrap(events)
	}
}

// TS is a timestamp optional parameter that specifies in which timestamp a query should be executed.
//
// Functions that accept this optional parameter are: Get, Insert, Remove, Exists, and Paginate.
func TS(timestamp interface{}) OptionalParameter {
	return func(fn unescapedObj) {
		fn["ts"] = wrap(timestamp)
	}
}

// After is an optional parameter used when cursoring that refers to the specified cursor's the next page, inclusive.
// For more information about pages, check https://docs.fauna.com/fauna/current/api/fql/types.
//
// Functions that accept this optional parameter are: Paginate.
func After(ref interface{}) OptionalParameter {
	return func(fn unescapedObj) {
		fn["after"] = wrap(ref)
	}
}

// Before is an optional parameter used when cursoring that refers to the specified cursor's previous page, exclusive.
// For more information about pages, check https://docs.fauna.com/fauna/current/api/fql/types.
//
// Functions that accept this optional parameter are: Paginate.
func Before(ref interface{}) OptionalParameter {
	return func(fn unescapedObj) {
		fn["before"] = wrap(ref)
	}
}

// Size is a numeric optional parameter that specifies the size of a pagination cursor.
//
// Functions that accept this optional parameter are: Paginate.
func Size(size interface{}) OptionalParameter {
	return func(fn unescapedObj) {
		fn["size"] = wrap(size)
	}
}

// Start is a numeric optional parameter that specifies the start of where to search.
//
// Functions that accept this optional parameter are: FindStr and FindStrRegex.
func Start(start interface{}) OptionalParameter {
	return func(fn unescapedObj) {
		fn["start"] = wrap(start)
	}
}

// StrLength is a numeric optional parameter that specifies the amount to copy.
//
// Functions that accept this optional parameter are: FindStr and FindStrRegex.
func StrLength(length interface{}) OptionalParameter {
	return func(fn unescapedObj) {
		fn["length"] = wrap(length)
	}
}

// OnlyFirst is a boolean optional parameter that only replace the first string
//
// Functions that accept this optional parameter are: ReplaceStrRegex
func OnlyFirst() OptionalParameter {
	return func(fn unescapedObj) {
		fn["first"] = BooleanV(true)
	}
}

// Sources is a boolean optional parameter that specifies if a pagination cursor should include
// the source sets along with each element.
//
// Functions that accept this optional parameter are: Paginate.
func Sources(sources interface{}) OptionalParameter {
	return func(fn unescapedObj) {
		fn["sources"] = wrap(sources)
	}
}

// Default is an optional parameter that specifies the default value for a select operation when
// the desired value path is absent.
//
// Functions that accept this optional parameter are: Select.
func Default(value interface{}) OptionalParameter {
	return func(fn unescapedObj) {
		fn["default"] = wrap(value)
	}
}

// Separator is a string optional parameter that specifies the separator for a concat operation.
//
// Functions that accept this optional parameter are: Concat.
func Separator(sep interface{}) OptionalParameter {
	return func(fn unescapedObj) {
		fn["separator"] = wrap(sep)
	}
}

// Precision is an optional parameter that specifies the precision for a Trunc and Round operations.
//
// Functions that accept this optional parameter are: Round and Trunc.
func Precision(precision interface{}) OptionalParameter {
	return func(fn unescapedObj) {
		fn["precision"] = wrap(precision)
	}
}

// ConflictResolver is an optional parameter that specifies the lambda for resolving Merge conflicts
//
// Functions that accept this optional parameter are: Merge
func ConflictResolver(lambda interface{}) OptionalParameter {
	return func(fn unescapedObj) {
		fn["lambda"] = wrap(lambda)
	}
}

// Normalizer is a string optional parameter that specifies the normalization function for casefold operation.
//
// Functions that accept this optional parameter are: Casefold.
func Normalizer(norm interface{}) OptionalParameter {
	return func(fn unescapedObj) {
		fn["normalizer"] = wrap(norm)
	}
}

// LetBuilder builds Let expressions
type LetBuilder struct {
	bindings unescapedArr
}

// Bind binds a variable name to a value and returns a LetBuilder
func (lb *LetBuilder) Bind(key string, in interface{}) *LetBuilder {
	binding := make(unescapedObj, 1)
	binding[key] = wrap(in)
	lb.bindings = append(lb.bindings, binding)
	return lb
}

// In sets the expression to be evaluated and returns the prepared Let.
func (lb *LetBuilder) In(in Expr) Expr {
	return fn2("let", lb.bindings, "in", in)
}

// Values

// Ref creates a new RefV value with the provided ID.
//
// Parameters:
//  id string - A string representation of a reference type.
//
// Returns:
//  Ref - A new reference type.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Ref
func Ref(id string) Expr { return fn1("@ref", id) }

// RefClass creates a new Ref based on the provided class and ID.
//
// Parameters:
//  classRef Ref - A class reference.
//  id string|int64 - The document ID.
//
// Deprecated: Use RefCollection instead, RefClass is kept for backwards compatibility
//
// Returns:
//  Ref - A new reference type.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/RefClass
func RefClass(classRef, id interface{}) Expr { return fn2("ref", classRef, "id", id) }

// RefCollection creates a new Ref based on the provided collection and ID.
//
// Parameters:
//  collectionRef Ref - A collection reference.
//  id string|int64 - The document ID.
//
// Returns:
//  Ref - A new reference type.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/RefCollection
func RefCollection(collectionRef, id interface{}) Expr { return fn2("ref", collectionRef, "id", id) }

// Null creates a NullV value.
//
// Returns:
//  Value - A null value.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Null
func Null() Expr { return NullV{} }

// Basic forms

// Abort aborts the execution of the query
//
// Parameters:
//  msg string - An error message.
//
// Returns:
//  Error
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Abort
func Abort(msg interface{}) Expr { return fn1("abort", msg) }

// Do sequentially evaluates its arguments, and returns the last expression.
// If no expressions are provided, do returns an error.
//
// Parameters:
//  exprs []Expr - A variable number of expressions to be evaluated.
//
// Returns:
//  Value - The result of the last expression in the list.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Do
func Do(exprs ...interface{}) Expr { return fn1("do", exprs) }

// If evaluates and returns then or elze depending on the value of cond.
// If cond evaluates to anything other than a boolean, if returns an “invalid argument” error
//
// Parameters:
//  cond bool - A boolean expression.
//  then Expr - The expression to run if condition is true.
//  elze Expr - The expression to run if condition is false.
//
// Returns:
//  Value - The result of either then or elze expression.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/If
func If(cond, then, elze interface{}) Expr { return fn3("if", cond, "then", then, "else", elze) }

// Lambda creates an anonymous function. Mostly used with Collection functions.
//
// Parameters:
//  varName string|[]string - A string or an array of strings of arguments name to be bound in the body of the lambda.
//  expr Expr - An expression used as the body of the lambda.
//
// Returns:
//  Value - The result of the body expression.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Lambda
func Lambda(varName, expr interface{}) Expr { return fn2("lambda", varName, "expr", expr) }

// At execute an expression at a given timestamp.
//
// Parameters:
//  timestamp time - The timestamp in which the expression will be evaluated.
//  expr Expr - An expression to be evaluated.
//
// Returns:
//  Value - The result of the given expression.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/At
func At(timestamp, expr interface{}) Expr { return fn2("at", timestamp, "expr", expr) }

// Let binds values to one or more variables.
//
// Returns:
//  *LetBuilder - Returns a LetBuilder.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Let
func Let() *LetBuilder { return &LetBuilder{nil} }

// Var refers to a value of a variable on the current lexical scope.
//
// Parameters:
//  name string - The variable name.
//
// Returns:
//  Value - The variable value.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Var
func Var(name string) Expr { return fn1("var", name) }

// Call invokes the specified function passing in a variable number of arguments
//
// Parameters:
//  ref Ref - The reference to the user defined functions to call.
//  args []Value - A series of values to pass as arguments to the user defined function.
//
// Returns:
//  Value - The return value of the user defined function.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Call
func Call(ref interface{}, args ...interface{}) Expr {
	return fn2("call", ref, "arguments", varargs(args...))
}

// Query creates an instance of the @query type with the specified lambda
//
// Parameters:
//  lambda Lambda - A lambda representation. See Lambda() function.
//
// Returns:
//  Query - The lambda wrapped in a @query type.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Query
func Query(lambda interface{}) Expr { return fn1("query", lambda) }

// Collections

// Map applies the lambda expression on each element of a collection or Page.
// It returns the result of each application on a collection of the same type.
//
// Parameters:
//  coll []Value - The collection of elements to iterate.
//  lambda Lambda - A lambda function to be applied to each element of the collection. See Lambda() function.
//
// Returns:
//  []Value - A new collection with elements transformed by the lambda function.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Map
func Map(coll, lambda interface{}) Expr { return fn2("map", lambda, "collection", coll) }

// Foreach applies the lambda expression on each element of a collection or Page.
// The original collection is returned.
//
// Parameters:
//  coll []Value - The collection of elements to iterate.
//  lambda Lambda - A lambda function to be applied to each element of the collection. See Lambda() function.
//
// Returns:
//  []Value - The original collection.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Foreach
func Foreach(coll, lambda interface{}) Expr { return fn2("foreach", lambda, "collection", coll) }

// Filter applies the lambda expression on each element of a collection or Page.
// It returns a new collection of the same type containing only the elements in which the
// function application returned true.
//
// Parameters:
//  coll []Value - The collection of elements to iterate.
//  lambda Lambda - A lambda function to be applied to each element of the collection. The lambda function must return a boolean value. See Lambda() function.
//
// Returns:
//  []Value - A new collection.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Filter
func Filter(coll, lambda interface{}) Expr { return fn2("filter", lambda, "collection", coll) }

// Take returns a new collection containing num elements from the head of the original collection.
//
// Parameters:
//  num int64 - The number of elements to take from the collection.
//  coll []Value - The collection of elements.
//
// Returns:
//  []Value - A new collection.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Take
func Take(num, coll interface{}) Expr { return fn2("take", num, "collection", coll) }

// Drop returns a new collection containing the remaining elements from the original collection
// after num elements have been removed.
//
// Parameters:
//  num int64 - The number of elements to drop from the collection.
//  coll []Value - The collection of elements.
//
// Returns:
//  []Value - A new collection.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Drop
func Drop(num, coll interface{}) Expr { return fn2("drop", num, "collection", coll) }

// Prepend returns a new collection that is the result of prepending elems to coll.
//
// Parameters:
//  elems []Value - Elements to add to the beginning of the other collection.
//  coll []Value - The collection of elements.
//
// Returns:
//  []Value - A new collection.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Prepend
func Prepend(elems, coll interface{}) Expr { return fn2("prepend", elems, "collection", coll) }

// Append returns a new collection that is the result of appending elems to coll.
//
// Parameters:
//  elems []Value - Elements to add to the end of the other collection.
//  coll []Value - The collection of elements.
//
// Returns:
//  []Value - A new collection.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Append
func Append(elems, coll interface{}) Expr { return fn2("append", elems, "collection", coll) }

// IsEmpty returns true if the collection is the empty set, else false.
//
// Parameters:
//  coll []Value - The collection of elements.
//
// Returns:
//   bool - True if the collection is empty, else false.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/IsEmpty
func IsEmpty(coll interface{}) Expr { return fn1("is_empty", coll) }

// IsNonEmpty returns false if the collection is the empty set, else true
//
// Parameters:
//  coll []Value - The collection of elements.
//
// Returns:
//   bool - True if the collection is not empty, else false.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/IsNonEmpty
func IsNonEmpty(coll interface{}) Expr { return fn1("is_nonempty", coll) }

// Read

// Get retrieves the document identified by the provided ref. Optional parameters: TS.
//
// Parameters:
//  ref Ref|SetRef - The reference to the object or a set reference.
//
// Optional parameters:
//  ts time - The snapshot time at which to get the document. See TS() function.
//
// Returns:
//  Object - The object requested.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Get
func Get(ref interface{}, options ...OptionalParameter) Expr { return fn1("get", ref, options...) }

// KeyFromSecret retrieves the key object from the given secret.
//
// Parameters:
//  secret string - The token secret.
//
// Returns:
//  Key - The key object related to the token secret.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/KeyFromSecret
func KeyFromSecret(secret interface{}) Expr { return fn1("key_from_secret", secret) }

// Exists returns boolean true if the provided ref exists (in the case of an document),
// or is non-empty (in the case of a set), and false otherwise. Optional parameters: TS.
//
// Parameters:
//  ref Ref - The reference to the object. It could be a document reference of a object reference like a collection.
//
// Optional parameters:
//  ts time - The snapshot time at which to check for the document's existence. See TS() function.
//
// Returns:
//  bool - true if the reference exists, false otherwise.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Exists
func Exists(ref interface{}, options ...OptionalParameter) Expr { return fn1("exists", ref, options...) }

// Paginate retrieves a page from the provided set.
//
// Parameters:
//  set SetRef - A set reference to paginate over. See Match() or MatchTerm() functions.
//
// Optional parameters:
//  after Cursor - Return the next page of results after this cursor (inclusive). See After() function.
//  before Cursor - Return the previous page of results before this cursor (exclusive). See Before() function.
//  sources bool - If true, include the source sets along with each element. See Sources() function.
//  ts time - The snapshot time at which to get the document. See TS() function.
//
// Returns:
//  Page - A page of elements.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Paginate
func Paginate(set interface{}, options ...OptionalParameter) Expr {
	return fn1("paginate", set, options...)
}

// Write

// Create creates an document of the specified collection.
//
// Parameters:
//  ref Ref - A collection reference.
//  params Object - An object with attributes of the document created.
//
// Returns:
//  Object - A new document of the collection referenced.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Create
func Create(ref, params interface{}) Expr { return fn2("create", ref, "params", params) }

// CreateClass creates a new class.
//
// Parameters:
//  params Object - An object with attributes of the class.
//
// Deprecated: Use CreateCollection instead, CreateClass is kept for backwards compatibility
//
// Returns:
//  Object - The new created class object.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/CreateClass
func CreateClass(params interface{}) Expr { return fn1("create_class", params) }

// CreateCollection creates a new collection.
//
// Parameters:
//  params Object - An object with attributes of the collection.
//
// Returns:
//  Object - The new created collection object.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/CreateCollection
func CreateCollection(params interface{}) Expr { return fn1("create_collection", params) }

// CreateDatabase creates an new database.
//
// Parameters:
//  params Object - An object with attributes of the database.
//
// Returns:
//  Object - The new created database object.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/CreateDatabase
func CreateDatabase(params interface{}) Expr { return fn1("create_database", params) }

// CreateIndex creates a new index.
//
// Parameters:
//  params Object - An object with attributes of the index.
//
// Returns:
//  Object - The new created index object.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/CreateIndex
func CreateIndex(params interface{}) Expr { return fn1("create_index", params) }

// CreateKey creates a new key.
//
// Parameters:
//  params Object - An object with attributes of the key.
//
// Returns:
//  Object - The new created key object.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/CreateKey
func CreateKey(params interface{}) Expr { return fn1("create_key", params) }

// CreateFunction creates a new function.
//
// Parameters:
//  params Object - An object with attributes of the function.
//
// Returns:
//  Object - The new created function object.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/CreateFunction
func CreateFunction(params interface{}) Expr { return fn1("create_function", params) }

// CreateRole creates a new role.
//
// Parameters:
//  params Object - An object with attributes of the role.
//
// Returns:
//  Object - The new created role object.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/CreateRole
func CreateRole(params interface{}) Expr { return fn1("create_role", params) }

// MoveDatabase moves a database to a new hierachy.
//
// Parameters:
//  from Object - Source reference to be moved.
//  to Object   - New parent database reference.
//
// Returns:
//  Object - instance.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/MoveDatabase
func MoveDatabase(from interface{}, to interface{}) Expr { return fn2("move_database", from, "to", to) }

// Update updates the provided document.
//
// Parameters:
//  ref Ref - The reference to update.
//  params Object - An object representing the parameters of the document.
//
// Returns:
//  Object - The updated object.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Update
func Update(ref, params interface{}) Expr { return fn2("update", ref, "params", params) }

// Replace replaces the provided document.
//
// Parameters:
//  ref Ref - The reference to replace.
//  params Object - An object representing the parameters of the document.
//
// Returns:
//  Object - The replaced object.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Replace
func Replace(ref, params interface{}) Expr { return fn2("replace", ref, "params", params) }

// Delete deletes the provided document.
//
// Parameters:
//  ref Ref - The reference to delete.
//
// Returns:
//  Object - The deleted object.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Delete
func Delete(ref interface{}) Expr { return fn1("delete", ref) }

// Insert adds an event to the provided document's history.
//
// Parameters:
//  ref Ref - The reference to insert against.
//  ts time - The valid time of the inserted event.
//  action string - Whether the event shoulde be a ActionCreate, ActionUpdate or ActionDelete.
//
// Returns:
//  Object - The deleted object.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Insert
func Insert(ref, ts, action, params interface{}) Expr {
	return fn4("insert", ref, "ts", ts, "action", action, "params", params)
}

// Remove deletes an event from the provided document's history.
//
// Parameters:
//  ref Ref - The reference of the document whose event should be removed.
//  ts time - The valid time of the inserted event.
//  action string - The event action (ActionCreate, ActionUpdate or ActionDelete) that should be removed.
//
// Returns:
//  Object - The deleted object.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Remove
func Remove(ref, ts, action interface{}) Expr { return fn3("remove", ref, "ts", ts, "action", action) }

// String

// Format formats values into a string.
//
// Parameters:
//  format string - format a string with format specifiers.
//
// Optional parameters:
//  values []string - list of values to format into string.
//
// Returns:
//  string - A string.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Format
func Format(format interface{}, values ...interface{}) Expr {
	return fn2("format", format, "values", varargs(values...))
}

// Concat concatenates a list of strings into a single string.
//
// Parameters:
//  terms []string - A list of strings to concatenate.
//
// Optional parameters:
//  separator string - The separator to use between each string. See Separator() function.
//
// Returns:
//  string - A string with all terms concatenated.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Concat
func Concat(terms interface{}, options ...OptionalParameter) Expr {
	return fn1("concat", terms, options...)
}

// Casefold normalizes strings according to the Unicode Standard section 5.18 "Case Mappings".
//
// Parameters:
//  str string - The string to casefold.
//
// Optional parameters:
//  normalizer string - The algorithm to use. One of: NormalizerNFKCCaseFold, NormalizerNFC, NormalizerNFD, NormalizerNFKC, NormalizerNFKD.
//
// Returns:
//  string - The normalized string.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Casefold
func Casefold(str interface{}, options ...OptionalParameter) Expr {
	return fn1("casefold", str, options...)
}

// StartsWith returns true if the string starts with the given prefix value, or false if otherwise
//
// Parameters:
//
//  value  string -  the string to evaluate
//  search string -  the prefix to search for
//
// Returns:
//   boolean       - does `value` start with `search
//
// See https://docs.fauna.com/fauna/current/api/fql/functions/startswith
func StartsWith(value interface{}, search interface{}) Expr {
	return fn2("startswith", value, "search", search)
}

// EndsWith returns true if the string ends with the given suffix value, or false if otherwise
//
// Parameters:
//
// value  string  -  the string to evaluate
// search  string -  the suffix to search for
//
// Returns:
// boolean       - does `value` end with `search`
//
// See https://docs.fauna.com/fauna/current/api/fql/functions/endswith
func EndsWith(value interface{}, search interface{}) Expr {
	return fn2("endswith", value, "search", search)
}

// ContainsStr returns true if the string contains the given substring, or false if otherwise
//
// Parameters:
//
// value string  -  the string to evaluate
// search string -  the substring to search for
//
// Returns:
// boolean      - was the search result found
//
// See https://docs.fauna.com/fauna/current/api/fql/functions/containsstr
func ContainsStr(value interface{}, search interface{}) Expr {
	return fn2("containsstr", value, "search", search)
}

// ContainsStrRegex returns true if the string contains the given pattern, or false if otherwise
//
// Parameters:
//
// value   string      -  the string to evaluate
// pattern string      -  the pattern to search for
//
// Returns:
// boolean      - was the search result found
//
// See https://docs.fauna.com/fauna/current/api/fql/functions/containsstrregex
func ContainsStrRegex(value interface{}, pattern interface{}) Expr {
	return fn2("containsstrregex", value, "pattern", pattern)
}

// RegexEscape It takes a string and returns a regex which matches the input string verbatim.
//
// Parameters:
//
// value  string     - the string to analyze
// pattern       -  the pattern to search for
//
// Returns:
// boolean      - was the search result found
//
// See https://docs.fauna.com/fauna/current/api/fql/functions/regexescape
func RegexEscape(value interface{}) Expr {
	return fn1("regexescape", value)
}

// FindStr locates a substring in a source string.  Optional parameters: Start
//
// Parameters:
//  str string  - The source string
//  find string - The string to locate
//
// Optional parameters:
//  start int - a position to start the search. See Start() function.
//
// Returns:
//  string - The offset of where the substring starts or -1 if not found
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/FindStr
func FindStr(str, find interface{}, options ...OptionalParameter) Expr {
	return fn2("findstr", str, "find", find, options...)
}

// FindStrRegex locates a java regex pattern in a source string.  Optional parameters: Start
//
// Parameters:
//  str string      - The sourcestring
//  pattern string  - The pattern to locate.
//
// Optional parameters:
//  start long - a position to start the search.  See Start() function.
//
// Returns:
//  string - The offset of where the substring starts or -1 if not found
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/FindStrRegex
func FindStrRegex(str, pattern interface{}, options ...OptionalParameter) Expr {
	return fn2("findstrregex", str, "pattern", pattern, options...)
}

// Length finds the length of a string in codepoints
//
// Parameters:
//  str string - A string to find the length in codepoints
//
// Returns:
//  int - A length of a string.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Length
func Length(str interface{}) Expr { return fn1("length", str) }

// LowerCase changes all characters in the string to lowercase
//
// Parameters:
//  str string - A string to convert to lowercase
//
// Returns:
//  string - A string in lowercase.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/LowerCase
func LowerCase(str interface{}) Expr { return fn1("lowercase", str) }

// LTrim returns a string wtih leading white space removed.
//
// Parameters:
//  str string - A string to remove leading white space
//
// Returns:
//  string - A string with all leading white space removed
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/LTrim
func LTrim(str interface{}) Expr { return fn1("ltrim", str) }

// Repeat returns a string wtih repeated n times
//
// Parameters:
//  str string - A string to repeat
//  number int - The number of times to repeat the string
//
// Returns:
//  string - A string concatendanted the specified number of times
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Repeat
func Repeat(str, number interface{}) Expr { return fn2("repeat", str, "number", number) }

// ReplaceStr returns a string with every occurence of the "find" string changed to "replace" string
//
// Parameters:
//  str string     - A source string
//  find string    - The substring to locate in in the source string
//  replace string - The string to replaice the "find" string when located
//
// Returns:
//  string - returns a string with every occurence of the "find" string changed to "replace"
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/ReplaceStr
func ReplaceStr(str, find, replace interface{}) Expr {
	return fn3("replacestr", str, "find", find, "replace", replace)
}

// ReplaceStrRegex returns a string with occurence(s) of the java regular expression "pattern" changed to "replace" string.   Optional parameters: OnlyFirst
//
// Parameters:
//  value string   - The source string
//  pattern string - A java regular expression to locate
//  replace string - The string to replace the pattern when located
//
// Optional parameters:
//  OnlyFirst - Only replace the first found pattern.  See OnlyFirst() function.
//
// Returns:
//  string - A string with occurence(s) of the java regular expression "pattern" changed to "replace" string
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/ReplaceStrRegex
func ReplaceStrRegex(value, pattern, replace interface{}, options ...OptionalParameter) Expr {
	return fn3("replacestrregex", value, "pattern", pattern, "replace", replace, options...)
}

// RTrim returns a string wtih trailing white space removed.
//
// Parameters:
//  str string - A string to remove trailing white space
//
// Returns:
//  string - A string with all trailing white space removed
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/RTrim
func RTrim(str interface{}) Expr { return fn1("rtrim", str) }

// Space function returns "N" number of spaces
//
// Parameters:
//  value int - the number of spaces
//
// Returns:
//  string - function returns string with n spaces
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Space
func Space(value interface{}) Expr { return fn1("space", value) }

// SubString returns a subset of the source string.   Optional parameters: StrLength
//
// Parameters:
//  str string - A source string
//  start int  - The position in the source string where SubString starts extracting characters
//
// Optional parameters:
//  StrLength int - A value for the length of the extracted substring. See StrLength() function.
//
// Returns:
//  string - function returns a subset of the source string
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/SubString
func SubString(str, start interface{}, options ...OptionalParameter) Expr {
	return fn2("substring", str, "start", start, options...)
}

// TitleCase changes all characters in the string to TitleCase
//
// Parameters:
//  str string - A string to convert to TitleCase
//
// Returns:
//  string - A string in TitleCase.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/TitleCase
func TitleCase(str interface{}) Expr { return fn1("titlecase", str) }

// Trim returns a string wtih trailing white space removed.
//
// Parameters:
//  str string - A string to remove trailing white space
//
// Returns:
//  string - A string with all trailing white space removed
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Trim
func Trim(str interface{}) Expr { return fn1("trim", str) }

// UpperCase changes all characters in the string to uppercase
//
// Parameters:
//  string - A string to convert to uppercase
//
// Returns:
//  string - A string in uppercase.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/UpperCase
func UpperCase(str interface{}) Expr { return fn1("uppercase", str) }

// Time and Date

// Time constructs a time from a ISO 8601 offset date/time string.
//
// Parameters:
//  str string - A string to convert to a time object.
//
// Returns:
//  time - A time object.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Time
func Time(str interface{}) Expr { return fn1("time", str) }

// TimeAdd returns a new time or date with the offset in terms of the unit
// added.
//
// Parameters:
// base        -  the base time or data
// offset      -  the number of units
// unit        -  the unit type
//
// Returns:
// Expr
//
//See: https://docs.fauna.com/fauna/current/api/fql/functions/timeadd
func TimeAdd(base interface{}, offset interface{}, unit interface{}) Expr {
	return fn3(
		"time_add", base,
		"offset", offset,
		"unit", unit,
	)
}

// TimeSubtract returns a new time or date with the offset in terms of the unit
// subtracted.
//
// Parameters:
// base        -  the base time or data
// offset      -  the number of units
// unit        -  the unit type
//
// Returns:
// Expr
//
//See: https://docs.fauna.com/fauna/current/api/fql/functions/timesubtract
func TimeSubtract(base interface{}, offset interface{}, unit interface{}) Expr {
	return fn3(
		"time_subtract", base,
		"offset", offset,
		"unit", unit,
	)
}

// TimeDiff returns the number of intervals in terms of the unit between
// two times or dates. Both start and finish must be of the same
// type.
//
// Parameters:
//   start the starting time or date, inclusive
//   finish the ending time or date, exclusive
//   unit the unit type//
// Returns:
// Expr
//
//See: https://docs.fauna.com/fauna/current/api/fql/functions/timediff
func TimeDiff(start interface{}, finish interface{}, unit interface{}) Expr {
	return fn3(
		"time_diff", start,
		"other", finish,
		"unit", unit,
	)
}

// Date constructs a date from a ISO 8601 offset date/time string.
//
// Parameters:
//  str string - A string to convert to a date object.
//
// Returns:
//  date - A date object.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Date
func Date(str interface{}) Expr { return fn1("date", str) }

// Epoch constructs a time relative to the epoch "1970-01-01T00:00:00Z".
//
// Parameters:
//  num int64 - The number of units from Epoch.
//  unit string - The unit of number. One of TimeUnitSecond, TimeUnitMillisecond, TimeUnitMicrosecond, TimeUnitNanosecond.
//
// Returns:
//  time - A time object.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Epoch
func Epoch(num, unit interface{}) Expr { return fn2("epoch", num, "unit", unit) }

// Set

// Singleton returns the history of the document's presence of the provided ref.
//
// Parameters:
//  ref Ref - The reference of the document for which to retrieve the singleton set.
//
// Returns:
//  SetRef - The singleton SetRef.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Singleton
func Singleton(ref interface{}) Expr { return fn1("singleton", ref) }

// Events returns the history of document's data of the provided ref.
//
// Parameters:
//  refSet Ref|SetRef - A reference or set reference to retrieve an event set from.
//
// Returns:
//  SetRef - The events SetRef.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Events
func Events(refSet interface{}) Expr { return fn1("events", refSet) }

// Match returns the set of documents for the specified index.
//
// Parameters:
//  ref Ref - The reference of the index to match against.
//
// Returns:
//  SetRef
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Match
func Match(ref interface{}) Expr { return fn1("match", ref) }

// MatchTerm returns th set of documents that match the terms in an index.
//
// Parameters:
//  ref Ref - The reference of the index to match against.
//  terms []Value - A list of terms used in the match.
//
// Returns:
//  SetRef
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/MatchTerm
func MatchTerm(ref, terms interface{}) Expr { return fn2("match", ref, "terms", terms) }

// Union returns the set of documents that are present in at least one of the specified sets.
//
// Parameters:
//  sets []SetRef - A list of SetRef to union together.
//
// Returns:
//  SetRef
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Union
func Union(sets ...interface{}) Expr { return fn1("union", varargs(sets...)) }

// Merge two or more objects..
//
// Parameters:
//   merge merge the first object.
//   with the second object or a list of objects
//   lambda a lambda to resolve possible conflicts
//
// Returns:
// merged object
//
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Merge
func Merge(merge interface{}, with interface{}, lambda ...OptionalParameter) Expr {
	return fn2("merge", merge, "with", with, lambda...)
}

// Reduce function applies a reducer Lambda function serially to each member of the collection to produce a single value.
//
// Parameters:
// lambda     Expr  - The accumulator function
// initial    Expr  - The initial value
// collection Expr  - The collection to be reduced
//
// Returns:
// Expr
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/reduce
func Reduce(lambda, initial interface{}, collection interface{}) Expr {
	return fn3("reduce", lambda, "initial", initial, "collection", collection)
}

// Intersection returns the set of documents that are present in all of the specified sets.
//
// Parameters:
//  sets []SetRef - A list of SetRef to intersect.
//
// Returns:
//  SetRef
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Intersection
func Intersection(sets ...interface{}) Expr { return fn1("intersection", varargs(sets...)) }

// Difference returns the set of documents that are present in the first set but not in
// any of the other specified sets.
//
// Parameters:
//  sets []SetRef - A list of SetRef to diff.
//
// Returns:
//  SetRef
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Difference
func Difference(sets ...interface{}) Expr { return fn1("difference", varargs(sets...)) }

// Distinct returns the set of documents with duplicates removed.
//
// Parameters:
//  set []SetRef - A list of SetRef to remove duplicates from.
//
// Returns:
//  SetRef
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Distinct
func Distinct(set interface{}) Expr { return fn1("distinct", set) }

// Join derives a set of resources by applying each document in the source set to the target set.
//
// Parameters:
//  source SetRef - A SetRef of the source set.
//  target Lambda - A Lambda that will accept each element of the source Set and return a Set.
//
// Returns:
//  SetRef
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Join
func Join(source, target interface{}) Expr { return fn2("join", source, "with", target) }

// Range filters the set based on the lower/upper bounds (inclusive).
//
// Parameters:
//  set SetRef - Set to be filtered.
//  from - lower bound.
//  to - upper bound
//
// Returns:
//  SetRef
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Range
func Range(set interface{}, from interface{}, to interface{}) Expr {
	return fn3("range", set, "from", from, "to", to)
}

// Authentication

// Login creates a token for the provided ref.
//
// Parameters:
//  ref Ref - A reference with credentials to authenticate against.
//  params Object - An object of parameters to pass to the login function
//    - password: The password used to login
//
// Returns:
//  Key - a key with the secret to login.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Login
func Login(ref, params interface{}) Expr { return fn2("login", ref, "params", params) }

// Logout deletes the current session token. If invalidateAll is true, logout will delete all tokens associated with the current session.
//
// Parameters:
//  invalidateAll bool - If true, log out all tokens associated with the current session.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Logout
func Logout(invalidateAll interface{}) Expr { return fn1("logout", invalidateAll) }

// Identify checks the given password against the provided ref's credentials.
//
// Parameters:
//  ref Ref - The reference to check the password against.
//  password string - The credentials password to check.
//
// Returns:
//  bool - true if the password is correct, false otherwise.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Identify
func Identify(ref, password interface{}) Expr { return fn2("identify", ref, "password", password) }

// Identity returns the document reference associated with the current key.
//
// For example, the current key token created using:
//	Create(Tokens(), Obj{"document": someRef})
// or via:
//	Login(someRef, Obj{"password":"sekrit"})
// will return "someRef" as the result of this function.
//
// Returns:
//  Ref - The reference associated with the current key.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Identity
func Identity() Expr { return fn1("identity", NullV{}) }

// HasIdentity checks if the current key has an identity associated to it.
//
// Returns:
//  bool - true if the current key has an identity, false otherwise.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/HasIdentity
func HasIdentity() Expr { return fn1("has_identity", NullV{}) }

// Miscellaneous

// NextID produces a new identifier suitable for use when constructing refs.
//
// Deprecated: Use NewId instead
//
// Returns:
//  string - The new ID.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/NextID
func NextID() Expr { return fn1("new_id", NullV{}) }

// NewId produces a new identifier suitable for use when constructing refs.
//
// Returns:
//  string - The new ID.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/NewId
func NewId() Expr { return fn1("new_id", NullV{}) }

// Database creates a new database ref.
//
// Parameters:
//  name string - The name of the database.
//
// Returns:
//  Ref - The database reference.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Database
func Database(name interface{}) Expr { return fn1("database", name) }

// ScopedDatabase creates a new database ref inside a database.
//
// Parameters:
//  name string - The name of the database.
//  scope Ref - The reference of the database's scope.
//
// Returns:
//  Ref - The database reference.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/ScopedDatabase
func ScopedDatabase(name interface{}, scope interface{}) Expr {
	return fn2("database", name, "scope", scope)
}

// Index creates a new index ref.
//
// Parameters:
//  name string - The name of the index.
//
// Returns:
//  Ref - The index reference.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Index
func Index(name interface{}) Expr { return fn1("index", name) }

// ScopedIndex creates a new index ref inside a database.
//
// Parameters:
//  name string - The name of the index.
//  scope Ref - The reference of the index's scope.
//
// Returns:
//  Ref - The index reference.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/ScopedIndex
func ScopedIndex(name interface{}, scope interface{}) Expr { return fn2("index", name, "scope", scope) }

// Class creates a new class ref.
//
// Parameters:
//  name string - The name of the class.
//
// Deprecated: Use Collection instead, Class is kept for backwards compatibility
//
// Returns:
//  Ref - The class reference.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Class
func Class(name interface{}) Expr { return fn1("class", name) }

// Collection creates a new collection ref.
//
// Parameters:
//  name string - The name of the collection.
//
// Returns:
//  Ref - The collection reference.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Collection
func Collection(name interface{}) Expr { return fn1("collection", name) }

// ScopedClass creates a new class ref inside a database.
//
// Parameters:
//  name string - The name of the class.
//  scope Ref - The reference of the class's scope.
//
// Deprecated: Use ScopedCollection instead, ScopedClass is kept for backwards compatibility
//
// Returns:
//  Ref - The collection reference.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/ScopedClass
func ScopedClass(name interface{}, scope interface{}) Expr {
	return fn2("class", name, "scope", scope)
}

// ScopedCollection creates a new collection ref inside a database.
//
// Parameters:
//  name string - The name of the collection.
//  scope Ref - The reference of the collection's scope.
//
// Returns:
//  Ref - The collection reference.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/ScopedCollection
func ScopedCollection(name interface{}, scope interface{}) Expr {
	return fn2("collection", name, "scope", scope)
}

// Function create a new function ref.
//
// Parameters:
//  name string - The name of the functions.
//
// Returns:
//  Ref - The function reference.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Function
func Function(name interface{}) Expr { return fn1("function", name) }

// ScopedFunction creates a new function ref inside a database.
//
// Parameters:
//  name string - The name of the function.
//  scope Ref - The reference of the function's scope.
//
// Returns:
//  Ref - The function reference.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/ScopedFunction
func ScopedFunction(name interface{}, scope interface{}) Expr {
	return fn2("function", name, "scope", scope)
}

// Role create a new role ref.
//
// Parameters:
//  name string - The name of the role.
//
// Returns:
//  Ref - The role reference.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Role
func Role(name interface{}) Expr { return fn1("role", name) }

// ScopedRole create a new role ref.
//
// Parameters:
//  name string - The name of the role.
//  scope Ref - The reference of the role's scope.
//
// Returns:
//  Ref - The role reference.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/ScopedRole
func ScopedRole(name, scope interface{}) Expr { return fn2("role", name, "scope", scope) }

// Classes creates a native ref for classes.
//
// Deprecated: Use Collections instead, Classes is kept for backwards compatibility
//
// Returns:
//  Ref - The reference of the class set.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Classes
func Classes() Expr { return fn1("classes", NullV{}) }

// Collections creates a native ref for collections.
//
// Returns:
//  Ref - The reference of the collections set.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Collections
func Collections() Expr { return fn1("collections", NullV{}) }

// ScopedClasses creates a native ref for classes inside a database.
//
// Parameters:
//  scope Ref - The reference of the class set's scope.
//
// Deprecated: Use ScopedCollections instead, ScopedClasses is kept for backwards compatibility
//
// Returns:
//  Ref - The reference of the class set.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/ScopedClasses
func ScopedClasses(scope interface{}) Expr { return fn1("classes", scope) }

// ScopedCollections creates a native ref for collections inside a database.
//
// Parameters:
//  scope Ref - The reference of the collections set's scope.
//
// Returns:
//  Ref - The reference of the collections set.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/ScopedCollections
func ScopedCollections(scope interface{}) Expr { return fn1("collections", scope) }

// Indexes creates a native ref for indexes.
//
// Returns:
//  Ref - The reference of the index set.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Indexes
func Indexes() Expr { return fn1("indexes", NullV{}) }

// ScopedIndexes creates a native ref for indexes inside a database.
//
// Parameters:
//  scope Ref - The reference of the index set's scope.
//
// Returns:
//  Ref - The reference of the index set.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/ScopedIndexes
func ScopedIndexes(scope interface{}) Expr { return fn1("indexes", scope) }

// Databases creates a native ref for databases.
//
// Returns:
//  Ref - The reference of the datbase set.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Databases
func Databases() Expr { return fn1("databases", NullV{}) }

// ScopedDatabases creates a native ref for databases inside a database.
//
// Parameters:
//  scope Ref - The reference of the database set's scope.
//
// Returns:
//  Ref - The reference of the database set.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/ScopedDatabases
func ScopedDatabases(scope interface{}) Expr { return fn1("databases", scope) }

// Functions creates a native ref for functions.
//
// Returns:
//  Ref - The reference of the function set.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Functions
func Functions() Expr { return fn1("functions", NullV{}) }

// ScopedFunctions creates a native ref for functions inside a database.
//
// Parameters:
//  scope Ref - The reference of the function set's scope.
//
// Returns:
//  Ref - The reference of the function set.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/ScopedFunctions
func ScopedFunctions(scope interface{}) Expr { return fn1("functions", scope) }

// Roles creates a native ref for roles.
//
// Returns:
//  Ref - The reference of the roles set.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Roles
func Roles() Expr { return fn1("roles", NullV{}) }

// ScopedRole creates a native ref for roles inside a database.
//
// Parameters:
//  scope Ref - The reference of the role set's scope.
//
// Returns:
//  Ref - The reference of the role set.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/ScopedRoles
func ScopedRoles(scope interface{}) Expr { return fn1("roles", scope) }

// Keys creates a native ref for keys.
//
// Returns:
//  Ref - The reference of the key set.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Keys
func Keys() Expr { return fn1("keys", NullV{}) }

// ScopedKeys creates a native ref for keys inside a database.
//
// Parameters:
//  scope Ref - The reference of the key set's scope.
//
// Returns:
//  Ref - The reference of the key set.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/ScopedKeys
func ScopedKeys(scope interface{}) Expr { return fn1("keys", scope) }

// Tokens creates a native ref for tokens.
//
// Returns:
//  Ref - The reference of the token set.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Tokens
func Tokens() Expr { return fn1("tokens", NullV{}) }

// ScopedTokens creates a native ref for tokens inside a database.
//
// Parameters:
//  scope Ref - The reference of the token set's scope.
//
// Returns:
//  Ref - The reference of the token set.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/ScopedTokens
func ScopedTokens(scope interface{}) Expr { return fn1("tokens", scope) }

// Credentials creates a native ref for credentials.
//
// Returns:
//  Ref - The reference of the credential set.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Credentials
func Credentials() Expr { return fn1("credentials", NullV{}) }

// ScopedCredentials creates a native ref for credentials inside a database.
//
// Parameters:
//  scope Ref - The reference of the credential set's scope.
//
// Returns:
//  Ref - The reference of the credential set.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/ScopedCredentials
func ScopedCredentials(scope interface{}) Expr { return fn1("credentials", scope) }

// Equals checks if all args are equivalents.
//
// Parameters:
//  args []Value - A collection of expressions to check for equivalence.
//
// Returns:
//  bool - true if all elements are equals, false otherwise.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Equals
func Equals(args ...interface{}) Expr { return fn1("equals", varargs(args...)) }

// Contains checks if the provided value contains the path specified.
//
// Parameters:
//  path Path - An array representing a path to check for the existence of. Path can be either strings or ints.
//  value Object - An object to search against.
//
// Returns:
//  bool - true if the path contains any value, false otherwise.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Contains
func Contains(path, value interface{}) Expr { return fn2("contains", path, "in", value) }

// Abs computes the absolute value of a number.
//
// Parameters:
//  value number - The number to take the absolute value of
//
// Returns:
//  number - The abosulte value of a number
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Abs
func Abs(value interface{}) Expr { return fn1("abs", value) }

// Acos computes the arccosine of a number.
//
// Parameters:
//  value number - The number to take the arccosine of
//
// Returns:
//  number - The arccosine of a number
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Acos
func Acos(value interface{}) Expr { return fn1("acos", value) }

// Asin computes the arcsine of a number.
//
// Parameters:
//  value number - The number to take the arcsine of
//
// Returns:
//  number - The arcsine of a number
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Asin
func Asin(value interface{}) Expr { return fn1("asin", value) }

// Atan computes the arctan of a number.
//
// Parameters:
//  value number - The number to take the arctan of
//
// Returns:
//  number - The arctan of a number
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Atan
func Atan(value interface{}) Expr { return fn1("atan", value) }

// Add computes the sum of a list of numbers.
//
// Parameters:
//  args []number - A collection of numbers to sum together.
//
// Returns:
//  number - The sum of all elements.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Add
func Add(args ...interface{}) Expr { return fn1("add", varargs(args...)) }

// BitAnd computes the and of a list of numbers.
//
// Parameters:
//  args []number - A collection of numbers to and together.
//
// Returns:
//  number - The and of all elements.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/BitAnd
func BitAnd(args ...interface{}) Expr { return fn1("bitand", varargs(args...)) }

// BitNot computes the 2's complement of a number
//
// Parameters:
//  value number - A numbers to not
//
// Returns:
//  number - The not of an element
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/BitNot
func BitNot(value interface{}) Expr { return fn1("bitnot", value) }

// BitOr computes the OR of a list of numbers.
//
// Parameters:
//  args []number - A collection of numbers to OR together.
//
// Returns:
//  number - The OR of all elements.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/BitOr
func BitOr(args ...interface{}) Expr { return fn1("bitor", varargs(args...)) }

// BitXor computes the XOR of a list of numbers.
//
// Parameters:
//  args []number - A collection of numbers to XOR together.
//
// Returns:
//  number - The XOR of all elements.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/BitXor
func BitXor(args ...interface{}) Expr { return fn1("bitxor", varargs(args...)) }

// Ceil computes the largest integer greater than or equal to
//
// Parameters:
//  value number - A numbers to compute the ceil of
//
// Returns:
//  number - The ceil of value
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Ceil
func Ceil(value interface{}) Expr { return fn1("ceil", value) }

// Cos computes the Cosine of a number
//
// Parameters:
//  value number - A number to compute the cosine of
//
// Returns:
//  number - The cosine of value
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Cos
func Cos(value interface{}) Expr { return fn1("cos", value) }

// Cosh computes the Hyperbolic Cosine of a number
//
// Parameters:
//  value number - A number to compute the Hyperbolic cosine of
//
// Returns:
//  number - The Hyperbolic cosine of value
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Cosh
func Cosh(value interface{}) Expr { return fn1("cosh", value) }

// Degrees computes the degress of a number
//
// Parameters:
//  value number - A number to compute the degress of
//
// Returns:
//  number - The degrees of value
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Degrees
func Degrees(value interface{}) Expr { return fn1("degrees", value) }

// Divide computes the quotient of a list of numbers.
//
// Parameters:
//  args []number - A collection of numbers to compute the quotient of.
//
// Returns:
//  number - The quotient of all elements.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Divide
func Divide(args ...interface{}) Expr { return fn1("divide", varargs(args...)) }

// Exp computes the Exp of a number
//
// Parameters:
//  value number - A number to compute the exp of
//
// Returns:
//  number - The exp of value
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Exp
func Exp(value interface{}) Expr { return fn1("exp", value) }

// Floor computes the Floor of a number
//
// Parameters:
//  value number - A number to compute the Floor of
//
// Returns:
//  number - The Floor of value
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Floor
func Floor(value interface{}) Expr { return fn1("floor", value) }

// Hypot computes the Hypotenuse of two numbers
//
// Parameters:
//  a number - A side of a right triangle
//  b number - A side of a right triangle
//
// Returns:
//  number - The hypotenuse of value
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Hypot
func Hypot(a, b interface{}) Expr { return fn2("hypot", a, "b", b) }

// Ln computes the natural log of a number
//
// Parameters:
//  value number - A number to compute the natural log of
//
// Returns:
//  number - The ln of value
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Ln
func Ln(value interface{}) Expr { return fn1("ln", value) }

// Log computes the Log of a number
//
// Parameters:
//  value number - A number to compute the Log of
//
// Returns:
//  number - The Log of value
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Log
func Log(value interface{}) Expr { return fn1("log", value) }

// Max computes the max of a list of numbers.
//
// Parameters:
//  args []number - A collection of numbers to find the max of.
//
// Returns:
//  number - The max of all elements.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Max
func Max(args ...interface{}) Expr { return fn1("max", varargs(args...)) }

// Min computes the Min of a list of numbers.
//
// Parameters:
//  args []number - A collection of numbers to find the min of.
//
// Returns:
//  number - The min of all elements.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Min
func Min(args ...interface{}) Expr { return fn1("min", varargs(args...)) }

// Modulo computes the reminder after the division of a list of numbers.
//
// Parameters:
//  args []number - A collection of numbers to compute the quotient of. The remainder will be returned.
//
// Returns:
//  number - The remainder of the quotient of all elements.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Modulo
func Modulo(args ...interface{}) Expr { return fn1("modulo", varargs(args...)) }

// Multiply computes the product of a list of numbers.
//
// Parameters:
//  args []number - A collection of numbers to multiply together.
//
// Returns:
//  number - The multiplication of all elements.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Multiply
func Multiply(args ...interface{}) Expr { return fn1("multiply", varargs(args...)) }

// Pow computes the Power of a number
//
// Parameters:
//  base number - A number which is the base
//  exp number  - A number which is the exponent
//
// Returns:
//  number - The Pow of value
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Pow
func Pow(base, exp interface{}) Expr { return fn2("pow", base, "exp", exp) }

// Radians computes the Radians of a number
//
// Parameters:
//  value number - A number which is convert to radians
//
// Returns:
//  number - The Radians of value
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Radians
func Radians(value interface{}) Expr { return fn1("radians", value) }

// Round a number at the given percission
//
// Parameters:
//  value number - The number to truncate
//  precision number - precision where to truncate, defaults is 2
//
// Returns:
//  number - The Rounded value.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Round
func Round(value interface{}, options ...OptionalParameter) Expr {
	return fn1("round", value, options...)
}

// Sign computes the Sign of a number
//
// Parameters:
//  value number - A number to compute the Sign of
//
// Returns:
//  number - The Sign of value
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Sign
func Sign(value interface{}) Expr { return fn1("sign", value) }

// Sin computes the Sine of a number
//
// Parameters:
//  value number - A number to compute the Sine of
//
// Returns:
//  number - The Sine of value
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Sin
func Sin(value interface{}) Expr { return fn1("sin", value) }

// Sinh computes the Hyperbolic Sine of a number
//
// Parameters:
//  value number - A number to compute the Hyperbolic Sine of
//
// Returns:
//  number - The Hyperbolic Sine of value
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Sinh
func Sinh(value interface{}) Expr { return fn1("sinh", value) }

// Sqrt computes the square root of a number
//
// Parameters:
//  value number - A number to compute the square root of
//
// Returns:
//  number - The square root of value
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Sqrt
func Sqrt(value interface{}) Expr { return fn1("sqrt", value) }

// Subtract computes the difference of a list of numbers.
//
// Parameters:
//  args []number - A collection of numbers to compute the difference of.
//
// Returns:
//  number - The difference of all elements.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Subtract
func Subtract(args ...interface{}) Expr { return fn1("subtract", varargs(args...)) }

// Tan computes the Tangent of a number
//
// Parameters:
//  value number - A number to compute the Tangent of
//
// Returns:
//  number - The Tangent of value
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Tan
func Tan(value interface{}) Expr { return fn1("tan", value) }

// Tanh computes the Hyperbolic Tangent of a number
//
// Parameters:
//  value number - A number to compute the Hyperbolic Tangent of
//
// Returns:
//  number - The Hyperbolic Tangent of value
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Tanh
func Tanh(value interface{}) Expr { return fn1("tanh", value) }

// Trunc truncates a number at the given percission
//
// Parameters:
//  value number - The number to truncate
//  precision number - precision where to truncate, defaults is 2
//
// Returns:
//  number - The truncated value.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Trunc
func Trunc(value interface{}, options ...OptionalParameter) Expr {
	return fn1("trunc", value, options...)
}

// LT returns true if each specified value is less than all the subsequent values. Otherwise LT returns false.
//
// Parameters:
//  args []number - A collection of terms to compare.
//
// Returns:
//  bool - true if all elements are less than each other from left to right.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/LT
func LT(args ...interface{}) Expr { return fn1("lt", varargs(args...)) }

// LTE returns true if each specified value is less than or equal to all subsequent values. Otherwise LTE returns false.
//
// Parameters:
//  args []number - A collection of terms to compare.
//
// Returns:
//  bool - true if all elements are less than of equals each other from left to right.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/LTE
func LTE(args ...interface{}) Expr { return fn1("lte", varargs(args...)) }

// GT returns true if each specified value is greater than all subsequent values. Otherwise GT returns false.
// and false otherwise.
//
// Parameters:
//  args []number - A collection of terms to compare.
//
// Returns:
//  bool - true if all elements are greather than to each other from left to right.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/GT
func GT(args ...interface{}) Expr { return fn1("gt", varargs(args...)) }

// GTE returns true if each specified value is greater than or equal to all subsequent values. Otherwise GTE returns false.
//
// Parameters:
//  args []number - A collection of terms to compare.
//
// Returns:
//  bool - true if all elements are greather than or equals to each other from left to right.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/GTE
func GTE(args ...interface{}) Expr { return fn1("gte", varargs(args...)) }

// And returns the conjunction of a list of boolean values.
//
// Parameters:
//  args []bool - A collection to compute the conjunction of.
//
// Returns:
//  bool - true if all elements are true, false otherwise.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/And
func And(args ...interface{}) Expr { return fn1("and", varargs(args...)) }

// Or returns the disjunction of a list of boolean values.
//
// Parameters:
//  args []bool - A collection to compute the disjunction of.
//
// Returns:
//  bool - true if at least one element is true, false otherwise.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Or
func Or(args ...interface{}) Expr { return fn1("or", varargs(args...)) }

// Not returns the negation of a boolean value.
//
// Parameters:
//  boolean bool - A boolean to produce the negation of.
//
// Returns:
//  bool - The value negated.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Not
func Not(boolean interface{}) Expr { return fn1("not", boolean) }

// Select traverses into the provided value, returning the value at the given path.
//
// Parameters:
//  path []Path - An array representing a path to pull from an object. Path can be either strings or numbers.
//  value Object - The object to select from.
//
// Optional parameters:
//  default Value - A default value if the path does not exist. See Default() function.
//
// Returns:
//  Value - The value at the given path location.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Select
func Select(path, value interface{}, options ...OptionalParameter) Expr {
	return fn2("select", path, "from", value, options...)
}

// SelectAll traverses into the provided value flattening all values under the desired path.
//
// Parameters:
//  path []Path - An array representing a path to pull from an object. Path can be either strings or numbers.
//  value Object - The object to select from.
//
// Returns:
//  Value - The value at the given path location.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/SelectAll
func SelectAll(path, value interface{}) Expr {
	return fn2("select_all", path, "from", value)
}

// ToString attempts to convert an expression to a string literal.
//
// Parameters:
//   value Object - The expression to convert.
//
// Returns:
//   string - A string literal.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/ToString
func ToString(value interface{}) Expr {
	return fn1("to_string", value)
}

// ToNumber attempts to convert an expression to a numeric literal -
// either an int64 or float64.
//
// Parameters:
//   value Object - The expression to convert.
//
// Returns:
//   number - A numeric literal.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/ToNumber
func ToNumber(value interface{}) Expr {
	return fn1("to_number", value)
}

// ToTime attempts to convert an expression to a time literal.
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   time - A time literal.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/ToTime
func ToTime(value interface{}) Expr {
	return fn1("to_time", value)
}

// Converts a time expression to seconds since the UNIX epoch.
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   time - A time literal.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/ToSeconds
func ToSeconds(value interface{}) Expr {
	return fn1("to_seconds", value)
}

// Converts a time expression to milliseconds since the UNIX epoch.
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   time - A time literal.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/ToMillis
func ToMillis(value interface{}) Expr {
	return fn1("to_millis", value)
}

// Converts a time expression to microseconds since the UNIX epoch.
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   time - A time literal.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/ToMicros
func ToMicros(value interface{}) Expr {
	return fn1("to_micros", value)
}

// Returns the time expression's year, following the ISO-8601 standard.
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   time - year.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Year
func Year(value interface{}) Expr {
	return fn1("year", value)
}

// Returns a time expression's month of the year, from 1 to 12.
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   time - Month.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Month
func Month(value interface{}) Expr {
	return fn1("month", value)
}

// Returns a time expression's hour of the day, from 0 to 23.
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   time - year.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Hour
func Hour(value interface{}) Expr {
	return fn1("hour", value)
}

// Returns a time expression's minute of the hour, from 0 to 59.
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   time - year.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Minute
func Minute(value interface{}) Expr {
	return fn1("minute", value)
}

// Returns a time expression's second of the minute, from 0 to 59.
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   time - year.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/Second
func Second(value interface{}) Expr {
	return fn1("second", value)
}

// Returns a time expression's day of the month, from 1 to 31.
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   time - day of month.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/DayOfMonth
func DayOfMonth(value interface{}) Expr {
	return fn1("day_of_month", value)
}

// Returns a time expression's day of the week following ISO-8601 convention, from 1 (Monday) to 7 (Sunday).
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   time - day of week.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/DayOfWeek
func DayOfWeek(value interface{}) Expr {
	return fn1("day_of_week", value)
}

// Returns a time expression's day of the year, from 1 to 365, or 366 in a leap year.
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   time - Day of the year.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/DateOfYear
func DayOfYear(value interface{}) Expr {
	return fn1("day_of_year", value)
}

// ToDate attempts to convert an expression to a date literal.
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   date - A date literal.
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/ToDate
func ToDate(value interface{}) Expr {
	return fn1("to_date", value)
}
