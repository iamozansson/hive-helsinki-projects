# Weather Station

A Go application completed during my Hive Helsinki programming sprint.

## Objective

The goal of this project was to build a simple weather station state manager that receives partial sensor updates and maintains the full known state of the station.

The application processes sensor data through standard input and can return, clear, or update the current state.

## Task

The program maintains nine meteorological sensors identified by numeric IDs:

| ID | Key             |
| -- | --------------- |
| 1  | `airTemp`       |
| 2  | `airPressure`   |
| 7  | `precipitation` |
| 11 | `windSpeed`     |
| 12 | `windDirection` |
| 13 | `humidity`      |
| 14 | `dewPoint`      |
| 15 | `soilMoisture`  |
| 22 | `cloudCover`    |

Each sensor starts with a `NULL` state.

The program accepts four types of input:

* `id,value` — update a sensor
* `get` — print the complete current state
* `clear` — reset all sensors to `NULL`
* `exit` — terminate the application

### Example

```text
11,15.5
13,32.3
get
```

Output:

```text
airTemp:NULL
airPressure:NULL
precipitation:NULL
windSpeed:15.5
windDirection:NULL
humidity:32.3
dewPoint:NULL
soilMoisture:NULL
cloudCover:NULL
```

## Implementation

The application uses a `Weather` struct to represent each sensor:

```go
type Weather struct {
    id    int
    key   string
    value float64
    check bool
}
```

The `check` field tracks whether a sensor currently has a known value or should be displayed as `NULL`.

The application reads input using `bufio.Reader`, parses sensor IDs and values with `strconv`, and processes commands using a `switch` statement.

The main operations are handled by separate functions:

* `Set` — updates a sensor value and marks it as known
* `SetNull` — marks a sensor as missing
* `Get` — prints the complete sensor state in ID order
* `Clear` — resets all sensors to `NULL`

## Usage

Run the application with:

```bash
go run .
```

The application starts with:

```text
--- Weather Station ---
```

Then enter commands through standard input.

To exit:

```text
exit
```

which prints:

```text
Exiting...
```

## What I Practiced

* Go structs
* Slices of structs
* Reading from standard input
* Buffered I/O with `bufio`
* String parsing
* `strconv` for numeric conversion
* `switch` statements
* State management
* Functions and slice manipulation
* Handling `NULL` values
* Command-line applications
