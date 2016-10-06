package faunadb

const (
	ActionCreate = "create"
	ActionDelete = "delete"
)

const (
	TimeUnitSecond      = "second"
	TimeUnitMillisecond = "millisecond"
	TimeUnitMicrosecond = "microsecond"
	TimeUnitNanosecond  = "nanosecond"
)

// Helper functions

func varargs(expr ...interface{}) interface{} {
	if len(expr) == 1 {
		return expr[0]
	}

	return expr
}

// Optional parameters

type OptionalParameter interface {
	params() map[string]interface{}
}

type params map[string]interface{}

func (p params) params() map[string]interface{} {
	return p
}

func withOptions(f fn, optionals []OptionalParameter) Expr {
	for _, option := range optionals {
		for k, v := range option.params() {
			f[k] = v
		}
	}

	return f
}

func Events(events interface{}) OptionalParameter   { return params{"events": events} }
func TS(timestamp interface{}) OptionalParameter    { return params{"ts": timestamp} }
func After(ref interface{}) OptionalParameter       { return params{"after": ref} }
func Before(ref interface{}) OptionalParameter      { return params{"before": ref} }
func Sources(sources interface{}) OptionalParameter { return params{"sources": sources} }
func Size(size interface{}) OptionalParameter       { return params{"size": size} }
func Separator(sep interface{}) OptionalParameter   { return params{"separator": sep} }

// Basic forms

func Do(exprs ...interface{}) Expr          { return fn{"do": varargs(exprs...)} }
func If(cond, then, elze interface{}) Expr  { return fn{"if": cond, "then": then, "else": elze} }
func Lambda(varName, expr interface{}) Expr { return fn{"lambda": varName, "expr": expr} }
func Let(bindings Obj, in interface{}) Expr { return fn{"let": fn(bindings), "in": in} }
func Var(name string) Expr                  { return fn{"var": name} }

// Collections

func Map(coll, lambda interface{}) Expr     { return fn{"map": lambda, "collection": coll} }
func Foreach(coll, lambda interface{}) Expr { return fn{"foreach": lambda, "collection": coll} }
func Filter(coll, lambda interface{}) Expr  { return fn{"filter": lambda, "collection": coll} }
func Take(num, coll interface{}) Expr       { return fn{"take": num, "collection": coll} }
func Drop(num, coll interface{}) Expr       { return fn{"drop": num, "collection": coll} }
func Prepend(elems, coll interface{}) Expr  { return fn{"prepend": elems, "collection": coll} }
func Append(elems, coll interface{}) Expr   { return fn{"append": elems, "collection": coll} }

// Read

func Get(ref interface{}, options ...OptionalParameter) Expr {
	return withOptions(fn{"get": ref}, options)
}

func Exists(ref interface{}, options ...OptionalParameter) Expr {
	return withOptions(fn{"exists": ref}, options)
}

func Count(set interface{}, options ...OptionalParameter) Expr {
	return withOptions(fn{"count": set}, options)
}

func Paginate(set interface{}, options ...OptionalParameter) Expr {
	return withOptions(fn{"paginate": set}, options)
}

// Write

func Create(ref, params interface{}) Expr    { return fn{"create": ref, "params": params} }
func CreateClass(params interface{}) Expr    { return fn{"create_class": params} }
func CreateDatabase(params interface{}) Expr { return fn{"create_database": params} }
func CreateIndex(params interface{}) Expr    { return fn{"create_index": params} }
func CreateKey(params interface{}) Expr      { return fn{"create_key": params} }
func Update(ref, params interface{}) Expr    { return fn{"update": ref, "params": params} }
func Replace(ref, params interface{}) Expr   { return fn{"replace": ref, "params": params} }
func Delete(ref interface{}) Expr            { return fn{"delete": ref} }

func Insert(ref, ts, action, params interface{}) Expr {
	return fn{"insert": ref, "ts": ts, "action": action, "params": params}
}

func Remove(ref, ts, action interface{}) Expr {
	return fn{"remove": ref, "ts": ts, "action": action}
}

// String

func Concat(terms interface{}, options ...OptionalParameter) Expr {
	return withOptions(fn{"concat": terms}, options)
}

func Casefold(str interface{}) Expr {
	return fn{"casefold": str}
}

// Time and Date

func Time(str interface{}) Expr        { return fn{"time": str} }
func Date(str interface{}) Expr        { return fn{"date": str} }
func Epoch(num, unit interface{}) Expr { return fn{"epoch": num, "unit": unit} }

// Others

func Ref(id string) Expr              { return RefV{id} }
func Null() Expr                      { return NullV{} }
func Add(args ...interface{}) Expr    { return fn{"add": varargs(args...)} }
func Modulo(args ...interface{}) Expr { return fn{"modulo": varargs(args...)} }
func Equals(args ...interface{}) Expr { return fn{"equals": varargs(args...)} }
func Match(ref interface{}) Expr      { return fn{"match": ref} }