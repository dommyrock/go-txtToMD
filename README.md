### Transform .txt file to MD and HTML

#### Initally made for generating simple docs from Team sync notes


#### Runing on Test data (test.txt file)

```
go run . .\testData\test.txt
```

#### Passing custom theme params
> light (default)

> mid 

> dark

```
go run . .\testData\test.txt dark
```
#### Get all available mappings (-o or -options)
```
go run . -options
```
---
## Demo

https://user-images.githubusercontent.com/32032778/156761329-e724cc70-f7e5-4c0b-a705-2e237283d0e7.mp4


![image](https://user-images.githubusercontent.com/32032778/155892201-6a589bf5-09c8-4603-902c-f435106ae65d.png)

---

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
