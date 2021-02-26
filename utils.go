package email

func uniqueStringList(strList ...string) []string {
	var (
		count         = len(strList)
		uniqueStrMap  = make(map[string]struct{}, count)
		uniqueStrList = make([]string, 0, count)
	)
	for _, str := range strList {
		if _, exist := uniqueStrMap[str]; !exist {
			uniqueStrList = append(uniqueStrList, str)
			uniqueStrMap[str] = struct{}{}
		}
	}
	return uniqueStrList
}
