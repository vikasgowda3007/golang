## Dynamic XML Processing Using GoLang

## Aim of this package is to dynamically process any XML data without having to pre-define a struct with the xml tags.


#### Package
```
    import github.com/vikasgowda3007/golang/dynamicXmlProcessing/core"
```

func GetXmlMap
```
    GetXmlMap(xmlStr string, sep string) (xmlMap map[string]string, err error)
```
GetXmlMap accepts a xml string the first parameter and a separator as the second parameter.
Will return a map of xml data where recursive xml tag keys separated by a separator passed as parameter.
If processing is unsuccessful xmlMap will be set to nil and err is returned.