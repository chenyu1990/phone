package schema

type PhoneInfo struct {
	Number          string
	Province        string
	City            string
	Area            string
	ServiceProvider string
}

type PhoneInfos []*PhoneInfo
