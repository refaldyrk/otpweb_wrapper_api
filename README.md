# OTPWEB-WRAPPER-API
Ini adalah unofficial [OtpWeb Wrapper API](https://otpweb.com/api-docs).


## Install
```bash
go get github.com/refaldyrk/otpweb_wrapper_api
```

## Usage

<h3>Create Wrapper...</h3>

```go
w := wrapper.NewWrapper("apikey")
```

### Get Balance
```go
result, err := w.GetBalance()
```

### Get Country
```go
country, err := w.GetCountry()
```

### Get Operator
```go
operator, _ := w.GetOperator()
```

### Get Service
```go
c, err := w.GetCountry()
s, err := w.GetService(c.Data[0].ID)
```
Param
- countryID -> int

### Get Special Service
```go
ss, err := w.GetSpecialService()
```

### Get Number
```go
n, err := w.GetNumber(6, "go", "")
```
Param
- countryID -> int
-  serviceID -> string
- operator (optional) -> string

### Get Special Number
```go
resp, err := w.GetSpecialNumber("8")
```
-  serviceID -> string

### Get Status
```go
resp, err := w.GetStatus("orderID")
```
- orderID -> string

### Change Status
```go
resp, err := w.ChangeStatus("orderID", "cancel")
```
- orderID -> string
- status -string
### Stay tune for update!