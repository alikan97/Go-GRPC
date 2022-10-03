# NOTES
 - Can `require` min. version of modules, `exclude` module versions, `replace using " => "` to replace module/version with other content or `retract` versions (deprecated versions)
 - Can install specific commit (taking the hash) version of a package
 - Can run multiple modules in a shared workspace `go work init <directory>`. Then apply keyword `go work use <directory>` to switch between modules.
 - <strong> Types IMPLICITLY implement from interfaces by implementing its methods </strong>
    ### Example
    
```
type I interface {
    M()
}

type T struct {
    S string
}

func (t T) M() {
    fmt.Println(t.S)
}

fmt.printf(i.M())
```

# DB Migrations
 - Migration files must be sequential with 6 digits, will throw error if files dont have this
 - Use `migrate create -ext sql -dir migrations/ -seq $(filename)` to create a new migration script
 - Use `migrate -source file://migrations/ -database "postgres://postgres:abcd1234@localhost:5432/GRPC?sslmode=disable" -verbose up` to run migrations

# Testing
 - Coverage -> coverage of code
 - Benchmarking -> Compares a metric/point of reference to a standard amount agreed by the business (performance...)
 - Parallelism -> Tests different versions/components
 - Teardown -> Cleanup after test cases

# Types
 - Initialising empty pointers will return null, not error
 - Type assertions to say that interface's value holds the concrete type T $(interfaceName).(T) 