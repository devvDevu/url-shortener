package custom

import (
	"fmt"
	"strings"

	"github.com/gorilla/mux"
)

func PrintRoutesTable(router *mux.Router) {
	var routes []struct {
		Method string
		Path   string
	}

	_ = router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, _ := route.GetPathTemplate()
		methods, _ := route.GetMethods()
		if len(methods) == 0 {
			methods = []string{"ANY"}
		}
		for _, method := range methods {
			routes = append(routes, struct {
				Method string
				Path   string
			}{
				Method: method,
				Path:   path,
			})
		}
		return nil
	})

	// Строим таблицу
	table := buildRoutesTable(routes)
	for _, line := range table {
		fmt.Println(line)
	}
}

func buildRoutesTable(routes []struct {
	Method string
	Path   string
}) []string {
	maxMethodLen := 7 // "METHOD".length
	maxPathLen := 4   // "PATH".length

	// Находим максимальные длины
	for _, r := range routes {
		if len(r.Method) > maxMethodLen {
			maxMethodLen = len(r.Method)
		}
		if len(r.Path) > maxPathLen {
			maxPathLen = len(r.Path)
		}
	}

	// Строим строки таблицы
	header := fmt.Sprintf("┌─%s─┬─%s─┐",
		strings.Repeat("─", maxMethodLen),
		strings.Repeat("─", maxPathLen))

	separator := fmt.Sprintf("├─%s─┼─%s─┤",
		strings.Repeat("─", maxMethodLen),
		strings.Repeat("─", maxPathLen))

	footer := fmt.Sprintf("└─%s─┴─%s─┘",
		strings.Repeat("─", maxMethodLen),
		strings.Repeat("─", maxPathLen))

	// Форматируем строки
	var lines []string
	lines = append(lines, header)
	lines = append(lines, fmt.Sprintf("│ %-*s │ %-*s │",
		maxMethodLen, "METHOD",
		maxPathLen, "PATH"))
	lines = append(lines, separator)

	for _, r := range routes {
		lines = append(lines, fmt.Sprintf("│ %-*s │ %-*s │",
			maxMethodLen, r.Method,
			maxPathLen, r.Path))
	}

	lines = append(lines, footer)
	return lines
}
