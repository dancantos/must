/*
Package must implements a generic version of the Must methods you see all over
golang libraries.

	// regex
	regexp.MustCompile()
	must.Must(regexp.Compile())

	// template
	template.Must(template.ParseFS())
	must.Must(template.ParseFS)
*/
package must
