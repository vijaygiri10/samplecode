package helpers

const (
	StatusAccepted            = 2002
	StatusUnauthorized        = 4001
	StatusInternalServerError = 5000
)

var statusText = map[int]string{
	StatusAccepted:            "Received",
	StatusUnauthorized:        "Unauthorized",
	StatusInternalServerError: "Some thing went wrong. We are looking into the issue!",
}

func StatusText(code int) string {
	return statusText[code]
}
