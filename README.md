# strpattern
The words found in your file will be allocated in the template defined by you, respecting the {?} icon.

## USE
strpattern --origin [path] --destiny [path] --regex [regex]

## INFO
The words found in your file will be allocated in the template defined by you, respecting the {?} icon.&nbsp;
	
## EXAMPLE
Origin: apple banana grape&nbsp;
Pattern: {?} is a fruit.&nbsp;
Destiny:&nbsp;
    apple is a fruit.&nbsp;
    banana is a fruit.&nbsp;
    grape is a fruit.&nbsp;

### COMPILE
Prerequisites: GO Lang Environment&nbsp;
Next step: go build strpattern.go&nbsp;
