# OTPWEB-WRAPPER-API
Ini adalah unofficial [OtpWeb Wrapper API](https://otpweb.com/api-docs).

### Still Development

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

### Stay tune for update!