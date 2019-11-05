package main

import "github.com/hardstylez72/pgsqltostruct/pkg"

var data = `
    id uuid primary key,
    user_id uuid references bb.users(id),
    created_at timestamp default now(),
    is_active bool,
    parent_email_id uuid null,
    address varchar(256)
`

func main() {
	out, err := pkg.Convert(data)
	if err != nil {
		panic(err)
	}

	println(out)
}