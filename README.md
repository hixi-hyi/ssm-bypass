# ssm-bypass
This binary will export the path of SSM ParameterStore as an environment variables.
For example, it can be used in entrypoint.sh.

SSM Parameters put
```
/your/pass/HOGE → hoge
/your/pass/FUGA → fuga
```

`$(ssm-bypass /your/pass)` do
```
export HOGE=hoge
export FUGA=fuga
```

# usage
1. copy binary for your system from bin
2. write the bellow code into your entrypoint.sh
```
$(ssm-bypass /your/path/)
```

# development
```
rm -rf bin/
go get github.com/mitchellh/gox
gox --output="bin/{{.Dir}}_{{.OS}}_{{.Arch}}"
```
