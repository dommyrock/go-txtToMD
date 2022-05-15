### Generate HTML and MD files from .txt file
[Blog post](https://dev.to/dompolzer/generating-hmtl-and-md-files-from-txt-in-go-59lh)

#### Running code locally: (with test files form ./testData directory)
> Generate Html/md files to %HOMEDIR%/Downloads directory:
```
go run . ./testData/test.txt
```
> selecting custom theme:
```
go run . ./testData/test.txt mid
```
> Getting all currently available mappings:
```
go run . -options (or -o)
```

> Note: Currently line break/newline is expected between each mapping element to be parsed into MD/HTML correctly!

---
#### Running CLI command globally:
If you want to run this tool globally you can place txtToMD.exe file to 'Path' in windows environment variables.
that way you can use it from anywhere in your system
> txtToMD {txt File Path} [options: light(default),dark,mid]
```
txtToMD D:\Desktop\test.txt dark
```
> Getting all currently available mappings:
```
txtToMD -options
```

#### Passing custom theme params
> light (default)

> mid 

> dark

---
## Demo

https://user-images.githubusercontent.com/32032778/168468294-d0e18c5a-ff80-427b-98a2-755679d43f39.mp4

![image](https://user-images.githubusercontent.com/32032778/155892201-6a589bf5-09c8-4603-902c-f435106ae65d.png)

---

![Available mappings](https://user-images.githubusercontent.com/32032778/168449622-ee2f7869-cf1b-4351-a40a-4874f0638adf.png)



### Currently available  mappings
> #h4

> #code

> #b

> #p

> #h3

> #bp

> #links

> #h5

> #h2

> #link

> #h1

> #img

> #table

> '-' (dash, for solid line break)
