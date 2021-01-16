package replacer

import ("strings")

//Replace in phrase where has {?}
func Replace(phrase string,word string) string {
	return strings.ReplaceAll(phrase,"{?}",word)
}