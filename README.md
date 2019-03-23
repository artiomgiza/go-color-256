# go-color-256

To see how the colors look on your term, run:
```
for code in {0..255}; do echo -e "\e[38;05;${code}m $code: COLOR"; done
```

Example:
![colors example](https://github.com/artiomgiza/go-color-256/blob/master/colors.png)

+ nice [stackoverflow](https://stackoverflow.com/questions/5947742/how-to-change-the-output-color-of-echo-in-linux/) Q/A about term colors
