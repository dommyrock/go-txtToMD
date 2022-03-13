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

https://user-images.githubusercontent.com/32032778/156761329-e724cc70-f7e5-4c0b-a705-2e237283d0e7.mp4


![image](https://user-images.githubusercontent.com/32032778/155892201-6a589bf5-09c8-4603-902c-f435106ae65d.png)

---

![Available mappings](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/noukkyfp0nsq5axm710x.png) 


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
