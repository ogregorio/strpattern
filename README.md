# strpattern
The words found in your file will be allocated in the template defined by you, respecting the {?} icon.

## USE
strpattern --origin [path] --destiny [path] --regex [regex]

## INFO
The words found in your file will be allocated in the template defined by you, respecting the {?} icon.\
	
## EXAMPLE
Origin: apple banana grape\
Pattern: {?} is a fruit.\
Destiny:\
&nbsp;apple is a fruit.\
&nbsp;banana is a fruit.\
&nbsp;grape is a fruit.

### COMPILE
Prerequisites: GO Lang Environment\
Next step: go build strpattern.go\
