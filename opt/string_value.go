package opt

type StringValue string

func (a *StringValue) Get() interface{} { return string(*a) }

func (a *StringValue) Set(s string) error {
	*a = StringValue(s)
	return nil
}

func (a *StringValue) String() string {
	return string(*a)
}
