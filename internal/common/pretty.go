package common

import (
	"github.com/muesli/termenv"
	"strings"
)

var profile *termenv.Profile

func Profile() termenv.Profile {
	if profile == nil {
		p := termenv.ColorProfile()
		profile = &p
	}
	return *profile
}

func StyleInfo() termenv.Style {
	return termenv.String("INFO |").Foreground(Profile().Color("0")).Background(Profile().Color("#71BEF2"))
}

func StyleWarn() termenv.Style {
	return termenv.String("WARN |").Foreground(Profile().Color("0")).Background(Profile().Color("#DBAB79"))
}

func StyleDebug() termenv.Style {
	return termenv.String("DBUG |").Foreground(Profile().Color("0")).Background(Profile().Color("#B9BFCA"))
}

func StyleUpdate() termenv.Style {
	return termenv.String("UPDT |").Foreground(Profile().Color("0")).Background(Profile().Color("#D290E4"))
}

func StyleServe() termenv.Style {
	return termenv.String("SRVE |").Foreground(Profile().Color("0")).Background(Profile().Color("#66C2CD"))
}

func StyleCache() termenv.Style {
	return termenv.String("CCHE |").Foreground(Profile().Color("0")).Background(Profile().Color("#D290E4"))
}

func WordClient() termenv.Style {
	return termenv.String("Client").Foreground(Profile().Color("#DBAB79"))
}

func WordServer() termenv.Style {
	return termenv.String("Server").Foreground(Profile().Color("#D290E4"))
}

func PrettyLimit(in string, max int) string {
	in = strings.ReplaceAll(in, "\n", "[lb]")
	if len(in) > max {
		in = in[:max-3] + "..."
	}
	return in
}

func Color(style, color string) termenv.Style {
	return termenv.String(style).Foreground(Profile().Color("#" + color))
}
