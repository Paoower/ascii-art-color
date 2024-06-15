package colors

import (
  "strings"
)

func GetColor(s string) string {
  Colors := make(map[string]string)
  Colors["RED"] = "\033[31;1m"
  Colors["GREEN"] = "\033[32;1m"
  Colors["YELLOW"] = "\033[33;1m"
  Colors["ORANGE"] = "\033[33;1m"
  Colors["BLUE"] = "\033[34;1m"
  Colors["PURPLE"] = "\033[35;1m"
  Colors["CYAN"] = "\033[36;1m"
  Colors["GREY"] = "\033[37;1m"
  
  s = strings.TrimPrefix(s, "--color=")

  if strings.Contains(s, "rgb") {
    co := strings.ReplaceAll(strings.TrimPrefix(s, "rgb("), " ", "")
    co = strings.ReplaceAll(co, ")", "")
    c := strings.Split(co, ",")
    return "\033[38;2;"+c[0]+";"+c[1]+";"+c[2]+"m"
  }

  return Colors[strings.ToUpper(s)]
}
